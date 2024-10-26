// prefix_sum.h

#ifndef PREFIX_SUM_H
#define PREFIX_SUM_H

#include <vector>

// 计算前缀和
std::vector<int> computePrefixSum(const std::vector<int>& arr);

// 根据前缀和计算区间和
int rangeSum(const std::vector<int>& sums, int left, int right);

#endif // PREFIX_SUM_H
