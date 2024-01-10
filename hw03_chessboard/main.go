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
	fmt.Println("Введите размерность доски:")
	fmt.Scanf("%d", &size)
	// Обход строк
	for row := 1; row <= size; row++ {
		// Обход колонок
		for col := 1; col < size; col++ {
			fmt.Print("|" + getSquare(&isBlack))
		}
		fmt.Println("|" + getSquare(&isBlack) + "|")
		// Если количетсво четное - то дополнительно переключаем тип клетки
		if size%2 == 0 {
			isBlack = !isBlack
		}
	}
}
