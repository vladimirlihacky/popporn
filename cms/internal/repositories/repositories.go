package repositories

import "cms/internal/models"

type UsersRepo interface {
	GetByID(id int) (*models.User, error)
	GetByName(name string) (*models.User, error)
	Create(data *models.User) error
	Update(id int, data *models.User) error
}

type CommentsRepo interface {
	GetByID(id int) (*models.Comment, error)
	Create(data *models.Comment) error
	Update(id int, data *models.Comment) error
	Delete(id int) error
}

type VideosRepo interface {
	GetByID(id int) (*models.Video, error)
	Create(data *models.Video) error
	Update(id int, data *models.Video) error
	Delete(id int) error
}
