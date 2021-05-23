package main

import "fmt"

const repeatValue = 5

func Repeat(word string) string {
	var result string
	for i := 0; i < repeatValue; i++ {
		result += word
	}
	return result
}

func main() {
	v := Repeat("b")
	fmt.Print(v)
	t := Add(1, 2)
	fmt.Print(t)

}
