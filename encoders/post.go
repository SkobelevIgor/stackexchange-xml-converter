package encoders

func getPostColumnMap() map[string]int {
	return map[string]int{
		"Id":                    0,
		"OwnerUserId":           1,
		"LastEditorUserId":      2,
		"PostTypeId":            3,
		"AcceptedAnswerId":      4,
		"Score":                 5,
		"ParentId":              6,
		"ViewCount":             7,
		"AnswerCount":           8,
		"CommentCount":          9,
		"OwnerDisplayName":      10,
		"LastEditorDisplayName": 11,
		"Title":                 12,
		"Tags":                  13,
		"ContentLicense":        14,
		"Body":                  15,
		"FavoriteCount":         16,
		"CreationDate":          17,
		"CommunityOwnedDate":    18,
		"ClosedDate":            19,
		"LastEditDate":          20,
		"LastActivityDate":      21}
}

// // Post entity
// type Post struct {
// 	ID                    string `xml:"Id,attr"`
// 	OwnerUserID           string `xml:"OwnerUserId,attr"`
// 	LastEditorUserID      string `xml:"LastEditorUserId,attr"`
// 	PostTypeID            string `xml:"PostTypeId,attr"`
// 	AcceptedAnswerID      string `xml:"AcceptedAnswerId,attr"`
// 	Score                 string `xml:"Score,attr"`
// 	ParentID              string `xml:"ParentId,attr"`
// 	ViewCount             string `xml:"ViewCount,attr"`
// 	AnswerCount           string `xml:"AnswerCount,attr"`
// 	CommentCount          string `xml:"CommentCount,attr"`
// 	OwnerDisplayName      string `xml:"OwnerDisplayName,attr"`
// 	LastEditorDisplayName string `xml:"LastEditorDisplayName,attr"`
// 	Title                 string `xml:"Title,attr"`
// 	Tags                  string `xml:"Tags,attr"`
// 	ContentLIcense        string `xml:"ContentLicense,attr"`
// 	Body                  string `xml:"Body,attr"`
// 	FavoriteCount         string `xml:"FavoriteCount,attr"`
// 	CreationDate          string `xml:"CreationDate,attr"`
// 	CommunityOwnedDate    string `xml:"CommunityOwnedDate,attr"`
// 	ClosedDate            string `xml:"ClosedDate,attr"`
// 	LastEditDate          string `xml:"LastEditDate,attr"`
// 	LastActivityDate      string `xml:"LastActivityDate,attr"`
// }

// func (p Post) GetCSVHeaderRow() []string {
// 	return []string{"Id", "OwnerUserId", "LastEditorUserId",
// 		"PostTypeId", "AcceptedAnswerId",
// 		"Score", "ParentId", "ViewCount",
// 		"AnswerCount", "CommentCount",
// 		"OwnerDisplayName", "LastEditorDisplayName",
// 		"Title", "Tags", "ContentLicense",
// 		"Body", "FavoriteCount", "CreationDate",
// 		"CommunityOwnedDate", "ClosedDate",
// 		"LastEditDate", "LastActivityDate"}
// }

// func (p *Post) GETCSVRow() []string {
// 	return []string{p.ID, p.OwnerUserID, p.LastEditorUserID,
// 		p.PostTypeID, p.AcceptedAnswerID,
// 		p.Score, p.ParentID, p.ViewCount,
// 		p.AnswerCount, p.CommentCount,
// 		p.OwnerDisplayName, p.LastEditorDisplayName,
// 		p.Title, p.Tags, p.ContentLIcense,
// 		p.Body, p.FavoriteCount, p.CreationDate,
// 		p.CommunityOwnedDate, p.ClosedDate,
// 		p.LastEditDate, p.LastActivityDate}
// }
