package encoders

func getCommentColumnsMap() map[string]int {
	return map[string]int{
		"Id":              0,
		"PostId":          1,
		"UserId":          2,
		"Score":           3,
		"ContentLicense":  4,
		"UserDisplayName": 5,
		"Text":            6,
		"CreationDate":    7}
}

// // Comment entity
// type Comment struct {
// 	ID              string `xml:"Id,attr"`
// 	PostID          string `xml:"PostId,attr"`
// 	UserID          string `xml:"UserId,attr"`
// 	Score           string `xml:"Score,attr"`
// 	ContentLicense  string `xml:"ContentLicense,attr"`
// 	UserDisplayName string `xml:"UserDisplayName,attr"`
// 	Text            string `xml:"Text,attr"`
// 	CreationDate    string `xml:"CreationDate,attr"`
// }

// func (c Comment) GetCSVHeaderRow() []string {
// 	return []string{"Id", "PostId", "UserId",
// 		"Score", "ContentLicense", "UserDisplayName", "Text", "CreationDate"}
// }

// func (c *Comment) GETCSVRow() []string {
// 	return []string{c.ID, c.PostID, c.UserID,
// 		c.Score, c.ContentLicense, c.UserDisplayName, c.Text, c.CreationDate}
// }
