package insertingsort

func InsertingSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if i > 0 {
			if arr[i] < arr[i-1] {
				element := arr[i]
				arr = append(arr[0:i], arr[i+1:]...)
				for j := i - 1; j >= 0; j-- {
					if element > arr[j] {
						arr = insert(arr, element, j+1)
						i = 0
						j = 0
					} else if j == 0 {
						arr = insert(arr, element, j)
						i = 0
					}
				}
			}
		}
	}
	return arr
}

func insert(array []int, element int, i int) []int {
	return append(array[:i], append([]int{element}, array[i:]...)...)
}
