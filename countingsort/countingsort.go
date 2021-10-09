package countingsort

import "fmt"

func CountSort(arr []int) {
	// fmt.Println(getMap(arr))
	fmt.Println(prefixSum(getMap(arr)))
}

func getMap(arr []int) map[int]int {
	var arrCount = make(map[int]int)
	for _, v := range arr {
		arrCount[v] += 1
	}
	return arrCount
}

func prefixSum(m map[int]int) map[int]int {
	fmt.Println(m)
	ele := 0
	for i, v := range m {
		if i != 0 {
			m[i] = ele + m[i]
		}
		ele = v
	}
	fmt.Println(m)
	return m
}
