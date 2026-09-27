package main

import "fmt"

func main(){
  x:= 10
  p:=&x   //it stores the address of x
  fmt.Println("value of x", x)
  fmt.Println("address of x", p)
  fmt.Println("value of pointer p", *p)

  *p=20       // *p stores the value at address p
  fmt.Println("value of x after modification", x)


  
}
