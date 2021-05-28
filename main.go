package main

import (
	"flag"
	"log"

	"github.com/SkobelevIgor/stackexchange-xml-converter/converter"
)

func main() {
	var cfg converter.Config

	flag.StringVar(&cfg.ResultFormat, "result-format", "", "Result format (csv or json)")
	flag.StringVar(&cfg.SourcePath, "source-path", "", "Path to XML file(s)")
	flag.StringVar(&cfg.StoreToDir, "store-to-dir", "", "Path where to store CSV file(s)")
	flag.BoolVar(&cfg.SkipHTMLDecoding, "skip-html-decoding", false, "Path where to store CSV file(s)")
	flag.StringVar(&cfg.FilterByTagId, "filter-by-tag-id", "", "Filter for tags, space sperated list")
	flag.BoolVar(&cfg.JsonOneLine, "json-one-line", false, "Save json file as one object per line")
	flag.Parse()

	var err error
	err = converter.Convert(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
