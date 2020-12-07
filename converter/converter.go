package converter

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// StatusMsg status message
type StatusMsg struct {
	Name    string
	Message string
	IsError bool
}

const (
	// Badges file name
	Badges = "Badges.xml"
	// Comments file name
	Comments = "Comments.xml"
	// PostHistory file name
	PostHistory = "PostHistory.xml"
	// PostLinks file name
	PostLinks = "PostLinks.xml"
	// Posts file name
	Posts = "Posts.xml"
	// Tags file name
	Tags = "Tags.xml"
	// Users file name
	Users = "Users.xml"
	// Votes file name
	Votes = "Votes.xml"
)

var sourceFiles []string

func init() {
	sourceFiles = []string{Badges, Comments, PostHistory, PostLinks, Posts, Tags, Users, Votes}
}

// Convert sourcePath xml to csv and store at storeToDir
func Convert(sourcePath string, storeToDir string) (err error) {
	sourcePathResolved, err := resolvePath(sourcePath)
	if err != nil {
		return
	}

	sourceFiles, err := getSourceFiles(sourcePathResolved)
	if err != nil {
		return
	}

	if len(sourceFiles) == 0 {
		err = errors.New("Nothing to convert. Please specify correct source path to the extracted XML files")
		return
	}

	if storeToDir != "" {
		storeTo, err := resolvePath(storeToDir)
		if err != nil {
			return err
		}
		var fi os.FileInfo
		fi, err = os.Stat(storeTo)
		if err != nil {
			return err
		}

		if fi.Mode().IsDir() == false {
			err = errors.New("Extract path has to be a directory, not a file")
			return err
		}
	} else {
		storeToDir = sourcePathResolved
	}

	var wg sync.WaitGroup
	for _, sf := range sourceFiles {
		f := filepath.Base(sf)
		fName := f[:len(f)-len(filepath.Ext(f))]
		csvFileName := fName + ".csv"
		go convertXMLFile(&wg, fName, sf, filepath.Join(storeToDir, csvFileName))
	}

	wg.Wait()

	return
}

func convertXMLFile(wg *sync.WaitGroup, typeName string, xmlFilePath string, csvFilePath string) {
	wg.Add(1)
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		// @TODO send message to chan
		fmt.Println(err)
	}
	defer xmlFile.Close()

	csvFile, err := os.Create(csvFilePath)
	if err != nil {
		// @TODO send message to chan
		fmt.Println(err)
	}
	defer csvFile.Close()

	iterate(typeName, xmlFile, csvFile)
	wg.Done()
}

func resolvePath(path string) (string, error) {
	p := strings.TrimSpace(path)
	if len(p) == 0 {
		return "", nil
	}

	if filepath.IsAbs(p) {
		return p, nil
	}

	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(workDir, p), nil
}

func isSourceFile(fileName string) bool {
	for _, sf := range sourceFiles {
		if sf == fileName {
			return true
		}
	}
	return false
}

func getSourceFiles(sourcePath string) (sourceFiles []string, err error) {
	fi, err := os.Stat(sourcePath)
	if err != nil {
		return
	}

	switch mode := fi.Mode(); {
	case mode.IsDir():
		files, _ := ioutil.ReadDir(sourcePath)
		for _, f := range files {
			if isSourceFile(f.Name()) {
				fullPath := filepath.Join(sourcePath, f.Name())
				sourceFiles = append(sourceFiles, fullPath)
			}
		}

	case mode.IsRegular():
		if isSourceFile(fi.Name()) {
			sourceFiles = append(sourceFiles, sourcePath)
		}

	}
	return
}
