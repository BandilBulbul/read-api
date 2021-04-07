package main

import "fmt"

func main(){
	fmt.Println("#########################################")
	arr :=[]float64{12.0,7.0,9.0}
	a:=average(arr)
	fmt.Println(a)
	fmt.Print("Enter  a number:")
	var input int
	fmt.Scanf("%v", &input) //another function  from fmt package to read the user input(scanf)
	output :=input*2
	fmt.Println(output)
	control(output)
	printDivisibleByThree()
	printFizzBuzz()
	arrayPrograming() //before index 4 and after that ..will be assign free  space
	sliceProgramming()
	mapsFunction()
    smallestNumberInArray()
	x,y :=multipleValues()
	fmt.Println(x,y)
	fmt.Println(variadicFunction(1,2,3,4,5,6,8))
}
func control(i int){
	if i % 2 == 0 {
		// divisible by 2
	} else if i % 3 == 0 {
		// divisible by 3
	} else if i % 4 == 0 {
		// divisible by 4
	}

}

func switchStatements(i int){
	switch i {
	case 0: fmt.Println("Zero")
	case 1: fmt.Println("One")
	case 2: fmt.Println("Two")
	case 3: fmt.Println("Three")
	case 4: fmt.Println("Four")
	case 5: fmt.Println("Five")
	default: fmt.Println("Unknown Number")
	}

}

func printDivisibleByThree(){
	for i:=1;i<=100;i++{
		if i%3==0 || i%6==0 || i%9==0{
			fmt.Println(i)
		}
	}
}

func printFizzBuzz(){
	i:=1
	for i<=100{
		if i%3==0{
			fmt.Println(i,"Fizz")
		} else if i%5==0{
			fmt.Println(i,"Buzz")
		}else if i%15==0{
			fmt.Println(i,"FizzBuzz")
		} else{
			fmt.Println(i,"neither divisible by 3 nor 5")

		}
		i++


	}
}

func arrayPrograming(){
	var studentName [6]string
	studentName[4]="Ramu"
	fmt.Println(studentName)
	for i:=0;i<len(studentName);i++{
		studentName[i] +="raja"
	}
	fmt.Println(studentName)

	var variablesNumbers [10]float64
	var  total float64=0
	for i:=0;i<len(variablesNumbers);i++{
		variablesNumbers[i]+=20
		total +=variablesNumbers[i]
	}
	fmt.Println(total/float64(len(variablesNumbers)))

	//Range
	for index,_ :=range variablesNumbers{
		variablesNumbers[index] =3.4
		total+=variablesNumbers[index]

	}
	fmt.Println(total)

}
//The only difference between this and an array is the
//missing length between the brackets.
/*ABOUT SLICE: A slice is a segment of an array. Like arrays slices are
indexable and have a length. Unlike arrays this length
is allowed to change.*/

func sliceProgramming(){
	//var x []float64
	/*If you want to create a slice you should use the built-in
	function make
	x := make([]float64, 5)*/
	//

	//SLICE FUNCTIONS ---APPEND, COPY
	slice1:=[]int{1,2,3}
	slice2:=append(slice1,4,5)
	fmt.Println(slice1,slice2)
	fmt.Println(len(slice1),cap(slice1))
	fmt.Println(len(slice2),cap(slice2))
	slice4 :=make([]int,5,10)
	copy(slice4,slice2)
	fmt.Println(slice4,len(slice4),cap(slice4))
	slice4=append(slice2,1,2,3,4,5,5)
	fmt.Println(slice4,len(slice4),cap(slice4))
	fmt.Println(slice1[2])

}
func mapsFunction() {
	//var x map[string]int
	//x["key"]=10
	//
	//fmt.Println(x)//	panic: runtime error: assignment to entry in nil
	y := make(map[string]int)
	y["key"] = 10
	fmt.Println(y["key"])
	//built-in function in the map  ---delete function
	delete(y, "key")
	fmt.Println(y)

	//Create a map elements
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"
	fmt.Println(elements)

	//check the Condition
	name, ok := elements["C"]
	fmt.Println(name, ok)
	name, ok = elements["cn"]
	fmt.Println(name, ok)
	if name, ok = elements["O"]; ok {
		fmt.Println(name, ok)
	}

	elementsValues := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elementsValues)

	elementss := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}
	fmt.Println(elementss)
}

func smallestNumberInArray(){
	arrayList :=[11]int {12,34,3,54,67,5,234,64,12,56,32}
	 small := arrayList[0]
	for i:=1;i<len(arrayList);i++{
		if arrayList[i] <= small{
			small =arrayList[i]
		}

	}
	fmt.Println(small)
	fmt.Println(arrayList[2:9],arrayList)

}

func average(xs []float64) float64{

	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total/float64(len(xs))

}

func multipleValues() (int,int){
	return 1,3
}

func variadicFunction( args ...int) int{
	sum :=0
	for index,_:=range args{
		sum +=args[index]
	}
	//CLOSURE FUNCTION EXAMPLE: CREATE FUNCTIONS INSIDE OF FUNCTIONS
	add :=func(x,y int) int{
		return x+y
	}
	fmt.Println(add(1,2))
	return sum

}

