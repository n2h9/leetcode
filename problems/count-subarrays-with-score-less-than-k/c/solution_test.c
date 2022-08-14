#include <stdio.h>
#include <assert.h>

extern long long countSubarrays(int* nums, int numsSize, long long k);

void testCase0() {
    int nums[] = {2,1,4,3,5};
    int numsSize = 5;
    long long k = 10;
    
    long long expected = 6;

    long long res = countSubarrays(nums, numsSize, k);

    assert(expected == res);
}

void testCase1() {
    int nums[] = {1,1,1};
    int numsSize = 3;
    long long k = 5;
    
    long long expected = 5;

    long long res = countSubarrays(nums, numsSize, k);

    assert(expected == res);
}

void testCase2() {
    int nums[] = {1,2,1,2,1,3,1,2};
    int numsSize = 8;
    long long k = 16;
    
    long long expected = 19;

    long long res = countSubarrays(nums, numsSize, k);

    assert(expected == res);
}

int main() {
    testCase0();
    testCase1();

    printf("ok\n");
    return 0;
}