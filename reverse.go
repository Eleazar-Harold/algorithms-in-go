package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	res := ""

	// first solution attempt
	// for i := len(word)-1; i >= 0; i -- {
	// 	res = res + string(word[i])
	// }

	// second solution attempt
	for _, r := range word {
		res = string(r) + res
	}

	return res
}
