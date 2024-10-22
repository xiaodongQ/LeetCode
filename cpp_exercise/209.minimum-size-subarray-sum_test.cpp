/*
 * @lc app=leetcode id=209 lang=cpp
 *
 * [209] Minimum Size Subarray Sum
 *
 * https://leetcode.com/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (48.00%)
 * Likes:    12857
 * Dislikes: 458
 * Total Accepted:    1.2M
 * Total Submissions: 2.5M
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * Given an array of positive integers nums and a positive integer target,
 * return the minimal length of a subarray whose sum is greater than or equal
 * to target. If there is no such subarray, return 0 instead.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: target = 7, nums = [2,3,1,2,4,3]
 * Output: 2
 * Explanation: The subarray [4,3] has the minimal length under the problem
 * constraint.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: target = 4, nums = [1,4,4]
 * Output: 1
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: target = 11, nums = [1,1,1,1,1,1,1,1]
 * Output: 0
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= target <= 10^9
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^4
 * 
 * 
 * 
 * Follow up: If you have figured out the O(n) solution, try coding another
 * solution of which the time complexity is O(n log(n)).
 */

// @lc code=start
class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int result = INT32_MAX;
        // 滑动窗口起止
        int left = 0;
        int right = 0;
        // 滑动窗口长度
        int sub_len = 0;
        // 窗口内求和，题目中数组成员是正整数
        int sum = 0;
        for (int right = 0; right < nums.size(); right++) {
            sum += nums[right];
            // 基于窗口求和循环判断来调整左侧（避免下标作为循环，判断多次边界容易出错）
            while (sum >= target) {
                sub_len = right - left + 1;
                if (sub_len < result) {
                    result = sub_len;
                }
                // 尝试收缩左侧
                sum -= nums[left];
                left++;
            }
        }
        if (result == INT32_MAX) {
            result = 0;
        }
        return result;
    }
/*
    // 收缩窗口的内层循环，用左右边界作为判断条件，不简洁且容易搞错
    int minSubArrayLen(int target, vector<int>& nums) {
        int result = INT32_MAX;
        // 滑动窗口起止
        int left = 0;
        int right = 0;
        // 滑动窗口长度
        int sub_len = 0;
        // 窗口内求和，题目中数组成员是正整数
        int sum = 0;
        for (int right = 0; right < nums.size(); right++) {
            sum += nums[right];
            if (sum >= target) {
                while (left <= right) {
                    sub_len = right - left + 1;
                    if ( sub_len < result ) {
                        result = sub_len;
                    }
                    sum -= nums[left++];
                    if (sum < target) {
                        break;
                    }
                }
            }
        }
        if (result == INT32_MAX) {
            result = 0;
        }
        return result;
    }
*/
};
// @lc code=end

