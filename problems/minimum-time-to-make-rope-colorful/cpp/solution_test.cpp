#include "solution.cpp"

#include <gtest/gtest.h>

#include <string>
#include <vector>

struct data_entity {
  int expected;

  string colors;
  vector<int> neededTime;
};

data_entity data_array[] = {{3, "abaac", {1, 2, 3, 4, 5}},
                            {0, "abc", {1, 2, 3}},
                            {2, "aabaa", {1, 2, 3, 4, 1}},
                            {0, "d", {10000}}};

TEST(SolutionTest, BasicAssertions) {
  auto solution = Solution{};

  for (auto data : data_array) {
    auto result = solution.minCost(data.colors, data.neededTime);
    EXPECT_EQ(data.expected, result);
  }
}
