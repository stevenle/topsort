package topsort

// Reverse Reverse the slice items.
func Reverse(a []string) {
	for l, i := len(a), len(a)/2-1; i >= 0; i-- {
		opp := l-1-i
		a[i], a[opp] = a[opp], a[i]
	}
}
