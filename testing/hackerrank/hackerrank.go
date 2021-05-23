package main

import "fmt"

func GetOddIndex(data []string) []string {
	var result []string
	for _, value := range data {
		var OddIndexData string = ""
		for i, v1 := range value {
			if i%2 != 0 {
				OddIndexData += string(v1)
			}
		}
		result = append(result, OddIndexData)
	}
	return result
}

func GetOneLine() int {
	//var input int
	//fmt.Scan(&input)
	input := 2
	return input
}

func main() {
	v := [2]string{"Hacker", "Rank"}
	var res [2][2]string = GetIndex(v)
	for _, val := range res {
		fmt.Println(val)
	}
	//fmt.Print(res)
}

func GetEvenIndex(data []string) []string {
	var result []string
	for _, value := range data {
		var EvenIndexData string = ""
		for i, v1 := range value {
			if i%2 == 0 {
				EvenIndexData += string(v1)
			}
		}
		result = append(result, EvenIndexData)
	}
	return result
}
func GetIndex(data [2]string) [2][2]string {
	var result [2][2]string
	k, h := 0, 0
	for _, value := range data {
		var EvenIndexData string = ""
		var OddIndexData string = ""
		for i, v1 := range value {
			if i%2 == 0 {
				EvenIndexData += string(v1)
			} else {
				OddIndexData += string(v1)

			}
		}
		//pair := EvenIndexData + " " + OddIndexData
		//result = append(result, EvenIndexData, OddIndexData)
		//result = append(result, pair)
		result[k][0] = EvenIndexData
		result[h][1] = OddIndexData
		k += 1
		h += 1
	}
	//fmt.Println(result)
	return result
}
