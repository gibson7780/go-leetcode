package leetcode

// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

// Constraints:

// n == nums.length
// 1 <= n <= 5 * 104
// -109 <= nums[i] <= 109
// The input is generated such that a majority element will exist in the array.

// Follow-up: Could you solve the problem in linear time and in O(1) space?

// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。

// 示例 1：

// 输入：nums = [3,2,3]
// 输出：3
// 示例 2：

// 输入：nums = [2,2,1,1,1,2,2]
// 输出：2

// 提示：
// n == nums.length
// 1 <= n <= 5 * 104
// -109 <= nums[i] <= 109
// 输入保证数组中一定有一个多数元素。

// 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。

// 方法一 hash map 時間O(n) 空間O(n)
// 方法二 先排序在計算 時間O(n log n) 空間O(log n)

// 方法三 moore voting algorithm
// 時間O(n) 空間O(1)
func majorityElement(nums []int) int {
	candidate := 0
	voting := 0
	for _, n := range nums {
		if candidate == n {
			voting++
		} else {
			if voting == 0 {
				candidate = n
				voting++
			}
			if voting > 0 {
				voting--
			}
		}

	}
	return candidate
}
