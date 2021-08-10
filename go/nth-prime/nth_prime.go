package prime

func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	list := sieve(n * 11) // anyway

	return list[n-1], true
}

func sieve(up int) []int {
	primes := []int{}
	stor := map[int]bool{}

	for i := 2; i <= up; i++ {
		if !stor[i] {
			primes = append(primes, i)
			for j := i + i; j <= up; j += i {
				stor[j] = true
			}
		}
	}
	//stor = map[int]bool{} // free up memory ????

	return primes
}
