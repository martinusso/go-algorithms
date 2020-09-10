// Package sieveofatkin is an implementation for the Sieve of Atkin, a modern algorithm for finding all prime numbers up to a specified integer
package sieveofatkin

import "math"

// GetPrimes returns all prime numbers up to a specified integer
func GetPrimes(limit int64) []int64 {
	return initSieveOfAtkin(limit).getPrimes()
}

type sieveOfAtkin struct {
	limit int64
	sieve []bool
}

func initSieveOfAtkin(limit int64) *sieveOfAtkin {
	soa := &sieveOfAtkin{limit: limit}
	soa.sieve = make([]bool, soa.limit)
	for i := int64(0); i < soa.limit; i++ {
		soa.sieve[i] = false
	}
	return soa
}

func (soa sieveOfAtkin) flip(prime int64) {
	soa.sieve[prime] = !soa.sieve[prime]
}

func (soa sieveOfAtkin) invalidate(prime int64) {
	if soa.sieve[prime] {
		soa.sieve[prime] = false
	}
}

func (soa sieveOfAtkin) isPrime(prime int64) bool {
	return soa.sieve[prime]
}

func (soa sieveOfAtkin) getPrimes() []int64 {
	testingLimit := int64(math.Ceil(math.Sqrt(float64(soa.limit))))

	for i := int64(0); i < testingLimit; i++ {
		for j := int64(0); j < testingLimit; j++ {
			var n int64
			iPow2 := int64(math.Pow(float64(i), 2))
			jPow2 := int64(math.Pow(float64(j), 2))

			// n = 4*i^2 + j^2
			n = 4*iPow2 + jPow2
			if n <= soa.limit && (n%12 == 1 || n%12 == 5) {
				soa.flip(n)
			}

			// n = 3*i^2 + j^2
			n = 3*iPow2 + jPow2
			if n <= soa.limit && n%12 == 7 {
				soa.flip(n)
			}

			// n = 3*i^2 - j^2
			n = 3*iPow2 - jPow2
			if (n <= soa.limit) && (i > j) && (n%12 == 11) {
				soa.flip(n)
			}
		}
	}

	for i := int64(5); i < testingLimit; i++ {
		if soa.isPrime(i) {
			k := int64(math.Pow(float64(i), 2))
			for j := k; j < soa.limit; j += k {
				soa.invalidate(j)
			}
		}
	}

	primes := []int64{2, 3}
	for i := int64(5); i < int64(len(soa.sieve)); i += 2 {
		if soa.isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}
