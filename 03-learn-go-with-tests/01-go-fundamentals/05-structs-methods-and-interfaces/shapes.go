package shapes

import "math"

func Perimeter(width, height float64) float64 {
	return (width + height) * 2
}

func Area(width, height float64) float64 {
	return width * height
}

type Rectangle struct {
	Width  float64
	Height float64
}

// func Perimeter(r Rectangle) float64 {
// 	return 0.0
// }

// func Area(r Rectangle) float64 {
// 	return 0.0
// }

// NOTE
// - this won't work, because go doesn't allow function overloading, where
//   multiple function share the same name, but have different parameter lists
// - so we define methods instead, which are functions with a specified receiver
//   - the receiver in this case is an instance of a struct

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Shape interface {
	Area() float64
}

// NOTE
// - an interface allows us to group structs by common methods that they have
// - in go interfaces that a struct matches is implicit - if all the methods
//   identified in the the interface have been defined for a struct, then
//   that struct is considered to of that interface even without any
//   declaration of this.

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height / 2
}
