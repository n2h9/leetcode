using namespace ::std;

#include <algorithm>
#include <vector>

class Solution {
 public:
  int findValueOfPartition(vector<int>& nums) {
    sort(nums.begin(), nums.end());

    auto min = nums[1] - nums[0];
    for (size_t i = 2; i < nums.size(); i++) {
      auto diff = nums[i] - nums[i - 1];
      if (diff < min) {
        min = diff;
      }
    }

    return min;
  }
};