package encoders

import "fmt"

// Encoder interface
type Encoder interface {
	// GetCSVHeaderRow returns CSV header for the correspondig encoder type.
	// WARNING! Order is crucial!
	// Fields will be written to the CSV file in the same order this function returns them.
	GetCSVHeaderRow() []string

	// GETCSVRow returns row values for the corresponding encoder type
	// WARNING! Order is crucial!
	// Values will be written to the CSV file in the same order this function returns them.
	GETCSVRow(skipHTMLDecoding bool) []string
}

// NewEncoder returns a pointer to the new encoder according to requested type
func NewEncoder(typeName string) (c Encoder, err error) {
	switch typeName {
	case "Badges":
		c = &Badge{}
	case "Comments":
		c = &Comment{}
	case "Posts":
		c = &Post{}
	case "PostLinks":
		c = &PostLink{}
	case "PostHistory":
		c = &PostHistory{}
	case "Tags":
		c = &Tag{}
	case "Users":
		c = &User{}
	case "Votes":
		c = &Vote{}
	default:
		err = fmt.Errorf("Undefined CSVEncoder type: %s", typeName)
	}
	return
}
