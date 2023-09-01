package model

type Bookmark struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
