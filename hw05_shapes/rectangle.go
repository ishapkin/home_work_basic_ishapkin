package main

type Rectangle struct {
	a, b float64
}

func (r *Rectangle) A() float64 {
	return r.a
}

func (r *Rectangle) SetA(a float64) {
	r.a = a
}

func (r *Rectangle) B() float64 {
	return r.b
}

func (r *Rectangle) SetB(b float64) {
	r.b = b
}

func (r Rectangle) area() float64 {
	return r.A() * r.B()
}
