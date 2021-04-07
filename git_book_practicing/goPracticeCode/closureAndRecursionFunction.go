package main

import "fmt"

func makeEvenGenerator() func() int{
	i :=0
	return func() (ret int){
		ret =i
		i +=2
		return
	}
}
/*Finally a function is able to call itself. Here is one way
to compute the factorial of a number:*/
func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}




func main(){
	nextEven :=makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(factorial(10))
}