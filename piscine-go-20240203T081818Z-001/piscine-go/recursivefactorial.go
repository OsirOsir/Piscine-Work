package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb < 0 || nb > 20 {
		return 0
	}
	return RecursiveFactorial(nb-1) * nb
}
