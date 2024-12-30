package models

type User struct {
	ID          int    `json:"id" validate:"required,min=1"`
	Name        string `json:"name" validate:"required,min=1,max=100"`
	LikedVideos []int  `json:"liked_videos"`
	Comments    []int  `json:"comments"`
}
