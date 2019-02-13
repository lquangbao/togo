package db

import (
	"entity"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	// "time"
)

const url = "mongodb://localhost:27017"
const database = "mydb"
const collection = "todo"

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func execute(callback func(*mgo.Collection)) {
	session, err := mgo.Dial(url)
	panicErr(err)
	c := session.DB(database).C(collection)
	defer session.Close()
	callback(c)
}

// Insert a Todo object
func Insert(t entity.Todo) {
	insertFunc := func(c *mgo.Collection) {
		err := c.Insert(&t)
		panicErr(err)
	}
	execute(insertFunc)
}

// SelectAll Todo objects
func SelectAll() (todos []entity.Todo) {
	selectFunc := func(c *mgo.Collection) {
		err := c.Find(bson.M{}).All(&todos)
		panicErr(err)
	}
	execute(selectFunc)
	return todos
}

// Update a Todo object
// func Update(t entity.Todo) {
// 	updateFunc := func(c *mgo.Collection) {

// 	}
// }