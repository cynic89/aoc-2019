package util

import "math"

type Coordinates struct {
	X, Y int
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func LCM(x, y int) int {
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	lim := x
	if y < x {
		lim = y
	}
	for i := lim; i >= 2; i-- {
		if x%i == 0 && y%i == 0 {
			return i
		}
	}

	return 1

}

func (u Coordinates) Angle(v Coordinates) float64 {
	delta := Coordinates{v.X - u.X, u.Y - v.Y}
	rad := math.Atan2(float64(delta.Y), float64(delta.X))
	return rad

}

func LCMOf(nos []int64) int64 {
	max := max(nos)
	factor := 2
	var lcm int64
	for {
		lcm = max * int64(factor)
		for i, n := range nos {
			if lcm%n == 0 {
				if i == (len(nos) - 1) {
					return lcm
				}
			} else {
				factor++
				break
			}
		}

	}
}

func max(nos []int64) int64 {
	var max int64 = 0
	for _, n := range nos {
		if n > max {
			max = n
		}
	}
	return max
}
