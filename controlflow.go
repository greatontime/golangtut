package main

import "fmt"

func main() {
	var name string
	fmt.Print("What's your name ?")
	fmt.Scan(&name)
	fmt.Println(getAge(name))
}
func getAge(name string) (int, string) {
	var age int
	var sex string

	switch name {
	case "pathomphong":
		age, sex = 33, "men"
	case "nathakorn":
		age, sex = 4, "boy"
	case "nathaporn":
		age, sex = 1, "girl"
	default:
		age, sex = 0, "-"
	}
	return age, sex
}
