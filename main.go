package main

import (
	"log"
	"net/http"

	"widgets-api/app/dao"
	"widgets-api/app/route"
	"widgets-api/app/util"
	"widgets-api/config"
)

var (
	db   = dao.MongoDAO{}
	conf = config.Config{}
)

func init() {
	// Read configuration from config.toml file
	conf.Read()

	// Set db info
	db.Addrs = conf.Database.ServerAddr
	db.Username = conf.Database.Username
	db.Password = conf.Database.Password
	db.Database = conf.Database.DbName
	db.Timeout = conf.Database.Timeout
	db.Connect()
}

func main() {
	cert := util.GetRootPath() + "/../src/widgets-api/cert/cert.pem"
	key := util.GetRootPath() + "/../src/widgets-api/cert/key.pem"
	log.Fatal(http.ListenAndServeTLS(":4000", cert, key, route.Load()))
	//log.Fatal(http.ListenAndServe(":4000", route.Load()))
}
