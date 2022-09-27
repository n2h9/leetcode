#include <stdlib.h>


int max(int a, int b) {
    if (a > b) {
        return a;
    }
    return b;
}

int rob(int* nums, int numsSize){
    if (numsSize >= 2) {
        nums[1] = max(nums[0], nums[1]); 
    }
    for (int i = 2; i < numsSize; i++) {
        nums[i] = max(nums[i-1], nums[i-2]+nums[i]);
    }
    return nums[numsSize-1];
}