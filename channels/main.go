package main

// channels are typed conduits through which you can send and receive values with the channel operator <-
// these are different from normal variables because we use this to communicate between goroutines

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("running")
		ch <- 2 + 2
	}()
	fmt.Println("----->", <-ch)

	// buffered channels
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	fmt.Println("----->", <-ch2)
	fmt.Println("----->", <-ch2)

	// worker goroutine using an anonymous function
	go func() {
		for i := 1; i < 5; i++ {
			
			ch3 := make(chan int, 1)
			ch3 <- i
			fmt.Println("worker ----->", <-ch3)

			// Wait for goroutine to finish execution
			time.Sleep(100 * time.Millisecond)
		}
	}()

  time.Sleep(4000 * time.Millisecond)

	
}

