package entity

type Post struct {
	ID         int
	Post       string `json:"post"`
	UserID     string
	PostBody   string
	PostDate   string
	Category   string
	LikeCounts int
}
