package sort

import (
	"math/rand"
	"testing"
)

// TestBubbleSort 测试冒泡排序法
func TestBubbleSort(t *testing.T) {
	arr := getArray()
	bubbleSort(arr)
	checkSort(arr, t)
}

func TestSelectSort(t *testing.T) {
	arr := getArray()
	selectSort(arr)
	checkSort(arr, t)
}

func TestQuickSort(t *testing.T) {
	arr := getArray()
	quickSort(arr)
	checkSort(arr, t)
}

// getArray 生成随机数组用于测试
func getArray() []int {
	arr := make([]int, 0, 20)
	for i := 0; i < 20; i++ {
		arr = append(arr, rand.Intn(128))
	}
	return arr
}

func checkSort(arr []int, t *testing.T) {
	maxIndex := len(arr) - 1
	for i := 0; i < maxIndex; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("排序出现错误. [index: %d,value: %d] > [index: %d,value: %d]", i, arr[i], i+1, arr[i+1])
		}
	}
}
