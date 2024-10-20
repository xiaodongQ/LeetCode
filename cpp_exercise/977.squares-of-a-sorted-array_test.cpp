/*
 * @lc app=leetcode id=977 lang=cpp
 *
 * [977] Squares of a Sorted Array
 *
 * https://leetcode.com/problems/squares-of-a-sorted-array/description/
 *
 * algorithms
 * Easy (72.98%)
 * Likes:    9419
 * Dislikes: 242
 * Total Accepted:    2M
 * Total Submissions: 2.7M
 * Testcase Example:  '[-4,-1,0,3,10]'
 *
 * Given an integer array nums sorted in non-decreasing order, return an array
 * of the squares of each number sorted in non-decreasing order.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [-4,-1,0,3,10]
 * Output: [0,1,9,16,100]
 * Explanation: After squaring, the array becomes [16,1,0,9,100].
 * After sorting, it becomes [0,1,9,16,100].
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [-7,-3,2,3,11]
 * Output: [4,9,9,49,121]
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 1 <= nums.length <= 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums is sorted in non-decreasing order.
 * 
 * 
 * 
 * Follow up: Squaring each element and sorting the new array is very trivial,
 * could you find an O(n) solution using a different approach?
 */

// @lc code=start
class Solution {
public:
    vector<int> sortedSquares(vector<int>& nums) {
        // 双指针法，原数组为非递减序，即递增序或者相等，首尾比较依次可得到每轮最大值
        int left = 0;
        int right = nums.size() - 1;
        // 指定容量
        vector<int> result(nums.size(), 0);
        // 用于记录最大值存储位置，从数组最后开始
        int k = nums.size() - 1;
        while (left <= right) {
            // 右边大或等于，都取右侧，并收缩右边界
            if (nums[right]*nums[right] >= nums[left]*nums[left]) {
                result[k--] = nums[right]*nums[right];
                right--;
            } else {
                // 取左侧值，并收缩左边界
                result[k--] = nums[left]*nums[left];
                left++;
            }
        }
        return result;
    }
};
// @lc code=end

