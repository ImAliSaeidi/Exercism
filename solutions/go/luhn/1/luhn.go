package luhn

func Valid(id string) bool {
	sum := 0
	digitCount := 0

	for i := len(id) - 1; i >= 0; i-- {
		r := id[i]
		switch {
		case r == ' ':
			continue
		case r >= '0' && r <= '9':
			digit := int(r - '0')
			digitCount++
			if digitCount%2 == 0 {
				digit *= 2
				if digit > 9 {
					digit -= 9
				}
			}
			sum += digit
		default:
			return false
		}
	}

	return digitCount > 1 && sum%10 == 0
}
