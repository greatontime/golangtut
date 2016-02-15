package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
//Movie struct
type Movie struct {
	Title  string "json:'Title'"
	Rating string "json:'Rating'"
	Year   string "json:'Year'"
	Customer int "json:'Customer'"
}
var movies = map[string] *Movie{
	"tt0076234": &Movie{Title:"AABBCCD",Rating:"10.00",Year:"2014",Customer:1000},
	"tt0076235": &Movie{Title:"AdsBCCD",Rating:"9.00",Year:"2016",Customer:1000},
	"tt0076236": &Movie{Title:"AABBCCD",Rating:"10.00",Year:"2014",Customer:1000},

}
type Customer struct {
	Id bson.ObjectId "json:'id,omitempty' bson:'_id,omitempty'"
	Status int "json:'status'"
	Name string "json:'name'"
	Detail string "json:'detail'"
}
type CustomerResource struct{
	Data Customer "json:'data'"
}
type CustomerRepo struct{
	coll *mgo.Collection
}
func (r *CustomerRepo) Find(id string)(CustomerResource,error){
	result := CustomerResource{}
	err := r.coll.FindId(bson.ObjectIdHex(id)).One(&result.Data)
	if err != nil{
		return result,err
	}
	return result,nil
}
type appContext struct{
	db *mgo.Database
}
func (c *appContext) customerHandler(req http.ResponseWriter, res *http.Request){
	params := context.Get(r, "params")
}

func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello/{name}",index).Methods("GET")
	router.HandleFunc("/movies",handleMovies).Methods("GET")
	router.HandleFunc("/movie/{id}",handleMovie).Methods("GET","DELETE","POST")
	log.Fatal(http.ListenAndServe(":8080",router))

}
func index(w http.ResponseWriter, r *http.Request){
	log.Println("Responseing to /hello request")
	log.Println(r.UserAgent())

	vars := mux.Vars(r)
	name := vars["name"]
	log.Println(vars)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w,"Hello:",name)
}
func handleMovies(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")

	outgoingJSON, error := json.Marshal(movies)

	if error != nil{
		log.Println(error.Error())
		http.Error(res,error.Error(),http.StatusInternalServerError)
		return
	}
	fmt.Fprint(res,string(outgoingJSON))
}
func handleMovie(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type","application/json")
	vars :=mux.Vars(req)
	id := vars["id"]

	log.Println("Request to ",id)

	movie, ok := movies[id]
	if !ok {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprint(res,string("movie not found"))
	}
	switch req.Method {
	case "GET":
		outgoingJSON, error := json.Marshal(movie)
		if error != nil{
			log.Println(error.Error())
			http.Error(res,error.Error(),http.StatusInternalServerError)
			return
		}
		fmt.Fprint(res,string(outgoingJSON))
	case "DELETE":
		delete(movies,id)
		res.WriteHeader(http.StatusNoContent)
	case "POST":
		movie := new(Movie)
		decoder := json.NewDecoder(req.Body)
		error := decoder.Decode(&movie)
		if error != nil {
			log.Println(error.Error())
			http.Error(res, error.Error(), http.StatusInternalServerError)
			return
		}
		movies[id] = movie
		outgoingJSON, err := json.Marshal(movie)
		if err != nil {
			log.Println(error.Error())
			http.Error(res,error.Error(),http.StatusInternalServerError)
			return
		}
		res.WriteHeader(http.StatusCreated)
		fmt.Fprint(res,string(outgoingJSON))
	}
}