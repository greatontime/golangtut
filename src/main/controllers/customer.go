package controllers
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"main/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
type(
	CustomerController struct{
		session *mgo.Session
	}
)
func NewCustomerController(s *mgo.Session) *CustomerController{
	return &CustomerController{s}
}
func (cc CustomerController) GetCustomer(res http.ResponseWriter,req *http.Request,params httprouter.Params){
	id := params.ByName("id")

	if !bson.IsObjectIdHex(id){
		res.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)
	u:=models.Customer{}
	if err := cc.session.DB("go_rest_tutorial").C("customers").FindId(oid).One(&u); err != nil{
		res.WriteHeader(404)
		return
	}
	uj, _ := json.Marshal(u)
	res.Header().Set("Content-Type","application/json")
	res.WriteHeader(200)
	fmt.Fprintf(res,"%s",uj)
	fmt.Printf("%s",oid)
}
func (cc CustomerController) CreateCustomer(res http.ResponseWriter, req *http.Request,params httprouter.Params){
	u := models.Customer{}
	json.NewDecoder(req.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	cc.session.DB("go_rest_tutorial").C("customers").Insert(u)
	uj, _ := json.Marshal(u)
	res.Header().Set("Content-Type","application/json")
	res.WriteHeader(201)
	fmt.Fprintf(res,"%s",uj)

}
func (cc CustomerController) RemoveCustomer(res http.ResponseWriter,req *http.Request,params httprouter.Params){
	id := params.ByName("id")
	if !bson.IsObjectIdHex(id){
		res.WriteHeader(404)
		return
	}
	oid := bson.ObjectIdHex(id)

	if err := cc.session.DB("go_rest_tutorial").C("customers").RemoveId(oid); err != nil{
		res.WriteHeader(404)
		return
	}
	res.WriteHeader(200)
}
func (cc CustomerController) UpdateCustomer(res http.ResponseWriter,req *http.Request,params httprouter.Params){
	id := params.ByName("id")
	if !bson.IsObjectIdHex(id){
		res.WriteHeader(404)
		return
	}
	u:=models.Customer{}
	u.Id = bson.ObjectIdHex(id)
	json.NewDecoder(req.Body).Decode(&u)
	cc.session.DB("go_rest_tutorial").C("customers").UpsertId(u.Id,u)
	uj, _ := json.Marshal(u)
	res.Header().Set("Content-Type","application/json")
	res.WriteHeader(200)
	fmt.Fprintf(res,"%s",uj)



}