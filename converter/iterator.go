package converter

import (
	"encoding/csv"
	"encoding/xml"
	"log"
	"os"

	"github.com/SkobelevIgor/stackexchange-xml-converter/encoders"
)

// XMLIterator struct
type XMLIterator struct {
	xmlDecoder *xml.Decoder
	element    *xml.StartElement
}

// NewIterator creates XML iterator
func NewIterator(file *os.File) (xi *XMLIterator) {
	xi = &XMLIterator{
		xmlDecoder: xml.NewDecoder(file),
	}
	return
}

// Next iterate through XML document
func (xi XMLIterator) Next() bool {
	for {
		t, _ := xi.xmlDecoder.Token()
		if t == nil {
			return false
		}

		switch ty := t.(type) {
		case xml.StartElement:
			if ty.Name.Local == "row" {
				xi.element = &ty
				return true
			}
		}
	}
}

// Decode fill encoder fields
func (xi *XMLIterator) Decode(encoder *encoders.Encoder) error {
	return xi.xmlDecoder.DecodeElement(encoder, xi.element)
}

func iterate(typeName string, xmlFile *os.File, csvFile *os.File, cfg Config) (totalCounter int64, convertedCounter int64, err error) {
	xmlDecoder := xml.NewDecoder(xmlFile)
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	encoderType, err := encoders.NewEncoder(typeName)
	if err != nil {
		return
	}

	err = csvWriter.Write(encoderType.GetCSVHeaderRow())
	if err != nil {
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
				totalCounter++
				encoder, _ := encoders.NewEncoder(typeName)
				err = xmlDecoder.DecodeElement(&encoder, &ty)
				if err != nil {
					log.Printf("[%s] Error: %s", typeName, err)
					continue
				}

				err = csvWriter.Write(encoder.GETCSVRow(cfg.SkipHTMLDecoding))
				if err != nil {
					log.Printf("[%s] Error: %s", typeName, err)
					continue
				}
				convertedCounter++
			}
		}
	}

	return
}
