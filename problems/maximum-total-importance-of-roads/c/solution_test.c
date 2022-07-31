#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern long long maximumImportance(int n, int** roads, int roadsSize, int* roadsColSize);

void testCase0() {
    int n = 5;
    int roads0[][2] = {{0,1},{1,2},{2,3},{0,2},{1,3},{2,4}};
    int rows = 6, cols = 2;

    int** roads = (int**)malloc(rows * sizeof(int*));
    for (int i = 0; i < rows; i++) {
        roads[i] = &roads0[i][0];
    }

    long long res = maximumImportance(n, roads, rows, &cols);

    assert(43 == res);
}

int main() {
    testCase0();

    printf("ok\n");
    return 0;
}