package binsearchloop

//Binary search Algorithm wich use a loop
func BinSearchLoop(array *[]int, searchValue int) int {
	index := len(*array) / 2
	low := 0
	hight := 0
	searchOneByOne := false
	arr := *array
	for index > 0 {
		if index == 1 {
			if arr[index] == searchValue {
				return index
			}
			index = 0
			searchOneByOne = true
		}
		if !searchOneByOne {
			if arr[index] == searchValue {
				return index
			} else if arr[index] > searchValue {
				//to left side of array
				low = 0
				hight = index
			} else if arr[index] < searchValue {
				//to right side of array
				low = index
				hight = len(*array)
			}
			index = index / 2
		} else {
			for i := low; i < hight; i++ {
				if arr[i] == searchValue {
					return i
				}
			}
		}

	}
	return -1
}
