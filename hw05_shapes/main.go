package main

import (
	"errors"
	"fmt"
)

func calculateArea(s any) (float64, error) {
	if shape, ok := s.(Shape); ok {
		return shape.area(), nil
	}
	return 0, errors.New("переданный объект не реализует интерфейс Shape")
}

func main() {
	c := Circle{5}
	cArea, cErr := calculateArea(c)
	if cErr != nil {
		fmt.Println(cErr)
	} else {
		fmt.Printf("Круг: радиус %v\nПлощадь: %v\n\n", c.R(), cArea)
	}

	r := Rectangle{10, 5}
	rArea, rErr := calculateArea(r)
	if rErr != nil {
		fmt.Println(rErr)
	} else {
		fmt.Printf("Прямоугольник: ширина %v, высота %v \nПлощадь: %v\n\n", r.A(), r.B(), rArea)
	}

	t := Triangle{8, 6}
	tArea, tErr := calculateArea(t)
	if tErr != nil {
		fmt.Println(tErr)
	} else {
		fmt.Printf("Треугольник: основание %v, высота %v \nПлощадь: %v\n\n", t.A(), t.H(), tArea)
	}

	_, err := calculateArea("")
	if err != nil {
		fmt.Println("Ошибка: " + err.Error())
	}
}
