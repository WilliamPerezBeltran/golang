package main 
import "fmt"

type person struct{
	name string
	age int
}

func newPerson(name string) *person{
	p := person{name:name}
	p.age = 42
	return &p

}


func main(){
	fmt.Println(person{"Bod",20})
	fmt.Println(person{name:"alice",age:30})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann", age: 40})
}
