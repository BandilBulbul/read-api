package main

import (
	"fmt"
)

//Swaping functions with pointers
var x int =10
var temp int=x

func pointersAddress(p int){
	var y *int=&x
	fmt.Println(x,y)
	*y=p//dereferencing
	fmt.Println(x)

}

func pointers(p *int){
	*p=temp
}

func main(){
	var p int=12
	pointersAddress(p)
	pointers(&p)
	fmt.Println(p)
}
