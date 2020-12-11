package encoders

// Vote entity
type Vote struct {
	ID           string `xml:"Id,attr"`
	UserID       string `xml:"UserId,attr"`
	PostID       string `xml:"PostId,attr"`
	VoteTypeID   string `xml:"VoteTypeId,attr"`
	BountyAmount string `xml:"BountyAmount,attr"`
	CreationDate string `xml:"CreationDate,attr"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (v Vote) GetCSVHeaderRow() []string {
	return []string{"Id", "UserId", "PostId",
		"VoteTypeId", "BountyAmount", "CreationDate"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (v *Vote) GETCSVRow(skipHTMLDecoding bool) []string {
	return []string{v.ID, v.UserID, v.PostID,
		v.VoteTypeID, v.BountyAmount, v.CreationDate}
}
