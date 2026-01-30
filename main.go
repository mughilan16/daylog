package main

import (
	"flag"

	"github.com/mughilan16/daylog/internal/git"
	"github.com/mughilan16/daylog/internal/logger"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug logs")
	flag.Parse()

	logger.Init(*debug)
	logger.Info("Starting DayLog :)")

	if *debug {
		logger.Debug("Debug Enabled")
	}
	isValidGitRepo := git.IsCurrentDirValidRepo()
	if !isValidGitRepo {
		logger.Fatal("Not a valid git repo")
	}

	repoRoot, err := git.RepoRoot()
	if err != nil {
		logger.Debug(err.Error())
		logger.Fatal("Failed to resolve repo root")
	}
	err = git.MostEditedFiles(repoRoot)
	if err != nil {
		logger.Debug(err.Error())
		logger.Fatal("Failed to analyze most edited files")
	}

	err = git.TodaysCommit(repoRoot)
	if err != nil {
		logger.Debug(err.Error())
		logger.Fatal("Failed to analyze today's commits")
	}
}
