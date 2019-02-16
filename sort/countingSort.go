package main

import "fmt"

func main() {

	//nums := []int{30, 29, 50, 2, 42, 60}

	nums := []int{1, 4, 2, 4, 6, 9, 8, 2}

	fmt.Print("排序前")
	fmt.Println(nums)

	tmp := countingSort(nums)

	fmt.Print("排序后")
	fmt.Println(tmp)
}

//打印
func printArray(arr []int) {
	for _, v := range arr {
		fmt.Print(v, ", ")
	}
	fmt.Println()
}

//计数排序
func countingSort(arr []int) []int {

	length := len(arr)
	if length < 2 {
		return arr
	}

	//初始化
	max := getMax(arr)
	fmt.Println("max: ", max)
	var sortArr = make([]int, length)
	var countArr = make([]int, max+1) // max+1 是为了防止 countsArr[] 计数时溢出

	// 元素计数
	for _, v := range arr {
		countArr[v]++
	}

	fmt.Println("count arr :", countArr)

	// 统计独特数字个数并累加
	for i := 1; i <= max; i++ {
		countArr[i] += countArr[i-1]
	}

	fmt.Println("count arr add:", countArr)

	for _, kv := range arr {
		sortArr[countArr[kv]-1] = kv
		countArr[kv]--
		fmt.Println("==>", sortArr)
	}

	return sortArr
}

func getMax(arr []int) (max int) {
	max = arr[0]
	for _, v := range arr {
		if max < v {
			max = v
		}
	}
	return
}
