package lc

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (60.59%)
 * Likes:    219
 * Dislikes: 0
 * Total Accepted:    117.3K
 * Total Submissions: 193.6K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
 *
 * 示例 1:
 *
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "rat", t = "car"
 * 输出: false
 *
 * 说明:
 * 你可以假设字符串只包含小写字母。
 *
 * 进阶:
 * 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 *
 */

// @lc code=start
/*
	异位词指的是构成字母相同，位置不同

	利用两个map，记录每个字符串各个字符出现次数，再一一比较两个map字符和个数是否匹配
	34/34 cases passed (8 ms)
	Your runtime beats 54.27 % of golang submissions
	Your memory usage beats 57.14 % of golang submissions (3 MB)
	时间O(n)，空间O(1) 固定大小的map
*/
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	// 小写字母，则容量可以小一点
// 	m1 := make(map[byte]int, 26)
// 	m2 := make(map[byte]int, 26)
// 	for i := 0; i < len(s); i++ {
// 		m1[s[i]]++
// 	}
// 	for i := 0; i < len(t); i++ {
// 		m2[t[i]]++
// 	}
// 	// 比较两个map
// 	if len(m1) != len(m2) {
// 		return false
// 	}
// 	for k, v := range m1 {
// 		if c, ok := m2[k]; ok && c == v {
// 			continue
// 		}
// 		return false
// 	}
// 	return true
// }

/*
	跟上面一样，但是并没有必要用两个map
	34/34 cases passed (12 ms)
	Your runtime beats 23.26 % of golang submissions
	Your memory usage beats 57.14 % of golang submissions (3 MB)
*/
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
// 	// 小写字母，则容量可以小一点
// 	m1 := make(map[byte]int, 26)
// 	for i := 0; i < len(s); i++ {
// 		m1[s[i]]++
// 	}

// 	for i := 0; i < len(t); i++ {
// 		m1[t[i]]--
// 		// 只要出现陌生字符则说明非异位词
// 		if m1[t[i]] < 0 {
// 			return false
// 		}
// 	}

// 	return true
// }

/*
	看题解里的讨论说数组速度快很多，试一下
	类似与桶排序的思想，26个桶

	34/34 cases passed (0 ms)
	Your runtime beats 100 % of golang submissions
	Your memory usage beats 100 % of golang submissions (3 MB)

	确实快了很多，对比map
		make时已经初始化了容量，应该不用考虑扩容？
		CPU Cache Line？数组这么小应该不是这个影响
	应该还是map里面的操作影响，查看map源码里，赋值时(不确定是不是 mapassign?) 里一大堆操作
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 小写字母，则容量可以小一点
	bucket := [26]int{}
	for i := 0; i < len(s); i++ {
		bucket[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {
		bucket[t[i]-'a']--
		// 只要出现陌生字符则说明非异位词
		if bucket[t[i]-'a'] < 0 {
			return false
		}
	}

	return true
}

/*
	利用排序，排序后的字符串是否相同
	这种思路也要考虑到

	34/34 cases passed (8 ms)
	Your runtime beats 54.27 % of golang submissions
	Your memory usage beats 57.14 % of golang submissions (3.2 MB)
	时间O(nlogn) 内部排序主要用的快排，空间O(n)
*/
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	s1 := []byte(s)
// 	s2 := []byte(t)
// 	// sort包没有提供 []byte 版本的排序，需要自己传入一个less函数，i,j为下标
// 	// 这个方法在sort包的slice.go里面，怪不得sort.go里没找到。。。
// 	sort.Slice(s1, func(i, j int) bool {
// 		return s1[i] < s1[j]
// 	})
// 	sort.Slice(s2, func(i, j int) bool {
// 		return s2[i] < s2[j]
// 	})
// 	return string(s1) == string(s2)
// }

// @lc code=end
