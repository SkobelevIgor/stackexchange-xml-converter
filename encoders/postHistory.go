package encoders

import "html"

// PostHistory entity
type PostHistory struct {
	ID                string `xml:"Id,attr"`
	PostID            string `xml:"PostId,attr"`
	UserID            string `xml:"UserId,attr"`
	PostHistoryTypeID string `xml:"PostHistoryTypeId,attr"`
	UserDisplayName   string `xml:"UserDisplayName,attr"`
	ContentLicense    string `xml:"ContentLicense,attr"`
	RevisionGUID      string `xml:"RevisionGUID,attr"`
	Text              string `xml:"Text,attr"`
	Comment           string `xml:"Comment,attr"`
	CreationDate      string `xml:"CreationDate,attr"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (ph PostHistory) GetCSVHeaderRow() []string {
	return []string{"Id", "PostId", "UserId",
		"PostHistoryTypeId", "UserDisplayName",
		"ContentLicense", "RevisionGUID",
		"Text", "Comment", "CreationDate"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (ph *PostHistory) GETCSVRow(skipHTMLDecoding bool) []string {

	text := ph.Text
	if skipHTMLDecoding {
		text = html.EscapeString(text)
	}

	return []string{ph.ID, ph.PostID, ph.UserID,
		ph.PostHistoryTypeID, ph.UserDisplayName,
		ph.ContentLicense, ph.RevisionGUID,
		text, ph.Comment, ph.CreationDate}
}
