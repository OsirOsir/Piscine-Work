package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 63 {
		return 0
	}
	results := 1
	for i := 1; i <= nb; i++ {
		if results < 0 {
			return 0
		}
		results *= i
	}
	return results
}
