package leetcode

// You are given a 0-indexed array of integers nums of length n. You are initially positioned at index 0.

// Each element nums[i] represents the maximum length of a forward jump from index i. In other words, if you are at index i, you can jump to any index (i + j) where:

// 0 <= j <= nums[i] and
// i + j < n
// Return the minimum number of jumps to reach index n - 1. The test cases are generated such that you can reach index n - 1.

// Example 1:

// Input: nums = [2,3,1,1,4]
// Output: 2
// Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
// Example 2:

// Input: nums = [2,3,0,1,4]
// Output: 2

// Constraints:

// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000
// It's guaranteed that you can reach nums[n - 1].

// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置在下标 0。

// 每个元素 nums[i] 表示从索引 i 向后跳转的最大长度。换句话说，如果你在索引 i 处，你可以跳转到任意 (i + j) 处：

// 0 <= j <= nums[i] 且
// i + j < n
// 返回到达 n - 1 的最小跳跃次数。测试用例保证可以到达 n - 1。

// 示例 1:

// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//      从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
// 示例 2:

// 输入: nums = [2,3,0,1,4]
// 输出: 2

// 提示:

// 1 <= nums.length <= 104
// 0 <= nums[i] <= 1000
// 题目保证可以到达 n - 1
// [2,2,2,1,4,3,0,0,1]
// 空間O(n)時間(1)
// 貪婪法
// 預先紀錄每一輪會跳的次數，也就是一開始還沒到終點所以一定會有1跳先記錄起來，而不數確定會跳才紀錄，i < len(nums)-1 不會執行終點那一次所以避免多跳一次
func jump(nums []int) int {
	jumpMaxDistance := 0
	jumpTimes := 0
	currentJumpMaxIndex := 0
	for i := 0; i < len(nums)-1; i++ {
		if jumpMaxDistance < nums[i]+i {
			jumpMaxDistance = nums[i] + i
		}
		if i == currentJumpMaxIndex {
			jumpTimes++
			currentJumpMaxIndex = jumpMaxDistance
		}
	}
	return jumpTimes
}
