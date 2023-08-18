/*Create a user defined struct with 
the identifier “person”
the fields:
first
last
age
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person*/

package staticinterface

import (
	"fmt"
)

type person struct{
	firstname string
	lastname string
	age int
}

func (a person) speak(){
fmt.Printf("My name is %v %v and i am %v years old\n",a.firstname,a.lastname, a.age)

}


func Runexercise3(){

	p1:=person{
		firstname: "john",
		lastname: "Doe",
		age: 25,
	}
	p1.speak()
}
