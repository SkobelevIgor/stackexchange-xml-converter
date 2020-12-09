package converter

import (
	"encoding/csv"
	"encoding/xml"
	"log"
	"os"

	"github.com/SkobelevIgor/stackexchange-xml-to-csv/encoders"
)

func iterate(typeName string, xmlFile *os.File, csvFile *os.File, skipHTMLDecoding bool) {
	xmlDecoder := xml.NewDecoder(xmlFile)
	csvWriter := csv.NewWriter(csvFile)

	encoderType, err := encoders.NewEncoder(typeName)
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
				encoder, _ := encoders.NewEncoder(typeName)

				if err = xmlDecoder.DecodeElement(&encoder, &ty); err != nil {
					log.Printf("File: %s. Error decoding item: %s",
						xmlFile.Name(), err)
					break
				}
				err = csvWriter.Write(encoder.GETCSVRow(skipHTMLDecoding))
			}
		}

	}

	csvWriter.Flush()
	err = csvWriter.Error()
	if err != nil {
		// @TODO write to chan
	}
}
