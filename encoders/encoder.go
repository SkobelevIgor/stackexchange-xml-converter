package encoders

import (
	"encoding/xml"
	"fmt"
)

// Entity interface
// type CSVEncoder interface {
// 	GetCSVHeaderRow() []string
// 	GETCSVRow() []string
// }

type CSVEncoder struct {
	data       []string
	columnsMap map[string]int
}

func NewEncoder(typeName string) (c CSVEncoder, err error) {
	var cols map[string]int

	switch typeName {
	case "Badges":
		cols = getBadgeColumnsMap()
	case "Comments":
		cols = getCommentColumnsMap()
	case "Posts":
		cols = getPostColumnMap()
	case "PostLinks":
		cols = getPostLinkColumnMap()
	case "PostHistory":
		cols = getPostHistoryColumnsMap()
	case "Tags":
		cols = getTagColumnMap()
	case "Users":
		cols = getUserColumnMap()
	case "Votes":
		cols = getVoteColumnMap()
	default:
		err = fmt.Errorf("Undefined CSVEncoder type: %s", typeName)
	}

	c.columnsMap = cols
	c.data = make([]string, len(cols))

	return
}

func (c *CSVEncoder) LoadData(xmlNode *xml.StartElement) (err error) {

	for _, attr := range xmlNode.Attr {
		if idx, ok := c.columnsMap[attr.Name.Local]; ok {
			c.data[idx] = attr.Value
		} else {
			err = fmt.Errorf("Unknown field %s", attr.Name.Local)
		}
	}

	return
}

func (c *CSVEncoder) GetCSVHeaderRow() (cols []string) {
	// @TODO add sorting
	for k := range c.columnsMap {
		cols = append(cols, k)
	}
	return
}

func (c *CSVEncoder) GETCSVRow() []string {
	return c.data
}
