package euler

import (
	"io/ioutil"
	"math"
)

// Problem 9
// Special Pythagorean triplet
// A Pythagorean triplet is a set of three natural numbers,
// a < b < c, for which, a^2 + b^2 = c^2
// For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
//
// There exists exactly one Pythagorean triplet for which a + b + c = 1000.
// Find the product abc.

func Problem009(n int) int {

	for b := 1; b < n/2; b++ {
		a := (float64(n*n)/2 - float64(n*b)) / float64(n-b)

		if int(a)%1 == 0 {
			return int(a) * b * (n - int(a) - b)
		}
	}

	return 0
}

func RelativelyPrime(a, b int) bool {

	for {
		if a %= b != 0 {
			return b ==1
		}
		if b %= a != 0 {
			return a == 1
		}
	}
}

// Problem 8
// Largest product in a series
// The four adjacent digits in the 1000-digit number
// that have the greatest product are 9 × 9 × 8 × 9 = 5832.
//
// Find the thirteen adjacent digits in the 1000-digit number
// that have the greatest product. What is the value of this product?

func Problem008(n int) int {

	// Read The 1000-digit number from disk
	digits := getSlice("1000.digit.number.txt")
	prod := 0

	for i := 0; i < 1000-n; i++ {
		tmp := 1
		for k := i; k < i+n; k++ {
			tmp *= digits[k]
		}
		if tmp > prod {
			prod = tmp
		}
	}

	return prod
}

func getSlice(file string) []int {
	var slice []int

	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return slice
	}

	input := string(dat)

	for _, a := range input {
		i := int(a - '0')
		if -1 < i && i < 10 {
			slice = append(slice, i)
		}
	}
	return slice
}

// Problem 7
// 10001st prime
// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13,
// we can see that the 6th prime is 13.
//
// What is the 10001st prime number?
func Problem007(n int) int {

	x := 3
	p := 0
	count := 1
	for count < n {
		if isPrime(x) {
			count++
			p = x
		}
		x++
	}

	return int(p)
}

func isPrime(n int) bool {
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Problem 6
// Sum square difference
// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385
// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 55^2 = 3025
// Hence the difference between the sum of the squares of the first ten
// natural numbers and the square of the sum is 3025 − 385 = 2640.
//
// Find the difference between the sum of the squares of the first
// one hundred natural numbers and the square of the sum.
func Problem006(n int) int {

	squares := (n * (n + 1) * (2*n + 1)) / 6
	sums := (n * (n + 1)) / 2
	sums *= sums

	return sums - squares
}

// Problem 5
// Smallest multiple
// 2520 is the smallest number that can be divided by each of the numbers
// from 1 to 10 without any remainder.
//
// What is the smallest positive number that is evenly divisible
// by all of the numbers from 1 to 20?
func Problem005(n int) int {

	primes := sieveOfEratosthenes(n)
	powers := make([]int, n)

	// find all prime factors of 2 through n
	// 2=2^1, 3=3^1, 4=2^2, 5=5^1, 6=2^1x3^1, 7=7^1, 8=2^3, 9=3^2, 10=2^1*5^1
	for i := 1; i <= n; i++ {
		val := i
		for _, p := range primes {
			pow := 0
			for val%p == 0 {
				val /= p
				pow += 1
			}
			if pow > powers[p] {
				powers[p] = pow
			}
		}
	}

	// multiply factors with greatest exponent
	// 2^3*3^2*5^1*7^1 = 2520
	product := 1
	for base, exponent := range powers {
		product *= int(math.Pow(float64(base), float64(exponent)))
	}

	return product
}

// Problem 4
// Largest palindrome product
// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers
// is 9009 = 91 × 99.
//
// Find the largest palindrome made from the product of two 3-digit numbers.
func Problem004(n int) int {
	max := 0
	a := n
	limit := (n + 1) / 10
	for a >= limit {
		b := n
		for b >= a {
			if a*b <= max {
				break
			}

			if isPalindrome(a * b) {
				max = a * b
			}

			b--
		}
		a--
	}
	return max
}

func isPalindrome(n int) bool {
	if reverse(n) == n {
		return true
	}
	return false
}

func reverse(val int) int {
	reversed := 0
	for val > 0 {
		reversed = 10*reversed + val%10
		val = val / 10
	}
	return reversed
}

// Problem 3
// Largest prime factor
// The prime factors of 13195 are 5, 7, 13 and 29.
//
// What is the largest prime factor of the number 600851475143 ?
//
func Problem003(n int) int {

	factor := 1
	primes := sieveOfEratosthenes(int(math.Sqrt(float64(n))))

	for _, x := range primes {
		if n%x == 0 {
			factor = x
		}
	}

	return factor
}

func sieveOfEratosthenes(N int) []int {
	// sieve
	c := make([]bool, N) // c for composite.  false means prime candidate
	c[1] = true          // 1 not considered prime
	p := 2
	for {
		// first allowed optimization:  outer loop only goes to sqrt(limit)
		p2 := p * p
		if p2 >= N {
			break
		}
		// second allowed optimization:  inner loop starts at sqr(p)
		for i := p2; i < N; i += p {
			c[i] = true // it's a composite

		}
		// scan to get next prime for outer loop
		for {
			p++
			if !c[p] {
				break
			}
		}
	}

	// sieve complete.  now save off primes.
	v := make([]int, 0)
	for n := 1; n < N; n++ {
		if c[n] == false {
			v = append(v, n)
		}
	}

	return v
}

// Problem 2
// Even Fibonacci numbers
// Each new term in the Fibonacci sequence is generated by adding the
// previous two terms. By starting with 1 and 2, the first 10 terms will be:
//
// 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
//
// By considering the terms in the Fibonacci sequence whose values
// do not exceed four million, find the sum of the even-valued terms.
//
func Problem002() int {

	x, y := 1, 1
	sum := 0

	for sum <= 4000000 {
		sum += x + y
		x, y = nextPair(x, y)
	}
	return sum
}

func nextPair(x, y int) (int, int) {
	return x + 2*y, 2*x + 3*y
}

// Problem 1
// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
//
func Problem001(count int) int {

	count -= 1

	return sumDivisibleBy(count, 3) +
		sumDivisibleBy(count, 5) -
		sumDivisibleBy(count, 15)
}

func sumDivisibleBy(count, n int) int {
	p := count / n
	return n * (p * (p + 1)) / 2
}
