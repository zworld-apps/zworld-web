package main

import (
	"context"
	"errors"

	"github.com/google/go-github/github"
)

func GetLatestRelease() (*github.RepositoryRelease, error) {
	ctx := context.Background()
	// cant use github.GetLatestRelease because it doesnt count pre-releases
	releases, _, err := client.Repositories.ListReleases(ctx,
		"xzebra", "zworld-client", nil)

	if err != nil || len(releases) == 0 {
		return nil, err
	}

	// get latest release according to max tag
	latest := releases[0]
	for _, release := range releases[1:] {
		if release.GetTagName() > latest.GetTagName() {
			latest = release
		}
	}
	return latest, nil
}

func GetReleaseList() ([]*github.RepositoryRelease, error) {
	ctx := context.Background()
	// cant use github.GetLatestRelease because it doesnt count pre-releases
	releases, _, err := client.Repositories.ListReleases(ctx,
		"xzebra", "zworld-client", nil)

	return releases, err
}

func GetDownloadURL(id int64) (string, error) {
	ctx := context.Background()
	_, redirect, err := client.Repositories.DownloadReleaseAsset(ctx,
		"xzebra", "zworld-client", id)
	if redirect == "" && err == nil {
		err = errors.New("couldn't get download url")
	}
	return redirect, err
}
