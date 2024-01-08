package main

import (
	"fmt"
	"math"
)

func main() {
	var size int
	fmt.Println("Введите размерность доски:")
	fmt.Scanf("%d\n", &size)

	loopSize := int(2 * math.Pow(float64(size), 2))
	for i := 1; i <= loopSize; i++ {
		if i%2 == 0 {
			if i%(loopSize/size) == 0 {
				fmt.Println("#")
			} else {
				fmt.Print("#")
			}
		} else {
			fmt.Print("  ")
		}
	}
}
