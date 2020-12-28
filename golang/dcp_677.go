package main

import (
	"fmt"
)

func main() {
	whichPrimes(100,algoEra(100))
}


func algoEra(max int) []bool{

	prime := make([]bool, max+1)

	for i:=0; i < max+1;i++ {
	  prime[i] = false
	}

	for i := 2; i*i <= max; i++ {
		if prime[i] == false {
				for j := i * 2; j <= max; j += i {
						prime[j] = true
				}
		}
	}
	return prime
}

func whichPrimes(max int, primes []bool) {
	for i := 2; i <= max; i++ {
			if primes[i] == false {
					fmt.Printf("%d ", i)
			}
	}
}

// The Sieve of Eratosthenes is an algorithm used to generate all prime numbers smaller than N. 
// The method is to take increasingly larger prime numbers, and mark their multiples as composite.
// For example, to find all primes less than 100, we would first mark [4, 6, 8, ...] (multiples of two), 
// then [6, 9, 12, ...] (multiples of three), and so on. 
// Once we have done this for all primes less than N, the unmarked numbers that remain will be prime.
// Implement this algorithm.
// Bonus: Create a generator that produces primes indefinitely (that is, without taking N as an input).

