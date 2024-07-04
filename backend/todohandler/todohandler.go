package todohandler

import (
	"context"
	"encoding/json"
	"net/http"
	"todo/model"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TODOService struct {
	MongoCollection *mongo.Collection
}

func (todo *TODOService) RetreiveALLTODOS(w http.ResponseWriter, r *http.Request) {
	todoList, err := todo.MongoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var todos []model.TODOList
	if err = todoList.All(context.Background(), &todos); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(todos) == 0 {
		http.Error(w, "NO TODOS", http.StatusLengthRequired)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (todo *TODOService) AddTODO(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var newTODO model.TODOList
	if err := json.NewDecoder(r.Body).Decode(&newTODO); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err := todo.MongoCollection.InsertOne(context.Background(), newTODO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (todo *TODOService) DelelteTODDO(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["id"]

	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}
	if todoID == "" {
		http.Error(w, "ID not found ", http.StatusInternalServerError)
	}
	_, err := todo.MongoCollection.DeleteOne(context.Background(), bson.M{"todo_id": todoID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (todo *TODOService) UpdateTODO(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["id"]
	if r.Method != http.MethodPut {
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
		return
	}
	upadte := bson.M{
		"$set": bson.M{
			"is_done": true,
		},
	}
	_, err := todo.MongoCollection.UpdateOne(context.Background(), bson.M{"todo_id": todoId}, upadte)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "applcation/json")
	w.WriteHeader(http.StatusOK)
}
