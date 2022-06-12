package infrastructure

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func NewConfig() Config {
	godotenv.Load()
	port := os.Getenv("PORT")
	dbname := os.Getenv("MONGO_DB_NAME")
	dbport := os.Getenv("MONGO_PORT")
	dbuser := os.Getenv("MONGO_DB_USER")
	dbpassword := os.Getenv("MONGO_PWD")
	dbhost := os.Getenv("MONGO_HOST")
	return Config{
		Server:   ServerConfig{Port: port, TrustedSources: []string{"http://localhost:3000"}},
		Database: DatabaseConfig{dbname: dbname, dbport: dbport, dbuser: dbuser, dbpassword: dbpassword, dbhost: dbhost},
	}
}

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port           string
	TrustedSources []string
}

type DatabaseConfig struct {
	dbname     string
	dbport     string
	dbuser     string
	dbpassword string
	dbhost     string
}

func (c DatabaseConfig) DbConnection() string {
	return strings.Join([]string{c.dbuser, c.dbname, c.dbport}, ":")
}

func (c DatabaseConfig) DbName() string {
	return c.dbname
}
func (c DatabaseConfig) DbPort() string {
	return c.dbport
}
func (c DatabaseConfig) DbUser() string {
	return c.dbuser
}
func (c DatabaseConfig) DbPassword() string {
	return c.dbpassword
}
func (c DatabaseConfig) DbHost() string {
	return c.dbhost
}
