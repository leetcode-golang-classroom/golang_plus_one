package sol

func plusOne(digits []int) []int {
	carry := 1
	lastPos := len(digits) - 1
	res := []int{}
	for pos := lastPos; pos >= 0; pos-- {
		result := digits[pos] + carry
		if result > 9 {
			carry = 1
		} else {
			carry = 0
		}
		digits[pos] = result % 10
	}
	if carry != 0 {
		res = append(res, carry)
	}
	for pos := 0; pos <= lastPos; pos++ {
		res = append(res, digits[pos])
	}
	return res
}
