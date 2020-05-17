package test

import (
	"../algo"
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	inputItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	outputItems := []int{0, 1, 2, 3, 4, 5, 6, 7}

	res := algo.BubbleSort(inputItems)
	if !reflect.DeepEqual(inputItems, outputItems) {
		fmt.Println(res)
		t.Error()
	}
}

func TestSelectSort(t *testing.T) {
	inputItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	outputItems := []int{0, 1, 2, 3, 4, 5, 6, 7}

	res := algo.SelectSort(inputItems)
	if !reflect.DeepEqual(inputItems, outputItems) {
		fmt.Println(res)
		t.Error()
	}
}

func TestSelectSortV2(t *testing.T) {
	inputItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	outputItems := []int{0, 1, 2, 3, 4, 5, 6, 7}

	res := algo.SelectSortV2(inputItems)
	if !reflect.DeepEqual(inputItems, outputItems) {
		fmt.Println(res)
		t.Error()
	}
}

func TestSelectSortV3(t *testing.T) {
	inputItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	outputItems := []int{0, 1, 2, 3, 4, 5, 6, 7}

	res := algo.SelectSortV3(inputItems)
	if !reflect.DeepEqual(inputItems, outputItems) {
		fmt.Println(res)
		t.Error()
	}
}

func TestInsertSort(t *testing.T) {
	inputItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	outputItems := []int{0, 1, 2, 3, 4, 5, 6, 7}

	res := algo.InsertSort(inputItems)
	if !reflect.DeepEqual(inputItems, outputItems) {
		fmt.Println(res)
		t.Error()
	}
}

func TestInsertSortV2(t *testing.T) {
	inputItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	outputItems := []int{0, 1, 2, 3, 4, 5, 6, 7}

	res := algo.InsertSortV2(inputItems)
	if !reflect.DeepEqual(inputItems, outputItems) {
		fmt.Println(res)
		t.Error()
	}
}

func TestShellSortV2(t *testing.T) {
	inputItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	outputItems := []int{0, 1, 2, 3, 4, 5, 6, 7}

	res := algo.ShellSortV2(inputItems)
	if !reflect.DeepEqual(inputItems, outputItems) {
		fmt.Println(res)
		t.Error()
	}
}
func TestShellSortV3(t *testing.T) {
	inputItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	outputItems := []int{0, 1, 2, 3, 4, 5, 6, 7}

	res := algo.ShellSortV3(inputItems)
	if !reflect.DeepEqual(inputItems, outputItems) {
		fmt.Println(res)
		t.Error()
	}
}

func TestMergeSort(t *testing.T) {
	inputItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	outputItems := []int{0, 1, 2, 3, 4, 5, 6, 7}

	res := algo.MergeSort(inputItems)
	if !reflect.DeepEqual(inputItems, outputItems) {
		fmt.Println(res)
		t.Error()
	}
}

func TestMergerSortBU(t *testing.T) {
	inputItems := []int{2, 1, 0, 6, 3, 5, 7, 4}
	outputItems := []int{0, 1, 2, 3, 4, 5, 6, 7}

	res := algo.MergeSortBU(inputItems)
	if !reflect.DeepEqual(inputItems, outputItems) {
		fmt.Println(res)
		t.Error()
	}
}
