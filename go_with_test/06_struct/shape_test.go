package main

import (
	"testing"
)

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0, 10.0}
// 	got := Perimeter(rectangle)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// func TestArea(t *testing.T) {
// 	rectangle := Rectangle{12.0, 6.0}
// 	got := Area(rectangle)
// 	want := 72.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)

// 	}
// }

// func TestArea(t *testing.T) {
// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		got := rectangle.Area()
// 		want := 72.0

// 		if got != want {
// 			t.Errorf("got %.2f, want %.2f", got, want)
// 		}
// 	})

// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		want := 314.1592653589793
// 		if got != want {
// 			t.Errorf("got %.2f, want %.2f", got, want)
// 		}
// 	})
// }

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %.2f, want %.2f", got, want)
// 		}
// 	}

// 	t.Run("rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		checkArea(t, rectangle, 72.0)
// 	})

// 	t.Run("circle", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

func TestArea(t *testing.T) {
	areaTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangle", Rectangle{12, 7}, 84.0},
		{"circle", Circle{10}, 314.1592653589793},
		{"square", Square{4}, 16},
		{"triangle", Triangle{12, 6}, 3.0},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v  got %g want %g", tt.shape, got, tt.want)
		}
	}
}
