package main

import (
	"fmt"
	"math"
)

//INTERFACE EXAMPLE :---It is collections of method signatures.

type geometry interface { //interface for geometric shapes.
	area() float64
	perimeter() float64
}
//implement this interface on rect and circle types.

/*
To implement an interface in Go,
we just need to implement all the methods in the interface.
Here we implement geometry on rects as well as circle.*/
type rect struct{
	width, height float64
}
type circle struct{
	radius float64
}
func (r rect) area() float64{
	return r.width*r.height
}

func (r rect) perimeter() float64{
	return 2*r.width+2*r.height
}

func (c circle) area() float64{
	return math.Pi*c.radius*c.radius
}
func (c circle) perimeter() float64{
	return 2*math.Pi*c.radius
}

func measure(g geometry){ /*If a variable has an interface
	type, then we can call methods that are in the
	named interface.*/
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main(){
	r :=rect{width:3,height:4}
	c :=circle{radius:5}
	measure(r)
	measure(c)
}
/*The circle and rect struct types both implement the geometry interface
so we can use instances of these structs as
arguments to measure.*/






