// cerner_2tothe5th_2021

package main

import "fmt"

// Function to find the sum of two numbers in a list that equals the target
func twoSums(nums []int, target int) []int {
	// Loop through entire slice as i
	for i, num := range nums {
		// Loop from i + 1 to the end as j
		for j, next := range nums[i+1:] {
			j += i + 1
			// Test if the two numbers equals target
			if num + next == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	// variables
	nums := []int{2, 7, 11, 15}
	target := 18

	// Print output
	fmt.Println(twoSums(nums, target))
}
