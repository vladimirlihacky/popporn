package config

type Config struct {
	SQLiteURL string
}

var Default *Config = &Config{
	SQLiteURL: "./repo.db",
}
