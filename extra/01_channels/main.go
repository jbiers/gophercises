package main

import "fmt"

func main() {
	oneChannel := make(chan int)

	go func() {
		oneChannel <- 234
	}()

	num := <-oneChannel

	fmt.Println(num)
}
