package ch04

import "testing"

// 4.2 思想不难理解，代码编写困难
// 采用分解矩阵，加减之后组合，把原来需要计算8次减少为只需要计算7次

func TestStrassen(t *testing.T) {
	arr := [][]int{}
	arr = append(arr, []int{1, 2, 3})
	arr = append(arr, []int{1, 2, 3})
	arr = append(arr, []int{1, 2, 3})

	t.Log(squareMatrixMultiply(arr, arr))
}

// n^2矩阵乘法
func squareMatrixMultiply(A, B [][]int) [][]int {
	n := len(A)
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
		for j := 0; j < n; j++ {
			temp := 0
			for k := 0; k < n; k++ {
				temp += A[i][k] * B[k][j]
			}
			result[i][j] = temp
		}
	}
	return result
}
