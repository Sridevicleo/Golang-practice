/* 1) Create a struct type called companyA with one field of type string called name.

2) Create a struct type called companyB with one field of type string called name.

3) Create a method for the companyA type (pointer-receiver) called SetName() 
that takes one argument of type string, and set the name field.

4) Create a method for the companyA type (pointer-receiver) called GetName() 
that takes no arguments, and print the name field.

5) Create a method for the companyB type (pointer-receiver) called SetName() 
that takes one argument of type string, and set the name field.

6) Create a method for the companyB type (pointer-receiver) called GetName() 
that takes no arguments, and print the name field.

7) Create a wrapper interface with two methods called GetName() and SetName().

8) Create a function called NewWrapper() that takes one argument of type wrapper.

9) Inside the NewWrapper() function call the SetName() method on the argument of 
type wrapper.

10) Inside the NewWrapper() function call the GetName() method on the argument of 
type wrapper.  *\

*/

package staticinterface

import (
	"fmt"
)

type companyA struct{
	name string
}

type companyB struct{
	name string
}

type wrapper interface{
	SetName(name string)
	GetName()
}

func Run(){
c1:=companyA{}
NewWrapper(&c1)
}

func (c *companyA) SetName(name string){

	c.name =name
}

func (c *companyA) GetName(){
fmt.Printf("The name of company A is %v",c.name)
}

func (c *companyB) SetName(name string){

	c.name =name
}

func (c *companyB) GetName(){
fmt.Printf("The name of company B is %v",c.name)
}

func NewWrapper(input wrapper){
	input.SetName("THIS IS THE COMPANY NAME\n")
	input.GetName()


} 