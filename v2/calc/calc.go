package calc

// Add - add two or more integers
func Add(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

// Sub - subtract two integers
func Sub(a, b int) int {
	return a - b
}
