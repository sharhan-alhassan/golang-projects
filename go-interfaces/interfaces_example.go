
package main
import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

// square struct
type square struct {
	length float64
}

// circle struct
type circle struct {
	radius float64
}

// square methods -- implement the shape interface
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods -- implement the shape interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s shape)  {
	fmt.Printf("Area of %T is: %0.2f \n ", s, s.area())
	fmt.Printf("Circumference of %T is: %0.2f \n ", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 23.4},
		square{length: 3.4},
		circle{radius: 4.4},
		circle{radius: 3.2},
	}

	for _, v := range shapes{
		printShapeInfo(v)
		fmt.Println("---")
	}
}