/*
Package selection_sort implements a selection sort algorithm.

Selection Sort is an in-place comparison sorting algorithm. It has an O(n2) time complexity, which makes it inefficient on large lists, and generally performs worse than the similar insertion sort. Selection sort is noted for its simplicity and has performance advantages over more complicated algorithms in certain situations, particularly where auxiliary memory is limited.

Following example explains the above steps:

arr[] = 64 25 12 22 11

// Find the minimum element in arr[0...4] and place it at beginning

11 25 12 22 64

// Find the minimum element in arr[1...4] and place it at beginning of arr[1...4]

11 12 25 22 64

// Find the minimum element in arr[2...4] and place it at beginning of arr[2...4]

11 12 22 25 64

// Find the minimum element in arr[3...4] and place it at beginning of arr[3...4]

11 12 22 25 64
```

References: https://en.wikipedia.org/wiki/Selection_sort
*/
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
