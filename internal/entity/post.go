package entity

type Post struct {
	ID         int
	Body       string   `json:"post_body,omitempty"`
	Title      string   `json:"post_title,omitempty"`
	UserName   string   `json:"username,omitempty"`
	UserID     string   `json:"-"`
	PostDate   string   `json:"post_date,omitempty"`
	Categories []string `json:"category,omitempty"`
	Comments
	Like int
}
