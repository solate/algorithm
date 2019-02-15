package main

import "fmt"

func main() {

	nums := []int{30, 29, 50, 2, 42, 60}

	fmt.Print("排序前")
	fmt.Println(nums)

	selectSort(nums)

	fmt.Print("排序后")
	fmt.Println(nums)
}

//打印
func printArray(arr []int) {
	for _, v := range arr {
		fmt.Print(v, ", ")
	}
	fmt.Println()
}

//选择排序
func selectSort(arr []int) {

	length := len(arr)
	if length < 2 {
		return
	}

	for i := 0; i < length; i++ {
		// 初始化未排序序列中最小数据数组下标
		min := i

		for j := i + 1; j < length; j++ {
			// 在未排序元素中继续寻找最小元素，并保存其下标
			if arr[j] < arr[min] {
				min = j
			}
		}

		// 将未排序列中最小元素放到已排序列末尾
		arr[i], arr[min] = arr[min], arr[i]

	}

}
