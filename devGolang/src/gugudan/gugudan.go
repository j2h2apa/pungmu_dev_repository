package main

import (
	"fmt"
)

func main() {
	// practice 1
	for dan := 1; dan < 10; dan++ {
		fmt.Printf(">>>>>>>>>>>>> %d단 <<<<<<<<<<<<<<\n", dan)

		if dan == 5 {
			fmt.Printf("%d 단은 continue\n", dan)
			fmt.Println()
			continue
		}

		for j := 1; j < 10; j++ {
			fmt.Printf("%d * %d = %d\n", dan, j, dan*j)
		}

		fmt.Println()
	}

	// practice 2
	for lines := 1; lines < 6; lines++ {
		for stars := 1; stars <= lines; stars++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	// practice 3
	for i := 0; i < 4; i++ {
		for j := 0; j < 3-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// practice 4
	for i := 0; i < 3; i++ {
		for j := 0; j < 2-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 3-i*2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
