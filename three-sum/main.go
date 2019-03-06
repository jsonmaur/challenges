package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func generateId(nums []int) string {
	var id strings.Builder

	for _, v := range nums {
		id.WriteString(strconv.Itoa(v))
	}

	return id.String()
}

func binarySearch(nums []int, num int) (int, bool) {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2

		if nums[mid] < num {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if start == len(nums) || nums[start] != num {
		return -1, false
	} else {
		return start, true
	}
}

func ThreeSum(nums []int) [][]int {
	var solutions [][]int
	solutionsMap := map[string]bool{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			pairSum := nums[i] + nums[j]
			index, found := binarySearch(nums, -pairSum)

			if found && index != i && index != j {
				solution := []int{nums[i], nums[j], nums[index]}
				sort.Ints(solution)

				id := generateId(solution)
				if !solutionsMap[id] {
					solutionsMap[id] = true
					solutions = append(solutions, solution)
				}
			}
		}
	}

	return solutions
}

func main() {
	fmt.Println(ThreeSum([]int{-1, 0, 1, 2, -1, -4}))
}
