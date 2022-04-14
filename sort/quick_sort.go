package sort

import "fmt"

// 快速排序法
// 快速排序法是一种以分而治之为思想的排序算法。
// 实现思路是首先找到一个基准数，取数组的左右两个指针，左指针的值必须小于等于基准数，右指针的值必须大于等于基准数。
// 然后首先保持左指针不动，右指针一直向左移动直到右指针值小于基准数，此时交换左右指针的值，左指针的值一定比基准数小。
// 此时保持右指针不动，左指针一直向右移动直到左指针值大于基准数，此时交换左右指针的值，右指针的值一定比基准数小。
// 如此往复，当左右指针相邻时，左指针的值一定小于右指针，然后递归将数组分段，以左指针为界分别对子数组再进行上述排序，最终能保证序列有序。

func quickSort(arr []int) {
	fmt.Println("正在使用快速排序法对arr进行排序.arr=", arr)
	doQuickSort(arr, 0, len(arr)-1)
	fmt.Printf("result: %v \n", arr)
}

func doQuickSort(arr []int, start, last int) {
	if start >= last {
		return
	}
	//默认选择第一个元素作为基准数
	benchmark := arr[start]
	left, right := start, last
	fmt.Printf("开始索引:%d\t结束索引:%d\t基准数:%d\t数组内容:%v\n", left, right, benchmark, arr)
	for left < right {
		//右指针一直往左移动，直到找到一个小于基准数的值.
		for arr[right] > benchmark && right > left {
			right--
		}
		//循环结束右两种情况，要么是直接执行完了，要么是找到了第一个右值小于基准数的.
		if arr[right] <= benchmark && right > left {
			//交换左右值
			arr[left], arr[right] = arr[right], arr[left]
			//因为此时的左值一定比基准数小，所以左指针向右移动一位
			left++
		}
		//左指针一直向右移动，直到找到一个大于基准数的值
		for arr[left] < benchmark && left < right {
			left++
		}

		if arr[left] >= benchmark && left < right {
			//交换左右值
			arr[left], arr[right] = arr[right], arr[left]
			//此时右值一定比基准数小，所以可以向左移动一位
			right--
		}
		if start < left {
			doQuickSort(arr, start, left-1)
		}
		if left < last {
			doQuickSort(arr, left+1, last)
		}
	}
}
