package model

import (
	"widgets-api/app/dao"

	"gopkg.in/mgo.v2/bson"
)

const (
	COLLECTION = "user"
)

// User struct model
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Gravatar string        `bson:"gravatar" json:"gravatar"`
	Password string        `bson:"password" json:"password"`
}

func FindUsers() (*[]User, error) {
	var results []User
	session := dao.MongoSession.Copy()
	defer session.Close()
	err := session.DB(dao.GetDbName()).C(COLLECTION).Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func FindUserById(id string) (*User, error) {
	var result User
	session := dao.MongoSession.Copy()
	defer session.Close()
	err := session.DB(dao.GetDbName()).C("user").FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func FindUserByName(name string) (*User, error) {
	var result User
	session := dao.MongoSession.Copy()
	defer session.Close()
	err := session.DB(dao.GetDbName()).C(COLLECTION).Find(bson.M{"name": name}).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
