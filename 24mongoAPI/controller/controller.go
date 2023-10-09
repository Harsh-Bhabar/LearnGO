// package controller

// import (
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"mongoApis/model"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// const ConnectionString = "mongodb+srv://user:user@cluster0.tzw5ea6.mongodb.net/?retryWrites=true&w=majority"
// const DBName = "netflix"
// const CollectionName = "watchlist"

// // most imp as a connection
// var collection *mongo.Collection

// // connect with mongo db
// func init() {
// 	//client options
// 	clientOption := options.Client().ApplyURI(ConnectionString)

// 	// connect to mongo db
// 	client, err := mongo.Connect(context.TODO(), clientOption)
// 	errHandle(err)

// 	fmt.Println("MongoDB connection success")

// 	collection = client.Database(DBName).Collection(CollectionName)

// 	fmt.Println("Collection reference is readt ")

// }

// func errHandle(err error) {
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return
// }

// // MONGO HELPERS

// // insert one record
// func insertOneMovie(movie model.Netflix) {
// 	inserted, err := collection.InsertOne(context.Background(), movie)
// 	errHandle(err)

// 	fmt.Println("Inserted one movie with the id-", inserted.InsertedID)
// }

// // update one record
// func updateOneMovie(movieId string) {
// 	id, _ := primitive.ObjectIdFromHex(movieId)
// 	filter := bson.M{"_id": id}
// 	update := bson.M{"$set": bson.M{"watched": true}}

// 	res, err := collection.UpdateOne(context.Background(), filter, update)
// 	errHandle(err)

// 	fmt.Println("modified cnt:", res.ModifiedCount())
// }

// //delete one record

// func deleteOneMovie(movieId string) {
// 	id, _ := primitive.ObjectIdFromHex(movieId)
// 	filter := bson.M{"_id": id}
// 	deleteCnt, err := collection.DeleteOne(context.Background(), filter)
// 	errHandle(err)

// 	fmt.Println("Deleted cnt:", deleteCnt)
// }

// // delete all records

// func deleteAllRecords() {
// 	deleteRes, err := collection.DeleteMany(context.Background(), bson.M{{}}, nil)
// 	errHandle(err)

// 	fmt.Println("Deleted all records", deleteRes.DeleteCount())
// 	return deleteRes.DeleteCount()
// }

// // get all records

// func getAllRecords() []primitive.M {
// 	cursor, err := collection.Find(context.Background(), bson.M{{}})
// 	errHandle(err)

// 	var movies []primitive.M

// 	for cursor.Next(context.Background()) {
// 		var movie bson.M
// 		err := cursor.Decode(&movie)
// 		errHandle(err)

// 		movies = append(movies, movie)
// 	}

// 	defer cursor.Close(context.Background())
// 	return movies
// }

// // actual controllers - files

// func GetAllRecords(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	allMovies := getAllRecords()
// 	json.NewEncoder(w).Encode(allMovies)
// }

// func CreateMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Allow-Control-Allow-Headers", "POST")

// 	singleMovie := model.Netflix

// 	_ = json.NewDecoder(r.Body).Decode(&singleMovie)

// 	insertOneMovie(singleMovie)
// 	json.NewEncoder(w).Encode(singleMovie)
// }

// func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Allow-Control-Allow-Origin", "PUT")

// 	params := mux.Vars(r)
// 	id := params["id"]

// 	updateOneMovie(id)
// 	json.NewEncoder(w).Encode(id)
// }

// func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Allow-Control-Allow-Origin", "DELETE")

// 	param := mux.Vars(r)
// 	id := param["id"]

// 	deleteOneMovie(id)
// 	json.NewEncoder(w).Encode(id)
// }

// func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Allow-Control-Allow-Origin", "DELETE")

// 	cnt := deleteAllRecords()
// 	json.NewEncoder(w).Encode(cnt)
// }

package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mongoApis/model"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ConnectionString = "mongodb+srv://user:user@cluster0.tzw5ea6.mongodb.net/?retryWrites=true&w=majority"
const DBName = "netflix"
const CollectionName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(ConnectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection success")
	collection = client.Database(DBName).Collection(CollectionName)
	fmt.Println("Collection reference is ready")
}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted one movie with the id-", inserted.InsertedID)
}

func updateOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count:", res.ModifiedCount)
}

func deleteOneMovie(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deleteCnt, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count:", deleteCnt.DeletedCount)
}

func deleteAllRecords() int64 {
	deleteRes, err := collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted all records", deleteRes.DeletedCount)
	return deleteRes.DeletedCount
}

func getAllRecords() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie primitive.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)
	}

	return movies
}

func GetAllRecordsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllRecords()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "POST")

	var singleMovie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&singleMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	insertOneMovie(singleMovie)
	json.NewEncoder(w).Encode(singleMovie)
}

func MarkAsWatchedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "PUT")

	params := mux.Vars(r)
	id := params["id"]

	updateOneMovie(id)
	json.NewEncoder(w).Encode(id)
}

func DeleteOneMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "DELETE")

	param := mux.Vars(r)
	id := param["id"]

	deleteOneMovie(id)
	json.NewEncoder(w).Encode(id)
}

func DeleteAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "DELETE")

	cnt := deleteAllRecords()
	json.NewEncoder(w).Encode(cnt)
}
