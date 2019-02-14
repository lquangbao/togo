package db

import (
	"entity"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
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

// Select one todo object
func Select(title string, todoChannel chan entity.Todo) {
	var todo entity.Todo
	selectFunc := func(c *mgo.Collection) {
		err := c.Find(bson.M{"title": title}).One(&todo)
		panicErr(err)
	}
	execute(selectFunc)
	time.Sleep(time.Second * 5) // test goroutine
	todoChannel <- todo
}

// SelectAll Todo objects
func SelectAll() (todos []entity.Todo) {
	selectFunc := func(c *mgo.Collection) {
		err := c.Find(nil).All(&todos)
		panicErr(err)
	}
	execute(selectFunc)
	return todos
}

// Update a Todo object
func Update(t entity.Todo) {
	query := bson.M{"title": t.Title}
	updateFunc := func(c *mgo.Collection) {
		err := c.Update(query, t)
		panicErr(err)
	}
	execute(updateFunc)
}

func Remove(title string) {
	removeFunc := func(c *mgo.Collection) {
		err := c.Remove(bson.M{"title": title})
		panicErr(err)
	}
	execute(removeFunc)
}
