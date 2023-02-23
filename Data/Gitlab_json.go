package Data

type GitlabJson struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	State        string `json:"state"`
	AvatarURL    string `json:"avatar_url"`
	WebURL       string `json:"web_url"`
	CreatedAt    string `json:"created_at"`
	Bio          string `json:"bio"`
	BioHTML      string `json:"bio_html"`
	Location     string `json:"location"`
	PublicEmail  string `json:"public_email"`
	Skype        string `json:"skype"`
	LinkedIn     string `json:"linkedin"`
	Twitter      string `json:"twitter"`
	WebsiteURL   string `json:"website_url"`
	Organization string `json:"organization"`
	JobTitle     string `json:"job_title"`
	Bot          bool   `json:"bot"`
	WorkInfo     string `json:"work_information"`
	Followers    int    `json:"followers"`
	Following    int    `json:"following"`
}
