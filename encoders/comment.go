package encoders

import "html"

// Comment entity
type Comment struct {
	ID              string `xml:"Id,attr" json:"Id"`
	PostID          string `xml:"PostId,attr" json:"PostId"`
	UserID          string `xml:"UserId,attr" json:"UserId,omitempty"`
	Score           string `xml:"Score,attr"`
	ContentLicense  string `xml:"ContentLicense,attr"`
	UserDisplayName string `xml:"UserDisplayName,attr" json:"UserDisplayName,omitempty"`
	Text            string `xml:"Text,attr" json:"Text,omitempty"`
	CreationDate    string `xml:"CreationDate,attr"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (c Comment) GetCSVHeaderRow() []string {
	return []string{"Id", "PostId", "UserId",
		"Score", "ContentLicense", "UserDisplayName", "Text", "CreationDate"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (c *Comment) GETCSVRow() []string {
	return []string{c.ID, c.PostID, c.UserID,
		c.Score, c.ContentLicense, c.UserDisplayName, c.Text, c.CreationDate}
}

// EscapeFields update fields to the original (escaped) state.
func (c *Comment) EscapeFields() {
	c.Text = html.EscapeString(c.Text)
}
