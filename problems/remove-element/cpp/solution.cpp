using namespace ::std;

#include <vector>

class Solution {
 public:
  int removeElement(vector<int>& nums, int val) {
    // explicitly specify index as int as index could go below 0
    auto search_for_next_not_val = [&nums, val](int index) -> auto {
      while (index >= 0 && nums[index] == val) {
        index--;
      }
      return index;
    };

    auto search_for_next_val = [&nums, val](int index) -> auto {
      while (index < (int)nums.size() && nums[index] != val) {
        index++;
      }
      return index;
    };

    auto do_swap = [&nums](auto i, auto j) { swap(nums[i], nums[j]); };

    auto next_from_end_not_val_index = search_for_next_not_val(nums.size() - 1);
    auto next_val_index = search_for_next_val(0);

    while (next_from_end_not_val_index > next_val_index) {
      do_swap(next_from_end_not_val_index, next_val_index);

      next_from_end_not_val_index =
          search_for_next_not_val(next_from_end_not_val_index);
      next_val_index = search_for_next_val(next_val_index);
    }

    // corener case, when all elements are eq to val
    // next_from_end_not_val_index = -1
    // res = -1 + 1 = 0
    return next_from_end_not_val_index + 1;
  }
};