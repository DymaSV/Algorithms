package countingsort

func CountingSort(arr []int) []int {
	return getSortedArray(prefixSum(getMap(arr)), arr)
}

func getMap(arr []int) (map[int]int, int) {
	var arrCount = make(map[int]int)
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
		arrCount[v] += 1
	}
	return arrCount, max
}

func prefixSum(m map[int]int, max int) map[int]int {
	var arrCount = make(map[int]int)
	for i := 0; i <= max; i++ {
		if i != 0 {
			arrCount[i] = arrCount[i-1] + m[i]
		} else {
			arrCount[i] = m[i]
		}
	}
	for i := max; i >= 0; i-- {
		arrCount[i] = arrCount[i-1]
	}
	arrCount[0] = 0
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
