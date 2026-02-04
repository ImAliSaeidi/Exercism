package chessboard

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	result := 0

	foundedFile, ok := cb[file]
	if !ok {
		return result
	}

	for _, square := range foundedFile {
		if square {
			result++
		}
	}

	return result
}

func CountInRank(cb Chessboard, rank int) int {
	result := 0

	if rank < 1 || rank > 8 {
		return result
	}

	for _, file := range cb {
		if file[rank-1] {
			result++
		}
	}

	return result
}

func CountAll(cb Chessboard) int {
	result := 0

	for _, file := range cb {
		for range file {
			result++
		}
	}

	return result
}

func CountOccupied(cb Chessboard) int {
	result := 0

	for _, file := range cb {
		for _, square := range file {
			if square {
				result++
			}
		}
	}

	return result
}
