#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern int orangesRotting(int** grid, int gridSize, int* gridColSize);

void testCase0() {
    int size = 3;
    int colSize = 3;
    int array[][3] = {{2,1,1}, {1,1,0},{0,1,1}};
    int expected = 4;

    int** grid = (int **) malloc(size * sizeof(int *));
    for (int i = 0; i < size; i++) {
        grid[i] = array[i];
    }

    int res = orangesRotting(grid, size, &colSize);

    free(grid);

    assert(expected == res);
}

void testCase1() {
    int size = 1;
    int colSize = 2;
    int array[][2] = {{0,2}};
    int expected = 0;

    int** grid = (int **) malloc(size * sizeof(int *));
    for (int i = 0; i < size; i++) {
        grid[i] = array[i];
    }

    int res = orangesRotting(grid, size, &colSize);

    free(grid);

    assert(expected == res);
}

void testCase2() {
    int size = 3;
    int colSize = 3;
    int array[][3] = {{2,1,1}, {0,1,1},{1,0,1}};
    int expected = -1;

    int** grid = (int **) malloc(size * sizeof(int *));
    for (int i = 0; i < size; i++) {
        grid[i] = array[i];
    }

    int res = orangesRotting(grid, size, &colSize);

    free(grid);

    assert(expected == res);
}

int main() {
    testCase0();
    testCase1();
    testCase2();

    printf("ok\n");
    return 0;
}