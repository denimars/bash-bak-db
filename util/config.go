package util

type Database struct {
	DbName     string `yaml:"db_name"`
	DbPassword string `yaml:"db_password"`
	DbUser     string `yaml:"db_user"`
}

type Data struct {
	Location string `yaml:"location"`
}

type Config struct {
	Database    []Database `yaml:"database"`
	DbBakFolder string     `yaml:"db_bak_folder"`
	Data        []Data     `yaml:"data"`
}
