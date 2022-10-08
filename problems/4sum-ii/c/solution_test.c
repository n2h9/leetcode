#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern int fourSumCount(int* nums1, int nums1Size, int* nums2, int nums2Size, int* nums3, int nums3Size, int* nums4, int nums4Size);

void testCase0() {
    int size = 2;
    int nums1[] = {1,2};
    int nums2[] = {-2,-1};
    int nums3[] = {-1,2};
    int nums4[] = {0,2};
    int expected = 2;

    int res = fourSumCount(nums1, size, nums2, size, nums3, size, nums4, size);

    assert(expected == res);
}

int main() {
    testCase0();

    printf("ok\n");
    return 0;
}