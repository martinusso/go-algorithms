package selection_sort

func Sort(arr []int) []int {
	l := len(arr)
	for i := 0; i < l; i++ {
		min := i
		for j := i; j < l; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
