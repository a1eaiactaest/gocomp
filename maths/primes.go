package maths

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
    if primalityTest(i) {
      primes = append(primes, i)
    }
  }
  return 
}

