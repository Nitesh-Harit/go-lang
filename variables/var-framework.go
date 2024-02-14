package main

import (
	"fmt"
)

/*
var (

	name   = "Nitesh"
	course = "Getting started with GO"
	module = "10"
	clip   = 1

)
*/
func main() {
	//name := "Nitesh"
	/*course := "Getting started with go"
	module := "10"
	clip := 1
	fmt.Println("this is the value of module", module)
	fmt.Println("Modeule is of type", reflect.TypeOf(module))

	iModule, err := strconv.Atoi(module)
	if err == nil {
		total := iModule + clip
		fmt.Println("the value of total is ", total)
	}

	ptr := &course
	fmt.Println("value of ptr", ptr)
	*/
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
