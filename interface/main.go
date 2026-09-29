//package interface
package main

import "fmt"
import "errors"

type Shape interface{
	area() float64
}

type rect struct{
	height, width int
}
type circle struct{
	radius float64
}

func(r rect) area() int{
	return r.height*r.width
}

func(c circle) area() float64{
	return 3.14 * c.radius * c.radius
}

func main(){
	// s1:=rect{10,20}
	// s2:=circle{10}
	// fmt.Println("area of s1", s1.area())
	// fmt.Println("area of s2", s2.area())

	// var i Shape
	// i=s1
	// fmt.Println("area of i", i.area())
	// i=s2
	// fmt.Println("area of i", i.area())


	result, err := divide(10,0)
	if err!=nil{
		fmt.Println("error:-", err)
		return
	}
	fmt.Println("result", result)
}
func divide(a,b int) (int, error){
	
	if(b==0){
       return 0,errors.New("can not be divided by zero")
	}
	return a/b,nil
}
