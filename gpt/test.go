package gpt

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	result := a / b
	return result, true
}
