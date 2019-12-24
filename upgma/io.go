package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//ReadMatrixFromFile takes a file name and reads the information in this file to produce
//a distance matrix and a slice of strings holding the species names.  The first line of the
//file should contain the number of species.  Each other line contains a species name
//and its distance to each other species.
func ReadMatrixFromFile(fileName string) (DistanceMatrix, []string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error: couldn't open the file")
		os.Exit(1)
	}
	var lines []string = make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		fmt.Println("Sorry: there was some kind of error during the file reading")
		os.Exit(1)
	}
	file.Close()

	mtx := make(DistanceMatrix, 0)
	speciesNames := make([]string, 0)

	for idx, _ := range lines {
		if idx >= 1 {
			row := make([]float64, 0)
			nums := strings.Split(lines[idx], "\t")
			for i, num := range nums {
				if i == 0 {
					speciesNames = append(speciesNames, num)
				} else {
					n, err := strconv.ParseFloat(num, 64)
					if err != nil {
						fmt.Println("Error: Wrong format of matrix!")
						os.Exit(1)
					}
					row = append(row, n)
				}
			}
			mtx = append(mtx, row)
		}
	}
	return mtx, speciesNames
}

// PrintGraphViz prints the tree in GraphViz format, where directed = true
// if we desire to print a directed graph and directed = false for an
// undirected graph.
func (t Tree) PrintGraphViz() {
	fmt.Println("strict digraph {")
	for i := range t {
		if t[i].child1 != nil && t[i].child2 != nil {
			//print first edge
			fmt.Print("\"", t[i].label, "\"")
			fmt.Print("->")
			fmt.Print("\"", t[i].child1.label, "\"")
			fmt.Print("[label = \"")
			fmt.Printf("%.2f", t[i].age-t[i].child1.age)
			fmt.Print("\"]")
			fmt.Println()

			//print second edge
			fmt.Print("\"", t[i].label, "\"")
			fmt.Print("->")
			fmt.Print("\"", t[i].child2.label, "\"")
			fmt.Print("[label = \"")
			fmt.Printf("%.2f", t[i].age-t[i].child2.age)
			fmt.Print("\"]")
			fmt.Println()
		}
	}
	fmt.Println("}")
}
