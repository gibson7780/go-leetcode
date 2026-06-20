package leetcode

import "strings"

// Given an input string s, reverse the order of the words.

// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.

// Return a string of the words in reverse order concatenated by a single space.

// Note that s may contain leading or trailing spaces or multiple spaces between two words. The returned string should only have a single space separating the words. Do not include any extra spaces.

// Example 1:

// Input: s = "the sky is blue"
// Output: "blue is sky the"
// Example 2:

// Input: s = "  hello world  "
// Output: "world hello"
// Explanation: Your reversed string should not contain leading or trailing spaces.
// Example 3:

// Input: s = "a good   example"
// Output: "example good a"
// Explanation: You need to reduce multiple spaces between two words to a single space in the reversed string.

// Constraints:

// 1 <= s.length <= 104
// s contains English letters (upper-case and lower-case), digits, and spaces ' '.
// There is at least one word in s.

// Follow-up: If the string data type is mutable in your language, can you solve it in-place with O(1) extra space?

// 给你一个字符串 s ，请你反转字符串中 单词 的顺序。

// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。

// 返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。

// 注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

// 示例 1：

// 输入：s = "the sky is blue"
// 输出："blue is sky the"
// 示例 2：

// 输入：s = "  hello world  "
// 输出："world hello"
// 解释：反转后的字符串中不能存在前导空格和尾随空格。
// 示例 3：

// 输入：s = "a good   example"
// 输出："example good a"
// 解释：如果两个单词间有多余的空格，反转后的字符串需要将单词间的空格减少到仅有一个。

// 提示：

// 1 <= s.length <= 104
// s 包含英文大小写字母、数字和空格 ' '
// s 中 至少存在一个 单词

// 进阶：如果字符串在你使用的编程语言中是一种可变数据类型，请尝试使用 O(1) 额外空间复杂度的 原地 解法。

// time O(n) space O(n)
// 從尾部開始loop 遇到有字就開始計算 直到下一個空白表示這個單字結束就取出
// 要想比較多edge case 有點複雜
func reverseWords(s string) string {
	var reverseStr strings.Builder
	endIndex := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if endIndex != -1 && i != 0 {
				if len(reverseStr.String()) != 0 {
					reverseStr.WriteString(" ")
				}
				reverseStr.WriteString(s[i+1 : endIndex+1])
				endIndex = -1

			}

		} else {
			if i == 0 {
				if len(reverseStr.String()) != 0 {
					reverseStr.WriteString(" ")
				}
				if endIndex == -1 {
					reverseStr.WriteString(s[:1])
				} else {
					reverseStr.WriteString(s[0 : endIndex+1])
				}
			} else {
				if endIndex == -1 {
					endIndex = i
				}
			}
		}

	}
	return reverseStr.String()
}

// go 不能直接操作ｓｔｒｉｎｇ，先轉成byte, 無法解決單字間多個去白
func reverseWords2(s string) string {
	s = strings.TrimSpace(s)
	b := []byte(s)
	reverseByte(b, 0, len(s)-1)

	start := 0
	for i := 0; i <= len(b); i++ {
		if i == len(b) || b[i] == ' ' {
			reverseByte(b, start, i-1)
			start = i + 1
		}

	}
	return string(b)
}
func reverseByte(b []byte, l, r int) {
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}
}
