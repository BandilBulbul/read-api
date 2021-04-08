package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main(){
	//define flags
	maxp :=flag.Int("max",6,"the max value")
	//parse
	flag.Parse()
	//Generate a number between  and max
	fmt.Println(rand.Intn(*maxp))
}
