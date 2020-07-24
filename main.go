package main

import (
	"fmt"
	"kdtree/kdtree"
)

func main() {
	tree := kdtree.Create([][]float64{{2,3}, {5,4}, {9,6}, {4,7}, {8,1}, {7,2}}, 0)

	best := tree.NearestNeighbour([]float64{10,10})

	fmt.Println(tree.ToJson())
	fmt.Println(best.Point)
}
