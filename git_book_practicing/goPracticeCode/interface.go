package main

import "fmt"

type Shape interface{
	area() float64
	//perimeter() float64
}

type MultiShape struct{
	shapes []Shape
}
//defining the method
func (m *MultiShape) area() float64{
	var area float64
	for _,s:=range m.shapes{
		area +=s.area()

	}
	return area
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}

	return area
}

func main(){
	var areValues=MultiShape
	//fmt.Println(totalArea(&c, &r))
	fmt.Println(areValues)



}