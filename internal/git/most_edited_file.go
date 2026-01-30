package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

type EditedFile struct {
	FileName              string
	AddedLineCount        int
	DeletedLineCount      int
	NumberOfBranchOccured int
}

func MostEditedFiles(repoRoot string) error {
	mostEditedFiles, err := getMostEditedFiles(repoRoot)
	if err != nil {
		return err
	}
	for _, file := range mostEditedFiles {
		fmt.Printf("%s +%d -%d\n", file.FileName, file.AddedLineCount, file.DeletedLineCount)
	}
	return nil
}

func getMostEditedFiles(repoRoot string) ([]EditedFile, error) {
	cmd := exec.Command("git", "log", "--since=midnight", "--numstat", "--pretty=")
	var out bytes.Buffer
	cmd.Dir = repoRoot
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	allFiles := strings.Split(out.String(), "\n")
	filteredFiles := make([]string, 0, len(allFiles))
	for _, f := range allFiles {
		if !strings.HasPrefix(f, "-") {
			filteredFiles = append(filteredFiles, f)
		}
	}
	files := make([]EditedFile, 0, len(filteredFiles))
	for _, f := range filteredFiles {
		temp := strings.Split(f, "\t")
		if len(temp) != 3 {
			continue
		}
		addedLineCount, err := strconv.Atoi(temp[0])
		if err != nil {
			return nil, err
		}
		deletedLineCount, err := strconv.Atoi(temp[1])
		if err != nil {
			return nil, err
		}
		files = append(files, EditedFile{
			FileName:              temp[2],
			AddedLineCount:        addedLineCount,
			DeletedLineCount:      deletedLineCount,
			NumberOfBranchOccured: 1,
		})
	}
	sort.Slice(files, func(i, j int) bool {
		file1LineCount := files[i].AddedLineCount + files[i].DeletedLineCount
		file2LineCount := files[j].AddedLineCount + files[j].DeletedLineCount
		if file1LineCount == file2LineCount {
			return files[i].FileName < files[j].FileName
		}
		return file1LineCount > file2LineCount
	})

	return files, nil
}
