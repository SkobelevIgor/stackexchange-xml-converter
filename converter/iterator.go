package converter

import (
	"encoding/csv"
	"encoding/xml"
	"log"
	"os"

	"github.com/SkobelevIgor/stackexchange-xml-to-csv/encoders"
)

func iterate(typeName string, xmlFile *os.File, csvFile *os.File) {
	xmlDecoder := xml.NewDecoder(xmlFile)
	csvWriter := csv.NewWriter(csvFile)

	enc, err := encoders.NewEncoder(typeName)
	if err != nil {
		// @TODO send to chan
		return
	}

	err = csvWriter.Write(enc.GetCSVHeaderRow())
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
				enc, _ := encoders.NewEncoder(typeName)

				if err = enc.LoadData(&ty); err != nil {
					log.Printf("File: %s. Error decoding item: %s",
						xmlFile.Name(), err)
					break
				}

				// if err = xmlDecoder.DecodeElement(&encoder, &ty); err != nil {
				// 	log.Printf("File: %s. Error decoding item: %s",
				// 		xmlFile.Name(), err)
				// 	break
				// }
				err = csvWriter.Write(enc.GETCSVRow())
			}
		}

	}

	csvWriter.Flush()
	err = csvWriter.Error()
	if err != nil {
		// @TODO write to chan
	}
}
