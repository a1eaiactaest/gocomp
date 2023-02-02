package maths

import "math"

func MaxInt(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func MinInt(nums ...int) int {
	ret := nums[0]

	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func AbsoluteInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SumIntSlice(nums []int) int {
	var acc int = 0
	for _, v := range nums {
		acc += v
	}
	return acc
}

func ProdIntSlice(nums []int) int {
	var ret int = 1
	for _, v := range nums {
		ret *= v
	}
	return ret
}


func PythagoreanDistance(x1, y1, x2, y2 int) float64 {
	x := float64(x1 - x2)
	y := float64(y1 - y2)

	squares := math.Pow(x, 2) + math.Pow(y, 2)
	return math.Sqrt(squares)
}

func ManhattanDistance(x1, y1, x2, y2 int) int {
	x := x1 - x2
	y := y1 - y2

	if x < 0 {
		x *= -1
	}

	if y < 0 {
		y *= -1
	}

	return x + y
}


func FactorInt(x int) (factors []int) {

  factors = append(factors, 1)

  for (x % 2 == 0) {
    factors = append(factors, 2)
    x /= 2
  }

  for d := 3; d*d <= x; d += 2 {
    for (x % d == 0) {
      factors = append(factors, d)
      x /= d
    }
  }

  if x >= 1 {
    factors = append(factors, x)
  }

  factors = append(factors, x)
  return factors
}
