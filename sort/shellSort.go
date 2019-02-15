package main

import "fmt"

func main() {

	nums := []int{30, 29, 50, 2, 42, 60}

	fmt.Print("排序前")
	fmt.Println(nums)

	shellSort2(nums)

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

//希尔排序
func shellSort(array []int) {

	n := len(array)
	if n < 2 {
		return
	}

	key := n / 2

	for key > 0 {
		for i := key; i < n; i++ {
			j := i
			for j >= key && array[j] < array[j-key] {
				array[j], array[j-key] = array[j-key], array[j]
				j = j - key
			}
		}
		key = key / 2
	}

	return

}

func shellSort2(arr []int) {

	length := len(arr)
	if length < 2 {
		return
	}

	inc := length / 2 //先取一半当步长

	for inc > 0 { //或者是 >= 1

		var i, j int
		for i = inc; i < length; i++ {
			temp := arr[i]
			for j = i; j >= inc && arr[j-inc] > temp; j -= inc { //减成步长
				arr[j] = arr[j-inc]
			}
			arr[j] = temp
		}

		inc = inc / 2
	}

	return

}
