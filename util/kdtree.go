package util

import "math"

type KdNode struct {
	value   Coordinates
	ltChild *KdNode
	gtChild *KdNode
}

func newKdNode(coordinates Coordinates) *KdNode {
	return &KdNode{value: coordinates}
}

// Create a K-D Tree from input coordinates.
func CreateKdTree(input []Coordinates, numDimensions int) *KdNode {
	sorted := make([][]Coordinates, numDimensions)

	for i := 0; i < numDimensions; i++ {
		ds := &dimensionSorter{
			dimension: i,
		}
		ds.Sort(input)
		sorted[i] = ds.coordinates
	}

	node := buildKdTree(sorted, 0, len(input)-1, 0)
	node.verifyTree(len(sorted), 0)
	return node
}

// Search returns a list of nodes which are within dist in all dimensions.
func (node *KdNode) Search(point Coordinates, dist float64, numDimensions, depth int) []Coordinates {
	curDim := depth % numDimensions
	var list []Coordinates

	coordPoints := point.GetCoordinates()
	coordNode := node.value.GetCoordinates()
	isWithinRange := true
	for i, c := range coordNode {
		if cmpFloat64(math.Abs(c-coordPoints[i]), dist) > 0 {
			isWithinRange = false
			break
		}
	}
	if isWithinRange {
		list = append(list, node.value)
	}

	if node.ltChild != nil && cmpFloat64(coordPoints[curDim]-dist, coordNode[curDim]) <= 0 {
		ltList := node.ltChild.Search(point, dist, numDimensions, depth+1)
		list = append(list, ltList...)
	}

	if node.gtChild != nil && cmpFloat64(coordPoints[curDim]+dist, coordNode[curDim]) >= 0 {
		gtList := node.gtChild.Search(point, dist, numDimensions, depth+1)
		list = append(list, gtList...)
	}

	return list
}

// buildKdTree compiles a K-D Tree for fast multi-dimensional distance calculation.
//
// It requires as input:
// sortedCoordinates - A 2-D slice of Coordinates. Assumes that the length of the three slices are equal. Detailed below.
// idxStart - The lower bound for finding the median.
// idxEnd - The upper bound for finding the median. NOTE: This index must contain the last node! In my experience,
//     the end usually points to the index _after_ the last node. This is not the case for this algorithm.
// depth - The current depth of the tree. The function will use this to determine the current dimension.
//
// Regarding the input sortedCoordinates, the inner slice is a sorted coordinate slice that is sorted on
// the dimension corresponding to the outer slice's index. Ties should compare the next dimension up until there is
// a difference. Duplicates must not exist. The slice will be permutated during the tree compilation.
//
// Example:
// Input: [{0, 1, 2}, {2, 1, 0}, {1, 2, 0}]
// sortedCoordinates: [
//     [{0, 1, 2}, {1, 2, 0}, {2, 1, 0}],
//     [{2, 1, 0}, {0, 1, 2}, {1, 2, 0}], // Here, {2, 1, 0} is before {0, 1, 2} because the third dimension is lower.
//     [{1, 2, 0}, {2, 1, 0}, {0, 1, 2}], // Here, {1, 2, 0} is before {2, 1, 0} because the first dimension is lower.
// ]
//
func buildKdTree(sortedCoordinates [][]Coordinates, idxStart, idxEnd, depth int) *KdNode {
	var node *KdNode
	dimCur := depth % len(sortedCoordinates)
	if idxEnd == idxStart {
		// Only one coordinate left
		node = newKdNode(sortedCoordinates[0][idxStart])
	} else if idxEnd == idxStart+1 {
		// Two coordinates left
		node = newKdNode(sortedCoordinates[0][idxStart])
		node.gtChild = newKdNode(sortedCoordinates[0][idxStart+1])
	} else if idxEnd == idxStart+2 {
		// Three coordinates left
		node = newKdNode(sortedCoordinates[0][idxStart+1])
		node.ltChild = newKdNode(sortedCoordinates[0][idxStart])
		node.gtChild = newKdNode(sortedCoordinates[0][idxStart+2])
	} else if idxEnd > idxStart+2 {

		idxMedian := idxStart + ((idxEnd - idxStart) / 2)
		node = newKdNode(sortedCoordinates[0][idxMedian])

		lower, upper := rotateCoordinates(
			sortedCoordinates, node.value,
			idxStart, idxMedian, idxEnd,
			dimCur)

		node.ltChild = buildKdTree(sortedCoordinates, idxStart, lower, depth+1)
		node.gtChild = buildKdTree(sortedCoordinates, idxMedian+1, upper, depth+1)

	} else if idxEnd < idxStart {
		panic("End is less than start?")
	}
	return node
}

