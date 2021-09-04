package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/google/go-github/github"
)

type RespError struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

func (rr RespError) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type RespRelease struct {
	TagName string      `json:"tag_name"`
	Name    string      `json:"name"`
	Date    string      `json:"release_date"`
	Body    string      `json:"body"`
	Assets  []RespAsset `json:"assets"`
}

func NewRespRelease(release *github.RepositoryRelease) RespRelease {

	releaseData := RespRelease{
		TagName: release.GetTagName(),
		Name:    release.GetName(),
		Date:    release.GetPublishedAt().Format("02 January 2006"),
		Body:    release.GetBody(),
	}

	for _, asset := range release.Assets {
		releaseData.Assets = append(releaseData.Assets, RespAsset{
			Name:        asset.GetName(),
			DownloadURL: fmt.Sprintf("/api/v1/download/%d", asset.GetID()),
		})
	}

	return releaseData
}

func (rr RespRelease) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type RespAsset struct {
	Name        string `json:"name"`
	DownloadURL string `json:"download_url"`
}

func apiV1Router() *chi.Mux {
	router := chi.NewRouter()
	router.Use(render.SetContentType(render.ContentTypeJSON))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.Render(w, r, RespError{
			Code: http.StatusNotFound,
			Desc: "not found",
		})
	})

	router.Get("/version", GameVersion)

	router.Get("/login", ServerLogin)

	router.Get("/releases", AllReleases)
	router.Get("/releases/latest", LatestRelease)
	router.Get("/releases/{version}", CustomRelease)

	router.Get("/download/{id}", DownloadRelease)

	return router
}

func GameVersion(w http.ResponseWriter, r *http.Request) {
	latest, err := GetLatestRelease()
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprint(w, latest.GetTagName())
}

func ServerLogin(w http.ResponseWriter, r *http.Request) {
	db.
		fmt.Fprint(w, latest.GetTagName())
}

func AllReleases(w http.ResponseWriter, r *http.Request) {

	releases, err := GetReleaseList()
	if err != nil {
		render.Render(w, r, RespError{
			Code: http.StatusServiceUnavailable,
			Desc: "couldn't get releases list",
		})
		return
	}

	releasesData := []render.Renderer{}
	for _, release := range releases {
		releasesData = append(releasesData, NewRespRelease(release))
	}

	render.RenderList(w, r, releasesData)
}

func LatestRelease(w http.ResponseWriter, r *http.Request) {

	release, err := GetLatestRelease()
	if err != nil {
		render.Render(w, r, RespError{
			Code: http.StatusServiceUnavailable,
			Desc: "couldn't get releases list",
		})
		return
	}

	err = render.Render(w, r, NewRespRelease(release))
	if err != nil {
		log.Println("failed to error response: ", err)
	}

}

func CustomRelease(w http.ResponseWriter, r *http.Request) {

}

func DownloadRelease(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		render.Render(w, r, RespError{
			Code: http.StatusBadRequest,
			Desc: "couldn't parse download id",
		})
		return
	}

	ctx := context.Background()
	file, redirect, err := client.Repositories.DownloadReleaseAsset(ctx,
		"zworld-apps", "zworld-client", id)
	if err != nil {
		render.Render(w, r, RespError{
			Code: http.StatusNotFound,
			Desc: "download not found",
		})
		return
	}

	if redirect != "" {
		// we have been redirected to download page, just redirect the user
		http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
	} else {
		defer file.Close()

		err := ServeZIP(w, file)
		if err != nil {
			render.Render(w, r, RespError{
				Code: http.StatusInternalServerError,
				Desc: "couldn't send release zip",
			})
		}
	}

}
