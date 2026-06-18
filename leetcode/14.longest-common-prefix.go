package leetcode

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters if it is non-empty.

// 编写一个函数来查找字符串数组中的最长公共前缀。

// 如果不存在公共前缀，返回空字符串 ""。

// 示例 1：

// 输入：strs = ["flower","flow","flight"]
// 输出："fl"
// 示例 2：

// 输入：strs = ["dog","racecar","car"]
// 输出：""
// 解释：输入不存在公共前缀。

// 提示：

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 如果非空，则仅由小写英文字母组成

// time O(n*m)
// 要loop陣列 也要loop字串 但單字的長度大多10個或更少 所以即使是m 也比預期小很多
// i >= len(strs[j]) 防止每個單字長度不同 會找到最短單字
// 逐一比對每個單字的字母 有一個不對表示這個字母不是共同prefix 就回傳他之前的字母[:i]
// 最後return strs[0] 如果走到這一步表示全部都跟strs[0]一樣 就回傳strs[0] 也代表他是最短單字 就不會觸發i >= len(strs[j])
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
