#include <stdlib.h>
#include <stdio.h>
#include "queue.c"

#define ORANGE 1
#define ORANGE_ROTTEN 2

int orangesRotting(int** grid, int gridSize, int* gridColSize) {
    queue * q = newQueue();
    
    // total number of orranges
    int num = 0;
    for (int i = 0; i < gridSize; i++) {
        for (int j = 0; j < *gridColSize; j++) {
            if (grid[i][j] == ORANGE_ROTTEN) {
                q_push(q, newQNode(newQNValue(i, j, 0)));
                num++;
            } else if (grid[i][j] == ORANGE) {
                num++;
            }
        }   
    }

    // total number of rotten orranges
    int rnum = 0;

    // current time tick
    int time = 0;
    for (;!q_isEmpty(q);) {
        qnode * node = q_top(q);
        qnvalue * value = node->value;
        q_pop(q);

        time = value->t;
        rnum++;

        int i = value->i, j = value->j;
        // possible coordinates in grid which could become rotten
        int c[4][2] = {{i,j-1}, {i,j+1},{i-1,j},{i+1,j}};
        for (int k = 0; k < 4; k++) {
            if (c[k][0] >= 0 && c[k][0] < gridSize && c[k][1] >= 0 && c[k][1] < *gridColSize) {
                if (grid[c[k][0]][c[k][1]] == ORANGE) {
                    q_push(q, newQNode(newQNValue(c[k][0], c[k][1], time+1)));
                    grid[c[k][0]][c[k][1]] = ORANGE_ROTTEN;
                }
            }
        }


        free(value);
        free(node);
    }

    free(q);

    // not all oranges became rotten
    if (rnum != num) {
        return -1;
    } 

    return time;
}