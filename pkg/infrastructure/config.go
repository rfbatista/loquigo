package infrastructure

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func NewConfig() Config {
	godotenv.Load()
	port := 3000
	dbname := os.Getenv("MONGO_DB_NAME")
	dbport := os.Getenv("MONGO_PORT")
	dbuser := os.Getenv("MONGO_DB_USER")
	dbpassword := os.Getenv("MONGO_PWD")
	dbhost := os.Getenv("MONGO_HOST")
	return Config{port: port, dbname: dbname, dbport: dbport, dbuser: dbuser, dbpassword: dbpassword, dbhost: dbhost}
}

type Config struct {
	port       int
	dbname     string
	dbport     string
	dbuser     string
	dbpassword string
	dbhost     string
}

func (c Config) DbConnection() string {
	return strings.Join([]string{c.dbuser, c.dbname, c.dbport}, ":")
}

func (c Config) DbName() string {
	return c.dbname
}
func (c Config) DbPort() string {
	return c.dbport
}
func (c Config) DbUser() string {
	return c.dbuser
}
func (c Config) DbPassword() string {
	return c.dbpassword
}
func (c Config) DbHost() string {
	return c.dbhost
}
