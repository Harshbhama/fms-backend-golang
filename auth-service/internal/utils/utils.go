package utils

const base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func IntToBase62(num int64) string {
	var result string
	for num > 0 {
		remainder := num % 62
		result = string(base62Alphabet[remainder]) + result
		num /= 62
	}
	return result
}

func Base62ToInt(base62 string) (int64, error) {
	var result int64
	for _, char := range base62 {
		var idx int
		for k, c := range base62Alphabet {
			if c == char {
				idx = k
				break
			}
		}
		result = result*62 + int64(idx)
	}
	return result, nil
}