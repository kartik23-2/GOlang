package main
// package channels

import "fmt"

func main(){
  ch:= make(chan int)
  go func(){
    fmt.Println("running")
    ch <- 2+2
  }()
  fmt.Println("----->", <-ch)
}
