package encoders

// Badge entity
type Badge struct {
	ID       string `xml:"Id,attr" json:"Id"`
	UserID   string `xml:"UserId,attr" json:"UserId"`
	Class    string `xml:"Class,attr"`
	Name     string `xml:"Name,attr"`
	TagBased string `xml:"TagBased,attr"`
	Date     string `xml:"Date,attr"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (b Badge) GetCSVHeaderRow() []string {
	return []string{"Id", "UserId", "Class", "Name", "TagBased", "Date"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (b *Badge) GETCSVRow() []string {
	return []string{b.ID, b.UserID, b.Class, b.Name, b.TagBased, b.Date}
}

// EscapeFields update fields to the original (escaped) state.
func (b *Badge) EscapeFields() {}
