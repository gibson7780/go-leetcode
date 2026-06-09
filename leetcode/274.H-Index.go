package leetcode

// Given an array of integers citations where citations[i] is the number of citations a researcher received for their ith paper, return the researcher's h-index.

// According to the definition of h-index on Wikipedia: The h-index is defined as the maximum value of h such that the given researcher has published at least h papers that have each been cited at least h times.

// Example 1:

// Input: citations = [3,0,6,1,5]
// Output: 3
// Explanation: [3,0,6,1,5] means the researcher has 5 papers in total and each of them had received 3, 0, 6, 1, 5 citations respectively.
// Since the researcher has 3 papers with at least 3 citations each and the remaining two with no more than 3 citations each, their h-index is 3.
// Example 2:

// Input: citations = [1,3,1]
// Output: 1

// Constraints:

// n == citations.length
// 1 <= n <= 5000
// 0 <= citations[i] <= 1000

// 给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。

// 根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。

// 示例 1：

// 输入：citations = [3,0,6,1,5]
// 输出：3
// 解释：给定数组表示研究者总共有 5 篇论文，每篇论文相应的被引用了 3, 0, 6, 1, 5 次。
//      由于研究者有 3 篇论文每篇 至少 被引用了 3 次，其余两篇论文每篇被引用 不多于 3 次，所以她的 h 指数是 3。
// 示例 2：

// 输入：citations = [1,3,1]
// 输出：1

// 提示：

// n == citations.length
// 1 <= n <= 5000
// 0 <= citations[i] <= 1000

// 1. 時間O(n^) 兩次for 外圈從1(hindex)開始計算 內圈有幾個大於等於外圈的數字直到內圈找到的總數小於外圈

// 2. 時間O(nlogn) 先排序大到小再loop檢查就知道h到哪

// 3. 時間O(n) 空間O(n)
// 先算出每個分數的數量 index是引用次數 value是有幾個論文達到這個引用次數
// 論文總數如果只有n 超過的引用次數算在最大值n裡面
// 跟使用排序最大到最小概念差不多 只是把排序的時間複雜度利用空間來代替
func hIndex(citations []int) int {
	n := len(citations)
	bucket := make([]int, n+1) // 0也要算進去
	// 轉成每個引用有幾篇 index: 引用 value: 論文篇數
	for _, c := range citations {
		if n <= c {
			bucket[n]++
		} else {
			bucket[c]++
		}
	}
	count := 0 // count — 代表引用次數 >= i 的論文篇數
	// i — 代表引用次數（h-index 的候選值）
	// 找出最大的h-index
	// 至少有h篇論文至少被引用至少h次 才能作為這個作者的最大h-index
	// 如果hindex是3 代表當i是3時發現至少有3篇引用至少3次
	// 論文篇數跟hindex上限是一起增加的, hindex想要變4就至少要有4篇論文被引用至少4次 以此類推
	for i := n; i >= 0; i-- { // [3,0,6,1,5], 引用次數i=5時
		count += bucket[i] // 只有2篇論文引用次數大於5
		if i <= count {    // 5的hindex <= 2篇 不符合定義
			return i // 引用次數i=3時 有3篇引用次數大於3 符合hindex定義
		}
	}
	return 0
}
