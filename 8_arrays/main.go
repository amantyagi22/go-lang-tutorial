package main

import "fmt"

func main() {
	// var nums [4]int;
	// nums[0] = 78;
	// declare in one line
	// nums := [4]int{1,2,3,4}
	// there is no need to specify the size of the array
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(nums)
}
