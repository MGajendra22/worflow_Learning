package main

import(
	"fmt"
)

// Struct is created

type helper struct{

	 Name string
	 Age int64
}

func (h helper) main_func () (helper) {

	return h
}

func main(){

	// Implementation of struct
	temp:=helper{"Khushi",21}
    fmt.Println(temp.main_func())

	// Array initialized
    arr :=[10]int{1,2,3,4,5,6,7,8,9,10}
    
	//slices
	vec :=arr[1:5]
	fmt.Println(vec)
    
	//Append element into slice : Result is -> {2,3,4,5,20,30,40,50,60} means changes the next elements in array
	vec=append(vec,20,30,40,50,60)
	fmt.Println(vec)

	fmt.Println(arr)
    
	// Fresh slice
	ext:=make([]int,7,10)
    
	// Added elements
	for i,_:=range ext{
		ext[i]=i
	}
    
	// ext ->{0,1,2,3,4,5,6} len: 7 and cap:10
	fmt.Println(ext,len(ext),cap(ext))
	
	ext=append(ext,7,8)
    fmt.Println(ext,len(ext),cap(ext))

	ext=append(ext,9)
    fmt.Println(ext,len(ext),cap(ext))

	ext=append(ext,10)
    fmt.Println(ext,len(ext),cap(ext))


	// Changing reference of Slice arr1 (dynamic array)
	arr1:=make([]int,2,3)
    fmt.Println(arr1)

	sl:=arr1[0:2]
    fmt.Println(sl)

	arr1[0]=12 // Changed the value 
    fmt.Println(sl)

	// Added elements
	arr1=append(arr1,1,2)
	arr1[0]=20 //Change

	sl=append(sl,2,3) //Append
    fmt.Println(arr1, sl)

}