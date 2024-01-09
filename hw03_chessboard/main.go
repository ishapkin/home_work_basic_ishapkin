package main

import (
	"fmt"
	"math"
)

func getSquare(blackCase bool) string {
	if blackCase {
		return "#"
	}
	return " "
}

func main() {
	var size int
	blackMode := false
	fmt.Println("Введите размерность доски:")
	fmt.Scanf("%d", &size)
	loopSize := int(math.Pow(float64(size), 2))
	for i := 1; i <= loopSize; i++ {
		blackMode = !blackMode
		if i%2 == 0 {
			if i%(loopSize/size) == 0 {
				fmt.Println("|" + getSquare(blackMode) + "|")
				blackMode = !blackMode
			} else {
				fmt.Print("|" + getSquare(blackMode))
			}
		} else {
			fmt.Print("|" + getSquare(blackMode))
		}
	}
}
