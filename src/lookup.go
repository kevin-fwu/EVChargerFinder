package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/kevin-fwu/EVChargerFinder/util"
)

type LocDist struct {
	Dist float64
	Loc  *Location
}

type Point struct {
	coords []float64
}

func (p *Point) GetCoordinates() []float64 {
	return p.coords
}

const (
	EARTH_RADIUS       float64 = 3963
	APPROX_DEGREE_DIST float64 = 69
)

var sortedTree *util.KdNode

// calc_dist calculates the distance between two points.
//
// Uses Haversine Formula (https://en.wikipedia.org/wiki/Haversine_formula)
// Returns the distance in miles.
func calcDist(lat1, lon1, lat2, lon2 float64) float64 {
	lat1Radians := lat1 * math.Pi / 180
	lon1Radians := lon1 * math.Pi / 180
	lat2Radians := lat2 * math.Pi / 180
	lon2Radians := lon2 * math.Pi / 180

	return 2 * EARTH_RADIUS *
		math.Asin(
			math.Sqrt(
				math.Pow(math.Sin((lat2Radians-lat1Radians)/2), 2)+
					math.Cos(lat1Radians)*
						math.Cos(lat2Radians)*
						math.Pow(math.Sin((lon2Radians-lon1Radians)/2), 2)))
}

// findClosest finds the closest EV Stations to the given point.
//
// For now, distance must be between 0 and 500 miles
func findClosest(point *Point, dist float64, limit int) []*LocDist {

	if dist < 0 {
		dist = 25
	} else if dist > 500 {
		dist = 500
	}

	if limit <= 0 {
		limit = 10
	}

	estimateDistance := dist / APPROX_DEGREE_DIST

	// The K-D Tree searches for nodes which are within the given degrees of latitude/longitude.
	//
	// The estimate distance uses an approximate 69 miles per point of latitude/longitude.
	estimateNodes := sortedTree.Search(point, estimateDistance, 2, 0)

	var locationsInRange []*LocDist

	// Given the estimated nodes, calculate the real distance and keep only the nodes which are within the given distance.
	for _, node := range estimateNodes {
		coords := node.GetCoordinates()
		if nodeDist := calcDist(point.coords[0], point.coords[1], coords[0], coords[1]); nodeDist <= dist {
			fmt.Printf("node dist: [%f, %f] to [%f, %f]: %f\n", point.coords[0], point.coords[1], coords[0], coords[1], nodeDist)
			locationsInRange = append(locationsInRange, &LocDist{Dist: nodeDist, Loc: node.(*Location)})
		}
	}

	sort.Slice(locationsInRange, func(i, j int) bool {
		return locationsInRange[i].Dist < locationsInRange[j].Dist
	})

	if len(locationsInRange) > limit {
		locationsInRange = locationsInRange[:limit]
	}
	return locationsInRange
}
