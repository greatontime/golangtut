package main
import (
	"routers"
	"settings"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main(){
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":8080",n)

}