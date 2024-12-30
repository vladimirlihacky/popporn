package sqlite

import (
	"cms/internal/models"
	"database/sql"
)

type Comments struct {
	db *sql.DB
}

func (repo *Comments) GetByID(id int) (*models.Comment, error) {
	var comment models.Comment

	err := repo.db.QueryRow("SELECT id, author, text, reply_to FROM comments WHERE id = ?", id).
		Scan(&comment.ID, &comment.Author, &comment.Text, &comment.ReplyTo)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (repo *Comments) Create(data *models.Comment) error {
	_, err := repo.db.Exec("INSERT INTO comments (author, text, reply_to) VALUES (?, ?, ?)",
		data.Author, data.Text, data.ReplyTo)
	return err
}

func (repo *Comments) Update(id int, data *models.Comment) error {
	_, err := repo.db.Exec("UPDATE comments SET text = ?, reply_to = ? WHERE id = ?",
		data.Text, data.ReplyTo, id)
	return err
}

func (repo *Comments) Delete(id int) error {
	_, err := repo.db.Exec("DELETE FROM comments WHERE id = ?", id)
	return err
}
