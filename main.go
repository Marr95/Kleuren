package main

import (
	"fmt"
)

func main() {
	var input string

	fmt.Println("Voer uw kleur in: ")
	// _, is een soort van index deze houd die vrij
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println("Error: verkeerde data.")
	} else {
		switch input {
		case "blauw":
			fmt.Println("Blauw.")
		case "rood":
			fmt.Println("Rood.")
		case "geel":
			fmt.Println("Geel.")
		case "groen":
			fmt.Println("Groen.")
		case "oranje":
			fmt.Println("Oranje.")
		case "wit":
			fmt.Println("Wit.")
		case "zwart":
			fmt.Println("Zwart.")
		default:
			fmt.Println("Kleur niet herkend, voer a.u.b. een andere kleur in.")
		}
	}
}
