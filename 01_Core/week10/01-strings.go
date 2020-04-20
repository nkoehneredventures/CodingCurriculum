package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Fact is Not a Reason"
	test := acronymMaker(s)
	fmt.Println(test)
}

func acronymMaker(s string) string {
	splits := strings.Split(s, " ")
	result := ""
	for _, str := range splits {
		result += str[0:1]
	}
	return strings.ToUpper(result)
}

// * Make a function called `acronymMaker()` that takes in a string of words separated by a space
// * `acronymMaker` should return the acronym of the string you put into it
//   * Ex. `acronymMaker(“Fact is not a reason”)` -> `"FINAR"`
// * Make sure that the returned string is all in uppercase
// * **HINT:** The `strings` package has a `split()` and `ToUpper()` function which could be helpful for this problem
