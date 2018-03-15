package dao

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var (
	MongoSession *mgo.Session
	dbInfo       string
)

type MongoDAO struct {
	Addrs    string
	Timeout  time.Duration
	Database string
	Username string
	Password string
}

func (m *MongoDAO) Connect() {
	dbInfo = m.Database
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{m.Addrs},
		Timeout:  60 * time.Second,
		Database: m.Database,
		Username: m.Username,
		Password: m.Password,
	}
	session, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	MongoSession = session
}

func GetDbName() string {
	return dbInfo
}
