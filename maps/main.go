package main
import "fmt"

func main(){
	var name string 
	var age int
	var height float64
	var isstudent bool
	name="kartik"
	age=21
	height=5.9
	isstudent=true
	fmt.Println("name:",name)
	fmt.Println("age:",age)
	fmt.Println("height:",height)
	fmt.Println("isstudent:",isstudent)


	m:=make(map[string]string)
	m["name"]="kartik"
	m["age"]="21"
	m["height"]="5.9"
	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(m["age"])
	fmt.Println(m["height"])

	//delete
	delete(m,"age")	
	fmt.Println(m)

}