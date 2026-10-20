package main

import "fmt"

func main() {
	birthYear := 2007
	firstName := "Sabyr"
	weight := 73.3
	isAdult := true

	fmt.Println()
	fmt.Println("birthYear:", birthYear)
	fmt.Println("firstName:", firstName)
	fmt.Println("weight:", weight)
	fmt.Println("isAdult:", isAdult)
	fmt.Println()

	birthYear -= 3
	firstName += "zhan"
	weight++

	fmt.Println("updated birthYear:", birthYear)
	fmt.Println("updated firstName:", firstName)
	fmt.Println("updated weight:", weight)
	fmt.Println("updated isAdult:", isAdult)
	fmt.Println()
}
