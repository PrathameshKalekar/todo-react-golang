package model

type TODOList struct {
	TODOId string `json:"todo_id" bson:"todo_id"`
	TODO   string `json:"todo" bson:"todo"`
	IsDone bool   `json:"is_done" bson:"is_done"`
}
