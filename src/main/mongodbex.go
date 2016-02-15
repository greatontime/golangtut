package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)
type Person struct{
	ID 	bson.ObjectId `bson:"_id,omitempty"`
	Name 	string
	Phone  	string
	Timestamp time.Time
}
var (
	IsDrop = true
)
func main(){
	session,err := mgo.Dial("mongodb://localhost")
	if err!= nil{
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic,true)
	if IsDrop{
		err = session.DB("test").DropDatabase()
		if err != nil{
			panic(err)
		}
	}
	c := session.DB("test").C("people")

	index := mgo.Index{
		Key : []string{"name","phone"},
		Unique:true,
		DropDups:true,
		Background:true,
		Sparse:true,
	}
	err = c.EnsureIndex(index)
	if err != nil{
		panic(err)
	}
	err = c.Insert(
		&Person{Name:"Pathomphong",Phone:"0885728468",Timestamp:time.Now()},
		&Person{Name:"Nathakorn",Phone:"0885728568",Timestamp:time.Now()})
	if err != nil{
		panic(err)
	}
	result := Person{}
	err = c.Find(bson.M{"name":"Pathomphong"}).Select(bson.M{"phone":0}).One(&result)
	if err != nil{
		panic(err)
	}
	fmt.Println("Phone",result)
	var results []Person
	err = c.Find(bson.M{}).Sort("-timestamp").All(&results)
	if err != nil{
		panic(err)
	}
	fmt.Println("Result All :",results)
	colQuerier := bson.M{"name":"Pathomphong"}
	change := bson.M{"$set":bson.M{"name":"Pathomphong Panmanee","phone":"08857284681","timestamp":time.Now()}}
	err = c.Update(colQuerier,change)
	if err != nil{
		panic(err)
	}
	err = c.Find(bson.M{}).Sort("-timestamp").All(&results)
	if err != nil {
		panic (err)
	}
	fmt.Println("Resust All :",results)
}
