/*
 * @lc app=leetcode.cn id=454 lang=cpp
 *
 * [454] 四数相加 II
 */

// @lc code=start
class Solution {
public:
    int fourSumCount(vector<int>& nums1, vector<int>& nums2, vector<int>& nums3, vector<int>& nums4) {
        int result = 0;
        // 前2个数组遍历求和并保存
        unordered_map<int, int> sum1_map;
        for (int n1 : nums1) {
            for (int n2 : nums2) {
                sum1_map[n1+n2]++;
            }
        }
        // 后2个数组遍历求和
        for (int n3 : nums3) {
            for (int n4 : nums4) {
                // 找 -(n3+n4)
                if (sum1_map.find(-n3-n4) != sum1_map.end()) {
                    result += sum1_map[-n3-n4];
                }
            }
        }
        return result;
    }
};
// @lc code=end

