package main

import "fmt"

// struct is a collection of variables
type User struct {
	uid float64
	age int
	name string
}


func main()  {
	var kimtree = User{1, 29, "Namwoo"}

	fmt.Println(kimtree.name)
}