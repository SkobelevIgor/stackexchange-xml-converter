package converter

import (
	"encoding/xml"
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
func (xi *XMLIterator) Next() bool {
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
