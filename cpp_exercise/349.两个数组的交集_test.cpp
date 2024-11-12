/*
 * @lc app=leetcode.cn id=349 lang=cpp
 *
 * [349] 两个数组的交集
 */

// @lc code=start
class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        // 结果还是利用set来去重
        unordered_set<int> result_set;
        // 通过构造函数进行vector到orderend_set的转换(C++11起)
        unordered_set<int> memb(nums1.begin(), nums1.end());
        for (int i = 0; i < nums2.size(); i++) {
            if (memb.end() != memb.find(nums2[i])) {
                result_set.insert(nums2[i]);
            }
        }

        // 迭代器构造进行转换
        return vector<int>(result_set.begin(), result_set.end());
    }
};
// @lc code=end

