package encoders

import "fmt"

// Entity interface
type CSVEncoder interface {
	GetCSVHeaderRow() []string
	GETCSVRow() []string
}

func NewEncoder(typeName string) (c CSVEncoder, err error) {
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
