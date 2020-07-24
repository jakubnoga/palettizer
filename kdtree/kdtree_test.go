package kdtree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNearestNeighbour(t *testing.T) {
	tree := Create([][]float64{{2, 3}, {5, 4}, {9, 6}, {4, 7}, {8, 1}, {7, 2}}, 0)

	cases := [][]float64{
		{10, 10}, {0, 4}, {3, 0}, {0, 10}, {6.9, 4},
	}
	expected := [][]float64{
		{9, 6}, {2, 3}, {2, 3}, {4, 7}, {5, 4},
	}

	for idx, val := range cases {
		t.Run(fmt.Sprintf("%v", val), testCaseRunnerProvider(tree, val, expected[idx]))
	}
}

func testCaseRunnerProvider(tree *KdTree, testCase []float64, expected []float64) func(t *testing.T) {
	return func(t *testing.T) {
		best := tree.NearestNeighbour(testCase)
		if !reflect.DeepEqual(best.Point, expected) {
			t.Logf("Expected %v but got %v", expected, best.Point)
			t.Fail()
		}
	}
}