package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/chandanpradhan-123/golang-project/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connnectionString = "mongodb+srv://chandanpradhanbxr:I0rBik4Fb86FHjOM@cluster0.xbvad91.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0&ipv6=false" //mongo db connection url
const dbName = "entflix"
const colName = "watchlist"

// Most Important
var collection *mongo.Collection

// Connect with Mongo db

func init() {
	// client options
	clientOptions := options.Client().ApplyURI(connnectionString)
	//coonect to mongo db
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")
	collection = client.Database(dbName).Collection(colName)

	// collection instance

	fmt.Println("Collection instant is ready")

}

// Mongo DB Helpers

//insert One Record

func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie in db with id:=", inserted.InsertedID)
}

// update one record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) // This method will start automatically loop and find given id;
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)

}

// delete one record

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId) //find that movie by providing movie id
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete id: ", deleteCount)

}

//delete all records from mongoDB

func deleteAllMovie() int64 {
	filter := bson.D{}
	deleteResult, err := collection.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("All Movie deleted with delete count: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Get all movies from MongoDB
func getAllMovieList() []primitive.M {
	curser, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for curser.Next(context.Background()) {
		var movie bson.M
		err := curser.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer curser.Close(context.Background())
	return movies

}

//Actuall controller - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovieList()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie models.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	fmt.Printf("body of request is %v",r.Body)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)

	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)

}
