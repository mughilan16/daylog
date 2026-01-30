package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func TodaysCommit(repoRoot string) error {
	todaysCommit, err := getTodaysCommit(repoRoot)
	if err != nil {
		return err
	}
	fmt.Println("TODAY'S Commits")
	for _, commit := range todaysCommit {
		fmt.Println(commit)
	}
	return nil
}

func getTodaysCommit(repoRoot string) ([]string, error) {
	author, err := getAuthor()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command(
		"git",
		"log",
		"--pretty=format:%s",
		"--since=midnight",
		"--author="+author,
	)
	cmd.Dir = repoRoot
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	raw := strings.TrimSpace(out.String())
	if raw == "" {
		return []string{}, nil
	}
	return strings.Split(raw, "\n"), nil
}
