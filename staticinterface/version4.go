/*Hands-on exercise #5
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/

package staticinterface

import (
	"fmt"
	"math"

)

type square struct{
	length float64
	width float64

}
type circle struct{
	radius float64
}
type shape interface{
	Area()float64
}


func (s *square) Area()float64{
area:= s.length*s.width
//fmt.Println("The area of the square is ",area)
return area
}

func (c *circle) Area()float64{
	//var area float64
	 area:=math.Pi*math.Pow(c.radius, 2)
	//fmt.Println("the area of the crcle is ", area)
	return area
}

func info(sh shape){
	//fmt.Println(sh.Area())
	fmt.Printf("The area of %T is %v\n",sh,sh.Area())
}

func Runexercise4(){
	s1:=square{
		length: 30,
		width: 40,
	}

	c1:=circle{
		radius: 10,
	}
	info(&s1)
	info(&c1)

	//fmt.Println("the area of square is "s1.squarearea)
	//fmt.Println("the area of the circle is "c1.circlearea)
}