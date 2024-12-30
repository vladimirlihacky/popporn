package sqlite

import (
	"cms/internal/models"
	"cms/pkg/utils"
	"database/sql"
)

type Videos struct {
	db *sql.DB
}

func (r *Videos) GetByID(id int) (*models.Video, error) {
	var video models.Video
	var comments string

	err := r.db.QueryRow("SELECT id, title, description, likes_count, views_count, comments, source FROM videos WHERE id = ?", id).
		Scan(&video.ID, &video.Title, &video.Description, &video.LikesCount, &video.ViewsCount, &comments, &video.Source)
	if err != nil {
		return nil, err
	}

	video.Comments = utils.DeserializeIntSlice(comments)

	return &video, nil
}

func (r *Videos) Create(data *models.Video) error {
	comments := utils.SerializeIntSlice(data.Comments)

	_, err := r.db.Exec("INSERT INTO videos (title, description, likes_count, views_count, comments, source) VALUES (?, ?, ?, ?, ?, ?)",
		data.Title, data.Description, data.LikesCount, data.ViewsCount, comments, data.Source)
	return err
}

func (r *Videos) Update(id int, data *models.Video) error {
	comments := utils.SerializeIntSlice(data.Comments)

	_, err := r.db.Exec("UPDATE videos SET title = ?, description = ?, likes_count = ?, views_count = ?, comments = ?, source = ? WHERE id = ?",
		data.Title, data.Description, data.LikesCount, data.ViewsCount, comments, data.Source, id)
	return err
}

func (r *Videos) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM videos WHERE id = ?", id)
	return err
}
