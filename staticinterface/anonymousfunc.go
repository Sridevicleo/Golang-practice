/*Hands-on exercise #6
Build and use an anonymous func 
code: https://play.golang.org/p/DQX3xEIcRe  
Topic: 107*/

package staticinterface

import (
	"fmt"
)
func main(){

	ch:=make (chan chan int)

	func(){
for i:=0; i<10; i++{
	ch <- i

}
close(ch)
for v :=range ch{
fmt.Println(v)

	}()
}
}