package entity

type Post struct {
	ID         int
	Body       string   `json:"post_body,omitempty"`
	Title      string   `json:"post_title"`
	UserName   string   `json:"username"`
	UserID     string   `json:"-"`
	PostDate   string   `json:"post_date"`
	Category   []string `json:"category"`
	LikeCounts int
}
