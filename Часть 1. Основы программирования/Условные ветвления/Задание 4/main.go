package main

import "fmt"

func main() {
	beerDelicious := true
	croutonsDelicious := false

	if beerDelicious && croutonsDelicious {
		fmt.Println("Мы пойдём гулять")
	} else {
		fmt.Println("Гулять мы не идём")
	}
}
