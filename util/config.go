package util

type Database struct {
	DbName     string `yaml:"db_name"`
	DbPassword string `yaml:"db_password"`
	DbUser     string `yaml:"db_user"`
}

type Config struct {
	Database    []Database `yaml:"database"`
	DbBakFolder string     `yaml:"db_bak_folder"`
}
