package encoders

// Badge entity
type Badge struct {
	ID       string `xml:"Id,attr"`
	UserID   string `xml:"UserId,attr"`
	Class    string `xml:"Class,attr"`
	Name     string `xml:"Name,attr"`
	TagBased string `xml:"TagBased,attr"`
	Date     string `xml:"Date,attr"`
}

func (b Badge) GetCSVHeaderRow() []string {
	return []string{"Id", "UserId", "Class", "Name", "TagBased", "Date"}
}

func (b *Badge) GETCSVRow(skipHtmlDecoding bool) []string {
	return []string{b.ID, b.UserID, b.Class, b.Name, b.TagBased, b.Date}
}
