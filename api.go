package main

import (
    "net/http"
)

func GameVersion(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hi"))
}
