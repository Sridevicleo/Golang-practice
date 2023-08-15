package main

import (
	"fmt"
	"mypractice/addition"
)

func main(){

	var a interface{}
	a=[3]int{1,2,3}
	
	
	switch a.(type) {

	case int:
		fmt.Printf("a interface is of type int %v",a.(int))
	case string:
		fmt.Printf("a interface is of type string %v", a.(string))
	case [3]int:
		fmt.Printf("a interface is an integer array %v\n", a.([3]int))
	}

	sum1:=addition.Adding(20,30)
	fmt.Println(sum1)
}