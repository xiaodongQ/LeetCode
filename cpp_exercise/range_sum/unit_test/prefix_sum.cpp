#include "prefix_sum.h"

// 计算前缀和
std::vector<int> computePrefixSum(const std::vector<int>& arr) {
    std::vector<int> sums(arr.size() + 1, 0); // 初始化前缀和数组
    for (size_t i = 0; i < arr.size(); ++i) {
        sums[i + 1] = sums[i] + arr[i];
    }
    return sums;
}

// 根据前缀和计算区间和
int rangeSum(const std::vector<int>& sums, int left, int right) {
    return sums[right + 1] - sums[left];
}
