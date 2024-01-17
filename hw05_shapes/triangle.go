package main

type Triangle struct {
	a, h float64
}

func (t *Triangle) A() float64 {
	return t.a
}

func (t *Triangle) SetA(a float64) {
	t.a = a
}

func (t *Triangle) H() float64 {
	return t.h
}

func (t *Triangle) SetH(h float64) {
	t.h = h
}

func (t Triangle) area() float64 {
	return 0.5 * t.A() * t.H()
}
