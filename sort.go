package sort

import "math/rand"

// Bubble sort algorithm
func Bubble(slice []int) {
	swapped := true

	for swapped {
		swapped = false
		for idx := 0; idx < len(slice)-1; idx++ {
			if slice[idx] > slice[idx+1] {
				slice[idx], slice[idx+1] = slice[idx+1], slice[idx]
				swapped = true
			}
		}
	}
}

// Selection sort algorithm
func Selection(slice []int) {
	for i := 0; i < len(slice); i++ {
		min := i

		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[min] {
				min = j
			}
		}

		slice[i], slice[min] = slice[min], slice[i]
	}
}

// Insertion sort algorithm
func Insertion(slice []int) {
	i, j := 0, 0

	for i = 1; i < len(slice); i++ {
		for j = 0; j < i; j++ {
			if slice[j] > slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
}

// part of Merge sort
func mer(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

// Merge sort algorithm
func Merge(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}

	middle := len(slice) / 2

	left := Merge(slice[:middle])
	right := Merge(slice[middle:])

	return mer(left, right)
}

// Quick sort algorithm
func Quick(slice []int) []int {

	if len(slice) <= 1 {
		return slice
	}

	median := slice[rand.Intn(len(slice))]

	low := make([]int, 0, len(slice))
	high := make([]int, 0, len(slice))
	middle := make([]int, 0, len(slice))

	for _, item := range slice {
		switch {
		case item < median:
			low = append(low, item)
		case item == median:
			middle = append(middle, item)
		case item > median:
			high = append(high, item)
		}
	}

	low = Quick(low)
	high = Quick(high)

	low = append(low, middle...)
	low = append(low, high...)

	return low
}

// part of Heap sort
func sift(slice []int, i int, sliceLen int) []int {
	done := false
	maxChild := 0

	for i*2+1 < sliceLen && !done {
		if i*2+1 == sliceLen-1 {
			maxChild = i*2 + 1
		} else if slice[i*2+1] > slice[i*2+2] {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}

		if slice[i] < slice[maxChild] {
			slice[i], slice[maxChild] = slice[maxChild], slice[i]
			i = maxChild
		} else {
			done = true
		}
	}

	return slice
}

// Heap sort algorithm
func Heap(slice []int) {
	i := 0

	for i = len(slice)/2 - 1; i >= 0; i-- {
		slice = sift(slice, i, len(slice))
	}

	for i = len(slice) - 1; i >= 1; i-- {
		slice[0], slice[i] = slice[i], slice[0]
		slice = sift(slice, 0, i)
	}
}

// Counting sort algorithm
func Counting(slice []int) {
	k := 0
	if len(slice) == 0 {
		k = 1
	}

	n := slice[0]

	for _, v := range slice {
		if v > n {
			n = v
		}
	}

	k = n + 1

	sliceCounts := make([]int, k)

	for i := 0; i < len(slice); i++ {
		sliceCounts[slice[i]]++
	}

	for i, j := 0, 0; i < k; i++ {
		for {
			if sliceCounts[i] > 0 {
				slice[j] = i
				j++
				sliceCounts[i]--
				continue
			}
			break
		}
	}

}

// Shell sort algorithm
func Shell(slice []int) {
	for d := int(len(slice) / 2); d > 0; d /= 2 {
		for i := d; i < len(slice); i++ {
			for j := i; j >= d && slice[j-d] > slice[j]; j -= d {
				slice[j], slice[j-d] = slice[j-d], slice[j]
			}
		}
	}
}

// Cocktail sort algorithm
func Cocktail(slice []int) {
	for i := 0; i < len(slice)/2; i++ {
		left := 0
		right := len(slice) - 1

		for left <= right {

			if slice[left] > slice[left+1] {
				slice[left], slice[left+1] = slice[left+1], slice[left]
			}

			left++

			if slice[right-1] > slice[right] {
				slice[right-1], slice[right] = slice[right], slice[right-1]
			}

			right--
		}
	}
}

// Comb sort algorithm
func Comb(slice []int) {
	gap := len(slice)

	for {
		if gap > 1 {
			gap = gap * 100 / 124
		}

		for i := 0; ; {

			if slice[i] > slice[i+gap] {
				slice[i], slice[i+gap] = slice[i+gap], slice[i]
			}

			i++

			if i+gap >= len(slice) {
				break
			}
		}

		if gap == 1 {
			break
		}
	}
}

// Gnome sort algorithm
func Gnome(slice []int) {
	i := 1
	for i < len(slice) {
		if slice[i] >= slice[i-1] {
			i++
		} else {
			slice[i], slice[i-1] = slice[i-1], slice[i]

			if i > 1 {
				i--
			}
		}
	}
}
