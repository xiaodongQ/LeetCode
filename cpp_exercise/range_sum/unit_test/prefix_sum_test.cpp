#include <gtest/gtest.h>
#include "prefix_sum.h"

TEST(PrefixSumTest, ComputePrefixSum) {
    std::vector<int> arr = {1, 2, 3, 4, 5};
    std::vector<int> expected_sums = {0, 1, 3, 6, 10, 15};
    std::vector<int> sums = computePrefixSum(arr);
    ASSERT_EQ(sums, expected_sums);
}

TEST(PrefixSumTest, RangeSum) {
    std::vector<int> arr = {1, 2, 3, 4, 5};
    std::vector<int> sums = computePrefixSum(arr);
    EXPECT_EQ(rangeSum(sums, 0, 1), 3); // 从0到1
    EXPECT_EQ(rangeSum(sums, 1, 3), 9); // 从1到3
    EXPECT_EQ(rangeSum(sums, 2, 4), 12); // 从2到4
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
