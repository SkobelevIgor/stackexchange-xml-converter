package encoders

import "html"

// Comment entity
type Comment struct {
	ID              string `xml:"Id,attr"`
	PostID          string `xml:"PostId,attr"`
	UserID          string `xml:"UserId,attr"`
	Score           string `xml:"Score,attr"`
	ContentLicense  string `xml:"ContentLicense,attr"`
	UserDisplayName string `xml:"UserDisplayName,attr"`
	Text            string `xml:"Text,attr"`
	CreationDate    string `xml:"CreationDate,attr"`
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
