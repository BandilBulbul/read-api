package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	for i:=0;i<10;i++{
		go f1(i)
	}
	var input string
	fmt.Scanln(&input)

}
func f(){
	fmt.Println("new goroutine")
}
func f1(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}
