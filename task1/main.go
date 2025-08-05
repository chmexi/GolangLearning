package main

import (
	"sort"
	"strconv"
)

func singleNumber(nums []int) int {
	var result int
	for _, v := range nums {
		result ^= v
	}
	return result
}

func singleNumberNew(nums []int) int {
	mapNum2Count := make(map[int]int)
	for _, num := range nums {
		mapNum2Count[num]++
	}
	for num, count := range mapNum2Count {
		if count == 1 {
			return num
		}
	}
	return 0
}

func isPalindrome(x int) bool {
	var numStr string = strconv.Itoa(x)
	println(numStr)
	leftIndex := 0
	rightIndex := len(numStr) - 1
	for leftIndex <= rightIndex {
		if numStr[leftIndex] == numStr[rightIndex] {
			leftIndex++
			rightIndex--
		} else {
			return false
		}
	}
	return true
}

func isPalindrome2(x int) bool {
	// 特殊情况：
	// 如上所述，当 x < 0 时，x 不是回文数。
	// 同样地，如果数字的最后一位是 0，为了使该数字为回文，
	// 则其第一位数字也应该是 0
	// 只有 0 满足这一属性
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}

	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。
	return x == revertedNumber || x == revertedNumber/10
}

func isValid(s string) bool {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			stack = append(stack, s[i])
			continue
		}
		if s[i] == ')' && stack[len(stack)-1] == '(' {
			stack = stack[:len(stack)-1]
		} else if s[i] == '}' && stack[len(stack)-1] == '{' {
			stack = stack[:len(stack)-1]
		} else if s[i] == ']' && stack[len(stack)-1] == '[' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}

func longestCommonPrefix(strs []string) string {
	for i := 0; i < len(strs[0]); i++ {
		var cur_char byte = strs[0][i]
		for _, s := range strs {
			if len(s) <= i || cur_char != s[i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func plusOne(digits []int) []int {
	// 每一位+1
	for _, digit := range digits {
		digit++
	}
	// 进位
	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] > 9 {
			digits[i] = digits[i] - 10
			digits[i-1]++
		}
	}
	// 看第一位
	if digits[0] > 9 {
		digits[0] = digits[0] - 10
		digits = append([]int{1}, digits...)
	}
	return digits
}

func removeDuplicates(nums []int) int {
	curReplaceIndex := 1 // 当前要替换的位置
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[curReplaceIndex] = nums[i]
			curReplaceIndex++
		}
	}
	return curReplaceIndex
}

func merge(intervals [][]int) [][]int {
	// 合并的条件 aright > bleft
	var res_intervals = [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for _, interval := range intervals {
		// res为空或者不相交，直接添加
		if len(res_intervals) == 0 || res_intervals[len(res_intervals)-1][1] < interval[0] {
			res_intervals = append(res_intervals, interval)
		}
		// 合并这一块
		res_intervals[len(res_intervals)-1][1] = max(res_intervals[len(res_intervals)-1][1], interval[1])
	}
	return res_intervals
}

func twoSum(nums []int, target int) []int {
	num2sub := map[int]int{}
	for cur_index, num := range nums {
		index, exist := num2sub[num]
		if exist {
			return []int{cur_index, index}
		} else {
			num2sub[target-num] = cur_index
		}
	}
	return []int{}
}

func main() {
	println(isPalindrome(121))
	println(isPalindrome(1))
	println(isPalindrome(12))
	println(isValid("()"))
	println(isValid("()}"))
	println(isValid("{()}"))

	println(singleNumberNew([]int{1, 2, 2, 3, 3}))
	println(singleNumberNew([]int{2}))
	println(singleNumberNew([]int{1, 1, 4, 2, 2, 3, 3}))
	println(singleNumberNew([]int{1}))
}
