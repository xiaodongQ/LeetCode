package lc

import "fmt"

/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 *
 * https://leetcode-cn.com/problems/bulls-and-cows/description/
 *
 * algorithms
 * Easy (48.04%)
 * Likes:    82
 * Dislikes: 0
 * Total Accepted:    16.1K
 * Total Submissions: 33.6K
 * Testcase Example:  '"1807"\n"7810"'
 *
 * 你在和朋友一起玩 猜数字（Bulls and Cows）游戏，该游戏规则如下：
 *
 *
 * 你写出一个秘密数字，并请朋友猜这个数字是多少。
 * 朋友每猜测一次，你就会给他一个提示，告诉他的猜测数字中有多少位属于数字和确切位置都猜对了（称为“Bulls”,
 * 公牛），有多少位属于数字猜对了但是位置不对（称为“Cows”, 奶牛）。
 * 朋友根据提示继续猜，直到猜出秘密数字。
 *
 *
 * 请写出一个根据秘密数字和朋友的猜测数返回提示的函数，返回字符串的格式为 xAyB ，x 和 y 都是数字，A 表示公牛，用 B 表示奶牛。
 *
 *
 * xA 表示有 x 位数字出现在秘密数字中，且位置都与秘密数字一致。
 * yB 表示有 y 位数字出现在秘密数字中，但位置与秘密数字不一致。
 *
 *
 * 请注意秘密数字和朋友的猜测数都可能含有重复数字，每位数字只能统计一次。
 *
 *
 *
 * 示例 1:
 *
 * 输入: secret = "1807", guess = "7810"
 * 输出: "1A3B"
 * 解释: 1 公牛和 3 奶牛。公牛是 8，奶牛是 0, 1 和 7。
 *
 * 示例 2:
 *
 * 输入: secret = "1123", guess = "0111"
 * 输出: "1A1B"
 * 解释: 朋友猜测数中的第一个 1 是公牛，第二个或第三个 1 可被视为奶牛。
 *
 *
 *
 * 说明: 你可以假设秘密数字和朋友的猜测数都只包含数字，并且它们的长度永远相等。
 *
 */

// @lc code=start
/*
	只是根据输入生成提示，要是反向就头大了
	map记录出现的次数，如果是相同位置相同值则不用记录
	152/152 cases passed (4 ms)
	Your runtime beats 61.58 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.3 MB)
	时间O(n)，空间O(1)
*/
// func getHint(secret string, guess string) string {
// 	bulls := 0
// 	cows := 0
// 	m := make(map[byte]int, 10)
// 	for i := 0; i < len(secret); i++ {
// 		m[secret[i]]++
// 	}
// 	// 相同位置相同值先判断，会消耗对应值的个数
// 	for i := 0; i < len(guess); i++ {
// 		if secret[i] == guess[i] {
// 			m[guess[i]]--
// 			bulls++
// 		}
// 	}
// 	for i := 0; i < len(guess); i++ {
// 		if secret[i] != guess[i] {
// 			if c, ok := m[guess[i]]; ok && c > 0 {
// 				// 存在
// 				cows++
// 				m[guess[i]]--
// 			}
// 		}
// 	}
// 	return fmt.Sprintf("%dA%dB", bulls, cows)
// }

/*
	用两个map记录出现的次数(只有0-9范围可以用数组)，如果是相同位置相同值则不用记录
	152/152 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2.3 MB)
	时间O(n)，空间O(1)
*/
func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0
	arr1, arr2 := [10]int{}, [10]int{}
	for i := 0; i < len(guess); i++ {
		if guess[i] == secret[i] {
			bulls++
			continue
		}

		arr1[guess[i]-'0']++
		arr2[secret[i]-'0']++
	}
	// 找重合配对部分
	for i := 0; i < 10; i++ {
		if arr1[i] > 0 && arr2[i] > 0 {
			c := arr1[i]
			if arr2[i] < c {
				c = arr2[i]
			}
			cows += c
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows)
}

// @lc code=end
