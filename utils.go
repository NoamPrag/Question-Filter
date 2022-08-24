package main

func transpose(slice [][]string) [][]string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]string, xl)
	for i := range result {
		result[i] = make([]string, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			if len(slice[j]) < i+1 {
				result[i][j] = ""
			} else {
				result[i][j] = slice[j][i]
			}
		}
	}
	return result
}
