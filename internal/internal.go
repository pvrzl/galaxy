package internal

import "strings"

// Contains is a function to check wether the x is in a or not
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func StringHas(text string, arrString []string) (string, bool) {
	for _, val := range arrString {
		if strings.Contains(text, val) {
			return val, true
		}
	}
	return "", false
}
