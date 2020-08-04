package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	val := 0
	if len(numbers) == 0 {
		return 0
	}
	for _, num := range numbers {
		val += num
	}
	return val
}
