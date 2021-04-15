package main

import "fmt"

func main(){
	for i, c := range "hello" {
		fmt.Println(i, c)
	}
}
