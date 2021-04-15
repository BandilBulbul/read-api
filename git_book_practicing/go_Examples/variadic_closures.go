package main

import "fmt"

func main(){
	add(1,2,4,5,6)
	arr :=[]int {2,4,5,7,89} //should be slice
	add(arr...)
/*
   If you already have multiple args in a slice, apply
them to a variadic function using func(slice...) like this.*/


//closures function:
	nextInt :=intSeq()
	fmt.Println(nextInt())
}

func add(i ...int){
	total:=0
	for _,val :=range i{
		total  +=val
	}
	fmt.Println(total)
}
/*Go supports anonymous functions, which can form closures.
Anonymous functions are useful when you want to define a
function inline without having to name it.*/
func intSeq() func() int{
	i:=0
	return func() int{
		i++
		return i
	}
}
