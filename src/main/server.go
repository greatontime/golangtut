package main

import (
	// Standard library packages
	"net/http"
	"fmt"
	"io/ioutil"
	"main/models"
	"encoding/json"

	// Third party packages
	"github.com/julienschmidt/httprouter"
	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
//	"github.com/swhite24/go-rest-tutorial/controllers"
//	"main/controllers"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

)

func main() {
	// Instantiate a new router
	r := httprouter.New()
	n := negroni.Classic()

	// Get a UserController instance
//	uc := controllers.NewUserController(getSession())
//	cc := controllers.NewCustomerController(getSession())
//
//	// Get a user resource
//	r.GET("/user/:id", uc.GetUser)
//	r.GET("/customer/:id",cc.GetCustomer)
//
//	// Create a new user
//	r.POST("/user", uc.CreateUser)
//	r.POST("/customer",cc.CreateCustomer)
//
//	r.PUT("/customer/:id",cc.UpdateCustomer)
//
//	// Remove an existing user
//	r.DELETE("/user/:id", uc.RemoveUser)
//	r.DELETE("/customer/:id",cc.RemoveCustomer)

	// Fire up the server
	r.HandlerFunc("GET","/login",getToken)
	r.Handler("GET","/api",
		negroni.New(negroni.HandlerFunc(AuthMiddleware1),negroni.HandlerFunc(APIHandler1),negroni.Handler(CreateCustomer1)))
//	r.Handler("POST","/customer",
//		negroni.New(negroni.HandlerFunc(AuthMiddleware1),negroni.HandlerFunc(APIHandler1),cc.CreateCustomer))
	n.UseHandler(r)
	http.ListenAndServe("localhost:8080", n)
}

// getSession creates a new mongo session and panics if connection error occursf
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost/go_rest_tutorial")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
var (
	signingKey = "greatontime"
	privateKay []byte
	publicKey []byte
)
func init(){
	privateKay, _ = ioutil.ReadFile("src/main/private_key")
	publicKey, _ = ioutil.ReadFile("src/main/public_key.pub")
}
func AuthMiddleware1(w http.ResponseWriter, r *http.Request,next http.HandlerFunc){
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token)(interface{},error){
		return publicKey,nil
	})
	if err == nil && token.Valid{
		next(w,r)
	}else{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w,"GO HOME")
	}
}
func APIHandler1(w http.ResponseWriter,r *http.Request,next http.HandlerFunc){
	w.WriteHeader(http.StatusOK)
	next(w,r)
}
func getToken(w http.ResponseWriter,r *http.Request){
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	tokenString, _ := token.SignedString(privateKay)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,tokenString)
}
func CreateCustomer1(res http.ResponseWriter, req *http.Request,params httprouter.Params,next http.HandlerFunc){
	u := models.Customer{}
	json.NewDecoder(req.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	err := mgo.Session.DB("go_rest_tutorial").C("customers").Insert(u)
	fmt.Printf("%s",err)
	uj, _ := json.Marshal(u)
	res.Header().Set("Content-Type","application/json")
	res.WriteHeader(201)
	fmt.Fprintf(res,"%s",uj)
	next(res,req)

}

