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

// Config Initial config handler
type Config struct {
	ResultFormat     string
	SourcePath       string
	StoreToDir       string
	SkipHTMLDecoding bool
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

var (
	sourceFiles     []string
	converterConfig Config
)

func init() {
	sourceFiles = []string{Badges, Comments, PostHistory, PostLinks, Posts, Tags, Users, Votes}
}

// Convert xml files from sourcePath and store csv result file(s) to the storeDir
// func Convert(sourcePath string, storeToDir string, skipHTMLDecoding bool) (err error) {
func Convert(cfg Config) (err error) {

	converterConfig = cfg

	if cfg.SourcePath == "" {
		return errors.New("--source-path flag is required")
	}

	if cfg.ResultFormat == "" {
		return errors.New("--result-format flag is required")
	}

	cfg.ResultFormat = strings.ToLower(cfg.ResultFormat)
	if cfg.ResultFormat != "csv" && cfg.ResultFormat != "json" {
		return fmt.Errorf("Unknown result format %s. Expected: csv or json", cfg.ResultFormat)
	}

	sourcePathResolved, err := resolvePath(cfg.SourcePath)
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
			cfg.SourcePath)
	}

	if cfg.StoreToDir != "" {
		storeTo, err := resolvePath(cfg.StoreToDir)
		if err != nil {
			return err
		}
		var fi os.FileInfo
		fi, err = os.Stat(storeTo)
		if err != nil {
			return err
		}

		if fi.Mode().IsDir() == false {
			return fmt.Errorf("Result path [%s] has to be a directory, not a file", cfg.StoreToDir)
		}
	} else {
		cfg.StoreToDir = sourcePathResolved
	}

	log.Printf("Total %d file(s) to convert", len(sourceFiles))

	var wg sync.WaitGroup
	for _, sf := range sourceFiles {
		f := filepath.Base(sf)
		typeName := f[:len(f)-len(filepath.Ext(f))]
		resultFile := filepath.Join(cfg.StoreToDir,
			fmt.Sprintf("%s.%s", typeName, cfg.ResultFormat))
		wg.Add(1)
		log.Printf("[%s] Converting is started", typeName)
		go convertXMLFile(&wg, typeName, sf, resultFile)
	}

	wg.Wait()

	return
}

func convertXMLFile(wg *sync.WaitGroup, typeName string, xmlFilePath string, resultFilePath string) {
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		log.Printf("[%s] Error: %s", typeName, err)
		return
	}
	defer xmlFile.Close()

	resultFile, err := os.Create(resultFilePath)
	if err != nil {
		log.Printf("[%s] Error: %s", typeName, err)
		return
	}
	defer resultFile.Close()

	var total, converted int64
	switch converterConfig.ResultFormat {
	case "csv":
		total, converted, err = convertToCSV(typeName, xmlFile, resultFile, converterConfig)
	case "json":
		total, converted, err = convertToJSON(typeName, xmlFile, resultFile, converterConfig)
	}

	if err != nil {
		log.Printf("[%s]. Error: %s. Skipping the file.", typeName, err)
	} else {
		log.Printf("[%s] File is converted. %s of %s row(s) has been processed successfully",
			typeName, humanize.Comma(converted), humanize.Comma(total))
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
