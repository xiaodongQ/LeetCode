/*
 * @lc app=leetcode.cn id=383 lang=cpp
 *
 * [383] 赎金信
 */

// @lc code=start
class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        int letters[26] = {0};

        // 小优化，如果ransomNote长度比magazine长度大，则肯定不满足
        if (ransomNote.size() > magazine.size()) {
            return false;
        }

        for (char c : magazine) {
            letters[c - 'a'] ++;
        }

        // 检查赎金信ransomNote中字符是否都能在magazine中，每个字符最多一次
        for (char c : ransomNote) {
            letters[c - 'a']--;
            // 如果出现字符为负，则说明magazine里没有足够的字符能覆盖ransomNote
            if (letters[c - 'a'] < 0) {
                return false;
            }
        }
        return true;
    }
};
// @lc code=end

