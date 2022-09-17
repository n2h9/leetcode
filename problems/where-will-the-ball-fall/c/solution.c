#include <stdlib.h>

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* findBall(int** grid, int gridSize, int* gridColSize, int* returnSize){
    *returnSize = *gridColSize;

    int* result = (int*) malloc(sizeof(int) * *returnSize);

    for (int n = 0; n < *returnSize; n++) {
        result[n] = -1;
        for (int i = 0, j = n; i < gridSize && j < *gridColSize;) {
            if (grid[i][j] == 1) { // case when current cell is eq to '\'
                if ((j+1) >= *gridColSize) { // case when right cell is border
                    // the ball will stuck
                    break;
                }
                if (grid[i][j+1] == -1) { // case when right cell is '/'
                    // the ball will stuck
                    break;
                }
                // the only case when the ball fall further is (grid[i][j+1] == 1)
                i++;
                j++;
            } else { // case when current cell is eq to '/'
                if ((j-1) < 0) { // case when left cell is border
                    // the ball will stuck
                    break;
                }
                if (grid[i][j-1] == 1) { // case when left cell is '\'
                    // the ball will stuck
                    break;
                }
                // the only case when the ball fall further is (grid[i][j-1] == -1)
                i++;
                j--;
            }
            // if it the last row, than save the column to the result array
            if (i >= gridSize) {
                result[n] = j;
                break;
            }
            // otherwise check the next cell
        }
    }

    return result;
}