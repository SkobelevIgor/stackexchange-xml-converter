package converter

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dustin/go-humanize"
)

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

// Convert xml files from sourcePath and store csv result file(s) to the storeDir
func Convert(sourcePath string, storeToDir string, skipHTMLDecoding bool) (err error) {

	if sourcePath == "" {
		return errors.New("--source-path flag is required")
	}

	sourcePathResolved, err := resolvePath(sourcePath)
	if err != nil {
		return
	}

	sourceFiles, err := getSourceFiles(sourcePathResolved)
	if err != nil {
		return
	}

	if len(sourceFiles) == 0 {
		return fmt.Errorf(
			"Nothing to convert from %s. Please specify the correct source path to extracted XML files",
			sourcePath)
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
			return fmt.Errorf("Result path [%s] has to be a directory, not a file", storeToDir)
		}
	} else {
		storeToDir = sourcePathResolved
	}

	log.Printf("Total %d file(s) to convert", len(sourceFiles))

	var wg sync.WaitGroup
	for _, sf := range sourceFiles {
		f := filepath.Base(sf)
		fName := f[:len(f)-len(filepath.Ext(f))]
		csvFileName := fName + ".csv"
		wg.Add(1)
		log.Printf("[%s] Converting is started", fName)
		go convertXMLFile(&wg, fName, sf, filepath.Join(storeToDir, csvFileName), skipHTMLDecoding)
	}

	wg.Wait()

	return
}

func convertXMLFile(wg *sync.WaitGroup, typeName string, xmlFilePath string, csvFilePath string, skipHTMLDecoding bool) {
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		log.Printf("[%s] Error: %s", typeName, err)
		return
	}
	defer xmlFile.Close()

	csvFile, err := os.Create(csvFilePath)
	if err != nil {
		log.Printf("[%s] Error: %s", typeName, err)
		return
	}
	defer csvFile.Close()

	total, converted, err := iterate(typeName, xmlFile, csvFile, skipHTMLDecoding)
	if err != nil {
		log.Printf("[%s]. Error: %s. Skipping the file.", typeName, err)
	} else {
		log.Printf("[%s] File is converted. %s of %s row(s) has been processed successfully",
			typeName, humanize.Comma(total), humanize.Comma(converted))
	}
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
