package leetcode

// Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

// Example 1:

// Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
// Output: 6
// Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
// Example 2:

// Input: height = [4,2,0,3,2,5]
// Output: 9

// Constraints:

// n == height.length
// 1 <= n <= 2 * 104
// 0 <= height[i] <= 105

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

// 示例 1：

// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
// 示例 2：

// 输入：height = [4,2,0,3,2,5]
// 输出：9

// 提示：

// n == height.length
// 1 <= n <= 2 * 104
// 0 <= height[i] <= 105

// 左到右, 右到左 取最低值
// 時間O(n) 空間O(n)
func trap(height []int) int {
	maxH := 0
	// 可儲存水量
	LTR := make([]int, len(height))
	RTL := make([]int, len(height))
	total := 0
	for i, h := range height {

		if h > maxH {
			maxH = h
		} else {
			LTR[i] = maxH - h
		}
	}
	maxH = 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > maxH {
			maxH = height[i]
		} else {
			RTL[i] = maxH - height[i]
		}
	}
	for i := 0; i < len(height); i++ {
		if LTR[i] > RTL[i] {
			total += RTL[i]
		} else {
			total += LTR[i]
		}
	}
	return total
}

// 雙指標
// 可以確保已經形成凹槽
// time O(n) space O(1)
// 比較高的一方不動 表示另一邊比較低就可以安心的跑另一邊
// 只要下降就加一格水 因為另一邊比較高已經保證可以形成凹槽
// 就算另一邊的位置不是同一個凹槽的另一邊也無所謂 如果不是那就會自己跑完 自己找到凹槽的另一邊 如果是同個凹槽 那就剛好會交會
func trap2(height []int) int {
	leftMax := 0
	left := 0
	rightMax := 0
	right := len(height) - 1
	total := 0
	for left < right {
		if height[left] <= height[right] {
			if leftMax <= height[left] {
				leftMax = height[left]
			} else {
				total += leftMax - height[left]
			}
			left++
		} else {
			if rightMax <= height[right] {
				rightMax = height[right]
			} else {
				total += rightMax - height[right]
			}
			right--
		}

	}
	return total
}
