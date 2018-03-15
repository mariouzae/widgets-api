package model

import (
	"fmt"
	"widgets-api/app/dao"

	"gopkg.in/mgo.v2/bson"
)

// Widget struct model
type Widget struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Color     string        `bson:"color" json:"color"`
	Price     float32       `bson:"price" json:"price"`
	Inventory uint64        `bson:"inventory" json:"inventory"`
	Melts     bool          `bson:"melts" json:"melts"`
}

const (
	COLLECTION_WIDGET = "widget"
)

func FindWidgets() (*[]Widget, error) {
	var results []Widget
	session := dao.MongoSession.Copy()
	defer session.Close()
	err := session.DB(dao.GetDbName()).C(COLLECTION_WIDGET).Find(bson.M{}).All(&results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func FindWidgetById(id string) (*Widget, error) {
	var result Widget
	session := dao.MongoSession.Copy()
	defer session.Close()
	err := session.DB(dao.GetDbName()).C(COLLECTION_WIDGET).FindId(bson.ObjectIdHex(id)).One(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func CreateWidget(widget Widget) (*Widget, error) {
	session := dao.MongoSession.Copy()
	defer session.Close()
	w := Widget{
		bson.NewObjectId(),
		widget.Name,
		widget.Color,
		widget.Price,
		widget.Inventory,
		widget.Melts,
	}
	err := session.DB(dao.GetDbName()).C(COLLECTION_WIDGET).Insert(w)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func UpdateWidget(widget Widget) error {
	session := dao.MongoSession.Copy()
	defer session.Clone()
	fmt.Println(widget.ID)
	err := session.DB(dao.GetDbName()).C(COLLECTION_WIDGET).Update(bson.M{"_id": widget.ID}, widget)
	if err != nil {
		return err
	}
	return nil
}
