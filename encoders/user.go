package encoders

import "html"

// User entity
type User struct {
	ID              string `xml:"Id,attr"`
	AccountID       string `xml:"AccountId,attr"`
	Reputation      string `xml:"Reputation,attr"`
	Views           string `xml:"Views,attr"`
	DownVotes       string `xml:"DownVotes,attr"`
	UpVotes         string `xml:"UpVotes,attr"`
	DisplayName     string `xml:"DisplayName,attr"`
	Location        string `xml:"Location,attr"`
	ProfileImageURL string `xml:"ProfileImageUrl,attr"`
	WebsiteURL      string `xml:"WebsiteUrl,attr"`
	AboutMe         string `xml:"AboutMe,attr"`
	CreationDate    string `xml:"CreationDate,attr"`
	LastAccessDate  string `xml:"LastAccessDate,attr"`
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
