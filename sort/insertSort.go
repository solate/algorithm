package main

import "fmt"

func main() {

	nums := []int{30, 29, 50, 2, 42, 60}

	fmt.Print("排序前")
	fmt.Println(nums)

	insertSort(nums)

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

//插入排序
func insertSort(arr []int) {

	length := len(arr)
	if length < 2 {
		return
	}

	var i, j int
	for i = 1; i < length; i++ {
		temp := arr[i]
		for j = i - 1; j >= 1 && arr[j] > temp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp

	}

	return

}

//第二种互换位置的方式
func insertSort2(array []int) {
	n := len(array)
	if n < 2 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			} else {
				break
			}
		}
	}
}
