package main

import (
	"flag"
	"github.com/keighl/barkup"
)

func main() {
	host := flag.String("host", "localhost", "The host name where the database is located")
	port := flag.String("port", "3306", "The port number of the database")
	db := flag.String("db", "db", "The name of the database")
	user := flag.String("user", "root", "The user name of the database")
	pass := flag.String("pass", "password", "The password for the database")
	flag.Parse()

	mysql := &barkup.MySQL{
		Host:     *host,
		Port:     *port,
		DB:       *db,
		User:     *user,
		Password: *pass,
	}
	result := mysql.Export()

	if result.Error != nil {
		panic(result.Error)
	}
}
