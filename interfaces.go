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
func (s student) String() string {

	 return fmt.Sprintf("Student Name : %s , Age : %d , Grade : %c ",s.name,s.age,s.grade)
}

func main(){

	var s1 info

	s1=student{"Gajendra",21,'B'}

	fmt.Println(s1)

	fmt.Println("Show the information of student : ")
	fmt.Println("Name : ",s1.Name())
	fmt.Println("Age : ",s1.Age())
	fmt.Println("Grade : ",string(s1.Grade()))
	fmt.Println("Eligible for driving  : ",s1.Access())

    fmt.Println("Interface inside main : ")

	var inter interface{}

	inter=10
    fmt.Printf("\n%v : %T",inter,inter)

	inter="Hello Zopdev"
    fmt.Printf("\n%v : %T",inter,inter)

	 
    // Type assertions 
	fmt.Println("Type assertions : ")

	var i interface {}="Range"

	s:=i.(string)
	fmt.Println(" : ",s)

	f,ok:=i.(float32)

	fmt.Println(f,ok)



}