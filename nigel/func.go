package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "function basics"
	author := "nigel poulton"

	fmt.Println(converter(module, author))
}

func converter(m, a string) (s1, s2 string) {
	m = strings.Title(m)
	a = strings.ToUpper(a)

	return m, a
}
