#include "solution.cpp"

#include <gtest/gtest.h>

#include <vector>

struct data_entity {
  int expected;

  vector<int> nums;
};

data_entity data_array[] = {
    1, {1, 3, 2, 4}, 9,   {100, 1, 10},
    0, {2, 2},       111, {1, 121212, 555, 444, 333, 666},
};

TEST(SolutionTest, BasicAssertions) {
  auto solution = Solution{};

  for (auto data : data_array) {
    auto result = solution.findValueOfPartition(data.nums);
    EXPECT_EQ(data.expected, result);
  }
}
