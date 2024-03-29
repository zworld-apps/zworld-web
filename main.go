package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var client *github.Client
var publicPath = "public"

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

func HandleNotFound(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(publicPath, "404.html"))
}

func main() {
	if len(os.Args) == 2 {
		publicPath = os.Args[1]
	}

	initGithubClient()

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	//router.Use(middleware.Logger)
	router.Use(middleware.DefaultCompress)

	router.NotFound(HandleNotFound)

	router.Mount("/api/v1", apiV1Router())
	router.Mount("/api", apiV1Router())

	workDir, _ := os.Getwd()
	publicDir := filepath.Join(workDir, publicPath)
	FileServer(router, "/", publicDir)

	port := getPort()
	fmt.Println("Server listening at", port)
	http.ListenAndServe(port, router)
}
