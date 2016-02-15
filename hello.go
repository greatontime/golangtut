package main

import "fmt"

func main() {
	//	var a int
	//	var b string
	//	var c float64
	//	var d bool
	//
	//	fmt.Println("%v \n",a)
	//	fmt.Println("%v \n",b)
	//	fmt.Println("%v \n",c)
	//	fmt.Println("%v \n",d)
	//	fmt.Println(increment())
	//	fmt.Println(increment())
	//	fmt.Println(increment())
	//	fmt.Println(increment())
	//
	//	decreatment := func() int{
	//		x--
	//		return x
	//	}
	//	fmt.Println(decreatment())
	//	update := wrapper()
	//	fmt.Println(update())
	//	fmt.Println(update())

	//	res, _ := http.Get("http://www.greatonshop.com/")
	//	page, _ := ioutil.ReadAll(res.Body)
	//	res.Body.Close()
	//	fmt.Printf("%s",page)
	//	a := 43
	//	fmt.Println("a - ",a)
	//	fmt.Println("a's memory address -",&a)
	//	var meters float64
	//	fmt.Print("Enter meters swam : ")
	//	fmt.Scan(&meters)
	//	yards := meters * metersToYards
	//	fmt.Println("meter is ",yards, " yards")
	//	a := 43
	//	fmt.Println(a)
	//	fmt.Println(&a)
	//
	//	var b *int = &a
	//	var c = a
	//	fmt.Println(b)
	//	fmt.Println(*b)
	//	fmt.Println(c)
	//	x := 5
	//	fmt.Println("%p\n",&x)
	//	fmt.Println(&x)
	//	zero(x)
	//	fmt.Println(x)
	//	evenOdd(100)
	fmt.Println([]byte("1"))
}

//func multiplyVar(x int) int{
//	return x+x
//}
//func increment()int {
//	x++
//	return x
//}
//func wrapper() func() int{
//	return func() int{
//		x++
//		return x
//	}
//}
func zero(x int) {
	fmt.Println("%p\n", &x)
	fmt.Println(&x)
	x = 0
}
func evenOdd(x int) {
	for i := 0; i < x; i++ {
		if i%2 == 1 {
			fmt.Println(i, "odd")
		} else {
			fmt.Println(i, "even")
		}
	}
}
