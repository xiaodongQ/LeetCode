/*
 * @lc app=leetcode.cn id=15 lang=cpp
 *
 * [15] 三数之和
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> result;
        // 记住先得排序
        sort(nums.begin(), nums.end());
        for (int i = 0; i < nums.size(); i++) {
            // 排序后的第一个元素>0则不可能满足求和=0，可提前退出
            if (nums[i] > 0) {
                break;
            }

            // 这里要做一道去重
            if (i > 0 && nums[i] == nums[i-1]) {
                continue;
            }

            int left = i+1;
            int right = nums.size() - 1;
            while (left < right) {
                if (nums[i] + nums[left] + nums[right] > 0) {
                    right--;
                } else if (nums[i] + nums[left] + nums[right] < 0) {
                    left++;
                } else {
                    // 满足求和为0
                    result.push_back( vector<int>{nums[i], nums[left], nums[right]} );
                    // 左右都去重处理
                    while (right > left && nums[left] == nums[left+1]) {
                        left++;
                    }
                    while (right > left && nums[right] == nums[right-1]) {
                        right--;
                    }
                    // 由于求和为0，左右一增一减才可能继续求和为0
                    left++;
                    right--;
                }
            }
        }
        return result;
    }
};
// @lc code=end

