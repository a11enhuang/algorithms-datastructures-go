package sort

import "fmt"

// 冒泡排序法
// 每一轮排序可以确定其中一个最值(最大值或最小值)

func bubbleSort(arr []int) {
	fmt.Println("正在使用冒泡排序法对arr进行排序.arr=", arr)
	//step1: 获取arr的长度
	length := len(arr)
	for i := 0; i < length; i++ {
		//有两种冒泡方案: 1:从前往后冒,每次冒的都是最大值 2:从后往前冒，每次冒的都是最小值
		//本案例中采用的是第二种方案，从后往前冒
		//所以本案例的思路是比较两个相邻索引的元素大小，如果后值小于前值，则这两个值需要交换位置.
		//这样每一次循环下来，最小的值始终都能跑到最前面去
		for j := length - 1; j > i; j-- {
			v1 := arr[j]
			v2 := arr[j-1]
			//如果后值小于前值，
			if v1 < v2 {
				arr[j], arr[j-1] = v2, v1
			}
		}
		fmt.Printf("step %d: arr: %v \n", i+1, arr)
	}
	fmt.Printf("result: %v \n", arr)
}
