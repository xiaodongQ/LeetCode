/*
 * @lc app=leetcode.cn id=1 lang=cpp
 *
 * [1] 两数之和
 */

// @lc code=start
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        // 由于要返回下标，用unordered_map<int, int>记录<target-值, 下标>
        unordered_map<int, int> tmp;
        for (int i = 0; i < nums.size(); i++) {
            // 看当前成员是否为之前需要的值
            if (tmp.find(nums[i]) != tmp.end()) {
                return vector<int>{tmp[ nums[i] ], i};
            }
            // 记录当前成员需要哪个值就能凑成target，即 target-num 和 它的下标
            tmp[ target-nums[i] ] = i;
        }
        return vector<int>{};
    }
};
// @lc code=end

