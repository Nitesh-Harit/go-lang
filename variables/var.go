package main

import (
	"fmt"
)

func main() {

	name := "Nitesh"

	fmt.Println("my name is ", name)
	fmt.Println("this is not my full name.")
	changeName(&name)
	fmt.Println("this is my full name", name)

}

func changeName(name *string) string {
	*name = "Nitesh Harit"
	return *name
}
