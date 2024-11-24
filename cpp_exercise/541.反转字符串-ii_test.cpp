/*
 * @lc app=leetcode.cn id=541 lang=cpp
 *
 * [541] 反转字符串 II
 */

// @lc code=start
class Solution {
public:
    string reverseStr(string s, int k) {
        // 每次移动2k
        for (int i = 0; i < s.size(); i += 2*k) {
            // 剩余字符少于 k 个，反转剩下所有字符
            if (s.size() - i < k) {
                reverse(s.begin() + i, s.end());
            } else {
                // 每k个进行反转
                reverse(s.begin() + i, s.begin() + i + k);
            }
        }

        return s;
    }
};
// @lc code=end

