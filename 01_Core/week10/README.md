# Golang
For these challenges, you'll find starter files marked with the name of the challenge itself. 

## 1. Making Acronyms (Strings)
### To-do
* Make a function called `acronymMaker()` that takes in a string of words separated by a space
* `acronymMaker` should return the acronym of the string you put into it
  * Ex. `acronymMaker(“Fact is not a reason”)` -> `"FINAR"`
* Make sure that the returned string is all in uppercase
* **HINT:** The `strings` package has a `split()` and `ToUpper()` function which could be helpful for this problem

## 2. Strings to Numbers (Error Handling)
### To-do
* Create a function called `stringToInt()`
* `stringToInt()` will take in a string as a parameter but will return the string as an integer
  * **HINT:** Use `strconv`
* If the string is not a number, catch the error and return `0`
  * Ex. `stringToInt("Fact is not a reason")` -> `0`
* Otherwise, return the number that was represented in the string
  * Ex. `stringToInt("9")` -> `9`

## 3. Relocation (Structs)
### To-do
* Make a `customer` struct
  * `customer` should have a name field as a string
  * `customer` should have an address field as a string
* Create a method for customer called `updateAddress()`
  * `updateAddress()` will take in a string as a parameter that will be the new address for the customer
  * `updateAddress()` should update the address of the customer to the passed in string

#### Example
If I have a `customer`...
```go
tom := customer{name: "Tom", address:"123 avenue st. Charlotte, NC 27483"}
```
... and I update his address...
```go
tom.updateAddress("123 go st. Raleigh, NC 38387")
```

... Tom's new address should be the new address I put in
```go
tom.address == 123 go st. Raleigh, NC 38387"
```

## 4. Shapes (Reflection)
### To-do
We have a couple of predefined shape structs, but we want them in an interface. Your task for this challenge is to:

* Create a shape interface united by the method `area()`
* Create a function called `myType()`
  * `myType()` takes in a shape type as a parameter and returns a string of the type of shape it is
  * Ex: If the shape is of type `square`, it will return `"square"`. If the shape is of type `rectangle`, it will return `"rectangle"`.
* **HINT:** Use the `TypeOf` method from the `reflect` library to find the type of the variable. `TypeOf` returns a `type`. To convert it to a string, you can use the `Name()` method.

#### Example 
```go
// Creates a new rectangle
rec := rectangle{2, 3} 
// This is a type
reflect.TypeOf(rec) -> main.rectangle 
```

## 5. Hello World (Channels)
### To-do
* Finish the `hello()` function so that it inputs in the passed in [channel](https://tour.golang.org/concurrency/2) the string `"hello"`
  * Do not remove the sleep function
* Finish the `world()` function so that it inputs in the passed in channel the string `"world"`
  * Do not remove the sleep function
* In the `helloWorld()` function:
  * Create a string channel
  * Call the `hello()` and `world()` functions concurrently ([Goroutine](https://tour.golang.org/concurrency/1)) with the created channel
  * Create a string using the 2 channel outputs and separate the outputs with a space
    * The order of the output will always be the same in this case because of the sleep function
    * Ex. `helloWorld()` -> `"hello world"`