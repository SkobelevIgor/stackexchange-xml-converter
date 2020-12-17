package encoders

import "html"

// Post entity
type Post struct {
	ID                    string `xml:"Id,attr" json:"Id"`
	OwnerUserID           string `xml:"OwnerUserId,attr" json:"OwnerUserId,omitempty"`
	LastEditorUserID      string `xml:"LastEditorUserId,attr" json:"LastEditorUserId,omitempty"`
	PostTypeID            string `xml:"PostTypeId,attr" json:"PostTypeId"`
	AcceptedAnswerID      string `xml:"AcceptedAnswerId,attr" json:"AcceptedAnswerId,omitempty"`
	Score                 string `xml:"Score,attr"`
	ParentID              string `xml:"ParentId,attr" json:"ParentId,omitempty"`
	ViewCount             string `xml:"ViewCount,attr" json:"ViewCount,omitempty"`
	AnswerCount           string `xml:"AnswerCount,attr" json:"AnswerCount,omitempty"`
	CommentCount          string `xml:"CommentCount,attr" json:"CommentCount,omitempty"`
	OwnerDisplayName      string `xml:"OwnerDisplayName,attr" json:"OwnerDisplayName,omitempty"`
	LastEditorDisplayName string `xml:"LastEditorDisplayName,attr"`
	Title                 string `xml:"Title,attr" json:"Title,omitempty"`
	Tags                  string `xml:"Tags,attr" json:"Tags,omitempty"`
	ContentLIcense        string `xml:"ContentLicense,attr"`
	Body                  string `xml:"Body,attr" json:"Body,omitempty"`
	FavoriteCount         string `xml:"FavoriteCount,attr" json:"FavoriteCount,omitempty"`
	CreationDate          string `xml:"CreationDate,attr"`
	CommunityOwnedDate    string `xml:"CommunityOwnedDate,attr" json:"CommunityOwnedDate,omitempty"`
	ClosedDate            string `xml:"ClosedDate,attr" json:"ClosedDate,omitempty"`
	LastEditDate          string `xml:"LastEditDate,attr" json:"LastEditDate,omitempty"`
	LastActivityDate      string `xml:"LastActivityDate,attr" json:"LastActivityDate,omitempty"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (p Post) GetCSVHeaderRow() []string {
	return []string{"Id", "OwnerUserId", "LastEditorUserId",
		"PostTypeId", "AcceptedAnswerId",
		"Score", "ParentId", "ViewCount",
		"AnswerCount", "CommentCount",
		"OwnerDisplayName", "LastEditorDisplayName",
		"Title", "Tags", "ContentLicense",
		"Body", "FavoriteCount", "CreationDate",
		"CommunityOwnedDate", "ClosedDate",
		"LastEditDate", "LastActivityDate"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (p *Post) GETCSVRow() []string {
	return []string{p.ID, p.OwnerUserID, p.LastEditorUserID,
		p.PostTypeID, p.AcceptedAnswerID,
		p.Score, p.ParentID, p.ViewCount,
		p.AnswerCount, p.CommentCount,
		p.OwnerDisplayName, p.LastEditorDisplayName,
		p.Title, p.Tags, p.ContentLIcense,
		p.Body, p.FavoriteCount, p.CreationDate,
		p.CommunityOwnedDate, p.ClosedDate,
		p.LastEditDate, p.LastActivityDate}
}

// EscapeFields update fields to the original (escaped) state.
func (p *Post) EscapeFields() {
	p.Tags = html.EscapeString(p.Tags)
	p.Body = html.EscapeString(p.Body)
}
