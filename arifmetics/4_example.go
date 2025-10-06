package main

import "fmt"

func main() {

	go func() {
		fmt.Println("I'm goroutine")
	}()
	fmt.Println("I'm main here!")

}
