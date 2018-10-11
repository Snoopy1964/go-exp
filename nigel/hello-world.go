package main

import (
	"fmt"
	"reflect"
)

func main() {

	name := "Ralf"
	// course := "Docker Deep Dive"
	module := 3.2
	ptr := &module

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	// fmt.Println("Course is", course, "and is of type", reflect.TypeOf(course))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	fmt.Println("memory address of *module* is", &module, "and is of type", reflect.TypeOf(&module))
	fmt.Println("memory address of *module* variable is", ptr, "and the value is", *ptr)
}
