/*
 * @lc app=leetcode.cn id=344 lang=cpp
 *
 * [344] 反转字符串
 */

// @lc code=start
class Solution {
public:
/*
    void reverseString(vector<char>& s) {
        int left = 0;
        int right = s.size() - 1;
        // 保持区间的开闭一致
        char tmp;
        while (left <= right) {
            tmp = s[left];
            s[left] = s[right];
            s[right] = tmp;
            left++;
            right--;
        }
    }
*/

    void reverseString(vector<char>& s) {
        // for循环更简洁
        char tmp;
        for (int left = 0, right = s.size() - 1; left < s.size()/2; left++, right--) {
            tmp = s[left];
            s[left] = s[right];
            s[right] = tmp;
        }
    }
};
// @lc code=end

