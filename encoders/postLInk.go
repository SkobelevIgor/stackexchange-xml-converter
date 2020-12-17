package encoders

// PostLink entity
type PostLink struct {
	ID            string `xml:"Id,attr" json:"Id"`
	RelatedPostID string `xml:"RelatedPostId,attr" json:"RelatedPostId"`
	PostID        string `xml:"PostId,attr" json:"PostId"`
	LinkTypeID    string `xml:"LinkTypeId,attr" json:"LinkTypeId"`
	CreationDate  string `xml:"CreationDate,attr"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (pl PostLink) GetCSVHeaderRow() []string {
	return []string{"Id", "RelatedPostId", "PostId",
		"LinkTypeId", "CreationDate"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (pl PostLink) GETCSVRow() []string {
	return []string{pl.ID, pl.RelatedPostID, pl.PostID,
		pl.LinkTypeID, pl.CreationDate}
}

// EscapeFields update fields to the original (escaped) state.
func (pl *PostLink) EscapeFields() {}