// rotateCoordinates permutates sortedCoordinates for the next round.
//
// The 0th index sortedCoordinates should be sorted on the dimCur dimension.
//
// Rotation Logic:
// The coordinates at the 0-th index is stored in a temporary slice.
// Then for each other index K, the coordinates are placed in the slice
// at K-1. The coordinates are compared to the median coordinate using
// cmpAllDimensions(). Those coordinates which are less than the median
// are stored in the 'lower half' of the slice. Those coordinates which
// are greater than the median are stored in the 'upper half' of the slice.
// There should only be 1 coordinate which equals self (which is the median itself).
// As we are comparing the current dimension against its median, the two halves
// are equal in length (+1 for odd). And since we are iterating sorted slices,
// the resultant slice is still properly sorted for the corresponding dimension.
//
func rotateCoordinates(sortedCoordinates [][]Coordinates, median Coordinates,
	idxStart, idxMedian, idxEnd, dimCur int) (int, int) {

	temp := make([]Coordinates, idxEnd-idxStart+1)

	for i := 0; i <= idxEnd-idxStart; i++ {
		temp[i] = sortedCoordinates[0][i+idxStart]
	}

	var lower, upper, lowerSave, upperSave int
	for i := 1; i < len(sortedCoordinates); i++ {
		lower = idxStart
		upper = idxMedian + 1
		for j := idxStart; j <= idxEnd; j++ {

			cmp := cmpAllDimensions(sortedCoordinates[i][j], median, dimCur)
			if cmp < 0 {
				sortedCoordinates[i-1][lower] = sortedCoordinates[i][j]
				lower++
			} else if cmp > 0 {
				sortedCoordinates[i-1][upper] = sortedCoordinates[i][j]
				upper++
			}
		}
		// fmt.Printf("i: %d lower: %d, idxStart %d, idxMedian %d, idxEnd %d, upper %d\n",
		// 	i, lower, idxStart, idxMedian, idxEnd, upper)

		if lower < idxStart || lower-1 > idxMedian {
			panic("Impossible, lower is less than start or GE than median?")
		} else if upper < idxMedian || upper-1 > idxEnd {
			panic("Impossible, upper is LE than median or greater than end?")
		}

		if i > 1 {
			if lower != lowerSave {
				panic("Impossible, lower bounds do not match")
			} else if upper != upperSave {
				panic("Impossible, upper bounds do not match")
			}
		}
		lowerSave = lower
		upperSave = upper
	}
	for i := 0; i <= idxEnd-idxStart; i++ {
		sortedCoordinates[len(sortedCoordinates)-1][i+idxStart] = temp[i]
	}
	return lower - 1, upper - 1
}

func (node *KdNode) verifyTree(numDimensions, depth int) int {
	cntNodes := 1
	if node.value == nil {
		panic("Have a node with no value")
	}
	dimCur := depth % numDimensions

	if node.ltChild != nil {
		if cmpAllDimensions(node.ltChild.value, node.value, dimCur) >= 0 {
			panic("Less Than Child is greater than or equal to self!")
		}
		cntNodes += node.ltChild.verifyTree(numDimensions, depth+1)
	}

	if node.gtChild != nil {
		if cmpAllDimensions(node.gtChild.value, node.value, dimCur) <= 0 {
			panic("Greater Than Child is less than or equal to self!")
		}
		cntNodes += node.gtChild.verifyTree(numDimensions, depth+1)
	}
	return cntNodes
}
