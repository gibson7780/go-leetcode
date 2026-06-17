package leetcode

// Given a string s consisting of words and spaces, return the length of the last word in the string.

// A word is a maximal substring consisting of non-space characters only.

// Example 1:

// Input: s = "Hello World"
// Output: 5
// Explanation: The last word is "World" with length 5.
// Example 2:

// Input: s = "   fly me   to   the moon  "
// Output: 4
// Explanation: The last word is "moon" with length 4.
// Example 3:

// Input: s = "luffy is still joyboy"
// Output: 6
// Explanation: The last word is "joyboy" with length 6.

// Constraints:

// 1 <= s.length <= 104
// s consists of only English letters and spaces ' '.
// There will be at least one word in s.

// 给你一个字符串 s，由若干单词组成，单词前后用一些空格字符隔开。返回字符串中 最后一个 单词的长度。

// 单词 是指仅由字母组成、不包含任何空格字符的最大子字符串。

// 示例 1：

// 输入：s = "Hello World"
// 输出：5
// 解释：最后一个单词是“World”，长度为 5。
// 示例 2：

// 输入：s = "   fly me   to   the moon  "
// 输出：4
// 解释：最后一个单词是“moon”，长度为 4。
// 示例 3：

// 输入：s = "luffy is still joyboy"
// 输出：6
// 解释：最后一个单词是长度为 6 的“joyboy”。

// 提示：

// 1 <= s.length <= 104
// s 仅有英文字母和空格 ' ' 组成
// s 中至少存在一个单词

func lengthOfLastWord(s string) int {
	total := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if total != 0 {
				break
			}
		} else {
			total++
		}
	}
	return total
}
