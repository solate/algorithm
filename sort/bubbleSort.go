package main

import "fmt"

func main() {

	nums := []int{30, 29, 50, 2, 42, 60}

	fmt.Print("排序前")
	fmt.Println(nums)

	bubbleSort(nums)

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

//冒泡排序
func bubbleSort(arr []int) {

	length := len(arr)
	if length < 2 {
		return
	}

	for i := 0; i < length-1; i++ { //6 个数 比较5趟
		fmt.Println(i)
		for j := 0; j < length-i-1; j++ { //每趟从5次递减
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}

}
