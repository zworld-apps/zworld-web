package main

import (
    "fmt"
    "net/http"
)

type RespError struct {
    Code int `json:"code"`
    Desc string `json:"desc"`
}

func GameVersion(w http.ResponseWriter, r *http.Request) {
    latest, err := GetLatestRelease()
    if err != nil {
        fmt.Fprint(w, err.Error())
        return
    }
    fmt.Fprint(w, latest.GetTagName())
}

func AllReleases(w http.ResponseWriter, r *http.Request) {
    
}

func LatestRelease(w http.ResponseWriter, r *http.Request) {
    
}

func CustomRelease(w http.ResponseWriter, r *http.Request) {
    
}
