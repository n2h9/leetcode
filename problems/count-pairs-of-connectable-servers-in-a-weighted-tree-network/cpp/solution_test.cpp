#include "solution.cpp"

#include <gtest/gtest.h>

#include <vector>

struct data_entity {
  vector<int> expected;

  vector<vector<int>> edges;
  int signal_speed;
};

data_entity data_array[] = {{
    {2, 0, 0, 0, 0, 0, 2},

    {{0, 6, 3}, {6, 5, 3}, {0, 3, 1}, {3, 2, 7}, {3, 1, 6}, {3, 4, 2}},
    3,
}};

TEST(SolutionTest, BasicAssertions) {
  auto solution = Solution{};

  for (auto data : data_array) {
    auto result =
        solution.countPairsOfConnectableServers(data.edges, data.signal_speed);
    EXPECT_EQ(data.expected, result);
  }
}
