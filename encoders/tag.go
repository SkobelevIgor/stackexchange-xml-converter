package encoders

func getTagColumnMap() map[string]int {
	return map[string]int{
		"Id":            0,
		"ExcerptPostId": 1,
		"WikiPostId":    2,
		"TagName":       3,
		"Count":         4}
}

// // Tag entity
// type Tag struct {
// 	ID            string `xml:"Id,attr"`
// 	ExcerptPostID string `xml:"ExcerptPostId,attr"`
// 	WikiPostID    string `xml:"WikiPostId,attr"`
// 	TagName       string `xml:"TagName,attr"`
// 	Count         string `xml:"Count,attr"`
// }

// func (t Tag) GetCSVHeaderRow() []string {
// 	return []string{"Id", "ExcerptPostId", "WikiPostId", "TagName", "Count"}
// }

// func (t *Tag) GETCSVRow() []string {
// 	return []string{t.ID, t.ExcerptPostID, t.WikiPostID, t.TagName, t.Count}
// }
