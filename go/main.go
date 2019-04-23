package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	theText := "gose what yu kn"
	theSubText := "eso"
	containsPermutation(theText, theSubText)
	theText = removeSpaces(theText)
	matrixReloaded := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	expectedMatrix := [][]int{
		{13, 9, 5, 1},
		{14, 10, 6, 2},
		{15, 11, 7, 3},
		{16, 12, 8, 4},
	}

	fmt.Println(rotate90(matrixReloaded, 4))
	fmt.Println(rotate90(expectedMatrix, 4))
	fmt.Println("isUnique: " + strconv.FormatBool(isUnique(theText)))
	fmt.Println("containsPermutation: " + strconv.FormatBool(containsPermutation(theText, theSubText)))
}

func isUnique(sampleText string) bool {
	s := make(map[string]bool)
	for i := range sampleText {
		if s[string(sampleText[i])] == false {
			s[string(sampleText[i])] = true
		} else {
			return false
		}
	}
	return true
}

func containsPermutation(base string, entry string) bool {
	if len(base) < len(entry) {
		return false
	}
	subSplit := strings.Split(entry, "")
	sort.Strings(subSplit)
	for i := range base[:len(base)-len(entry)+1] {
		mainSplit := strings.Split(base[i:len(entry)+i], "")
		sort.Strings(mainSplit)
		if equal(subSplit, mainSplit) {
			return true
		}
	}
	return false
}

func removeSpaces(sampleText string) string {
	return strings.Replace(sampleText, " ", "", -1)
}

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func rotate90(image [][]int, n int) [][]int {
	for level := 0; level < n/2; level++ {
		for x := 0 + level; x < n-1-level; x++ {
			limit := (n - 1) - level
			TLpivot := image[level][x]
			TRpivot := image[level+x][limit]
			BRpivot := image[limit][limit-x]
			BLpivot := image[limit-x][level]
			// [0][0]->[3][0]->[3][3]->[0][3]
			// [1][0]->[3][1]->[2][3]->[0][2]
			// [2][0]->[3][2]->[1][3]->[0][1]
			// [1][1]->[2][1]->[2][2]->[1][2]
			fmt.Println(TLpivot, TRpivot, BRpivot, BLpivot)
			// for y := range image[x] {
			// 	fmt.Println(image[x][y])
			// }
		}
	}
	return image
}

//Node object
type Node struct {
	next *Node
	prev *Node
	val  int
}

//TreeNode object
type TreeNode struct {
	right *TreeNode
	left  *TreeNode
	val   int
}

//GraphNode object
type GraphNode struct {
	name  string
	edges []Edge
}

//Edge object
type Edge struct {
	start  GraphNode
	end    GraphNode
	weight int
}
