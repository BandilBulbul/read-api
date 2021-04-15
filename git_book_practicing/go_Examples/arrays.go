package main

import "fmt"

func main(){
	var a =[5]int {1,2,4,5,6} //declare and assign value at index
	fmt.Print(a)
	a[3]=10
	fmt.Println(a) //show updated array

	var array [8]int
	fmt.Println(array) //empty array shows with zero default values
	//2d demension
	var c [4][4]int//declare only
	for i:=0;i<4;i++{
		for j:=0;j<4;j++{
			c[i][j]=i*j
			}
	}
	fmt.Println(c)//2d arrays
}