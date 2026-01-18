using namespace ::std;

#include <string>
#include <vector>

class Solution {
 public:
  int minCost(string colors, vector<int>& neededTime) {
    auto totalTime = 0;

    auto isSameColor = false;
    auto recentSameColorSum = neededTime[0];
    auto recentMaxTime = neededTime[0];
    for (auto i = 1; i < colors.size(); i++) {
      auto wasSameColor = isSameColor;
      isSameColor = colors[i] == colors[i - 1];
      if (isSameColor) {
        // if balloons are of the same color, track max time and total sum.
        recentMaxTime = max(recentMaxTime, neededTime[i]);
        recentSameColorSum += neededTime[i];
      } else {
        if (wasSameColor) {
          // if there was a strick of the same color, add recentSameColorSum
          // without recentMaxTime to total time.
          totalTime += recentSameColorSum - recentMaxTime;
        }
        // prepare for a new strick
        recentMaxTime = neededTime[i];
        recentSameColorSum = neededTime[i];
      }
    }

    // if we ended up with the strick of the same color, add recentSameColorSum
    // without recentMaxTime to total time.
    if (isSameColor) {
      totalTime += recentSameColorSum - recentMaxTime;
    }

    return totalTime;
  }
};