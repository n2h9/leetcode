#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern int longestCycle(int* edges, int edgesSize);

void testCase0() {
    int edgesSize = 5;
    int edges[] = {3,3,4,2,3};
    int expected = 3;

    int res = longestCycle(edges, edgesSize);

    assert(expected == res);
}

void testCase1() {
    int edgesSize = 4;
    int edges[] = {2,-1,3,1};
    int expected = -1;

    int res = longestCycle(edges, edgesSize);

    assert(expected == res);
}

int main() {
    testCase0();
    testCase1();

    printf("ok\n");
    return 0;
}