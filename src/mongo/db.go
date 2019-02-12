package db

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

const url = "mongodb://localhost:27017"
const database = "mydb"
const collection = "todo"

type todo struct {
	ID          bson.ObjectId `bson:"_id, omitempty`
	Title       string
	Description string
	Create      time.Time
	Expire      time.Time
	Complete    bool
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func execute(callback func(*mgo.Collection) bool) {
	session, err := mgo.Dial(url)
	panicErr(err)
	c := session.DB(database).C(collection)
	callback(c)
	defer session.Close()
}
