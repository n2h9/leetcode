#include "solution.cpp"

#include <gtest/gtest.h>

#include <vector>

struct data_entity {
  int expected;

  int a;
  int b;
};

data_entity data_array[] = {
    {42, 40, 2},
    {42, 38, 4},
    {42, 36, 6},
};

TEST(SolutionTest, BasicAssertions) {
  auto solution = Solution{};

  for (auto data : data_array) {
    auto result = solution.solve(data.a, data.b);
    EXPECT_EQ(data.expected, result);
  }
}
