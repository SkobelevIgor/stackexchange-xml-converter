package encoders

import "html"

// User entity
type User struct {
	ID              string `xml:"Id,attr" json:"Id"`
	AccountID       string `xml:"AccountId,attr" json:"AccountId,omitempty"`
	Reputation      string `xml:"Reputation,attr" json:"Reputation"`
	Views           string `xml:"Views,attr" json:"Views,omitempty"`
	DownVotes       string `xml:"DownVotes,attr" json:"DownVotes,omitempty"`
	UpVotes         string `xml:"UpVotes,attr" json:"UpVotes,omitempty"`
	DisplayName     string `xml:"DisplayName,attr" json:"DisplayName"`
	Location        string `xml:"Location,attr" json:"Location,omitempty"`
	ProfileImageURL string `xml:"ProfileImageUrl,attr" json:"ProfileImageUrl,omitempty"`
	WebsiteURL      string `xml:"WebsiteUrl,attr" json:"WebsiteUrl,omitempty"`
	AboutMe         string `xml:"AboutMe,attr" json:"AboutMe,omitempty"`
	CreationDate    string `xml:"CreationDate,attr" json:"CreationDate"`
	LastAccessDate  string `xml:"LastAccessDate,attr" json:"LastAccessDate"`
}

// GetCSVHeaderRow returns CSV header for the correspondig encoder type
func (u User) GetCSVHeaderRow() []string {
	return []string{"Id", "AccountId", "Reputation", "Views",
		"DownVotes", "UpVotes", "DisplayName", "Location", "ProfileImageUrl",
		"WebsiteUrl", "AboutMe", "CreationDate", "LastAccessDate"}
}

// GETCSVRow returns row values for the corresponding encoder type
func (u *User) GETCSVRow(skipHTMLDecoding bool) []string {

	aboutMe := u.AboutMe
	if skipHTMLDecoding {
		aboutMe = html.EscapeString(aboutMe)
	}

	return []string{u.ID, u.AccountID, u.Reputation, u.Views,
		u.DownVotes, u.UpVotes, u.DisplayName, u.Location, u.ProfileImageURL,
		u.WebsiteURL, aboutMe, u.CreationDate, u.LastAccessDate}
}
