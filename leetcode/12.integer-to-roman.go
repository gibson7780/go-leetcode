package leetcode

import (
	"strings"
)

// Seven different symbols represent Roman numerals with the following values:

// Symbol	Value
// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000
// Roman numerals are formed by appending the conversions of decimal place values from highest to lowest. Converting a decimal place value into a Roman numeral has the following rules:

// If the value does not start with 4 or 9, select the symbol of the maximal value that can be subtracted from the input, append that symbol to the result, subtract its value, and convert the remainder to a Roman numeral.
// If the value starts with 4 or 9 use the subtractive form representing one symbol subtracted from the following symbol, for example, 4 is 1 (I) less than 5 (V): IV and 9 is 1 (I) less than 10 (X): IX. Only the following subtractive forms are used: 4 (IV), 9 (IX), 40 (XL), 90 (XC), 400 (CD) and 900 (CM).
// Only powers of 10 (I, X, C, M) can be appended consecutively at most 3 times to represent multiples of 10. You cannot append 5 (V), 50 (L), or 500 (D) multiple times. If you need to append a symbol 4 times use the subtractive form.
// Given an integer, convert it to a Roman numeral.

// Example 1:

// Input: num = 3749

// Output: "MMMDCCXLIX"

// Explanation:

// 3000 = MMM as 1000 (M) + 1000 (M) + 1000 (M)
//  700 = DCC as 500 (D) + 100 (C) + 100 (C)
//   40 = XL as 10 (X) less of 50 (L)
//    9 = IX as 1 (I) less of 10 (X)
// Note: 49 is not 1 (I) less of 50 (L) because the conversion is based on decimal places
// Example 2:

// Input: num = 58

// Output: "LVIII"

// Explanation:

// 50 = L
//  8 = VIII
// Example 3:

// Input: num = 1994

// Output: "MCMXCIV"

// Explanation:

// 1000 = M
//  900 = CM
//   90 = XC
//    4 = IV

// Constraints:

// 1 <= num <= 3999

// 七个不同的符号代表罗马数字，其值如下：

// 符号	值
// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000
// 罗马数字是通过添加从最高到最低的小数位值的转换而形成的。将小数位值转换为罗马数字有以下规则：

// 如果该值不是以 4 或 9 开头，请选择可以从输入中减去的最大值的符号，将该符号附加到结果，减去其值，然后将其余部分转换为罗马数字。
// 如果该值以 4 或 9 开头，使用 减法形式，表示从以下符号中减去一个符号，例如 4 是 5 (V) 减 1 (I): IV ，9 是 10 (X) 减 1 (I)：IX。仅使用以下减法形式：4 (IV)，9 (IX)，40 (XL)，90 (XC)，400 (CD) 和 900 (CM)。
// 只有 10 的次方（I, X, C, M）最多可以连续附加 3 次以代表 10 的倍数。你不能多次附加 5 (V)，50 (L) 或 500 (D)。如果需要将符号附加4次，请使用 减法形式。
// 给定一个整数，将其转换为罗马数字。

// 示例 1：

// 输入：num = 3749

// 输出： "MMMDCCXLIX"

// 解释：

// 3000 = MMM 由于 1000 (M) + 1000 (M) + 1000 (M)
//  700 = DCC 由于 500 (D) + 100 (C) + 100 (C)
//   40 = XL 由于 50 (L) 减 10 (X)
//    9 = IX 由于 10 (X) 减 1 (I)
// 注意：49 不是 50 (L) 减 1 (I) 因为转换是基于小数位
// 示例 2：

// 输入：num = 58

// 输出："LVIII"

// 解释：

// 50 = L
//  8 = VIII
// 示例 3：

// 输入：num = 1994

// 输出："MCMXCIV"

// 解释：

// 1000 = M
//  900 = CM
//   90 = XC
//    4 = IV

// 提示：

// 1 <= num <= 3999

func intToRoman(num int) string {
	var roman strings.Builder
	for num > 0 {
		if num >= 1000 {
			num = num - 1000
			roman.WriteString("M")
		} else if num >= 900 {
			num = num - 900
			roman.WriteString("CM")
		} else if num >= 500 {
			num = num - 500
			roman.WriteString("D")
		} else if num >= 400 {
			num = num - 400
			roman.WriteString("CD")
		} else if num >= 100 {
			num -= 100
			roman.WriteString("C")
		} else if num >= 90 {
			num -= 90
			roman.WriteString("XC")
		} else if num >= 50 {
			num -= 50
			roman.WriteString("L")
		} else if num >= 40 {
			num -= 40
			roman.WriteString("XL")
		} else if num >= 10 {
			num -= 10
			roman.WriteString("X")
		} else if num >= 9 {
			num -= 9
			roman.WriteString("IX")
		} else if num >= 5 {
			num -= 5
			roman.WriteString("V")
		} else if num == 4 {
			num -= 4
			roman.WriteString("IV")
		} else if num >= 1 {
			num -= 1
			roman.WriteString("I")
		}
	}
	return roman.String()
}

// array的長度每次固定的話時間跟空間就會是O(1)
func intToRoman2(num int) string {
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	var result strings.Builder
	for i, n := range values {
		for num >= n {
			num -= n
			result.WriteString(romans[i])
		}
	}
	return result.String()
}
