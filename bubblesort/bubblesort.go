package bubblesort

func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	index := 1
	changed := 0
	first := 0
	for changed == 0 {
		if arr[index-1] > arr[index] {
			arr[index], arr[index-1] = arr[index-1], arr[index]
			changed++
		}
		if changed > 0 {
			changed = 0
			first = 0
		} else if changed == 0 && index == len(arr)-1 {
			if first != 0 {
				changed = 1
			}
			first++
		}
		index++
		if index == len(arr) {
			index = 1
		}
	}
	return arr
}
