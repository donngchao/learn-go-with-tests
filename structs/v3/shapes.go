package main

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter returns the perimeter of the rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Area returns the area of the rectangle
func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}

// Circle has radius
type Circle struct{
	radius float64
}

//perimeterCircle returns the perimeter of the circle
func perimeterCircle(circle Circle) float64 {
	return 3.14 * 2 * (circle.radius)
}

//areaCircle returns the area of the circle
func areaCircle(circle Circle) float64{
	return 3.14 * circle.radius * circle.radius
}

