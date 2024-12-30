package main

import (
	"cms/internal/config"
	"cms/internal/repositories/sqlite"
	"fmt"
)

func main() {
	config := config.Default

	repo, _ := sqlite.NewSQLite(config.SQLiteURL)
	defer repo.Close()

	usersRepo := repo.NewUsersRepo()
	videosRepo := repo.NewVideosRepo()
	commentsRepo := repo.NewCommentsRepo()

	// usersService := services.NewUsersService()

	fmt.Printf("%v %v %v", usersRepo, videosRepo, commentsRepo)
}
