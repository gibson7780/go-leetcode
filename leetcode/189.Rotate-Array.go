package leetcode

import (
	"log/slog"
)

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

// Example 1:

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]
// Example 2:

// Input: nums = [-1,-100,3,99], k = 2
// Output: [3,99,-1,-100]
// Explanation:
// rotate 1 steps to the right: [99,-1,-100,3]
// rotate 2 steps to the right: [3,99,-1,-100]

// Constraints:

// 1 <= nums.length <= 105
// -231 <= nums[i] <= 231 - 1
// 0 <= k <= 105

// Follow up:

// Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
// Could you do it in-place with O(1) extra space?

// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

// 示例 1:

// 输入: nums = [1,2,3,4,5,6,7], k = 3
// 输出: [5,6,7,1,2,3,4]
// 解释:
// 向右轮转 1 步: [7,1,2,3,4,5,6]
// 向右轮转 2 步: [6,7,1,2,3,4,5]
// 向右轮转 3 步: [5,6,7,1,2,3,4]
// 示例 2:

// 输入：nums = [-1,-100,3,99], k = 2
// 输出：[3,99,-1,-100]
// 解释:
// 向右轮转 1 步: [99,-1,-100,3]
// 向右轮转 2 步: [3,99,-1,-100]

// 提示：

// 1 <= nums.length <= 105
// -231 <= nums[i] <= 231 - 1
// 0 <= k <= 105

// 进阶：

// 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
// 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？

// 空間O(n) slice 合併會產生新的記憶體 不符合O(1)
func Rotate(nums []int, k int) {

	for i := 0; i < k; i++ {
		length := len(nums)
		nums = append(nums[length-1:], nums[:length-1]...)
	}

}

// 空間O(1) 時間O(n^)
// 從後面把要替換的直往右移
func Rotate2(nums []int, k int) {

	for i := 0; i < k; i++ {
		temp := nums[len(nums)-1]
		for j := len(nums) - 1; j >= 0; j-- {
			if j == 0 {
				nums[0] = temp
			} else {
				nums[j] = nums[j-1]
			}
		}
	}
	slog.Info("rotate", "nums", nums)
}

// 時間O(n) 空間(k) k: 1 < k < n 複雜度比n小
func Rotate3(nums []int, k int) {
	moveNums := []int{}
	moveNums = append(moveNums, nums[len(nums)-k:]...)
	for j := len(nums) - 1; j >= 0; j-- {
		if j <= k-1 {
			nums[j] = moveNums[j]
		} else {
			nums[j] = nums[j-k]
		}
	}
	slog.Info("rotate", "nums", nums)
}

// 時間O(n) 空間(1)
// 3次反轉
func Rotate4(nums []int, k int) {
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
	slog.Info("rotate", "nums", nums)
}

func reverse(nums []int, start, end int) {
	// 只要start還小於end 就繼續交換
	for start < end {
		// 利用go的多變數賦值特性 會一次取出塞入 交換不用像其他語言需要temp暫存
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
