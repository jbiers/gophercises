package main

import "fmt"

func main() {
	oneChannel := make(chan int)
	anotherChannel := make(chan int)

	go func() {
		oneChannel <- 4
	}()

	go func() {
		anotherChannel <- 14
	}()

	var num int

	select {
	case num = <-oneChannel:
		fmt.Println(num)

	case num = <-anotherChannel:
		fmt.Println(num)

	}

	fmt.Println(num)
}
