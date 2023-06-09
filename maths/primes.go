package maths

import (
  "math"
)

// TODO: describe functions

func GetNthPrime(n int) (int, []int) {
  primes := []int{2,3}

  for i := primes[len(primes)-1] + 2; len(primes) <= n; i += 2 {
    for _, v := range primes {
      // not a prime
      if i % v == 0 {
        break
      }

      if math.Sqrt(float64(i)) < float64(v) {
        primes = append(primes, i)
        break
      }
    }
  }
  
  return primes[n-1], primes
}

func PrimalityTest(x int) bool {
  if x == 0 || x == 1 {
    return false
  }

  for i := 2; i <= x / 2; i++ {
    if (x % i == 0) {
      return false
    }
  }
  return true
}

func PrimesRange(n int) (primes []int) {
  for i := 0; i <= n; i++ {
    if PrimalityTest(i) {
      primes = append(primes, i)
    }
  }
  return 
}

func PrimeFactors(n int) (ret []int) {
  for n % 2 == 0 {
    pfs = append(pfs, 2)
    n = n / 2
  }

  for i := 3; i*i <= n; i+=2 {
    for n % i == 0 {
      pfs = append(pfs, i)
      n = n / i 
    }
  }

  if n > 2 {
    pfs = append(pfs, n)
  }

  return
}

func SumOfProperDivisors(n int) int {
  pfs := PrimeFactors(n)

  // {prime: prime_exponents}
  m := map(map[int]int)
  for _, prime := range pfs {
    _, ok := m[prime]
    if ok {
      m[prime] += 1
    } else {
      m[prime] -= 1
    }
  }

  ret := 1

  for prime, exponents := range m {
    ret *= (int(math.Pow(prime, exponents+1)-1) / (prime-1)
  }
  return ret - n

}












