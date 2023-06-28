package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Spirolina/todolist-restapi/models"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.iz1vu.mongodb.net/?retryWrites=true&w=majority", username, password)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	db := client.Database("todo-list")
	collection = db.Collection("todos")

}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		handleError(w, err)
	}

	_, err = collection.InsertOne(context.Background(), todo)
	if err != nil {
		handleError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	tempId := getTodoID(r)
	if tempId == "" {
		handleError(w, fmt.Errorf("missing todo ID"))
		return
	}

	id, err := primitive.ObjectIDFromHex(tempId)
	if err != nil {
		handleError(w, err)
		return
	}

	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		handleError(w, err)
		return
	}

	if result.DeletedCount == 0 {
		handleError(w, fmt.Errorf("todo not found"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, `{"error": %s}`, err.Error())
}

func getTodoID(r *http.Request) string {
	id := r.URL.Path[len("/todos/"):]
	return id
}
