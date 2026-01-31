package git

import (
	"os/exec"
	"strings"
)

func getAuthor() (string, error) {
	cmd, err := exec.Command("git", "config", "user.name").Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(cmd)), nil
}

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

func getRepoNameFromRoot(repoRoot string) string {
	parts := strings.Split(repoRoot, "/")
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}
