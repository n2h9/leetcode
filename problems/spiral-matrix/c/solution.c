#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* spiralOrder(int** matrix, int matrixSize, int* matrixColSize, int* returnSize){  
    // calculate return size
    (*returnSize) = matrixSize * *matrixColSize;

    int* res = (int*) malloc(sizeof(int) * *returnSize);
    // left bound, right bound, top bound, bottom bound.
    int lb = 0, rb = *matrixColSize - 1, tb = 0, bb = matrixSize - 1;
    
    for (int i = 0; i < *returnSize;) {
        // from left to right
        for (int j = lb; j <= rb && i < *returnSize; j++, i++) {
            res[i] = matrix[tb][j];
        }
        tb++;

        // from top to bottom
        for (int j = tb; j <= bb && i < *returnSize; j++, i++) {
            res[i] = matrix[j][rb];
        }
        rb--;
        // from right to left
        for (int j = rb; j >= lb && i < *returnSize; j--, i++) {
            res[i] = matrix[bb][j];
        }
        bb--;
        // from bottom to top
        for (int j = bb; j >= tb && i < *returnSize; j--, i++) {
            res[i] = matrix[j][lb];
        }
        lb++;    
    }

    return res;
}