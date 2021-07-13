package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"player-leaderboard/database"
	"player-leaderboard/types"
	"time"
)

// Example of sending a json with some data
func ReturnJsonExampleHandler(w http.ResponseWriter, r *http.Request){
	truck := types.Truck{Melancia: "aaaaaaaa", Gatinha: "KiwiTruck", Picole: "KIWI-001"}
	response, err := json.Marshal(truck)
	if err!= nil{
		panic(err)
	}
	//Setting header to retuurn json
	w.Header().Set("Content-Type","application/json")
	//HTTP Stattus 200
	w.WriteHeader(http.StatusOK)
	//returning a response in json
	w.Write(response)

}

//Example of getting queries form URL
func ParamsFromUrlExampleHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "plain/text")
	w.WriteHeader(http.StatusOK)
	//Geting all the queries in thee url
	queries := r.URL.RawQuery
	//Geting a single Query in the url
	id := r.URL.Query()["id"]
	//Printing all the queries
	fmt.Fprint(w, queries+"\n")
	//Printing the value of a key query
	fmt.Fprint(w, id)
}

//Example of decoding json from the body
func JsonBodytoJsonResponseexampleHandler(w http.ResponseWriter, r *http.Request){
	user := types.User{} //initialize empty user

	//Parse json request body and use it to set fields on user
	//Note that user is passed as a pointer variable so that it's fields can be modified
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		panic(err)
	}

	//Set CreatedAt field on user to current local time
	user.CreatedAt = time.Now().Local()

	//Marshal or convert user object back to json and write to response
	userJson, err := json.Marshal(user)
	if err != nil{
		panic(err)
	}

	//Set Content-Type header so that clients will know how to read response
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(userJson)
}

func InsertIntodatabaseCrapDataExample(w http.ResponseWriter, r *http.Request){
	truck := types.Truck{}

	err := json.NewDecoder(r.Body).Decode(&truck)
	if err != nil {
		panic(err)
	}

	database.InsertIntoDatabse("hello", truck)

	truckJson, err := json.Marshal(truck)
	if err != nil{
		panic(err)
	}


	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(truckJson)
}
