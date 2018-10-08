package greeting

import (
	"fmt"
)

// -------------- Structs -------------------------

// Salutation is my special function for salutation
type Salutation struct {
	Name     string
	Greeting string
}

// Salutations is a slice of Salution
type Salutations []Salutation

// Printer returns a function definition
type Printer func(string)

//----------- Interfaces -----------------------

// Renamable interface
type Renamable interface {
	Rename(name string)
}

// ------------------ Methods -------------------

// Rename a name of Salutation (example for methods)
func (salut *Salutation) Rename(newName string) {
	salut.Name = newName
}

// Write implementation
func (salut *Salutation) Write(p []byte) (n int, err error) {
	s := string(p)
	salut.Rename(s)
	n = len(s)
	err = nil
	return
}

// ---------------------- Functions ------------------

// Greet is my personal greeting message
func (saluts Salutations) Greet(do Printer, isFormal bool, times int) {

	for _, s := range saluts {
		message, alternate := CreateMessage(s.Name, s.Greeting)
		// _, alternate := CreateMessage(salutation.greeting, salutation.name)
		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

// ChannelGreeter is same as above but with a channel input
func (saluts Salutations) ChannelGreeter(c chan Salutation) {
	for _, s := range saluts {
		c <- s
	}
	close(c)
}

//GetPrefix selfexplaining: get prefix
func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{

		"Bob":    "Mr. ",
		"Ralf":   "Dr. ",
		"Aileen": "Dr. ",
		"Marie":  "Mrs. ",
		"Mary":   "Frl. ",
	}

	prefixMap["Aileen"] = "Jr. "

	delete(prefixMap, "Mary")

	if value, exists := prefixMap[name]; exists {
		return value
	}

	return "Dude "
}

// Print ...
func Print(s string) {
	fmt.Print(s)
}

// PrintLine ...
func PrintLine(s string) {
	fmt.Println(s)
}

// CreatePrintFunction is a custom printer
func CreatePrintFunction(c1 string) Printer {
	return func(s string) {
		fmt.Println(s + c1)
	}
}

// CreateMessage creates a greeeting message
func CreateMessage(name string, greeting string) (message string, alternate string) {
	// fmt.Println(len(greeting))
	message = name + ", " + greeting
	alternate = "HEY! " + name
	return
}

//TypeSwitchTest : test types within switche
func TypeSwitchTest(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("String")
	case Salutation:
		fmt.Println("Salutation")
	default:
		fmt.Println("unknown")
	}
}
