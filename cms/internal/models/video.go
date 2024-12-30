package models

type Video struct {
	ID          int    `json:"id" validate:"required,min=1"`
	Title       string `json:"title" validate:"required,min=1,max=100"`
	Description string `json:"description" validate:"max=500"`
	LikesCount  int    `json:"likes_count" validate:"min=0"`
	ViewsCount  int    `json:"views_count" validate:"min=0"`
	Comments    []int  `json:"comments"`
	Source      string `json:"source"`
}
