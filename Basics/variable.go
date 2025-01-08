
package main

import "fmt"

// func main() {
// 	// varriable()
// 	// defint()
// 	defintwotype()
// }

func varriable() {
	a := 20
	fmt.Println(a)
}

func defint() {
	var x int = 10
	fmt.Println(x)
}

func defintwotype() {
	xx := "hello boy" // can not assign different type of value
	var x = 11
	fmt.Println(x)
	fmt.Println(xx)
	x = x+1
	fmt.Println(x)

	const b = "Maruf" // can not change 
	
}



/*
	Variables in go
	int
	float32
	bool
	string
*/
