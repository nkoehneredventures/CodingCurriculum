package main

import (
	"fmt"
	"strconv"
)

func stringToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return result
}

func main() {
	s := "9"
	test := stringToInt(s)
	fmt.Printf("%T", test)
}

// ## 2. Strings to Numbers (Error Handling)
// ### To-do
// * Create a function called `stringToInt()`
// * `stringToInt()` will take in a string as a parameter but will return the string as an integer
//   * **HINT:** Use `strconv`
// * If the string is not a number, catch the error and return `0`
//   * Ex. `stringToInt(e"Fact is not a reason")` -> `0`
// * Otherwise, return the number that was represented in the string
//   * Ex. `stringToInt("9")` -> `9`
