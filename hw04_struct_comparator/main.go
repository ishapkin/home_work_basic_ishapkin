package main

import (
	"fmt"
	"reflect"
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

// Getters and setters for Book

func (b *Book) SetID(id int) {
	b.id = id
}

func (b *Book) ID() int {
	return b.id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) Year() int {
	return b.year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) Size() int {
	return b.size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) Rate() float64 {
	return b.rate
}

const (
	Year = "year"
	Size = "size"
	Rate = "rate"
)

// Compare struct

type CompareBook struct {
	compareMode string
}

// Compare constructor

func newCompareBook(compareMode string) *CompareBook {
	if compareMode == Year || compareMode == Size || compareMode == Rate {
		return &CompareBook{compareMode}
	}
	return &CompareBook{Year}
}

func (c *CompareBook) isCompare(firstBook Book, secondBook Book) bool {
	mode := c.compareMode
	firstVal := reflect.ValueOf(firstBook).FieldByName(mode)
	secondVal := reflect.ValueOf(secondBook).FieldByName(mode)
	if mode == Rate {
		return firstVal.Float() > secondVal.Float()
	}
	return firstVal.Int() > secondVal.Int()
}

func printResult(fBook Book, sBook Book, compareBook *CompareBook) {
	adv := ""
	switch compareBook.compareMode {
	case Rate:
		if compareBook.isCompare(fBook, sBook) {
			adv = "выше рейтингом"
		} else {
			adv = "ниже рейтингом"
		}
	case Size:
		if compareBook.isCompare(fBook, sBook) {
			adv = "больше"
		} else {
			adv = "меньше"
		}
	default:
		if compareBook.isCompare(fBook, sBook) {
			adv = "старше"
		} else {
			adv = "младше"
		}
	}
	fmt.Println(fBook.Title() + " " + adv + " чем " + sBook.Title())
}

func main() {
	fBook := Book{1, "«Война и Мир»", "Лев Толстой", 1873, 1300, 4.5}
	sBook := Book{2, "«Идиот»", "Федор Достоевский", 1869, 350, 4.9}
	printResult(fBook, sBook, newCompareBook("year"))
	printResult(fBook, sBook, newCompareBook("rate"))
	printResult(fBook, sBook, newCompareBook("size"))
}
