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
        // 先对vector排序，便于去重
        sort(nums.begin(), nums.end());
        for( int i = 0; i < nums.size() - 2; i++) {
            // 去重处理，前面已经处理过该值的情况了
            if (i > 0 && nums[i] == nums[i-1]) {
                continue;
            }
            int sum = -nums[i];
            // 处理两数之和
            // 只要用set记录需要的配对值即可，获取到之后就可以移除掉
            unordered_set<int> tmp_set;
            for (int j = i+1; j < nums.size(); j++) {
                // 去重
                if (j > i+2 && nums[j] == nums[j-1] && nums[j] == nums[j-2]) {
                    continue;
                }

                if (tmp_set.find(nums[j]) != tmp_set.end()) {
                    // 满足三数之和为0，直接组装结果，第2个数（a/b/c中的b），用0-a-c即可
                    result.push_back( vector<int>{nums[i], -nums[i]-nums[j], nums[j]} );
                    tmp_set.erase(nums[j]);
                } else {
                    tmp_set.insert(sum - nums[j]);
                }
            }
        }
        
        return result;
    }
};
// @lc code=end

