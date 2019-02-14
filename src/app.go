package main

import (
	"controller"
	"flag"
)

var (
	port = flag.String("p", "5000", "Port number (default 5000)")
)

func main() {
	flag.Parse()
	// testQuery()
	// testQueryAll()
	// testInsert()
	// testUpdate()
	// testDelete()
	controller.Serve(*port)
}

// func testInsert() {
// 	todo := entity.Todo{
// 		Title:       "another title",
// 		Description: "also no description",
// 		Complete:    false,
// 	}
// 	db.Insert(todo)
// }

// func testQueryAll() {
// 	todos := db.SelectAll()
// 	fmt.Print(todos)
// }

// func testQuery() {
// 	todo := db.Select("a title")
// 	fmt.Print(todo)
// }

// func testUpdate() {
// 	todo := entity.Todo{
// 		Title:       "a title",
// 		Description: "Now it got some description",
// 	}
// 	db.Update(todo)
// }

// func testDelete() {
// 	db.Remove("another title")
// }
