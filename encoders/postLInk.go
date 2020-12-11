package encoders

// PostLink entity
type PostLink struct {
	ID            string `xml:"Id,attr"`
	RelatedPostID string `xml:"RelatedPostId,attr"`
	PostID        string `xml:"PostId,attr"`
	LinkTypeID    string `xml:"LinkTypeId,attr"`
	CreationDate  string `xml:"CreationDate,attr"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (pl PostLink) GetCSVHeaderRow() []string {
	return []string{"Id", "RelatedPostId", "PostId",
		"LinkTypeId", "CreationDate"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (pl PostLink) GETCSVRow(skipHTMLDecoding bool) []string {
	return []string{pl.ID, pl.RelatedPostID, pl.PostID,
		pl.LinkTypeID, pl.CreationDate}
}
