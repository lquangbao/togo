package main

import (
	"db"
	"entity"
	"fmt"
)

func main() {
	testQuery()
	// testQueryAll()
	// testInsert()
}

func testInsert() {
	todo := entity.Todo{
		Title:       "another title",
		Description: "also no description",
		Complete:    false,
	}
	db.Insert(todo)
}

func testQueryAll() {
	todos := db.SelectAll()
	fmt.Print(todos)
}

func testQuery() {
	todo := entity.Todo{
		Title: "a title",
	}
	todo = db.Select(todo)
	fmt.Print(todo)
}
