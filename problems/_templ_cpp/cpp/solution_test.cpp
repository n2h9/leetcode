#include "solution.cpp"

#include <gtest/gtest.h>

TEST(SolutionTest, BasicAssertions) {
  // Expect two strings not to be equal.
  EXPECT_STRNE("hello", "world");

  auto solution = Solution{};

  EXPECT_EQ(42, solution.solve(40, 2));
}
