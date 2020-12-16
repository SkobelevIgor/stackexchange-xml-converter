package encoders

import "html"

// Comment entity
type Comment struct {
	ID              string `xml:"Id,attr" json:"Id"`
	PostID          string `xml:"PostId,attr" json:"PostId"`
	UserID          string `xml:"UserId,attr" json:"UserId,omitempty"`
	Score           string `xml:"Score,attr" json:"Score"`
	ContentLicense  string `xml:"ContentLicense,attr" json:"ContentLicense"`
	UserDisplayName string `xml:"UserDisplayName,attr" json:"UserDisplayName,omitempty"`
	Text            string `xml:"Text,attr" json:"Text,omitempty"`
	CreationDate    string `xml:"CreationDate,attr" json:"CreationDate"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (c Comment) GetCSVHeaderRow() []string {
	return []string{"Id", "PostId", "UserId",
		"Score", "ContentLicense", "UserDisplayName", "Text", "CreationDate"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (c *Comment) GETCSVRow(skipHTMLDecoding bool) []string {
	text := c.Text

	if skipHTMLDecoding {
		text = html.EscapeString(text)
	}

	return []string{c.ID, c.PostID, c.UserID,
		c.Score, c.ContentLicense, c.UserDisplayName, text, c.CreationDate}
}
