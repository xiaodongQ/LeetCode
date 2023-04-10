package lc

/*
 * @lc app=leetcode.cn id=292 lang=golang
 *
 * [292] Nim 游戏
 *
 * https://leetcode-cn.com/problems/nim-game/description/
 *
 * algorithms
 * Easy (69.43%)
 * Likes:    341
 * Dislikes: 0
 * Total Accepted:    52K
 * Total Submissions: 74.9K
 * Testcase Example:  '4'
 *
 * 你和你的朋友，两个人一起玩 Nim 游戏：桌子上有一堆石头，每次你们轮流拿掉 1 - 3 块石头。 拿掉最后一块石头的人就是获胜者。你作为先手。
 *
 * 你们是聪明人，每一步都是最优解。 编写一个函数，来判断你是否可以在给定石头数量的情况下赢得游戏。
 *
 * 示例:
 *
 * 输入: 4
 * 输出: false
 * 解释: 如果堆中有 4 块石头，那么你永远不会赢得比赛；
 * 因为无论你拿走 1 块、2 块 还是 3 块石头，最后一块石头总是会被你的朋友拿走。
 *
 *
 */

// @lc code=start
/*
	我为先手，若最后剩4个，轮到我拿时肯定是我输
	我拿a个，对方拿时，一直凑成拿完之后是4的倍数那我就输了

	对于总块数是4的倍数，没法避免留给我拿时还是4的倍数(对方只要每次凑4)
	对于总块数非4的倍数，我还应该尽力把4的倍数留给对方选择(即我拿完后应该是4的倍数)
	e.g.总块数10：
		若第一步我拿1，对方拿1，那后面只要对方一直凑4，我就输了
		若第一步我拿2(剩余8)，对方选完后我凑4，我就赢了
	e.g. 总块数8：
		我没得选。。。对方一直凑4我必输

	由于每步都最优(按上面规则凑4的倍数留给对方)，所以只要判断是否被4整除
	60/60 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (2 MB)
*/
func canWinNim(n int) bool {
	// 不能被4整除我才能赢
	return n%4 != 0
}

// @lc code=end
