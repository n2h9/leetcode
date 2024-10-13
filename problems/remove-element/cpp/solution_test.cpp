#include "solution.cpp"

#include <gtest/gtest.h>

#include <vector>

struct data_entity {
  int expected;

  std::vector<int> nums;
  int val;
};

data_entity data_array[] = {
    {2, std::vector<int>{3, 2, 2, 3}, 3},
    {5, std::vector<int>{0, 1, 2, 2, 3, 0, 4, 2}, 2},
    {0, std::vector<int>{8, 8, 8, 8, 8, 8, 8, 8}, 8},
    {6, std::vector<int>{8, 7, 7, 7, 7, 7, 7, 8}, 8},
    {6, std::vector<int>{7, 7, 7, 8, 8, 7, 7, 7}, 8},
    {8, std::vector<int>{9, 9, 9, 9, 9, 9, 9, 9}, 8},
    {0, std::vector<int>{}, 2},
};

TEST(SolutionTest, BasicAssertions) {
  auto solution = Solution{};

  for (auto data : data_array) {
    auto k = solution.removeElement(data.nums, data.val);
    EXPECT_EQ(data.expected, k);
    for (auto i = 0; i < k; i++) {
      // assert that first k elements != val
      EXPECT_NE(data.val, data.nums[i]);
    }
  }
}
