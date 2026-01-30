package git

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
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
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "MOST EDITED FILES")
	fmt.Fprintln(w, "FILE\tADD\tDEL\tCOMMITS")
	for _, file := range mostEditedFiles {
		fmt.Fprintf(w, "%s\t+%d\t-%d\t%d\n",
			file.FileName,
			file.AddedLineCount,
			file.DeletedLineCount,
			file.NumberOfBranchOccured,
		)
	}
	w.Flush()
	return nil
}

func getMostEditedFiles(repoRoot string) ([]EditedFile, error) {
	author, err := getAuthor()
	if err != nil {
		return nil, err
	}
	cmd := exec.Command("git", "log", "--since=midnight",
		"--numstat", "--pretty=", "--author="+author)
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
	fileMap := make(map[string]*EditedFile)
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
		file := fileMap[temp[2]]
		if file != nil {
			file.AddedLineCount += addedLineCount
			file.DeletedLineCount += deletedLineCount
			file.NumberOfBranchOccured += 1
		} else {
			fileMap[temp[2]] = &EditedFile{
				FileName:              temp[2],
				AddedLineCount:        addedLineCount,
				DeletedLineCount:      deletedLineCount,
				NumberOfBranchOccured: 1,
			}
		}
	}
	files := make([]EditedFile, 0, len(fileMap))
	for _, f := range fileMap {
		files = append(files, *f)
	}

	sort.Slice(files, func(i, j int) bool {
		file1LineCount := files[i].AddedLineCount + files[i].DeletedLineCount
		file2LineCount := files[j].AddedLineCount + files[j].DeletedLineCount
		if file1LineCount == file2LineCount {
			return files[i].FileName < files[j].FileName
		}
		return file1LineCount > file2LineCount
	})

	limit := min(5, len(files))
	return files[:limit], nil
}
