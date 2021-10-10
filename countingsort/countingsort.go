package countingsort

import "math"

func CountingSort(arr []int) []int {
	return getSortedArray(prefixSum(getMap(arr)), arr)
}

func getMap(arr []int) (map[int]int, int, int) {
	var arrCount = make(map[int]int)
	max := math.MinInt32
	min := math.MaxInt32
	for _, v := range arr {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
		arrCount[v] += 1
	}
	return arrCount, max, min
}

func prefixSum(m map[int]int, max int, min int) map[int]int {
	var arrCount = make(map[int]int)
	for i := min; i <= max; i++ {
		if i != min {
			arrCount[i] = arrCount[i-1] + m[i]
		} else {
			arrCount[i] = m[i]
		}
	}

	// Shift to the right
	for i := max; i >= min; i-- {
		arrCount[i] = arrCount[i-1]
	}
	arrCount[min] = 0
	return arrCount
}

func getSortedArray(m map[int]int, arr []int) []int {
	var result = make([]int, len(arr))
	for _, v := range arr {
		result[m[v]] = v
		m[v] = m[v] + 1
	}
	return result
}
