package main

//defer which schedules a function call to be run after the function completes.
import (
	"fmt"
)
func first(){
	fmt.Println("first")
}
func second(){
	fmt.Println("second")
}
func main(){
	//Panic nad recover topic
	/*panic("PANIC")
	str := recover()
	fmt.Println(str)*/
    panicRecover()


	defer second()
	first()
	fmt.Println(evenOdd(3))
	fmt.Println(variadicFunctionn(1,2,43,56,76,87,34,766,65,34,12,1))
}

//This program prints followed by 2nd 1st,Basically defer moves the call to second to the end of the function:


func evenOdd(num int) (int,bool){
	if num%2==0{
		return 1,true
	}else{
		return 0,false
	}

}

func variadicFunctionn(args ...int) int{
	greater:=args[0]
	for i:=1;i<len(args);i++{
		if args[i]>greater{
			greater=args[i]
		}
	}
	return greater
}

func panicRecover(){
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}