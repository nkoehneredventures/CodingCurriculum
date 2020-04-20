package main

import (
	"fmt"
	"reflect"
)

type shape interface {
	area() int
}

type square struct {
	side int
}

type rectangle struct {
	width, height int
}

func (s square) area() int {
	return s.side * s.side
}

func (r rectangle) area() int {
	return r.width * r.height
}

func myType(s shape) string {
	bol := "rectangle" == reflect.TypeOf(s).Name()
	if bol {
		return "rectangle"
	}
	return "square"
}

func main() {
	rec := rectangle{2, 3}
	test := myType(rec)
	fmt.Println(test)
	sq := square{5}
	test2 := myType(sq)
	fmt.Println(test2)
}

// ## 4. Shapes (Reflection)
// ### To-do
// We have a couple of predefined shape structs, but we want them in an interface. Your task for this challenge is to:

// * Create a shape interface united by the method `area()`
// * Create a function called `myType()`
//   * `myType()` takes in a shape type as a parameter and returns a string of the type of shape it is
//   * Ex: If the shape is of type `square`, it will return `"square"`. If the shape is of type `rectangle`, it will return `"rectangle"`.
// * **HINT:** Use the `TypeOf` method from the `reflect` library to find the type of the variable. `TypeOf` returns a `type`. To convert it to a string, you can use the `Name()` method.

// #### Example
// ```go
// // Creates a new rectangle
// rec := rectangle{2, 3}
// // This is a type
// reflect.TypeOf(rec) -> main.rectangle
