/*
 * @lc app=leetcode.cn id=15 lang=cpp
 *
 * [15] 三数之和
 */

// @lc code=start
class Solution {
public:
    // 自定义哈希函数
    struct VecHash {
        size_t operator()(const std::vector<int>& v) const {
            std::hash<int> hasher;
            size_t seed = 0;
            for (int i : v) {
                seed ^= hasher(i) + 0x9e3779b9 + (seed<<6) + (seed>>2);
            }
            return seed;
        }
    };
    // 自定义比较函数，用于确保向量相等时哈希值也相同
    struct VecEqual {
        bool operator()(const std::vector<int>& v1, const std::vector<int>& v2) const {
            return v1 == v2;
        }
    };

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
            unordered_map<int, int> tmp_map;
            for (int j = i+1; j < nums.size(); j++) {
                if (tmp_map.find(nums[j]) != tmp_map.end()) {
                    // 满足三数之和为0，按大小顺序添加记录
                    result.push_back(vector<int>{nums[i], nums[ tmp_map[nums[j]] ], nums[j]});
                } else {
                    tmp_map[sum - nums[j]] = j;
                }
            }
        }
        // 当前结果还需要去重处理
        unordered_set<vector<int>, VecHash, VecEqual> tmp_set(result.begin(), result.end());
        // unordered_set再转换为vector
        vector<vector<int>> result_nums(tmp_set.begin(), tmp_set.end());

        return result_nums;
    }
};

/*
class Solution {
public:
    // 自定义哈希函数
    struct VecHash {
        size_t operator()(const std::vector<int>& v) const {
            std::hash<int> hasher;
            size_t seed = 0;
            for (int i : v) {
                seed ^= hasher(i) + 0x9e3779b9 + (seed<<6) + (seed>>2);
            }
            return seed;
        }
    };
    // 自定义比较函数，用于确保向量相等时哈希值也相同
    struct VecEqual {
        bool operator()(const std::vector<int>& v1, const std::vector<int>& v2) const {
            return v1 == v2;
        }
    };

    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> result;
        // 外层遍历，里层在剩下的成员里找2数之和
        for (int i = 0; i < nums.size(); i++) {
            int sum = -nums[i];
            unordered_map<int, int> tmp_map;
            for (int j = i+1; j < nums.size(); j++) {
                if (tmp_map.find(nums[j]) != tmp_map.end()) {
                    // 满足求和为0，记录成员
                    result.push_back(vector<int>{nums[i], nums[tmp_map[nums[j]]], nums[j]});
                } else {
                    tmp_map[sum - nums[j]] = j;
                }
            }
        }

        // 三值排序
        for (int i = 0; i < result.size(); i++) {
            sort(result[i].begin(), result[i].end());
        }
        // 自定义去重
        unordered_set<vector<int>, VecHash, VecEqual> tmp_set(result.begin(), result.end());
        // unordered_set和vector转换
        vector<vector<int>> result_nums(tmp_set.begin(), tmp_set.end());

        return result_nums;
    }
};
*/
// @lc code=end

