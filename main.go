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
	todayCommits, err := git.TodaysCommit(repoRoot)
	if err != nil {
		logger.Debug(err.Error())
		logger.Fatal("Failed to calculate today's commits")
	}
	for i := 0; i < len(todayCommits); i++ {
		logger.Info(todayCommits[i])
	}
}
