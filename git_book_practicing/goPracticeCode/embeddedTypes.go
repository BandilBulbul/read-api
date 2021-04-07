package main

import "fmt"

//Suppose we had a person struct

type Person struct{
	Name string
}
func (p *Person) Talk(){
	fmt.Println("hi,my name is ",p.Name)
}
//create a new android struct
type Android struct{
	Person
	Model string
}
func main(){
	var a=new(Android)
	a.Person.Name="chew"
	fmt.Println(a.Person.Name)
	a.Person.Talk()
}

