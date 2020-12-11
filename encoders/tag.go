package encoders

// Tag entity
type Tag struct {
	ID            string `xml:"Id,attr"`
	ExcerptPostID string `xml:"ExcerptPostId,attr"`
	WikiPostID    string `xml:"WikiPostId,attr"`
	TagName       string `xml:"TagName,attr"`
	Count         string `xml:"Count,attr"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (t Tag) GetCSVHeaderRow() []string {
	return []string{"Id", "ExcerptPostId", "WikiPostId", "TagName", "Count"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (t *Tag) GETCSVRow(skipHTMLDecoding bool) []string {
	return []string{t.ID, t.ExcerptPostID, t.WikiPostID, t.TagName, t.Count}
}
