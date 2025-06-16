package main

import (
	"fmt"
	"time"
)

type myerr struct{

	when time.Time
	what string
}

func (e* myerr) Error() string {

	return fmt.Sprintf("Error happend at :%v , %s",e.when,e.what)
}

func help() error{
     
	return &myerr{time.Now(),"Something failed"} //struct element
}

func main(){

	 err:=help()
	 if err!=nil {
		fmt.Println(err)
	 }


}