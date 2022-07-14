package solution

const base = (1000*1000*1000 + 7)

func peopleAwareOfSecret(n int, delay int, forget int) int {
	a := make([]int, 2*n+2) // who knows the secret for day [i]
	b := make([]int, 2*n+2) // who could share the secret for day [i]
	a[1] = 1

	for i := 1; i < 1+delay; i++ {
		a[i] = 1
	}
	for i := 1 + delay; i < 1+forget; i++ {
		a[i] = 1
		b[i] = 1
	}

	for i := 2; i <= n; i++ {
		for j := i; j < i+delay; j++ {
			// a[j] += b[i]
			a[j] = (a[j] + b[i]) % base
		}
		for j := i + delay; j < i+forget; j++ {
			// b[j] += b[i]
			b[j] = (b[j] + b[i]) % base
			// a[j] += b[i]
			a[j] = (a[j] + b[i]) % base
		}
	}

	return a[n]
}
