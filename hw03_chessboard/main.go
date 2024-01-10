package main

import (
	"fmt"
)

func getSquare(isBlack *bool) string {
	*isBlack = !*isBlack
	if *isBlack {
		return "#"
	}
	return " "
}

func main() {
	var size int
	isBlack := false
	fmt.Println("Введите размер доски:")
	fmt.Scanf("%d", &size)
	for row := 1; row <= size; row++ {
		for col := 1; col < size; col++ {
			fmt.Print("|" + getSquare(&isBlack))
		}
		fmt.Println("|" + getSquare(&isBlack) + "|")
		if size%2 == 0 {
			isBlack = !isBlack
		}
	}
}
