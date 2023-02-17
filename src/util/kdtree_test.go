package util

import "testing"

type kdTreeTestLoc struct {
	coords []float64
}

func (t *kdTreeTestLoc) GetCoordinates() []float64 {
	return t.coords
}

func TestOneNode(t *testing.T) {

	locs := []Coordinates{
		&kdTreeTestLoc{[]float64{21, 25}},
	}
	actual := CreateKdTree(locs, 2)

	if actual == nil {
		t.Fatal("Failed to create node")
	}

	expected := &KdNode{
		value: locs[0],
	}

	verifyTree(t, actual, expected, 2, 0)
}
func TestTwoNodes(t *testing.T) {

	locs := []Coordinates{
		&kdTreeTestLoc{[]float64{21, 25}},
		&kdTreeTestLoc{[]float64{25, 21}},
	}
	actual := CreateKdTree(locs, 2)

	if actual == nil {
		t.Fatal("Failed to create node")
	}

	expected := &KdNode{
		value: locs[0],
		gtChild: &KdNode{
			value: locs[1],
		},
	}

	verifyTree(t, actual, expected, 2, 0)
}
func TestThreeNodes(t *testing.T) {

	locs := []Coordinates{
		&kdTreeTestLoc{[]float64{21, 25}},
		&kdTreeTestLoc{[]float64{25, 21}},
		&kdTreeTestLoc{[]float64{23, 23}},
	}
	actual := CreateKdTree(locs, 2)

	if actual == nil {
		t.Fatal("Failed to create node")
	}

	expected := &KdNode{
		value: locs[2],
		ltChild: &KdNode{
			value: locs[0],
		},
		gtChild: &KdNode{
			value: locs[1],
		},
	}

	verifyTree(t, actual, expected, 2, 0)
}

func TestFourNodes(t *testing.T) {

	locs := []Coordinates{
		&kdTreeTestLoc{[]float64{21, 25}},
		&kdTreeTestLoc{[]float64{25, 21}},
		&kdTreeTestLoc{[]float64{23, 23}},
		&kdTreeTestLoc{[]float64{25, 19}},
	}
	actual := CreateKdTree(locs, 2)

	if actual == nil {
		t.Fatal("Failed to create node")
	}

	expected := &KdNode{
		value: locs[2],
		ltChild: &KdNode{
			value: locs[0],
		},
		gtChild: &KdNode{
			value: locs[3],
			gtChild: &KdNode{
				value: locs[1],
			},
		},
	}

	verifyTree(t, actual, expected, 2, 0)
}

func TestMultiDepth(t *testing.T) {

	locs := []Coordinates{
		&kdTreeTestLoc{[]float64{12, 25, 34}},
		&kdTreeTestLoc{[]float64{15, 21, 39}},
		&kdTreeTestLoc{[]float64{19, 23, 35}},
		&kdTreeTestLoc{[]float64{18, 24, 33}},

		&kdTreeTestLoc{[]float64{16, 27, 39}},
		&kdTreeTestLoc{[]float64{18, 25, 39}},
		&kdTreeTestLoc{[]float64{13, 26, 38}},
		&kdTreeTestLoc{[]float64{12, 29, 33}},

		&kdTreeTestLoc{[]float64{16, 28, 32}},
		&kdTreeTestLoc{[]float64{10, 22, 37}},
		&kdTreeTestLoc{[]float64{18, 24, 38}},
		&kdTreeTestLoc{[]float64{14, 20, 31}},

		&kdTreeTestLoc{[]float64{15, 22, 32}},
		&kdTreeTestLoc{[]float64{11, 28, 30}},
		&kdTreeTestLoc{[]float64{11, 26, 37}},
		&kdTreeTestLoc{[]float64{12, 23, 32}},
	}
	actual := CreateKdTree(locs, 3)
	// Expected results:
	// 9.10, 14.11, 13.11, 15.12, 0.12, 7.12, 6.13, 11.14, 1.15, 12.15, 4.16, 8.16, 3.18, 10.18, 5.18, 2.19
	// 11.20, 1.21, 12.22, 9.22, 15.23, 2.23, 3.24, 10.24, 0.25, 5.25, 14.26, 6.26, 4.27, 13.28, 8.28, 7.29
	// 13.30, 11.31, 15.32, 12.32, 8.32, 7.33, 3.33, 0.34, 2.35, 9.37, 14.37, 6.38, 10.38, 1.39, 4.39, 5.39

	// Rotate:
	// 9.22, 15.23, 0.25, 14.26, 6.26, 13.28, 7.29 upper: 1.21, 12.22, 2.23, 3.24, 10.24, 5.25, 4.27, 8.28
	// 13.30, 15.32, 7.33, 0.34, 9.37, 14.37, 6.38 upper: 12.32, 8.32, 3.33, 2.35, 10.38, 1.39, 4.39, 5.39
	// 9.10, 14.11, 13.11, 15.12, 0.12, 7.12, 6.13 upper: 1.15, 12.15, 4.16, 8.16, 3.18, 10.18, 5.18, 2.19

	// Lower Rotate:
	// 15.32, 0.34, 9.37 upper: 13.30, 7.33, 6.38
	// 9.10, 15.12, 0.12 upper: 13.11, 7.12, 6.13

	// Upper Rotate:
	// 12.32, 2.35, 1.39 upper: 8.32, 10.38, 4.39, 5.39
	// 1.15, 12.15, 2.19 upper: 4.16, 8.16, 10.18, 5.18

	// Upper Upper Rotate:
	// 8.16 upper: 4.16 5.18

	if actual == nil {
		t.Fatal("Failed to create node")
	}

	expected := &KdNode{
		value: locs[11],
		ltChild: &KdNode{
			value: locs[14],
			ltChild: &KdNode{
				value: locs[0],
				ltChild: &KdNode{
					value: locs[15],
				},
				gtChild: &KdNode{
					value: locs[9],
				},
			},
			gtChild: &KdNode{
				value: locs[7],
				ltChild: &KdNode{
					value: locs[13],
				},
				gtChild: &KdNode{
					value: locs[6],
				},
			},
		},
		gtChild: &KdNode{
			value: locs[3],
			ltChild: &KdNode{
				value: locs[2],
				ltChild: &KdNode{
					value: locs[12],
				},
				gtChild: &KdNode{
					value: locs[1],
				},
			},
			gtChild: &KdNode{
				value: locs[10],
				ltChild: &KdNode{
					value: locs[8],
				},
				gtChild: &KdNode{
					value: locs[4],
					gtChild: &KdNode{
						value: locs[5],
					},
				},
			},
		},
	}

	verifyTree(t, actual, expected, 3, 0)
}

func verifyTree(t *testing.T, actual, expected *KdNode, numDimensions, depth int) {
	if cmpAllDimensions(actual.value, expected.value, numDimensions) != 0 {
		t.Fatalf("Depth %d: node value did not match. Actual: %+v. expected: %+v\n", depth, actual.value, expected.value)
	}

	if (actual.ltChild != nil) != (expected.ltChild != nil) ||
		(actual.gtChild != nil) != (expected.gtChild != nil) {
		t.Fatalf("Depth %d: children existence did not match. Actual: %+v. expected: %+v\n", depth, actual, expected)
	}

	if expected.ltChild != nil {
		verifyTree(t, actual.ltChild, expected.ltChild, numDimensions, depth+1)
	}

	if expected.gtChild != nil {
		verifyTree(t, actual.gtChild, expected.gtChild, numDimensions, depth+1)
	}
}
