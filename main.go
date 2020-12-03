package main

import (
	"flag"
	"log"

	"github.com/SkobelevIgor/stackexchange-xml-to-csv/converter"
)

// Config Initial config
type Config struct {
	SourcePath string
	StoreToDir string
}

func main() {
	var cfg Config

	flag.StringVar(&cfg.SourcePath, "source-path", "", "Path to XML file(s)")
	flag.StringVar(&cfg.StoreToDir, "store-to-dir", "", "Path where to store CSV file(s)")
	flag.Parse()

	var err error
	err = converter.Convert(cfg.SourcePath, cfg.StoreToDir)
	if err != nil {
		log.Fatal(err)
	}
}
