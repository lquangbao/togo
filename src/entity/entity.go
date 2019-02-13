package entity

import (
	"fmt"
	"github.com/globalsign/mgo/bson"
	// "time"
)

// Todo entity
type Todo struct {
	ID          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string
	Description string
	Complete    bool
}

func (t *Todo) String() string {
	return fmt.Sprintf("ID: %v -- %v -- %v\n%v \n", t.ID, t.Title, t.Complete, t.Description)
}
