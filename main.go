package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var client *github.Client

func initGithubClient() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
}

func getPort() (port string) {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No port variable detected, setting to", port)
	}
	return ":" + port
}

// Got from https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
//
// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(router *chi.Mux, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		router.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	router.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(fmt.Sprintf("%s%s", root, r.RequestURI)); os.IsNotExist(err) {
			router.NotFoundHandler().ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	}))
}

func ServeZIP(w http.ResponseWriter, file io.ReadCloser) error {
	w.Header().Set("Content-Type", "application/zip")
	_, err := io.Copy(w, file)
	return err
}

func main() {
	initGithubClient()

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/404.html")
	})

	router.Mount("/api/v1", apiV1Router())
	router.Mount("/api", apiV1Router())

	workDir, _ := os.Getwd()
	publicDir := filepath.Join(workDir, "public")
	FileServer(router, "/", http.Dir(publicDir))

	port := getPort()
	fmt.Println("Server listening at", port)
	http.ListenAndServe(port, router)
}
