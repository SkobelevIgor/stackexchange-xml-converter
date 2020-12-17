package encoders

// Vote entity
type Vote struct {
	ID           string `xml:"Id,attr" json:"Id"`
	UserID       string `xml:"UserId,attr" json:"UserId,omitempty"`
	PostID       string `xml:"PostId,attr" json:"PostId"`
	VoteTypeID   string `xml:"VoteTypeId,attr" json:"VoteTypeId"`
	BountyAmount string `xml:"BountyAmount,attr" json:"BountyAmount,omitempty"`
	CreationDate string `xml:"CreationDate,attr"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (v Vote) GetCSVHeaderRow() []string {
	return []string{"Id", "UserId", "PostId",
		"VoteTypeId", "BountyAmount", "CreationDate"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (v *Vote) GETCSVRow() []string {
	return []string{v.ID, v.UserID, v.PostID,
		v.VoteTypeID, v.BountyAmount, v.CreationDate}
}

// EscapeFields update fields to the original (escaped) state.
func (v *Vote) EscapeFields() {}