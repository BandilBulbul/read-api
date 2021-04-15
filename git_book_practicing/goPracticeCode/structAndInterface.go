package main

import (
	"fmt"
	"math"
)

//Struct type
type Circle struct{
	x,y,r float64
}
func main(){
	var c Circle//create  an instance  of our circle type
	c =Circle{2,4,3} //Allocates memory for all the fields
	fmt.Println(c.r)
	fmt.Println(circleArea(c))
	fmt.Println(newCircleArea(&c))
}

func circleArea(c Circle) float64{
	c.r=5
	 return math.Pi*c.r*c.r
}/*If we attempted to modify one of the fields inside of the function, it would not circleArea
modify the original variable*/
func newCircleArea( c *Circle) float64{
	*c=Circle{r:4}
	return math.Pi*c.r*c.r
}
func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r /*In between the keyword
	and the name of the
	func
	function we've added a “receiver”. The receiver is like a parameter*/
}