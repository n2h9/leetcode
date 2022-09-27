#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern int rob(int* nums, int numsSize);

void testCase0() {
    int size = 5;
    int array[] = {2,7,9,3,1};
    int expected = 12;

    int res = rob(array, size);

    assert(expected == res);
}

int main() {
    testCase0();

    printf("ok\n");
    return 0;
}