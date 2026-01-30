package git

import (
	"bytes"
	"os/exec"
	"strings"
)

func RepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func IsCurrentDirValidRepo() bool {
	_, err := RepoRoot()
	return err == nil
}

func TodaysCommit(repoRoot string) ([]string, error) {
	cmd := exec.Command("git", "log", "--pretty=format:%H", "--since=midnight")
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
