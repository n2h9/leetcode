#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern int* findBall(int** grid, int gridSize, int* gridColSize, int* returnSize);

void compareResults(int* expected, int* result, int arraySize) {
    assert(result != NULL);
    
    for (int i = 0; i < arraySize; i++) {
        printf("%d ", result[i]);
    }
    printf("\n\n"); 
    
    for (int i = 0; i < arraySize; i++) {
        if (expected[i] != result[i]) {
            printf("i = %d: expected[i] (%d) != res[i] (%d)\n", i, expected[i], result[i]);
        }
        assert(expected[i] == result[i]);
    }
}

void testCase0() {
    int rows = 5, cols = 5;
    int matrix0[][5] = {
        {1,1,1,-1,-1},
        {1,1,1,-1,-1},
        {-1,-1,-1,1,1},
        {1,1,1,1,-1},
        {-1,-1,-1,-1,-1}
    };
    int expected[] = {1,-1,-1,-1,-1};

    int** matrix = (int**)malloc(rows * sizeof(int*));
    for (int i = 0; i < rows; i++) {
        matrix[i] = &matrix0[i][0];
    }
    int returnSize;

    int* res = findBall(matrix, rows, &cols, &returnSize);

    compareResults(expected, res, returnSize);

    free(matrix);
    free(res);
}

int main() {
    testCase0();

    printf("ok\n");
    return 0;
}