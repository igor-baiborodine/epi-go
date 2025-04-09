package arrays

// GeneratePrimes returns all prime numbers less than or equal to n
// using the Sieve of Eratosthenes algorithm.
//
// Parameters:
//   - n: An integer specifying the inclusive upper bound up to which prime numbers are to be generated.
//
// Returns:
//   - []int: A slice containing all prime numbers less than or equal to n in ascending order.
//
// Time Complexity: O(n log log n), where n is the input parameter, due to the sieving process.
// Space Complexity: O(n), for storing the boolean array indicating primality.
func GeneratePrimes(n int) []int {
	primes := []int{}
	isPrime := make([]bool, n+1)

	// isPrime[p] represents if p is prime or not. Initially, set each
	// to true, excepting 0 and 1. Then use sieving to eliminate non-primes.
	for i := range isPrime {
		if i > 1 {
			isPrime[i] = true
		}
	}
	for p := 2; p <= n; p++ {
		if isPrime[p] {
			primes = append(primes, p)
		}
		for j := p; j <= n; j += p {
			isPrime[j] = false
		}
	}
	return primes
}
