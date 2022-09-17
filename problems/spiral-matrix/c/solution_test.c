#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern int* spiralOrder(int** matrix, int matrixSize, int* matrixColSize, int* returnSize);

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
    int rows = 3, cols = 3;
    int matrix0[][3] = {{1,2,3},{4,5,6},{7,8,9}};
    int expected[] = {1,2,3,6,9,8,7,4,5};

    int** matrix = (int**)malloc(rows * sizeof(int*));
    for (int i = 0; i < rows; i++) {
        matrix[i] = &matrix0[i][0];
    }
    int returnSize = rows * cols;

    int* res = spiralOrder(matrix, rows, &cols, &returnSize);

    compareResults(expected, res, returnSize);

    free(matrix);
    free(res);
}

void testCase1() {
    int rows = 3, cols = 4;
    int matrix0[][4] = {{1,2,3,4},{5,6,7,8},{9,10,11,12}};
    int expected[] = {1,2,3,4,8,12,11,10,9,5,6,7};

    int** matrix = (int**)malloc(rows * sizeof(int*));
    for (int i = 0; i < rows; i++) {
        matrix[i] = &matrix0[i][0];
    }
    int returnSize = rows * cols;

    int* res = spiralOrder(matrix, rows, &cols, &returnSize);

    compareResults(expected, res, returnSize);

    free(matrix);
    free(res);
}

int main() {
    testCase0();
    testCase1();

    printf("ok\n");
    return 0;
}