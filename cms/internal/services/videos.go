package services

import (
	"cms/internal/models"
	"cms/internal/repositories"
)

type CDN interface {
	UploadFile([]byte) (string, error)
	UpdateFile(string, []byte) error
	DeleteFile(string) error
}

type Videos struct {
	cdn  CDN
	repo repositories.Videos
}

type VideoWithFile struct {
	*models.Video
	File []byte
}

func (videos *Videos) Create(data *VideoWithFile) error {

	source, err := videos.cdn.UploadFile(data.File)

	if err != nil {
		return err
	}

	video := data.Video
	video.Source = source

	if err := videos.repo.Create(video); err != nil {
		return err
	}

	return nil
}

func (videos *Videos) GetByID(id int) (*models.Video, error) {
	return videos.repo.GetByID(id)
}

func (videos *Videos) Delete(id int) error {
	video, err := videos.repo.GetByID(id)

	if err != nil {
		return err
	}

	return videos.cdn.DeleteFile(video.Source)
}

func (videos *Videos) Update(data *VideoWithFile) error {

	if err := videos.repo.Update(data.ID, data.Video); err != nil {
		return err
	}

	video, err := videos.repo.GetByID(data.ID)

	if err != nil {
		return err
	}

	if len(data.File) != 0 {
		err = videos.cdn.UpdateFile(video.Source, data.File)
	}

	return err
}
