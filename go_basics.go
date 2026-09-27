package main

import (
	"fmt"
	"time"
)

func add(a int, b int) int{
	return a+b
}

func main() {
	// // Variables :- 3 ways to declare/initialize variables in Go
	// //---------------------------------------------------------------------

	// // 1. Explicit type declaration: var var_name data_type = value
	// var name1 string = "kartik"

	// // 2. Type inference declaration: var var_name = value
	// var name2 = "kartik"

	// // 3. Short declaration operator: var_name := value
	// name3 := "kartik"

	// // using fmt.Printf
	// fmt.Println("name1:", name1)

	// fmt.Println(name1, "\n", name2, "\n", name3)

	// // var age float32 = 21.0
	// // var istrue bool

	// // fmt.Printf("Age is :-%f\n", age)
	// // fmt.Println(istrue)

	// var experience = "20 months"
	// fmt.Println("experience is", experience)

	// //constants
	// const PI = 3.14
	// fmt.Println("the value of PI is :-", PI)

	// //VERBS
	// //--------------

	age := 21.4
	fmt.Printf("age is %v and data type is %T", age, age) // %v :- print the value 	 %T :- print the type
	var date time.Time
	fmt.Println("\n", date)

	// add(3, 6)

	// // CONTROL STRUCTURE(LOOPS, IF, SWITCH)
	// //-----------------------------------------

	// age_new := 21
	// if age_new < 18 {
	// 	fmt.Println("you are minor")
	// } else if age_new > 18 && age_new < 60 {
	// 	fmt.Println("you are major")
	// } else {
	// 	fmt.Println("you are old")
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("the square of %d is %d\n", i, i*i)
	// }

	// day:= "Monday"

	// switch day{
	// case "Monday":
	// 	fmt.Println("today is monday")
	// case "Tuesday":
	// 	fmt.Println("today is tuesday")
	// case "Wednesday":
	// 	fmt.Println("today is wednesday")
	// case "Thursday":
	// 	fmt.Println("today is thursday")
	// case "Friday":
	// 	fmt.Println("today is friday")
	// case "Saturday":
	// 	fmt.Println("today is saturday")
	// case "Sunday":
	// 	fmt.Println("today is sunday")
	// default:
	// 	fmt.Println("invalid day")
	// }

	nums := []int{1, 2, 3, 4, 5}

	for _, i := range nums { //using RANGE
		fmt.Println(i)
	}

	fmt.Println("sum is:- ",add(3,6))

	var n int
	fmt.Println("enter value of n :- ")
	fmt.Scan(&n)
	fmt.Println("you entered:- ",n)

}
