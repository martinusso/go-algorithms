package sieveofatkin

import "testing"

var (
	primesLowerThan10 = []int64{2, 3, 5, 7}

	primesLowerThan100 = []int64{2, 3, 5, 7,
		11, 13, 17, 19, 23, 29, 31, 37, 41, 43,
		47, 53, 59, 61, 67, 71, 73, 79, 83, 89,
		97}

	primesLowerThan1000 = []int64{2, 3, 5, 7,
		11, 13, 17, 19, 23, 29, 31, 37, 41, 43,
		47, 53, 59, 61, 67, 71, 73, 79, 83, 89,
		97, 101, 103, 107, 109, 113, 127, 131,
		137, 139, 149, 151, 157, 163, 167, 173,
		179, 181, 191, 193, 197, 199, 211, 223,
		227, 229, 233, 239, 241, 251, 257, 263,
		269, 271, 277, 281, 283, 293, 307, 311,
		313, 317, 331, 337, 347, 349, 353, 359,
		367, 373, 379, 383, 389, 397, 401, 409,
		419, 421, 431, 433, 439, 443, 449, 457,
		461, 463, 467, 479, 487, 491, 499, 503,
		509, 521, 523, 541, 547, 557, 563, 569,
		571, 577, 587, 593, 599, 601, 607, 613,
		617, 619, 631, 641, 643, 647, 653, 659,
		661, 673, 677, 683, 691, 701, 709, 719,
		727, 733, 739, 743, 751, 757, 761, 769,
		773, 787, 797, 809, 811, 821, 823, 827,
		829, 839, 853, 857, 859, 863, 877, 881,
		883, 887, 907, 911, 919, 929, 937, 941,
		947, 953, 967, 971, 977, 983, 991, 997}
)

func Test_Flip(t *testing.T) {
	soa := initSieveOfAtkin(10)
	soa.flip(2)
	if !soa.sieve[2] {
		t.Errorf("Expected true got false")
	}
	soa.flip(2)
	if soa.sieve[2] {
		t.Errorf("Expected false got true")
	}
}

func Test_Invalidate(t *testing.T) {
	soa := initSieveOfAtkin(10)

	soa.invalidate(2)
	if soa.sieve[2] {
		t.Errorf("Expected false got true")
	}

	soa.flip(2)
	if !soa.sieve[2] {
		t.Errorf("Expected true got false")
	}
}

func Test_IsPrime(t *testing.T) {
	soa := initSieveOfAtkin(100)

	if soa.isPrime(2) {
		t.Errorf("Expected false got true")
	}

	soa.flip(2)
	if !soa.sieve[2] {
		t.Errorf("Expected true got false")
	}
}

func Test_PrimesLowerThan10(t *testing.T) {
	primes := GetPrimes(10)
	if len(primes) != len(primesLowerThan10) {
		t.Errorf("Expected %d primes lower than 10 got %d", len(primesLowerThan10), len(primes))
	}
	for _, prime := range primesLowerThan10 {
		if !contains(primes, prime) {
			t.Errorf("Should contains %d as prime lower than 10", prime)
		}
	}
}

func Test_PrimesLowerThan100(t *testing.T) {
	primes := GetPrimes(100)
	if len(primes) != len(primesLowerThan100) {
		t.Errorf("Expected %d primes lower than 100 got %d", len(primesLowerThan100), len(primes))
	}
	for _, prime := range primesLowerThan100 {
		if !contains(primes, prime) {
			t.Errorf("Should contains %d as prime lower than 100", prime)
		}
	}
}

func Test_PrimesLowerThan1000(t *testing.T) {
	primes := GetPrimes(1000)
	if len(primes) != len(primesLowerThan1000) {
		t.Errorf("Expected %d primes lower than 1000 got %d", len(primesLowerThan1000), len(primes))
	}
	for _, prime := range primesLowerThan1000 {
		if !contains(primes, prime) {
			t.Errorf("Should contains %d as prime lower than 1000", prime)
		}
	}
}

func contains(s []int64, e int64) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
