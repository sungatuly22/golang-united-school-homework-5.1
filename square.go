package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (receiver Square) End() Point {
	x := receiver.start.x - int(receiver.a)
	y := receiver.start.y + int(receiver.a)
	return Point{x, y}
}

func (receiver Square) Area() uint {
	return receiver.a * receiver.a
}

func (receiver Square) Perimeter() uint {
	return receiver.a * 4
}
