package utils

func CheckBoundsByte(mat [][]byte, x, y int) bool {
	if x < 0 || x >= len(mat) {
		return false
	}
	if y < 0 || y >= len(mat[0]) {
		return false
	}
	return true
}

func CheckBoundsInt(mat [][]int, x, y int) bool {
	if x < 0 || x >= len(mat) {
		return false
	}
	if y < 0 || y >= len(mat[0]) {
		return false
	}
	return true
}

func MatCopyByte(mat [][]byte) [][]byte {
	duplicate := make([][]byte, len(mat))
	for i := range mat {
		duplicate[i] = make([]byte, len(mat[i]))
		copy(duplicate[i], mat[i])
	}
	return duplicate
}

func MatCopyInt(mat [][]int) [][]int {
	duplicate := make([][]int, len(mat))
	for i := range mat {
		duplicate[i] = make([]int, len(mat[i]))
		copy(duplicate[i], mat[i])
	}
	return duplicate
}

func EqMatByte(mat1 [][]byte, mat2 [][]byte) bool {
	for i := 0; i < len(mat1); i++ {
		for j := 0; j < len(mat1[0]); j++ {
			if mat1[i][j] != mat2[i][j] {
				return false
			}
		}
	}
	return true
}

func EqMatInt(mat1 [][]int, mat2 [][]int) bool {
	for i := 0; i < len(mat1); i++ {
		for j := 0; j < len(mat1[0]); j++ {
			if mat1[i][j] != mat2[i][j] {
				return false
			}
		}
	}
	return true
}
