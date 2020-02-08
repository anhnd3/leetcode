package main

import (
	"fmt"
)

// === Problem 1 ===
func twoSum(nums []int, target int) []int {
	mapData := make(map[int]int)

	for i, num := range nums {
		if val, ok := mapData[num]; ok {
			return []int{val, i}
		}
		mapData[target-num] = i
	}

	return []int{}
}

// === End Problem 1 ===

// === Problem 2 ===
type listNode struct {
	Val  int
	Next *listNode
}

func addTwoNumbers(l1 *listNode, l2 *listNode) *listNode {
	var (
		p, q, result, tail *listNode
		carry              int
	)
	p = l1
	q = l2
	for p != nil || q != nil {
		var x, y int
		if p != nil {
			x = p.Val
			p = p.Next

		}
		if q != nil {
			y = q.Val
			q = q.Next
		}
		sum := x + y + carry
		carry = sum / 10

		currNode := &listNode{Val: sum % 10}
		if result == nil {
			result = currNode
			tail = result
		} else {
			tail.Next = currNode
			tail = currNode
		}
	}
	if carry > 0 {
		tail.Next = &listNode{Val: carry}
	}
	return result
}

// === End Problem 2 ===

// === Problem 3 ===
func lengthOfLongestSubstring(s string) int {
	var length, curr int

	mapChar := make(map[rune]int)
	for pos, char := range s {
		if _, ok := mapChar[char]; !ok {
			curr++
		} else {
			curr = min(pos-mapChar[char], curr+1)
		}
		length = max(length, curr)
		mapChar[char] = pos
	}

	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// === End Problem 3 ===

// === Problem 4 ===
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int

	for len(nums1) > 0 && len(nums2) > 0 {
		if nums1[0] < nums2[0] {
			nums = append(nums, nums1[0])
			nums1 = nums1[1:]
		} else {
			nums = append(nums, nums2[0])
			nums2 = nums2[1:]
		}
	}
	if len(nums1) == 0 {
		nums = append(nums, nums2...)
	} else if len(nums2) == 0 {
		nums = append(nums, nums1...)
	}

	if len(nums)%2 != 0 {
		middle := (len(nums) - 1) / 2
		return float64(nums[middle])
	}
	middle := len(nums) / 2
	return float64(nums[middle]+nums[middle-1]) / 2
}

// === End Problem 4 ===

func main() {
	fmt.Println(findMedianSortedArrays([]int{1}, []int{1}))
}
