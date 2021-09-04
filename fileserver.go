package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi"
)

// Got from https://stackoverflow.com/questions/49589685/good-way-to-disable-directory-listing-with-http-fileserver-in-go
//
// Creates a custom filesystem for the FileServer function so it doesn't serve folders as a listing of files and,
// instead serve a 404 error
type CustomFilesystem struct {
	http.FileSystem
	readDirBatchSize int
}

func (fs *CustomFilesystem) Open(name string) (http.File, error) {
	f, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return &NeuteredStatFile{File: f, readDirBatchSize: fs.readDirBatchSize}, nil
}

type NeuteredStatFile struct {
	http.File
	readDirBatchSize int
}

func (e *NeuteredStatFile) Stat() (os.FileInfo, error) {
	s, err := e.File.Stat()
	if err != nil {
		return nil, err
	}
	if s.IsDir() {
	LOOP:
		for {
			fl, err := e.File.Readdir(e.readDirBatchSize)
			switch err {
			case io.EOF:
				break LOOP
			case nil:
				for _, f := range fl {
					if f.Name() == "index.html" {
						return s, err
					}
				}
			default:
				return nil, err
			}
		}
		return nil, os.ErrNotExist
	}
	return s, err
}

// Got from https://github.com/go-chi/chi/blob/master/_examples/fileserver/main.go
//
// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(router *chi.Mux, path string, root string) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(
		&CustomFilesystem{
			FileSystem:       http.Dir(root),
			readDirBatchSize: 2,
		},
	))

	// redirect to / terminated urls
	if path != "/" && path[len(path)-1] != '/' {
		router.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	router.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check if url has GET parameters
		if strings.Contains(r.RequestURI, "?") {
			// trim parameters as server is not gonna parse them
			r.RequestURI = r.RequestURI[:strings.LastIndex(r.RequestURI, "?")]
			fmt.Println(r.RequestURI)
		}

		info, err := os.Stat(fmt.Sprintf("%s%s", root, r.RequestURI))
		if err == nil && info.IsDir() {
			_, err = os.Stat(fmt.Sprintf("%s%s/index.html", root, r.RequestURI))
		}

		if os.IsNotExist(err) {
			router.NotFoundHandler().ServeHTTP(w, r)
		} else {
	        w.Header().Set("Cache-Control", "max-age=3600")
			fs.ServeHTTP(w, r)
		}
	}))
}

func ServeZIP(w http.ResponseWriter, file io.ReadCloser) error {
	w.Header().Set("Content-Type", "application/zip")
	_, err := io.Copy(w, file)
	return err
}
