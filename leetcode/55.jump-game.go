package leetcode

// You are given an integer array nums. You are initially positioned at the array's first index, and each element in the array represents your maximum jump length at that position.

// Return true if you can reach the last index, or false otherwise.

// Example 1:

// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:

// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.

// Constraints:

// 1 <= nums.length <= 104
// 0 <= nums[i] <= 105

// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

// 示例 1：

// 输入：nums = [2,3,1,1,4]
// 输出：true
// 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
// 示例 2：

// 输入：nums = [3,2,1,0,4]
// 输出：false
// 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。

// 提示：

// 1 <= nums.length <= 104
// 0 <= nums[i] <= 105

// [1, 0, 1, 0, 1]

func CanJump(nums []int) bool {
	lastIndex := len(nums) - 1 // 最終位置
	isArrive := false

	maxDistance := 0 // 最遠可到達的index
	for i := 0; i <= lastIndex; i++ {
		// 途中遇到0 一定要大於0的index位置
		if nums[i] == 0 && maxDistance <= i {
			return false
		}

		if maxDistance < nums[i]+i {
			maxDistance = nums[i] + i
			if maxDistance >= lastIndex {
				isArrive = true
			}
		}

	}
	return isArrive
}
