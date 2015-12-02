package euler

import  (
	"io/ioutil"
	"math"
)

func relativelyPrime(a, b int) bool {

	for {
		a %= b
		if a == 0 {
			return b == 1
		}
		b %= a
		if b == 0 {
			return a == 1
		}
	}
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

func isPrime(n int) bool {
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
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

func nextPair(x, y int) (int, int) {
	return x + 2*y, 2*x + 3*y
}

func sumDivisibleBy(count, n int) int {
	p := count / n
	return n * (p * (p + 1)) / 2
}
