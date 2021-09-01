package main

import (
	"github.com/gookit/config"
	"github.com/gookit/slog"
)

func main() {
	slog.Info("hello")

	// Inner ENV see: https://docs.github.com/en/actions/configuring-and-managing-workflows/using-environment-variables
	// GITHUB_REF
	// - on PR: "refs/pull/:prNumber/merge"
	// - on push: "refs/heads/master"
	// - on push tag: "refs/tags/v0.0.1"
	ghRefer := config.Getenv("GITHUB_REF")
	// ghRefer := "refs/pull/34/merge"
	prNumber := getPrNumber(ghRefer)
	if prNumber == 0 {
		slog.Fatalf("parse PR number failed, GITHUB_REF: %s", ghRefer)
	}

	slog.Info("Successful")
}
