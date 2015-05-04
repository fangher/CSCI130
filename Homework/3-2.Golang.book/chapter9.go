package main

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func (c *Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	return ((r.y2 - r.y1) * (r.x2 - r.x1))
}
func (r *Rectangle) perimeter() float64 {
	return (2 * ((r.y2 - r.y1) + (r.x2 - r.x1)))
}
