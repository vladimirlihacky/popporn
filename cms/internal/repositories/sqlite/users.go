package sqlite

import (
	"cms/internal/models"
	"cms/pkg/utils"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

func (repo *Users) Create(data *models.User) error {

	_, err := repo.db.Exec(
		"INSERT INTO users (name, liked_videos, comments) VALUES (?, ?, ?)",
		data.Name, data.Comments, data.LikedVideos,
	)

	return err
}

func (repo *Users) GetByID(id int) (*models.User, error) {
	var user *models.User
	var likedVideos, comments string

	err := repo.db.QueryRow("SELECT id, name, liked_videos, comments FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &likedVideos, &comments)
	if err != nil {
		return nil, err
	}

	user.LikedVideos = utils.DeserializeIntSlice(likedVideos)
	user.Comments = utils.DeserializeIntSlice(comments)

	return user, nil
}

func (repo *Users) GetByName(name string) (*models.User, error) {
	var user *models.User
	var likedVideos, comments string

	err := repo.db.QueryRow("SELECT id, name, liked_videos, comments FROM users WHERE name = ?", name).
		Scan(&user.ID, &user.Name, &likedVideos, &comments)
	if err != nil {
		return nil, err
	}

	user.LikedVideos = utils.DeserializeIntSlice(likedVideos)
	user.Comments = utils.DeserializeIntSlice(comments)

	return user, nil
}

func (repo *Users) Update(id int, data *models.User) error {
	likedVideos := utils.SerializeIntSlice(data.LikedVideos)
	comments := utils.SerializeIntSlice(data.Comments)

	_, err := repo.db.Exec("UPDATE users SET name = ?, liked_videos = ?, comments = ? WHERE id = ?",
		data.Name, likedVideos, comments, id)
	return err
}
