package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 15.0)
	want := 50.0

	if got != want {
		t.Errorf("want %.2f got %.2f", want, got)
	}
}

// NOTE
// - `%.2f` is the format string used to denote 2 decimal places
//   for floating point numbers

func TestArea(t *testing.T) {
	got := Area(10.0, 15.0)
	want := 150.0

	if got != want {
		t.Errorf("want %.2f got %.2f", want, got)
	}
}

func TestShapeStructMethods(t *testing.T) {

	t.Run("rectangle perimeter", func(t *testing.T) {
		r := Rectangle{10.0, 15.0}
		got := r.Perimeter()
		want := 50.0

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("rectangle area", func(t *testing.T) {
		r := Rectangle{10.0, 15.0}
		got := r.Area()
		want := 150.0

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("circle area", func(t *testing.T) {
		c := Circle{100.0}
		got := c.Area()
		want := 31415.926535897932

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

	t.Run("circle perimeter", func(t *testing.T) {
		c := Circle{100.0}
		got := c.Perimeter()
		want := 628.3185307179587

		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	})

}

func TestShapeInterfaceMethods(t *testing.T) {

	checkAreaHelper := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if want != got {
			t.Errorf("want %.2f got %.2f", want, got)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		r := Rectangle{10.0, 15.0}
		checkAreaHelper(t, r, 150.0)
	})

	t.Run("circle area", func(t *testing.T) {
		c := Circle{10.0}
		checkAreaHelper(t, c, 314.1592653589793)
	})

}

func TestShapeInterfaceMethodsTableDriven(t *testing.T) {
	testTable := []struct {
		testName string
		shape    Shape
		want     float64
	}{
		{"Rectangle", Rectangle{10.0, 15.0}, 150.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{10.0, 5.0}, 25.0},
	}

	for _, testItem := range testTable {
		t.Run(testItem.testName, func(t *testing.T) {
			got := testItem.shape.Area()
			if testItem.want != got {
				t.Errorf("want %.2f got %.2f for %#v", testItem.want, got, testItem.shape)
			}
		})

	}
}

// NOTE
// - create an array of expected test inputs and outputs, then iterate over them
// - this avoid the need to write tests repetitively
// - the `[]struct { ... }` here is an anonymous struct
// - to print a full struct, use the format string `%#v`, e.g.
//   `want 25.00 got 0.00 for shapes.Triangle{Base:10, Height:5}`
// - using structs, defining methods for structs, and grouping structs together
//   by means of interfaces are all part of designing your own types,
//   which is critical to being able to write in a statically typed language
