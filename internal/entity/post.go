package entity

type Post struct {
	ID         int
	Body       string   `json:"post_body"`
	Title      string   `json:"post_title"`
	UserName   string   `json:"username"`
	UserID     string   `json:"-"`
	PostDate   string   `json:"-"`
	Category   []string `json:"category"`
	LikeCounts int
}
