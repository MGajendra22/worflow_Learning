package main

import(
	"fmt"
)

type student struct{
    name string
	age int
	grade rune
}


type info interface{
	Age() int
	Name() string
	Grade() rune
	Access() bool

}

func (s student )Age() int{
	return s.age
}

func (s student) Name() string{
	return s.name
}

func(s student) Grade() rune{
	return s.grade
}

func(s student) Access() bool{
	if(s.Age()>=18) {
		return true
	} 
	return false
}
func main(){

	var s1 info

	s1=student{"Gajendra",21,'B'}

	fmt.Println("Show the information of student : ")
	fmt.Println("Name : ",s1.Name())
	fmt.Println("Age : ",s1.Age())
	fmt.Println("Grade : ",string(s1.Grade()))
	fmt.Println("Eligible for driving  : ",s1.Access())

}