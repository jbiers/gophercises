package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 234
	}()

	num := <-c

	fmt.Println(num)
}
