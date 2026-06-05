package leetcode

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.

// Constraints:

// 1 <= prices.length <= 105
// 0 <= prices[i] <= 104

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

// 示例 1：

// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
// 示例 2：

// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

// 提示：

// 1 <= prices.length <= 105
// 0 <= prices[i] <= 104

// 時間O(n^) 空間O(1)
func maxProfit(prices []int) int {
	maximum := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] > prices[i] && maximum < (prices[j]-prices[i]) {
				maximum = prices[j] - prices[i]
			}
		}
	}

	return maximum
}

// 時間O(n) 空間O(1)
func maxProfit2(prices []int) int {
	maximum := 0
	minimum := 0
	currentProfit := 0
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			minimum = prices[i]
		}
		if (prices[i]-minimum) > 0 && maximum < prices[i] {
			maximum = prices[i]
		}
		if (prices[i]-minimum) < 0 && minimum > prices[i] {
			minimum = prices[i]
			maximum = 0
		}

		if maximum-minimum > currentProfit {
			currentProfit = maximum - minimum
		}
	}

	return currentProfit
}

func maxProfit3(prices []int) int {
	minimumPrice := prices[0]
	profit := 0
	for _, price := range prices {
		if (price - minimumPrice) > profit {
			profit = price - minimumPrice
		}
		if price < minimumPrice {
			minimumPrice = price
		}
	}
	return profit
}
