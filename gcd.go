package module01

func GCD(a, b int) int {
	// step 1: if b == 0, return a
	if b == 0 {
		return a
	}
	// step 2: a becomes b and b becomes the remainder of dividing a by b
	a, b = b, a % b
	// step 3: go to step 1
	return GCD(a, b) // recursion

	// solution 2 using euclidean algorithm
	// for b != 0 {
	// 	a, b = b, a%b
	// }
	// return a
}
