package main

import (
	"fmt"
)

func main(){
	//declare slice
	s :=make([]string,3)
	fmt.Println("empty slice",s) //make a slice of strings of length 3 (initially zero-valued).

	s[0]="abc"
	fmt.Println(s)
	s[1]="def"
	s[2]="0to9"
	fmt.Println(s)

	fmt.Println("set",s)
	fmt.Println("get",s[2])
	fmt.Println("len",len(s))

	//appending values into existing arrays
	s=append(s,"123","456")
	fmt.Println(s)
	s=append(s,"abc","def")
	fmt.Println("appened slice",s)



	//declare the slice with in-built make  function
	var slice1=make([]string,len(s))
	copy(slice1,s)
	fmt.Println(slice1,"  # ",s)
	//slicing the existing slice into different shape
	l:=s[1:]
	fmt.Println(l)
	l =s[3:]
	fmt.Println(l)

	t :=[]string{"g","h","i"}//declare slice
	fmt.Println(t)

	//2d demension slice
	twoD :=make([][]int,3)
	for i:=0;i<3;i++{
		innerLen :=i+1
		twoD[i]=make([]int,innerLen)
		for j:=0;j<innerLen;j++{
			twoD[i][j]=i*j
			/*Slices can be composed into multi-dimensional data structures.
			The length of the inner slices can vary,
			unlike with multi-dimensional arrays.*/
		}
	}
	fmt.Println("2D:",twoD)





}
