package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	path string
	db   *sql.DB
}

func (repo *Repo) NewUsersRepo() *Users {
	return &Users{repo.db}
}

func (repo *Repo) NewCommentsRepo() *Comments {
	return &Comments{repo.db}
}

func (repo *Repo) NewVideosRepo() *Videos {
	return &Videos{repo.db}
}

func (repo *Repo) initTables() error {
	_, err := repo.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL CHECK(length(name) <= 100),
			liked_videos TEXT, -- JSON-массив для хранения списка liked_videos
			comments TEXT -- JSON-массив для хранения списка comments
		);

		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			author INTEGER NOT NULL,
			text TEXT NOT NULL CHECK(length(text) <= 500),
			reply_to INTEGER,
			FOREIGN KEY (author) REFERENCES users(id)
		);

		CREATE TABLE IF NOT EXISTS videos (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL CHECK(length(title) <= 100),
			description TEXT CHECK(length(description) <= 500),
			likes_count INTEGER DEFAULT 0 CHECK(likes_count >= 0),
			views_count INTEGER DEFAULT 0 CHECK(views_count >= 0),
			comments TEXT, -- JSON-массив для хранения списка comments
			source TEXT
		);
	`)

	return err
}

func (repo *Repo) Close() {
	repo.db.Close()
}

func NewSQLite(path string) (*Repo, error) {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		return nil, err
	}

	repo := &Repo{
		path: path,
		db:   db,
	}

	err = repo.initTables()

	return repo, err
}
