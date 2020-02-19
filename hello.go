package main

import "fmt"

func main() {

	var a int
	var b string
	var c float64
	var d bool

	type user struct {
		Name string
		Umur int64
	}

	u1 := user{"user1", 21}

	fmt.Println(u1.Name)

	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)

	fmt.Println()
}
