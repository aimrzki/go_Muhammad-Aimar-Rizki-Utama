package config

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func LoadConfig() Config {
	return Config{
		DB_Username: "root",
		DB_Password: "admin123",
		DB_Port:     "3306",
		DB_Host:     "35.238.117.77",
		DB_Name:     "crud_go",
	}
}
