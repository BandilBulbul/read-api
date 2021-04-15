package main

import (
	"fmt"
	"sort"
)

type Personn struct{
	Name string
	Age int
}
type ByName []Personn

type ByAge []Personn

func (this ByName) Len() int{
	return len(this)
}
func (this ByName) Less(i,j int)bool{
	return this[i].Name<this[j].Name
}
func (this ByName) Swap(i,j int){
	this[i],this[j]=this[j],this[i]
}

func(this ByAge) Len()int{
	return len(this)
}
func (this ByAge) Less(i,j int) bool{
	return this[i].Age<this[j].Age
}

func(this ByAge) Swap(i,j int){
	this[i],this[j]=this[j],this[i]
}
func main(){
	kids :=[]Personn{
		{"Jill",9},
		{"Jack",10},
		{"siri",12},
		{"ram",5},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}