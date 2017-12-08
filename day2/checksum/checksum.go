package day2

import ()

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinMax(a []int) (min, max int) {
	min, max = a[0], a[0]

	for _, v := range a {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func RowSum(r []int) int {
	a, b := MinMax(r)
	return a + b
}

func CodeOne(str string) int {
	return 0

}
func CodeTwo() {
}
