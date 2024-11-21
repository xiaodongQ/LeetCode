/*
 * @lc app=leetcode.cn id=18 lang=cpp
 *
 * [18] 四数之和
 */

// @lc code=start
class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        vector<vector<int>> result;
        // 先排序
        sort(nums.begin(), nums.end());
        // 利用双指针法从从有序数组里找四元组
        for (int i = 0; i < nums.size() - 3; i++) {
            // 优化：提前退出
            if (nums[i] > target) {
                break;
            }
            // 第一层去重
            if (i > 0 && nums[i] == nums[i-1]) {
                continue;
            }

            for (int j = i+1; j < nums.size() - 2; j++) {
                // 优化：提前退出
                if (nums[i] + nums[j] > target) {
                    break;
                }
                // 第二层去重
                if (j > i+1 && nums[j] == nums[j-1]) {
                    continue;
                }

                // 通过双指针法（而不是哈希法）
                int left = j + 1;
                int right = nums.size() - 1;
                while (left < right) {
                    if (nums[i] + nums[j] + nums[left] + nums[right] > target) {
                        right--;
                    } else if (nums[i] + nums[j] + nums[left] + nums[right] < target) {
                        left++;
                    } else {
                        // 满足条件，组织结果记录
                        result.push_back(vector<int>{nums[i], nums[j], nums[left], nums[right]});
                        // 为下一步做好去重
                        while (left < right && nums[left] == nums[left + 1]) {
                            left++;
                        }
                        while (left < right && nums[right] == nums[right - 1]) {
                            right--;
                        }
                        // 两侧都向内收缩才有可能继续满足求和条件
                        left++;
                        right--;
                    }
                }
            }
        }
        return result;
    }
};
// @lc code=end

