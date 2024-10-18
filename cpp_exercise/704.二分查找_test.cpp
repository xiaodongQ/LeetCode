/*
 * @lc app=leetcode.cn id=704 lang=cpp
 *
 * [704] 二分查找
 *
 * https://leetcode.cn/problems/binary-search/description/
 *
 * algorithms
 * Easy (54.49%)
 * Likes:    1276
 * Dislikes: 0
 * Total Accepted:    958.1K
 * Total Submissions: 1.8M
 * Testcase Example:  '[-1,0,3,5,9,12]\n9'
 *
 * 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的
 * target，如果目标值存在返回下标，否则返回 -1。
 * 
 * 
 * 示例 1:
 * 
 * 输入: nums = [-1,0,3,5,9,12], target = 9
 * 输出: 4
 * 解释: 9 出现在 nums 中并且下标为 4
 * 
 * 
 * 示例 2:
 * 
 * 输入: nums = [-1,0,3,5,9,12], target = 2
 * 输出: -1
 * 解释: 2 不存在 nums 中因此返回 -1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 你可以假设 nums 中的所有元素是不重复的。
 * n 将在 [1, 10000]之间。
 * nums 的每个元素都将在 [-9999, 9999]之间。
 * 
 * 
 */

// @lc code=start
class Solution {
public:
    // 左闭右闭区间：[]
/*
    int search(vector<int>& nums, int target) {
        int left = 0;
        int right = nums.size() - 1;
        // left == right是有意义的，所以使用 <=
        while (left <= right) {
            // 避免溢出
            int mid = left + (right-left)/2;
            if (target == nums[mid]) {
                return mid;
            }
            // [left, mid-1]
            if (nums[mid] > target) {
                right = mid - 1;
            }else{
                // [mid+1, right]
                left = mid + 1;
            }
        }
        return -1;
    }
*/

    // 左闭右开区间，[)
    int search(vector<int>& nums, int target) {
        int left = 0;
        int right = nums.size();
        while (left < right) {
            int mid = left + (right-left)/2;
            if (nums[mid] == target) {
                return mid;
            }
            // 左区间 [)
            if (nums[mid] > target) {
                right = mid;
            } else {
                // 右区间
                left = mid + 1;
            }
        }

        return -1;
    }

};
// @lc code=end

