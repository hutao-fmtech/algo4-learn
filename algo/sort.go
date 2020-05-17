package algo

func swap(items []int, i, j int) {
	tmp := items[i]
	items[i] = items[j]
	items[j] = tmp
}

// 1. 冒泡排序
// 算法描述
// 1.1  从 0 到 (待排序数组的长度 -2) ，依次比较相邻的一对元素大小，如果items[j] < items [j+1], 交换两个元素的值。
// 1.2 当 j = (待排序数组的长度 -2) 时，一轮排序完成，此时数组最后一个元素为最大值，待排序的数组长度减一。
// 1.3 重复上述步骤，已完成排序轮数，可以用i 记录，完成一轮后，i 自加 1。
func BubbleSort(items []int) []int {

	for i, _ := range items {
		for j := 0; j < len(items)-i-1; j++ {
			if items[j] > items[j+1] {
				tmp := items[j]
				items[j] = items[j+1]
				items[j+1] = tmp
			}
		}
	}
	return items
}

// 2. 选择排序
// 算法描述
// 2.1 将待排序数组分成两类，已排序数组和未排序数组，初始已排序数组长度为1，下标为0。
// 2.2 从已排序数组中最后一位 和 待排序的数组中选择一个最小值，记录下标
// 2.3 交换已完成数组最后一位 和 最小值位。
// 2.4 已完成排序数组长度加1，重复上述步骤。
func SelectSort(items []int) []int {
	for i, _ := range items {
		minIndex := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[minIndex] {
				minIndex = j
			}
		}
		tmp := items[i]
		items[i] = items[minIndex]
		items[minIndex] = tmp
	}

	return items
}

func SelectSortV2(items []int) []int {
	for i := 0; i < len(items)-1; i++ {
		minIndex := i + 1
		for j := minIndex + 1; j < len(items); j++ {
			if items[j] < items[minIndex] {
				minIndex = j
			}
		}
		if items[i] > items[minIndex] {
			tmp := items[i]
			items[i] = items[minIndex]
			items[minIndex] = tmp
		}
	}

	return items
}

func SelectSortV3(items []int) []int {
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[i] {
				swap(items, j, i)
			}
		}
	}
	return items
}

// 3. 插入排序

func InsertSort(items []int) []int {
	for i := 1; i < len(items); i++ {
		processedIndex := i - 1
		current := items[i]

		// 注意 processedIndex 可能会被自减到 -1
		for ; processedIndex >= 0 && items[processedIndex] > current; {
			items[processedIndex+1] = items[processedIndex]
			processedIndex--
		}
		// 注意 index 要加 1
		items[processedIndex+1] = current
	}
	return items
}

func InsertSortV2(items []int) []int {

	for i := 1; i < len(items); i++ {
		for j := i; j > 0 && items[j] < items[j-1]; j-- {
			swap(items, j, j-1)
		}
	}

	return items
}

// 4. 希尔排序
func ShellSort(items []int) []int {

	l := len(items)
	h := 1
	for h < l/3 {
		h = 3*h + 1
	}

	for ; h >= 1; h = h / 3 {
		for i := h; i < l; i++ {
			for j := i; j >= h && items[j] > items[j-h]; j -= h {
				swap(items, j, j-h)
			}
		}
	}

	return items
}

func ShellSortV2(items []int) []int {
	l := len(items)
	for gap := l / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < l; i++ {

			processedIndex := i - gap
			current := items[i]

			// 注意 processedIndex 可能会被自减到 -gap
			for ; processedIndex >= 0 && items[processedIndex] > current; {
				items[processedIndex+gap] = items[processedIndex]
				processedIndex -= gap
			}
			// 注意 index 要加 gap
			items[processedIndex+gap] = current
		}
	}
	return items
}

func ShellSortV3(items []int) []int {
	l := len(items)
	for gap := l / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < l; i++ {
			for j := i; j >= gap && items[j] < items[j-gap]; j -= gap {
				swap(items, j, j-gap)
			}
		}
	}
	return items
}

func merge(items, aux []int, lo, mid, hi int) {

	for k := lo; k <= hi; k++ {
		aux[k] = items[k]
	}

	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		// 左边已取完，只能取右边
		if i > mid {
			items[k] = aux[j]
			j++

		} else if j > hi { // 右边已取完，只能取左边
			items[k] = aux[i]
			i++

		} else if aux[i] < aux[j] { // 左边比右边小，取左边
			items[k] = aux[i]
			i++
		} else { // 左边比右边大，取右边
			items[k] = aux[j]
			j++
		}
	}
}

// 5。归并排序，自顶向下递归

func mergeSort(items, aux []int, lo, hi int) {

	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort(items, aux, lo, mid)
	mergeSort(items, aux, mid+1, hi)
	merge(items, aux, lo, mid, hi)
}

func MergeSort(items []int) []int {
	aux := make([]int, len(items))
	mergeSort(items, aux, 0, len(items)-1)
	return items
}

func getMin(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

func MergeSortBU(items []int) []int {
	l := len(items)
	aux := make([]int, len(items))
	for sz := 1; sz < l; sz = sz + sz {
		for lo := 0; lo < l-sz; lo += sz + sz {
			merge(items, aux, lo, lo+sz-1, getMin(lo+sz+sz-1, l-1))
		}
	}
	return items
}
