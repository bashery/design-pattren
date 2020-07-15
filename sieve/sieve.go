package sieve

// Sieve teturn list of prime number
func Sieve(up int) []int {
	primes := []int{}
	stor := map[int]string{}

	for i := 2; i <= up; i++ {
		if stor[i] == "" {
			primes = append(primes, i)
			for j := i + i; j <= up; j += i {
				stor[j] = "1"
			}
		}
	}
	return primes
}
