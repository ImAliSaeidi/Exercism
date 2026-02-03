package strand

func ToRNA(dna string) string {
	result := ""

	for _, c := range dna {
		switch c {
		case 'G':
			result += "C"
		case 'C':
			result += "G"
		case 'T':
			result += "A"
		case 'A':
			result += "U"
		}
	}

	return result
}
