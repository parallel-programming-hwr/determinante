package main

import (
	"fmt"
	"os"
)

func main() {
	matrix := [][]int{{2, 6, 1, 6}, {3, 0, 7, 3}, {3, 1, 9, 6}, {1, 3, 0, 3}}
	checkMat(matrix)
	fmt.Printf("Matrix:\n")
	outMat(matrix)
	fmt.Printf("Determinante: %d", getDet(matrix))

}

func checkMat(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		if len(mat) != len(mat[i]) {
			fmt.Printf("Matrix ungueltig")
			os.Exit(1)
		}

	}
}

func getDet(mat [][]int) int {
	if len(mat) > 1 {
		erg := 0
		neg := 0
		for i := 0; i < len(mat); i++ {
			mat1 := [][]int{}
			for j := 0; j < len(mat); j++ {
				if j == i {
					continue
				}
				mat1 = append(mat1, []int{})
				for k := 1; k < len(mat[j]); k++ {
					if j < i {
						mat1[j] = append(mat1[j], mat[j][k])
					} else {
						mat1[j-1] = append(mat1[j-1], mat[j][k])
					}
				}

			}
			if neg == 0 {
				erg = erg + mat[i][0]*getDet(mat1)
				neg = 1
			} else {
				erg = erg - mat[i][0]*getDet(mat1)
				neg = 0
			}
		}
		return erg
	} else {
		return mat[0][0]
	}

}

func outMat(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Printf("%d ", mat[i][j])
		}
		fmt.Printf("\n")
	}
}
