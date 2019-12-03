package main

import (
	"fmt"
	"os"
    "strings"
	"net/http"
    "path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

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
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

func main() {
    router := chi.NewRouter()
    router.Use(middleware.Recoverer)
    router.Use(middleware.Logger)

    router.Get("/api/version", GameVersion)
    router.Get("/api/releases", AllReleases)
    router.Get("/api/releases/latest", LatestRelease)
    router.Get("/api/releases/{version}", CustomRelease)

    workDir, _ := os.Getwd()
    publicDir := filepath.Join(workDir, "public")
    FileServer(router, "/", http.Dir(publicDir))

    port := getPort()
    fmt.Println("Server listening at", port)
    http.ListenAndServe(port, router)
}

