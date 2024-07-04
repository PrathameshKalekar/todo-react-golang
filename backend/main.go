package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"todo/todohandler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading end failed ", err)
	}
	MongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URL")))
	if err != nil {
		log.Fatal("Error connecting mongo client ", err)
	} else {
		log.Println("Database connected")
	}
	if err := MongoClient.Database("todo").RunCommand(context.TODO(), bson.D{{"ping", 0}}).Err(); err != nil {
		panic(err)
	}

	log.Println("Ping Success")
}
func main() {
	defer MongoClient.Disconnect(context.Background())
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested_With", "Content-Type", "Autherization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origin := handlers.AllowedOrigins([]string{"*"})

	collection := MongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	TODOService := todohandler.TODOService{MongoCollection: collection}

	router.HandleFunc("/api/todo/retrive", TODOService.RetreiveALLTODOS).Methods(http.MethodGet)
	router.HandleFunc("/api/todo/addtodo", TODOService.AddTODO).Methods(http.MethodPost)
	router.HandleFunc("/api/todo/delete/{id}", TODOService.DelelteTODDO).Methods(http.MethodDelete)
	router.HandleFunc("/api/todo/update/{id}", TODOService.UpdateTODO).Methods(http.MethodPut)
	log.Println("Server is running at port 1906 ...")
	http.ListenAndServe(os.Getenv("PORT"), handlers.CORS(headers, methods, origin)(router))
}
