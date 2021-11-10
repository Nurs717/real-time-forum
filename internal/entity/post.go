package entity

type Post struct {
	ID         int
	Body       string `json:"post_body"`
	Title      string `json:"post_title"`
	UserID     string
	PostDate   string
	Category   string `json:"category"`
	LikeCounts int
}
