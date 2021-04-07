package main

type Shape interface{
	area() float64
	perimeter() float64
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

func (m *MultiShape) Perimeter() float64{
	var perimeter float64
	for _,s:=range m.shapes{
		perimeter +=m.Perimeter()
	}
	return perimeter
}

func main(){
	var rect=new(MultiShape)


}