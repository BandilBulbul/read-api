package main

import (
	"fmt"
	"time"
)
type T interface {

}
func main(){
	i :=2
	fmt.Print("Write",i,"as")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("Thrid")

	}

	switch time.Now().Weekday() {
	case time.Saturday ,time.Sunday: //You can use commas to separate multiple expressions in the same case statement.
		fmt.Println("It is the Weekend")
	default:
		fmt.Println("It's a weekday")
	
	}

	t :=time.Now()
	switch  { //switch without an expression is an alternate way to express if/else logic.
	case t.Hour()<12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")

	}


    whatAmI:= func(i T){
    	switch t:=i.(type){
		case bool:
			fmt.Println("boo")
		case string:
			fmt.Println("string")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("Dont know type %T\n",t)
		}
	}
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
    /*A type switch compares types instead of values.
    You can use this to discover the type of an interface value.
    the variable t will have the type corresponding to its clause.*/
}
