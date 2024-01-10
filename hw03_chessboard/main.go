package main

import (
	"fmt"
)

// Функция типа клетки
func getSquare(isBlack *bool) string {
	*isBlack = !*isBlack
	if *isBlack {
		return "#"
	}
	return " "
}

func main() {
	var size int
	isBlack := true
	fmt.Println("Введите размерность доски:")
	fmt.Scanf("%d", &size)
	// Обход строки
	for row := 1; row <= size; row++ {
		// Обход колонки
		for col := 1; col < size; col++ {
			fmt.Print("|" + getSquare(&isBlack))
			if (col + 1) == size {
				fmt.Println("|" + getSquare(&isBlack) + "|")
			}
		}
		// Если количетсво четное - то дополнительно переключаем тип клетки
		if size%2 == 0 {
			isBlack = !isBlack
		}
	}
}
