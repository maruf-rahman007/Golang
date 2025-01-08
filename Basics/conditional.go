package main

import "fmt"

func main() {
	casee:= 0
	age := 12
	gender := "female"
	if age > 21 && gender == "male" {
		fmt.Println("You can marry")
		casee = 1
	}else if age < 21 && gender == "female" && age >= 18{
		fmt.Println("You can marry")
		casee = 2
	}else {
		fmt.Println("You can not marry")
		casee = 3
	}

	switch casee {
	case 1:
		fmt.Println("Falls under first case!")
	case 2:
		fmt.Println("Falls under second case!")
	case 3:
		fmt.Println("Falls under third case!")
	}

}