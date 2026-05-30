package leetcode

import (
	"log/slog"
	"sort"
)

// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.

// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

// Example 1:

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
// Example 2:

// Input: nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]
// Explanation: The arrays we are merging are [1] and [].
// The result of the merge is [1].
// Example 3:

// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
// Output: [1]
// Explanation: The arrays we are merging are [] and [1].
// The result of the merge is [1].
// Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

// Constraints:

// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200
// -109 <= nums1[i], nums2[j] <= 109

// Follow up: Can you come up with an algorithm that runs in O(m + n) time?

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

// 注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

// 示例 1：

// 输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// 输出：[1,2,2,3,5,6]
// 解释：需要合并 [1,2,3] 和 [2,5,6] 。
// 合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
// 示例 2：

// 输入：nums1 = [1], m = 1, nums2 = [], n = 0
// 输出：[1]
// 解释：需要合并 [1] 和 [] 。
// 合并结果是 [1] 。
// 示例 3：

// 输入：nums1 = [0], m = 0, nums2 = [1], n = 1
// 输出：[1]
// 解释：需要合并的数组是 [] 和 [1] 。
// 合并结果是 [1] 。
// 注意，因为 m = 0 ，所以 nums1 中没有元素。nums1 中仅存的 0 仅仅是为了确保合并结果可以顺利存放到 nums1 中。

// 提示：

// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200
// -109 <= nums1[i], nums2[j] <= 109

// 进阶：你可以设计实现一个时间复杂度为 O(m + n) 的算法解决此问题吗？

// 方法1 sort
// 填入 — O(n)
// sort.Ints — O((m+n)log(m+n))
func Merge1(nums1 []int, m int, nums2 []int, n int) {
	currentNums2Index := 0
	for i, item := range nums1 {
		if item == 0 {
			nums1[i] = nums2[currentNums2Index]
			currentNums2Index += 1
		}
	}
	// sort
	sort.Ints(nums1)
	slog.Info("result:", "nums1", nums1)
}

// 方法2 填入時比較i插入
// O(m+n)
func Merge2(nums1 []int, m int, nums2 []int, n int) {

}

// 方法2 後面往前填 遞迴
// 時間O(m+n) 空間O(m+n)
func Merge3(nums1 []int, m int, nums2 []int, n int) {
	compareNums1Index := m - 1
	compareNums2Index := n - 1
	totalLength := m + n
	currentSetIndex := totalLength - 1
	loop(nums1, nums2, compareNums1Index, compareNums2Index, currentSetIndex)
	slog.Info("result:", "nums1", nums1)
}

func loop(nums1, nums2 []int, compareNums1Index, compareNums2Index, currentSetIndex int) {
	if len(nums1) == 1 && nums1[0] == 0 {
		nums1[0] = nums2[0]
	}
	if currentSetIndex < 0 {
		return
	}

	if compareNums1Index < 0 {
		for compareNums2Index >= 0 {
			nums1[currentSetIndex] = nums2[compareNums2Index]
			compareNums2Index--
			currentSetIndex--
		}
		return
	}

	if compareNums2Index < 0 {
		return
	}
	if nums1[compareNums1Index] >= nums2[compareNums2Index] {
		nums1[currentSetIndex] = nums1[compareNums1Index]
		loop(nums1, nums2, compareNums1Index-1, compareNums2Index, currentSetIndex-1)
	} else if nums1[compareNums1Index] < nums2[compareNums2Index] {
		nums1[currentSetIndex] = nums2[compareNums2Index]
		loop(nums1, nums2, compareNums1Index, compareNums2Index-1, currentSetIndex-1)
	}
}

// 方法3 後面往前填 for迴圈
// 時間O(m+n) 空間O(1)
func merge3(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1

	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	// 處理 nums2 還有剩餘元素的情況
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
