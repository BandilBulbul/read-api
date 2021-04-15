package main

import "fmt"

//Maps are Go’s built-in associative data type
func main(){
	//To create an empty map,
	//use the builtin make: make(map[key-type]val-type).

	mapValues :=make(map[string]int)
	mapValues["k"]=9
	fmt.Println(mapValues)
	fmt.Println(len(mapValues))

	maps:=map[string]map[string]string{
		"foo":{"boo":"me"},
		"goo":{"hoo":"you"},
	}
	fmt.Println(maps)










}