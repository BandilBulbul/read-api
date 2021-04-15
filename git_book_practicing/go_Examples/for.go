package main

import "fmt"

func main(){
	for n:=0;n<100;n++{
		if n%2==0{
			continue
		}
		fmt.Print(n)
		if n==51{
			fmt.Println("Loop break!!")
			break
		}
	}
}
