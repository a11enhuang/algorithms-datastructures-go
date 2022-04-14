package sort

import "fmt"

// 选择排序法
// 排序思路: 每一轮选择最小的一个数，将此数放在未确认序列的列首

func selectSort(arr []int) {
	fmt.Println("正在使用选择排序法对arr进行排序.arr=", arr)
	length := len(arr)
	for i := 0; i < length; i++ {
		tmp := arr[i]
		//每一轮的i能够确定一个数字排在正确的序列中,所以起始位置可以是i
		for j := i + 1; j < length; j++ {
			v := arr[j]
			if v < tmp {
				arr[j] = tmp
				tmp = v
			}
		}
		arr[i] = tmp
		fmt.Printf("step %d: arr: %v \n", i+1, arr)
	}
	fmt.Printf("result: %v \n", arr)
}
