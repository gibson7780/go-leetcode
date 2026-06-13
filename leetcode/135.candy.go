package leetcode

// There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

// You are giving candies to these children subjected to the following requirements:

// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
// Return the minimum number of candies you need to have to distribute the candies to the children.

// Example 1:

// Input: ratings = [1,0,2]
// Output: 5
// Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
// Example 2:

// Input: ratings = [1,2,2]
// Output: 4
// Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
// The third child gets 1 candy because it satisfies the above two conditions.

// Constraints:

// n == ratings.length
// 1 <= n <= 2 * 104
// 0 <= ratings[i] <= 2 * 104

// 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。

// 你需要按照以下要求，给这些孩子分发糖果：

// 每个孩子至少分配到 1 个糖果。
// 相邻两个孩子中，评分更高的那个会获得更多的糖果。
// 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目 。

// 示例 1：

// 输入：ratings = [1,0,2]
// 输出：5
// 解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
// 示例 2：

// 输入：ratings = [1,2,2]
// 输出：4
// 解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
//      第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。

// 提示：

// n == ratings.length
// 1 <= n <= 2 * 104
// 0 <= ratings[i] <= 2 * 104

// 時間O(n) 空間O(n)
// 左到右 右到左 最後合併選大的
func candy(ratings []int) int {
	total := 0
	candiesLTR := make([]int, len(ratings))
	candiesRTL := make([]int, len(ratings))

	for i, n := range ratings {
		candiesLTR[i] += 1
		if i > 0 && ratings[i-1] < n {
			candiesLTR[i] += candiesLTR[i-1]
		}
	}
	for i := len(ratings) - 1; i >= 0; i-- {
		candiesRTL[i] += 1
		if i < len(ratings)-1 && ratings[i+1] < ratings[i] {
			candiesRTL[i] += candiesRTL[i+1]
		}
	}
	for i, _ := range ratings {
		if candiesLTR[i] > candiesRTL[i] {
			total += candiesLTR[i]
		} else {
			total += candiesRTL[i]
		}
	}

	return total
}

// 只用公式寫成 up*(up+1)/2 + down*(down+1)/2 + max(up,down) + 1, +1是指高峰
// 用上升跟下降的方式算出結果 時間O(n) 空間(1)
// 上升時每次多加一up後相加，下降時同理 down加一 持續下降時down也會累加，可以想成是幫前面持續下降的位置補加一
// peak是山頂的值 每次下降都會多加一次最高峰的位置，所以要扣掉，直到最高峰不再由上升控制變為下降主導高峰，就不用再扣，因為下降超越上升時高峰的值也需要再往上加
func candy2(ratings []int) int {
	total := 1
	up, down, peak := 0, 0, 0

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			up += 1
			peak = up
			down = 0
			total += up + 1
		} else if ratings[i] == ratings[i-1] {
			up, down, peak = 0, 0, 0
			total += 1
		} else {
			down += 1
			up = 0
			total += down + 1
			if peak >= down {
				total--
			}
		}
	}
	return total
}
