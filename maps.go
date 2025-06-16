package main

import (
	"fmt"
)

func maps(){
    

	// Creation of maps
	fmt.Println("\nMaps: (Method 1)")

	mp1:=map[string]int{
		"one":1,
		"two":2,
		"three":3,
	}

	for i,val:=range mp1{
		fmt.Println(i,":",val)
	}

	fmt.Println("\nMaps: (Method 2)")

	mp2:=make(map[int]int)

	mp2[1]=100
	mp2[2]=200
	mp2[3]=300

	for i,val:=range mp2{
		fmt.Println(i,":",val)
	}

	fmt.Println("\nMaps: (Method 3)")

	var mp3 map[rune]int

	mp3=make(map[rune]int)

	mp3['A']=1
	mp3['B']=2
	mp3['C']=3

	for i,val:=range mp3{
		fmt.Println(i,":",val)
	}
    
	//{value,bool}
	val,is :=mp1["one"]
	fmt.Println("value: ",val," is present?: ",is)

	delete(mp1,"two")

	ele,ok :=mp1["two"]
	fmt.Println("value: ",ele," is present?: ",ok)

	fmt.Println("Size of map m1 : ",len(mp1))
	fmt.Println("Size of map m2 : ",len(mp2))
	fmt.Println("Size of map m3 : ",len(mp3))

}

func main(){

	maps()

}