package leetcode

import "strings"

// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

// P   A   H   N
// A P L S I I G
// Y   I   R
// And then read line by line: "PAHNAPLSIIGYIR"

// Write the code that will take a string and make this conversion given a number of rows:

// string convert(string s, int numRows);

// Example 1:

// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
// Example 2:

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// Example 3:

// Input: s = "A", numRows = 1
// Output: "A"

// Constraints:

// 1 <= s.length <= 1000
// s consists of English letters (lower-case and upper-case), ',' and '.'.
// 1 <= numRows <= 1000

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

// 请你实现这个将字符串进行指定行数变换的函数：

// string convert(string s, int numRows);

// 示例 1：

// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"
// 示例 2：
// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
// 示例 3：

// 输入：s = "A", numRows = 1
// 输出："A"

// 提示：

// 1 <= s.length <= 1000
// s 由英文字母（小写和大写）、',' 和 '.' 组成
// 1 <= numRows <= 1000

// time O(n) space O(n)
// 重點: 控制每排的陣列index來回跳 最後再展開
// for range 在loop字串時每個元素會是rune 所以builder接收時會是WriteRune()
// for i 在loop時s[i]取出會是byte 所以builder接收時會是WriteByte()
// for i 在loop時 做slice片段操作s[i : i+1] 會回傳string 所以builder接收時會是WriteString()
func zConvert(s string, numRows int) string {
	rows := make([]strings.Builder, numRows)
	trigger := -1
	row := 0

	for i := 0; i < len(s); i++ {
		// r := rune(s[i])
		// rows[row].WriteString(s[i : i+1])
		rows[row].WriteByte(s[i])
		if row == 0 || row == numRows-1 {
			trigger = trigger * -1
		}
		row += trigger
	}

	var result strings.Builder
	for _, n := range rows {
		result.WriteString(n.String())
	}
	return result.String()
}
