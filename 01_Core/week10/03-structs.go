package main

import "fmt"

type customer struct {
	name    string
	address string
}

func main() {
	nick := customer{name: "Nick", address: "1331 W Morehead St. Charlotte, NC 28208"}
	fmt.Println(nick)
	nick.updateAddress("205 E. 14th Ave Columbus, OH 43201")
	fmt.Println(nick)
}

func (c *customer) updateAddress(a string) {
	c.address = a
}

// ## 3. Relocation (Structs)
// ### To-do
// * Make a `customer` struct
//   * `customer` should have a name field as a string
//   * `customer` should have an address field as a string
// * Create a method for customer called `updateAddress()`
//   * `updateAddress()` will take in a string as a parameter that will be the new address for the customer
//   * `updateAddress()` should update the address of the customer to the passed in string

// #### Example
// If I have a `customer`...
// ```go
// tom := customer{name: "Tom", address:"123 avenue st. Charlotte, NC 27483"}
// ```
// ... and I update his address...
// ```go
// tom.updateAddress("123 go st. Raleigh, NC 38387")
// ```

// ... Tom's new address should be the new address I put in
// ```go
// tom.address == 123 go st. Raleigh, NC 38387"
