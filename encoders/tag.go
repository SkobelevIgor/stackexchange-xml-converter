package encoders

// Tag entity
type Tag struct {
	ID            string `xml:"Id,attr" json:"Id"`
	ExcerptPostID string `xml:"ExcerptPostId,attr" json:"ExcerptPostId,omitempty"`
	WikiPostID    string `xml:"WikiPostId,attr" json:"WikiPostId,omitempty"`
	TagName       string `xml:"TagName,attr"`
	Count         string `xml:"Count,attr" json:"Count,omitempty"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (t Tag) GetCSVHeaderRow() []string {
	return []string{"Id", "ExcerptPostId", "WikiPostId", "TagName", "Count"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (t *Tag) GETCSVRow() []string {
	return []string{t.ID, t.ExcerptPostID, t.WikiPostID, t.TagName, t.Count}
}

// EscapeFields update fields to the original (escaped) state.
func (t *Tag) EscapeFields() {}
