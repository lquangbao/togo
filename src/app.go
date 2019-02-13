package main

import (
	"db"
	"entity"
	"fmt"
)

func main() {
	testQueryAll()
}

func testInsert() {
	todo := entity.Todo{
		Title:       "a title",
		Description: "no description",
		Complete:    false,
	}
	db.Insert(todo)
}

func testQueryAll() {
	todos := db.SelectAll()

	fmt.Print(todos)
}
