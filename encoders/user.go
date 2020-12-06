package encoders

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

func (u User) GetCSVHeaderRow() []string {
	return []string{"Id", "AccountId", "Reputation", "Views",
		"DownVotes", "UpVotes", "DisplayName", "Location", "ProfileImageUrl",
		"WebsiteUrl", "AboutMe", "CreationDate", "LastAccessDate"}
}

func (u *User) GETCSVRow() []string {
	return []string{u.ID, u.AccountID, u.Reputation, u.Views,
		u.DownVotes, u.UpVotes, u.DisplayName, u.Location, u.ProfileImageURL,
		u.WebsiteURL, u.AboutMe, u.CreationDate, u.LastAccessDate}
}
