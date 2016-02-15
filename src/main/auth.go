package main
import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codegangsta/negroni"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)
var (
	signingKey = "greatontime"
	privateKay []byte
	publicKey []byte
)
func init(){
	privateKay, _ = ioutil.ReadFile("src/main/private_key")
	publicKey, _ = ioutil.ReadFile("src/main/public_key.pub")
}
func AuthMiddleware(w http.ResponseWriter, r *http.Request,next http.HandlerFunc){
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
func APIHandler(w http.ResponseWriter,r *http.Request,next http.HandlerFunc){
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Got it")
}
func main(){
	router := mux.NewRouter()
	n := negroni.Classic()

	router.HandleFunc("/login",func(w http.ResponseWriter,r *http.Request){
		token := jwt.New(jwt.GetSigningMethod("RS256"))
		tokenString, _ := token.SignedString(privateKay)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,tokenString)
	})
	router.Handle("/api",negroni.New(negroni.HandlerFunc(AuthMiddleware),negroni.HandlerFunc(APIHandler)))
	n.UseHandler(router)
	http.ListenAndServe(":3000",n)
}