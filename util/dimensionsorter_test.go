package util

import (
	"fmt"
	"testing"
)

type testCoordinates struct {
	coords []float64
}

func (t *testCoordinates) GetCoordinates() []float64 {
	return t.coords
}

func TestDimensionSorter(t *testing.T) {

	ds := &dimensionSorter{}

	// Test 1D sorting
	fmt.Println("Test 1D sorting.")
	ds.Sort([]Coordinates{
		&testCoordinates{coords: []float64{5}},
		&testCoordinates{coords: []float64{3}},
		&testCoordinates{coords: []float64{-1}},
		&testCoordinates{coords: []float64{1}},
		&testCoordinates{coords: []float64{0}},
	})

	verifySort(t, ds.coordinates, []Coordinates{
		&testCoordinates{coords: []float64{-1}},
		&testCoordinates{coords: []float64{0}},
		&testCoordinates{coords: []float64{1}},
		&testCoordinates{coords: []float64{3}},
		&testCoordinates{coords: []float64{5}},
	})

	// Test 2D sorting
	input2d := []Coordinates{
		&testCoordinates{coords: []float64{5, 1}},
		&testCoordinates{coords: []float64{3, 2}},
		&testCoordinates{coords: []float64{-1, 3}},
		&testCoordinates{coords: []float64{1, 4}},
		&testCoordinates{coords: []float64{0, 5}},
	}
	fmt.Println("Test 2D sorting.")
	ds.Sort(input2d)

	verifySort(t, ds.coordinates, []Coordinates{
		&testCoordinates{coords: []float64{-1, 3}},
		&testCoordinates{coords: []float64{0, 5}},
		&testCoordinates{coords: []float64{1, 4}},
		&testCoordinates{coords: []float64{3, 2}},
		&testCoordinates{coords: []float64{5, 1}},
	})

	// Test 2D sorting by second dimension
	fmt.Println("Test 2D sorting by the second dimension.")
	ds.dimension = 1
	ds.Sort(input2d)

	verifySort(t, ds.coordinates, []Coordinates{
		&testCoordinates{coords: []float64{5, 1}},
		&testCoordinates{coords: []float64{3, 2}},
		&testCoordinates{coords: []float64{-1, 3}},
		&testCoordinates{coords: []float64{1, 4}},
		&testCoordinates{coords: []float64{0, 5}},
	})

	// Test 3D sorting
	input3d := []Coordinates{
		&testCoordinates{coords: []float64{5, 1, 3}},
		&testCoordinates{coords: []float64{1, 2, 8}},
		&testCoordinates{coords: []float64{-1, 2, 8}},
		&testCoordinates{coords: []float64{1, 4, 4}},
		&testCoordinates{coords: []float64{0, 5, 2}},
	}
	fmt.Println("Test 3D sorting.")
	ds.dimension = 0
	ds.Sort(input3d)

	verifySort(t, ds.coordinates, []Coordinates{
		&testCoordinates{coords: []float64{-1, 2, 8}},
		&testCoordinates{coords: []float64{0, 5, 2}},
		&testCoordinates{coords: []float64{1, 2, 8}},
		&testCoordinates{coords: []float64{1, 4, 4}},
		&testCoordinates{coords: []float64{5, 1, 3}},
	})

	// Test 3D sorting by second dimension. Matches should roll over.
	fmt.Println("Test 3D sorting by the second dimension.")
	ds.dimension = 1
	ds.Sort(input3d)

	verifySort(t, ds.coordinates, []Coordinates{
		&testCoordinates{coords: []float64{5, 1, 3}},
		&testCoordinates{coords: []float64{-1, 2, 8}},
		&testCoordinates{coords: []float64{1, 2, 8}},
		&testCoordinates{coords: []float64{1, 4, 4}},
		&testCoordinates{coords: []float64{0, 5, 2}},
	})

	// Test 3D sorting by third dimension. Matches should roll over.
	fmt.Println("Test 3D sorting by the third dimension.")
	ds.dimension = 2
	ds.Sort(input3d)

	verifySort(t, ds.coordinates, []Coordinates{
		&testCoordinates{coords: []float64{0, 5, 2}},
		&testCoordinates{coords: []float64{5, 1, 3}},
		&testCoordinates{coords: []float64{1, 4, 4}},
		&testCoordinates{coords: []float64{-1, 2, 8}},
		&testCoordinates{coords: []float64{1, 2, 8}},
	})
}

func verifySort(t *testing.T, actual, expected []Coordinates) {
	if len(actual) != len(expected) {
		t.Fatalf("Failed sort, actual slice length (%d) does not match expected (%d)\n", len(actual), len(expected))
	}
	for k, v := range expected {
		coordExpected := v.GetCoordinates()
		coordActual := actual[k].GetCoordinates()

		for i, val := range coordExpected {
			if val != coordActual[i] {
				t.Fatalf("Failed sort, index %d has actual coord %+v, expected %+v.\n", k, coordActual, coordExpected)
			}
		}
	}
}
