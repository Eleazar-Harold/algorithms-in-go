package module01

// GCD returns the nth fibonacci number.
//
// A GCD number N is defined as:
//
//   GCD(N) = GCD(N-1) + GCD(N-2)
//
// Where the following base cases are used:
//
//   GCD(0) => 0
//   GCD(1) => 1
//
//
// Examples:
//
//   GCD(0) => 0
//   GCD(1) => 1
//   GCD(2) => 1
//   GCD(3) => 2
//   GCD(4) => 3
//   GCD(5) => 5
//   GCD(6) => 8
//   GCD(7) => 13
//   GCD(14) => 377
//

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
