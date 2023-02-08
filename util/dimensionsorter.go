package util

import (
	"math"
	"sort"
)

type Coordinates interface {
	GetCoordinates() []float64
}

type dimensionSorter struct {
	dimension   int
	coordinates []Coordinates
}

func (ds *dimensionSorter) Sort(coordinates []Coordinates) {
	ds.coordinates = make([]Coordinates, len(coordinates))
	copy(ds.coordinates, coordinates)
	sort.Sort(ds)
}

func (ds *dimensionSorter) Swap(i, j int) {
	ds.coordinates[i], ds.coordinates[j] = ds.coordinates[j], ds.coordinates[i]
}

func (ds *dimensionSorter) Len() int {
	return len(ds.coordinates)
}

func (ds *dimensionSorter) Less(i, j int) bool {
	cmp := cmpAllDimensions(ds.coordinates[i], ds.coordinates[j], ds.dimension)
	return cmp < 0
}

func cmpAllDimensions(lhs, rhs Coordinates, idxStart int) int {
	coordL := lhs.GetCoordinates()
	coordR := rhs.GetCoordinates()

	numIdx := len(coordL)

	for iter := 0; iter < numIdx; iter++ {
		idxCheck := idxStart + iter
		if idxCheck >= numIdx {
			idxCheck -= numIdx
		}
		if cmp := cmpFloat64(coordL[idxCheck], coordR[idxCheck]); cmp != 0 {
			return cmp
		}
	}
	return 0
}

func cmpFloat64(lhs, rhs float64) int {
	// NaN is considered less than not-NaN

	if lhs < rhs {
		return -1
	} else if lhs > rhs {
		return 1
	}
	lNaN := math.IsNaN(lhs)
	rNaN := math.IsNaN(rhs)

	if lNaN == rNaN {
		return 0
	} else if lNaN {
		return -1
	} else {
		return 1
	}
}
