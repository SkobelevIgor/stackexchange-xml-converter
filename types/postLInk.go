package types

// PostLink entity
type PostLink struct {
	ID            string `xml:"Id,attr"`
	RelatedPostID string `xml:"RelatedPostId,attr"`
	PostID        string `xml:"PostId,attr"`
	LinkTypeID    string `xml:"LinkTypeId,attr"`
	CreationDate  string `xml:"CreationDate,attr"`
}

func (pl PostLink) GetCSVHeaderRow() []string {
	return []string{"Id", "RelatedPostId", "PostId",
		"LinkTypeId", "CreationDate"}
}

func (pl PostLink) GETCSVRow() []string {
	return []string{pl.ID, pl.RelatedPostID, pl.PostID,
		pl.LinkTypeID, pl.CreationDate}
}
