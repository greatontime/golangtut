package main
import "fmt"

type User struct {
	Id  int
	name  string
	location string
}
type Player struct {
	User
	position string
}

func main(){
	p :=Player{}
	p.Id = 1144
	p.name = "pathomphong panmanee"
	p.location = "1121232131,123131234"
	p.position = "golang"
	fmt.Printf("%+v",p)
}
