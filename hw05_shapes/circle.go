package main

import "math"

type Circle struct {
	r float64
}

func (c *Circle) R() float64 {
	return c.r
}

func (c *Circle) SetR(r float64) {
	c.r = r
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}
