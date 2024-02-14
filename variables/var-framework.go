package main

import (
	"fmt"
	"reflect"
	"strconv"
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
	course := "Getting started with go"
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

}
