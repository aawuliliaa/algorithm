package main

import "fmt"

func sort(arr []int) []int {

	for i:=0;i< len(arr);{
		lastSwappedIndex :=0
		for j :=0;j<len(arr)-i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
				lastSwappedIndex =j+1
			}
		}
		//i表示已经有多少个排好序了
		i = len(arr)-lastSwappedIndex
	}
	return arr

}

func main()  {
	arr :=[]int{4,5,6,7,8,9}
	arr = sort(arr)
	fmt.Println(arr)

}
