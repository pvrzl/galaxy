package internal

// Contains is a function to check wether the x is in a or not
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
