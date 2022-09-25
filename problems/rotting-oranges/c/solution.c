#include <stdlib.h>
#include <stdio.h>
#include "queue.c"

// ----
// _value_c is a value container, holds the value of node in queue

struct _value_c {
    int i, j;
    int t;
};
typedef struct _value_c value_c;

value_c * new_value_c(int i, int j, int t) {
    value_c* v = (value_c *) malloc(sizeof(value_c));
    v->i = i;
    v->j = j;
    v->t = t;

    return v;
}

// ----

#define ORANGE 1
#define ORANGE_ROTTEN 2

int orangesRotting(int** grid, int gridSize, int* gridColSize) {
    queue * q = new_queue();
    
    // total number of orranges
    int num = 0;
    for (int i = 0; i < gridSize; i++) {
        for (int j = 0; j < *gridColSize; j++) {
            if (grid[i][j] == ORANGE_ROTTEN) {
                q_push(q, new_q_node((void *)new_value_c(i, j, 0)));
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
    for (;!q_is_empty(q);) {
        qnode * node = q_top(q);
        value_c * v = (value_c *)node->value;
        q_pop(q);

        time = v->t;
        rnum++;

        int i = v->i, j = v->j;
        // possible coordinates in grid which could become rotten
        int c[4][2] = {{i,j-1}, {i,j+1},{i-1,j},{i+1,j}};
        for (int k = 0; k < 4; k++) {
            if (c[k][0] >= 0 && c[k][0] < gridSize && c[k][1] >= 0 && c[k][1] < *gridColSize) {
                if (grid[c[k][0]][c[k][1]] == ORANGE) {
                    q_push(q, new_q_node((void *)new_value_c(c[k][0], c[k][1], time+1)));
                    grid[c[k][0]][c[k][1]] = ORANGE_ROTTEN;
                }
            }
        }


        free(v);
        free(node);
    }

    free(q);

    // not all oranges became rotten
    if (rnum != num) {
        return -1;
    } 

    return time;
}