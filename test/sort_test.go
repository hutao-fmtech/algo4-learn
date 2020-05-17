package test

import (
	"../algo"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	digitItems := []int{2, 1, 0, 6, 3, 5, 7, 4}

	res := algo.BubbleSort(digitItems)

	fmt.Println(res)
}

func TestSelectSort(t *testing.T) {

	digitItems := []int{2, 1, 0, 6, 3, 5, 7, 4}

	res := algo.SelectSort(digitItems)

	fmt.Println(res)
}

func TestSelectSortV2(t *testing.T) {

	digitItems := []int{2, 1, 0, 6, 3, 5, 7, 4}

	res := algo.SelectSortV2(digitItems)

	fmt.Println(res)
}

func TestInsertSort(t *testing.T) {
	digitItems := []int{2, 1, 0, 6, 3, 5, 7, 4}

	res := algo.InsertSort(digitItems)

	fmt.Println(res)
}

func TestInsertSortV2(t *testing.T) {
	digitItems := []int{2, 1, 0, 6, 3, 5, 7, 4}

	res := algo.InsertSortV2(digitItems)

	fmt.Println(res)
}

func TestShellSortV2(t *testing.T) {
	digitItems := []int{2, 1, 0, 6, 3, 5, 7, 4}

	res := algo.ShellSortV2(digitItems)

	fmt.Println(res)
}
func TestShellSortV3(t *testing.T) {
	digitItems := []int{2, 1, 0, 6, 3, 5, 7, 4}

	res := algo.ShellSortV3(digitItems)

	fmt.Println(res)
}

func TestMergeSort(t *testing.T) {
	digitItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	res := algo.MergeSort(digitItems)
	fmt.Println(res)
}

func TestMergerSortBU(t *testing.T) {

	digitItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	res := algo.MergeSortBU(digitItems)

	fmt.Println(res)
}
