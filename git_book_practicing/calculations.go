package main

import (
	"fmt"
)

func add(i,j float64) float64{
	return i+j
}
func sub(i,j float64) float64{
	return i-j
}
func mul(i,j float64) float64{
	return i*j
}
func div(i,j float64) float64{
	return i/j
}

func main(){
	var input1 float64
	var input2 float64
	var str string
	//currentDate :=time.Now()

	fmt.Print("Enter any two number ")
	fmt.Scanln(&input1,&input2)
	fmt.Print("Select one operator like '+' ,'-' ,'*','/' ")
	fmt.Scanln(&str)
	if str=="+"{
		fmt.Println(add(input1,input2))
	} else if str=="-"{
		fmt.Println(sub(input1,input2))
	}else if str=="*"{
		fmt.Println(mul(input1,input2))
	}else if str=="/"{
		fmt.Println(div(input1,input2))
	}
    //fmt.Println(currentDate.Format(time.RubyDate))
}