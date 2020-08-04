package module01

import (
	"fmt"
	"strconv"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	res, err := strconv.ParseInt(value, base, 0)
	if err != nil {
		fmt.Println(err)
	}
	return int(res)
}
