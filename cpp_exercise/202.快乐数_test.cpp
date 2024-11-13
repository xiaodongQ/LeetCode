/*
 * @lc app=leetcode.cn id=202 lang=cpp
 *
 * [202] 快乐数
 */

// @lc code=start
class Solution {
public:
    bool isHappy(int n) {
        unordered_set<int> sums;
        int num = n;
        int sum = 0;
        int bit = 0;
        while (true) {
            // 计算平方和
            while (num > 0) {
                bit = num % 10;
                sum += bit * bit;
                num = num / 10;
            }
            if (sum == 1) {
                return true;
            }

            // 找到重复的平方和，说明会无限循环，非快乐数
            if (sums.find(sum) != sums.end()) {
                return false;
            }
            sums.insert(sum);
            // 下一轮
            num = sum;
            sum = 0;
        }
    }
};
// @lc code=end

