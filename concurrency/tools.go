package main

import "fmt"

func main() {
	c := make(chan string)
	go func(input chan string) {
		input <- "i am from another channe;"
	}(c)

	greeting := <-c
	fmt.Println(greeting)
	helloWorld()
}

func helloWorld() {
	fmt.Println("Hello world123")
}
