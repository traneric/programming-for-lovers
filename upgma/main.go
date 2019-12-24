// Unweighted Pair Group Method with Arithmetic Mean (UPGMA) program.

package main

import (
	"fmt"
	"os"
	"strconv"
)

type DistanceMatrix [][]float64

type Tree []*Node

// we also think of a cluster as a *Node

type Node struct {
	age            float64
	label          string
	child1, child2 *Node // if at leaf, both will be nil
}

func main() {
	fmt.Println("UPGMA.")

	fileName := os.Args[1]
	mtx, speciesNames := ReadMatrixFromFile(fileName)

	var t Tree = UPGMA(mtx, speciesNames)

	t.PrintGraphViz()
}

//UPGMA function goes here
func UPGMA(mtx DistanceMatrix, speciesNames []string) Tree {
	t := InitializeTree(speciesNames)
	// will create all nodes needed, and not point any node at any node as a child.

	clusters := t.InitializeClusters()
	//clusters will start out as just the leaves of t (slice of node pointers)
	// a cluster is a pointer to a node!

	// now for the UPGMA engine ... apply steps of the algorithm numLeaves - 1 steps
	numLeaves := len(speciesNames)
	for p := numLeaves; p < 2*numLeaves-1; p++ {
		// first, find minimum element of mtx
		row, col, minVal := FindMinElement(mtx)
		// big assumption: col > row
		// set the age of current node.
		t[p].age = minVal / 2.0

		// set children of t[p]
		t[p].child1 = clusters[row]
		t[p].child2 = clusters[col]

		// now, update matrix and clusters
		mtx = AddRowCol(mtx, clusters, row, col)

		// add t[p] to the end of clusters
		clusters = append(clusters, t[p])

		//now, hack out row and col from distance matrix and from clusters

		mtx = DelRowCol(mtx, row, col)
		clusters = DelClusters(clusters, row, col)

		fmt.Println(clusters)
	}
	// now, we are ready to return tree
	return t
}

//FindMinElement takes a DistanceMatrix and returns the row index, column index,
//and value corresponding to a minimum element.
//Assumption that col > row
func FindMinElement(mtx DistanceMatrix) (int, int, float64) {
	if len(mtx) <= 1 || len(mtx[0]) <= 1 {
		panic("One row or one column!")
	}

	// can now assume that matrix is at least 2 x 2
	row := 0
	col := 1
	minVal := mtx[row][col]

	// range over matrix, and see if we can do better than minVal.
	for i := 0; i < len(mtx)-1; i++ {
		// start column ranging at i + 1
		for j := i + 1; j < len(mtx[i]); j++ {
			// do we have a winner?
			if mtx[i][j] < minVal {
				// update all three variables
				minVal = mtx[i][j]
				row = i
				col = j
				// col will still always be > row.
			}
		}
	}
	return row, col, minVal
}

//DelRowCol takes a distance matrix and a row/col index and deletes the row and
//column indicated, returning the resulting matrix
func DelRowCol(mtx DistanceMatrix, row, col int) DistanceMatrix {
	// first, let's delete appropriate rows
	// remember that col > row, we should delete col-th row first
	mtx = append(mtx[:col], mtx[col+1:]...)
	mtx = append(mtx[:row], mtx[row+1:]...)

	//now, delete columns row and col as well.
	for i := range mtx {
		mtx[i] = append(mtx[i][:col], mtx[i][col+1:]...)
		mtx[i] = append(mtx[i][:row], mtx[i][row+1:]...)
	}

	return mtx
}

//DelClusters takes a slice of Node pointers along with a row/col index and
//deletes the clusters in the slice corresponding to these indices.
//Assume col > row
func DelClusters(clusters []*Node, row, col int) []*Node {
	clusters = append(clusters[:col], clusters[col+1:]...)
	clusters = append(clusters[:row], clusters[row+1:]...)
	return clusters
}

//AddRowCol takes a DistanceMatrix, a slice of current clusters, and a row/col
//index (col > row).
//It returns the matrix corresponding to "gluing" clusters[row] and clusters[col]
//together and forming a new row/col of the matrix (no deletions yet).
func AddRowCol(mtx DistanceMatrix, clusters []*Node, row, col int) DistanceMatrix {
	AssertRowCol(row, col)

	n := len(mtx)
	newRow := make([]float64, n+1) // last element will be 0 by default :)

	// all values 0.0 by default, let's set the values that need to be set.
	for j := 0; j < n; j++ {
		if j != row && j != col {
			// now compute newRow[j]
			//broke: newRow[j] = (mtx[row][j] + mtx[col][j]/2.0)
			//woke: weighted average based on number of elements in each cluster
			size1 := CountLeaves(clusters[row])

			size2 := CountLeaves(clusters[col])

			newRow[j] = (float64(size1)*mtx[row][j] + float64(size2)*mtx[col][j]) / float64(size1+size2)
		}
	}

	//let's append new row to matrix
	mtx = append(mtx, newRow)

	//we need to add last column as well to the matrix
	for i := 0; i < n; i++ {
		mtx[i] = append(mtx[i], newRow[i])
	}

	return mtx
}

func AssertRowCol(row, col int) {
	if row >= col { // want col > row
		panic("Minimum column index not bigger than row.")
	}
}

//CountLeaves is a recursive Node function that counts the number of leaves in
//the tree rooted at the node. It returns 1 at a leaf.
func CountLeaves(v *Node) int {
	// base case: if we are at a leaf, return 1
	if v.child1 == nil || v.child2 == nil {
		return 1
	} else { // know v.child1 and v.child2 are not nil
		// inductive step: count leaves of each child, and sum
		return CountLeaves(v.child1) + CountLeaves(v.child2)
	}
}

//InitializeClusters is a Tree method that returns a slice of pointers to the
//leaves of the Tree
func (t Tree) InitializeClusters() []*Node {
	numNodes := len(t)
	numLeaves := (numNodes + 1) / 2

	clusters := make([]*Node, numLeaves)
	// clusters[i] should point to the i-th leaf node of t
	for i := range clusters {
		clusters[i] = t[i]
	}

	return clusters
}

//InitializeTree takes the n names of our present-day species (leaves) and
//returns a rooted binary tree with 2n-1 total nodes, where the leaves are the
//first n and have the associated species names.
func InitializeTree(speciesNames []string) Tree {
	numLeaves := len(speciesNames)
	var t Tree // a Tree is []*Node

	t = make([]*Node, 2*numLeaves-1)
	// all of these pointers have default value of nil; we need to point them at nodes

	// we should create 2n-1 nodes.
	for i := range t {
		var vx Node
		// let's label the first numLeaves nodes with the appropriate species name.
		// by default, vx.age = 0.0, and its children are nil.
		if i < numLeaves {
			//at a leaf ... let's assign its label.
			vx.label = speciesNames[i]
		} else {
			// let's just give it an unspecific name
			vx.label = "Ancestor species " + strconv.Itoa(i)
		}
		// one thing to do: point t[i] at vx
		t[i] = &vx
	}

	return t
}
