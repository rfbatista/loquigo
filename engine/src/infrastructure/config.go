package infrastructure

import (
	"strings"
)

type Config struct {
	port       int
	dbname     string
	dbport     string
	dbuser     string
	dbpassword string
}

func (c *Config) Init() {
	c.port = 3000
	c.dbname = ""
	c.dbport = ""
	c.dbuser = ""
	c.dbpassword = ""
}

func (c Config) DbConnection() string {
	return strings.Join([]string{c.dbuser, c.dbname, c.dbport}, ":")
}
