#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern int numberOfPaths(int** grid, int gridSize, int* gridColSize, int k);

void testCase0() {
    int grid_size = 3;
    int grid_size_col = 3;
    int k = 3;
    int array[][3] = {{5,2,4,},{3,0,5}, {0,7,2}};
    
    int ** grid = (int **) malloc(sizeof(int*) * grid_size);
    for (int i = 0; i < grid_size; i++) {
        grid[i] = (int*) malloc(sizeof(int) * grid_size_col);
        for (int j = 0; j < grid_size_col; j++) {
            grid[i][j] = array[i][j];
        }
    }

    int expected = 2;

    int res = numberOfPaths(grid, grid_size, &grid_size_col, k);

    for (int i = 0; i < grid_size; i++) {
        free(grid[i]);
    }
    free(grid);

    assert(expected == res);
    printf("expected = %d; res = %d\n", expected, res);
}

void testCase1() {
    int grid_size = 4;
    int grid_size_col = 4;
    int k = 3;
    int array[][4] = {{5,2,4,8},{3,0,5,100}, {0,7,2,29}, {1,2,3,4}};
    
    int ** grid = (int **) malloc(sizeof(int*) * grid_size);
    for (int i = 0; i < grid_size; i++) {
        grid[i] = (int*) malloc(sizeof(int) * grid_size_col);
        for (int j = 0; j < grid_size_col; j++) {
            grid[i][j] = array[i][j];
        }
    }

    int expected = 8;

    int res = numberOfPaths(grid, grid_size, &grid_size_col, k);

    for (int i = 0; i < grid_size; i++) {
        free(grid[i]);
    }
    free(grid);

    // assert(expected == res);
    printf("expected = %d; res = %d\n", expected, res);
}

int main() {
    testCase0();
    testCase1();

    printf("ok\n");
    return 0;
}