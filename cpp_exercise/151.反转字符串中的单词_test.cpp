/*
 * @lc app=leetcode.cn id=151 lang=cpp
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
class Solution {
public:
    // 移除多余空格
    void removeExtraSpace(string &s) {
        int left = 0;
        int right = s.size() - 1;
        while (left <= right) {
            if (left > 0 && s[left] == ' ' && s[left] == s[left-1]) {
                s[left-1] = s[left];
                continue;
            }
            left++;
        }
    }

    // 反转
    void reverse()

    // 方式3：自行实现反转逻辑，辅助空间O(1)
    string reverseWords(string s) {
        // 1）先去除多余空格
        // 2）反转所有字符串
        // 3）反转单词
    }
};
// @lc code=end

