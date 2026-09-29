package main

import "fmt"
type student struct{
    name string
    age int
    
  }

type Rect struct{
  width, height int
}

func main(){
  x:= 10
  p:=&x   //it stores the address of x
  fmt.Println("value of x", x)
  fmt.Println("address of x", p)
  fmt.Println("value of pointer p", *p)

  *p=20       // *p stores the value at address p
  fmt.Println("value of x after modification", x)


  //  STRUCT
  //--------------------

  s1:= student{"kartik", 21}
  fmt.Println(s1)

  s2:=&s1
  fmt.Println("address of s2", s2)
  fmt.Println("value of pointer s2", *s2)

  s2.name = "karthik"
  fmt.Println(s2.name)
  fmt.Println(s1.name)


  
}

//value receiver
func(r Rect)area() int {
  return r.width*r.height
}

//pointer receiver
func(r *Rect)scale(factor int){
  r.width*=factor
  r.height*=factor
}
