package main

import (
	"fmt"
	"time"
)

func helloWorld() string {
	ch1 := make(chan string)

	go hello(ch1)
	go world(ch1)

	st := <-ch1 + " " + <-ch1
	return st
}

func hello(c chan string) {
	time.Sleep(1 * time.Second)
	c <- "hello"
}

func world(c chan string) {
	time.Sleep(2 * time.Second)
	c <- "world"
}

func main() {
	fmt.Println(helloWorld())
}

// ## 5. Hello World (Channels)
// ### To-do
// * Finish the `hello()` function so that it inputs in the passed in [channel](https://tour.golang.org/concurrency/2) the string `"hello"`
//   * Do not remove the sleep function
// * Finish the `world()` function so that it inputs in the passed in channel the string `"world"`
//   * Do not remove the sleep function
// * In the `helloWorld()` function:
//   * Create a string channel
//   * Call the `hello()` and `world()` functions concurrently ([Goroutine](https://tour.golang.org/concurrency/1)) with the created channel
//   * Create a string using the 2 channel outputs and separate the outputs with a space
//     * The order of the output will always be the same in this case because of the sleep function
//     * Ex. `helloWorld()` -> `"hello world"`
