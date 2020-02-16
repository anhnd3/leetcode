package main

import (
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
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

// === Problem 5 ===
func longestPalindrome(s string) string {
	if s == "" || len(s) < 1 {
		return ""
	}

	var start, end int
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		lenDiff := max(len1, len2)
		if lenDiff > end-start {
			start = i - (lenDiff-1)/2
			end = i + lenDiff/2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

// === End Problem 5 ===

// === Problem 6 ===
// func convert(s string, numRows int) string {
// 	if len(s) == 1 || numRows == 1 {
// 		return s
// 	}
// 	var result string
// 	matrix := make([][]string, numRows)
// 	matrixCol := len(s)/2 + 1
// 	for i := 0; i < numRows; i++ {
// 		matrix[i] = make([]string, matrixCol)
// 	}

// 	var row, col int
// 	var incr bool

// 	for i := 0; i < len(s); i++ {
// 		matrix[row][col] = string(s[i])
// 		if row == 0 {
// 			incr = true
// 		} else if row == numRows-1 {
// 			incr = false
// 		}

// 		if incr {
// 			row++
// 		} else {
// 			row--
// 			col++
// 		}
// 	}

// 	for i := 0; i < numRows; i++ {
// 		for j := 0; j < matrixCol; j++ {
// 			result += matrix[i][j]
// 		}
// 	}

// 	return result
// }

func convert(s string, numRows int) string {
	if s == "" || numRows == 1 || len(s) <= numRows {
		return s
	}
	zigzagPeriod := 2*numRows - 2
	stringLen := len(s)
	var output strings.Builder

	for i := 0; i < numRows; i++ {
		output.WriteByte(s[i])
		period1 := zigzagPeriod - i*2
		period2 := zigzagPeriod - period1
		pointer := i

		for pointer < stringLen {
			pointer += period1
			if pointer < stringLen && period1 != 0 {
				output.WriteByte(s[pointer])
			}
			pointer += period2
			if pointer < stringLen && period2 != 0 {
				output.WriteByte(s[pointer])
			}
		}
	}
	return output.String()
}

// === End Problem 6 ===

// === Problem 7 ===
func reverse(x int) int {
	if x > 0 && x < 10 {
		return x
	}

	var result, sign int64

	sign = 1
	if x < 0 {
		sign = -1
		x = -x
	}

	y := int64(x)
	for y > 0 {
		result = result*10 + y%10
		y = y / 10
	}

	result *= sign

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return int(result)
}

// === End Problem 7 ===

// === Problem 8 ===
func myAtoi(str string) int {
	strTrim := strings.Trim(str, " ")
	re := regexp.MustCompile(`^[-|+|\d]\d*`)
	strResult := string(re.Find([]byte(strTrim)))

	result, _ := strconv.ParseInt(strResult, 10, 64)
	if result > math.MaxInt32 {
		return math.MaxInt32
	}

	if result < math.MinInt32 {
		return math.MinInt32
	}

	return int(result)
}

// === End Problem 8 ===

// === Problem 9 ===
func isPalindrome(x int) bool {
	strX := strconv.Itoa(x)
	var left int
	var right int = len(strX) - 1

	for left < right {
		if strX[left] == strX[right] {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

// === Problem 10 ===
func isMatch(s string, p string) bool {
	p += "$"
	result, _ := regexp.MatchString(p, s)

	if result == true {
		re := regexp.MustCompile(p)
		strRexp := string(re.Find([]byte(s)))
		if strRexp != s {
			return false
		}
		return true
	}

	return false
}

// === End Problem 10 ===

// === Problem 11 ===
func maxArea(height []int) int {
	var left, right, max int
	right = len(height) - 1

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

// === End Problem 11 ===

// === Problem 12 ===
func intToRoman(num int) string {
	var result string
	if num < 1 || num > 3999 {
		return ""
	}

	for num > 0 {
		if num >= 1000 {
			repeat := num / 1000
			for i := 0; i < repeat; i++ {
				result += "M"
			}
			num -= repeat * 1000
		} else if num >= 900 {
			result += "CM"
			num -= 900
		} else if num >= 500 {
			result += "D"
			num -= 500
		} else if num >= 400 {
			result += "CD"
			num -= 400
		} else if num >= 100 {
			repeat := num / 100
			for i := 0; i < repeat; i++ {
				result += "C"
			}
			num -= repeat * 100
		} else if num >= 90 {
			result += "XC"
			num -= 90
		} else if num >= 50 {
			result += "L"
			num -= 50
		} else if num >= 40 {
			result += "XL"
			num -= 40
		} else if num >= 10 {
			repeat := num / 10
			for i := 0; i < repeat; i++ {
				result += "X"
			}
			num -= repeat * 10
		} else if num >= 9 {
			result += "IX"
			num -= 9
		} else if num >= 5 {
			result += "V"
			num -= 5
		} else if num >= 4 {
			result += "IV"
			num -= 4
		} else if num >= 1 {
			for i := 0; i < num; i++ {
				result += "I"
			}
			num -= num
		}
	}
	return result
}

// === End Problem 12 ===

// === Problem 13 ===
func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	mapRoman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000}
	var result int
	for idx, ch := range s {
		var curr, next int
		curr = mapRoman[string(ch)]
		if idx < len(s)-1 {
			next = mapRoman[string(s[idx+1])]
		}

		if curr < next {
			result -= curr
		} else {
			result += curr
		}
	}
	return result
}

// === End Problem 13 ===

// === Problem 14 ===
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var prefix string

	pattern := strs[0]
	for _, ch := range pattern {
		tmpPrefix := prefix + string(ch)
		for _, str := range strs {
			if !strings.HasPrefix(str, tmpPrefix) {
				return prefix
			}
		}
		prefix = tmpPrefix
	}
	return prefix
}

// === End Problem 14 ===

// === Problem 15 ===
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			twoSumUniqueParis(nums, nums[i], i+1, &res)
		}
	}

	return res
}

func twoSumUniqueParis(nums []int, n, i int, res *[][]int) {
	l := i
	h := len(nums) - 1

	for l < h {
		if nums[l]+nums[h]+n == 0 {
			*res = append(*res, []int{n, nums[l], nums[h]})
			for l < h && nums[l] == nums[l+1] {
				l++
			}
			for l < h && nums[h] == nums[h-1] {
				h--
			}
			l++
			h--
		} else if nums[l]+nums[h]+n > 0 {
			h--
		} else {
			l++
		}
	}
}

// === End Problem 15 ===

func main() {
	threeSum([]int{-1, 0, 1, 2, -1, -4})
}
