package converter

import (
	"encoding/csv"
	"encoding/xml"
	"log"
	"os"

	"github.com/SkobelevIgor/stackexchange-xml-to-csv/types"
)

const flushBatchOnSize = 10000

func iterate(typeName string, xmlFile *os.File, csvFile *os.File) {
	xmlDecoder := xml.NewDecoder(xmlFile)
	csvWriter := csv.NewWriter(csvFile)

	var batchSizeLimiter int

	encoderType, err := types.CreateEntity(typeName)
	if err != nil {
		// @TODO send to chan
		return
	}

	err = csvWriter.Write(encoderType.GetCSVHeaderRow())
	if err != nil {
		// @TODO send to chan
		return
	}

	for {
		t, _ := xmlDecoder.Token()
		if t == nil {
			break
		}

		switch ty := t.(type) {
		case xml.StartElement:
			if ty.Name.Local == "row" {
				encoderType, _ := types.CreateEntity(typeName)

				if err = xmlDecoder.DecodeElement(&encoderType, &ty); err != nil {
					log.Printf("File: %s. Error decoding item: %s",
						xmlFile.Name(), err)
					break
				}
				err = csvWriter.Write(encoderType.GETCSVRow())
				batchSizeLimiter++
			}
		}

		if batchSizeLimiter >= flushBatchOnSize {
			csvWriter.Flush()
			reportFlushError(csvWriter)
			batchSizeLimiter = 0
		}
	}

	csvWriter.Flush()
	reportFlushError(csvWriter)
}

// @TODO write to status chan on error
func reportFlushError(csvWriter *csv.Writer) {
	err := csvWriter.Error()
	if err != nil {
		// @TODO write to chan
	}
}
