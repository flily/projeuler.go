package p0011

const Length = 4
const Size = 20

func SolveNaive() int64 {
	result := int64(0)

	for i := range Size {
		for j := range Size - Length + 1 {
			r1 := matrix[i][j] * matrix[i][j+1] * matrix[i][j+2] * matrix[i][j+3]
			r2 := matrix[j][i] * matrix[j+1][i] * matrix[j+2][i] * matrix[j+3][i]

			result = max(result, r1, r2)
		}
	}

	for i := range Size - Length + 1 {
		for j := range Size - Length + 1 {
			r3 := matrix[i][j] * matrix[i+1][j+1] * matrix[i+2][j+2] * matrix[i+3][j+3]
			r4 := matrix[i+0][Size-j-1] * matrix[i+1][Size-j-2] * matrix[i+2][Size-j-3] * matrix[i+3][Size-j-4]

			result = max(result, r3, r4)
		}
	}

	return result
}
