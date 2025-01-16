package main

import "fmt"

func add1(num1 int , num2 int)  {
	var sum int = num1 + num2
	fmt.Println(sum)
}


func main () {
	// a:= 10
	// b:= 20
	// var sum int = a + b

	// fmt.Println(sum)
	// add1(a, b)

	// sum:= add2(11,12)
	// fmt.Println(sum)
	// sum,mul:=twoops(a,b)
	// fmt.Println(sum)
	// fmt.Println(mul)

	fmt.Println("CASIO ")

	fmt.Print("Enter your name : ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Hello",name ,"Welcome to the application")

}

func add2(num1 int , num2 int) int {
	return num1 + num2
}

func twoops (num1 int, num2 int) (int, int) {
	return num1+num2 , num1*num2
}