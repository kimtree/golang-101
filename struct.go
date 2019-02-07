package main

import "fmt"

// struct is a collection of variables
type User struct {
	uid float64
	age int
	name string
}

func (u User) printInformations() string {
	return fmt.Sprintf("%.4f: %s(%d)", u.uid, u.name, u.age)
}

func main()  {
	var kimtree = User{1.1024, 29, "Namwoo"}

	fmt.Println(kimtree.name)
	fmt.Println(kimtree.printInformations())
}