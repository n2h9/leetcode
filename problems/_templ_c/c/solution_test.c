#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern int someFunc(int* array, int size);

void testCase0() {
    int size = 5;
    int array[] = {3,3,4,2,3};
    int expected = 0;

    int res = someFunc(array, size);

    assert(expected == res);
}

int main() {
    testCase0();

    printf("ok\n");
    return 0;
}