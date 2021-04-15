package main

import (
	"fmt"
	"strings"
)

func main(){
	str :="12+23-4"
	arrInterger :=[]string {}
	//arrOperator :=[]string {}

	arrInterger =strings.Split(str,"[+-*/]")
	fmt.Println(arrInterger)

}
