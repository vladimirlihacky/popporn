package models

type Comment struct {
	ID      int    `json:"id" validate:"required,min=1"`
	Author  int    `json:"author" validate:"required,min=1"`
	Text    string `json:"text" validate:"required,min=1,max=500"`
	ReplyTo *int   `json:"reply_to,omitempty"`
}
