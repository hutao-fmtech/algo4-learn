package algo

func swap(items []int, i, j int) {
	if i == j {
		return
	}
	tmp := items[i]
	items[i] = items[j]
	items[j] = tmp
}

// 1. 冒泡排序
// 算法描述
// 1.1 index 从 0 到 (待排序数组的长度 -2) ，依次比较相邻的一对元素大小，如果items[j] < items [j+1], 交换两个元素的值。
// 1.2 当 j = (待排序数组的长度 -2) 时，一轮排序完成，此时数组最后一个元素为最大值，待排序的数组长度减一。
// 1.3 重复上述步骤，已完成排序轮数，可以用 i 记录，完成一轮后，i 自加 1。
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
		swap(items, i, minIndex)
	}

	return items
}

// 3. 插入排序

func InsertSort(items []int) []int {

	// 将 a[i] 插入到 a[i-1], a[i-2], a[i-3] ... 中
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

// 快速排序
// 1. 使用左右指针（i，j），分别寻找大于，小于 pivot 的位置。
// 2. 左指针的值小于等于 pivot 时，左指针++（i++），往右移动，直至值大于 pivot。
// 3. 右指针的值大于 pivot 值时，右指针-- （j--），往左移动，直至值小于 pivot。
// 4. 当出现 a[i] > pivot, a[j] < pivot 是时交换 i，j，此时依然满足 a[i] <= pivot , a[j] > pivot, 2，3 步骤继续。
// 5. 当 j <= i 时，整个数组扫描完成，交换 j 值 和 pivot 值。

// 一定 return j 指针。此时分区为 [lo，j-1],[j],[j+1,hi]
// 再次 partition 区间，[lo，j-1] ,[j+1,hi]
func partition(items []int, lo, hi int) int {

	i := lo
	j := hi + 1
	v := items[lo]

	for {

		for {
			i++
			if items[i] > v || i >= hi {
				break
			}
		}

		for {
			j--
			if items[j] < v || j <= lo {
				break
			}
		}

		if j <= i {
			break
		}
		// 此时 i 找一个一个比 pivot 大的，j 找一个比 pivot 小的，两者交换
		// 注意：交换后，i 位置的数值比 pivot 小，而 j 位置的 数值又比 pivot 大，进而又可以开始第二轮寻宝。
		swap(items, i, j)

	}

	swap(items, lo, j)
	return j
}

func QuickSort(items []int, lo, hi int) {

	if hi <= lo {
		return
	}

	pivot := partition(items, lo, hi)
	QuickSort(items, lo, pivot-1)
	QuickSort(items, pivot+1, hi)

}
