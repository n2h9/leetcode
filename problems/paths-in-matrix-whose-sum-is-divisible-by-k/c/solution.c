#include <stdlib.h>
#include <stdio.h>

#define modulo 1000000007

int * _init_reminder_cache(int k) {
    int * cache = (int *) malloc(sizeof(int) * k);
    for (int i = 0; i < k; i++) {
        cache[i] = 0;
    }
    return cache;
}

int ** _init_raw_cache(int col_size, int k) {
    // the same as left but for the whole top raw with relation to the current processeing row
    int ** cache = (int **) malloc(sizeof(int *) * col_size);
    for (int i = 0; i < col_size; i++) {
        cache[i] = _init_reminder_cache(k);
    }

    return cache;
}

void free_raw_cache(int ** cache, int size) {
    for (int i = 0; i < size; i++) {
        free(cache[i]);
    }
    free(cache);
}

int numberOfPaths(int** grid, int gridSize, int* gridColSize, int k) {
    int ** top = NULL;
    //  [i][r] -> v, of the current row 
    //  where 
    //       i - index in the current row
    //       r = reminder of the path sum divided to k 
    //       v = number of pathes which give this value of a reminder
    int ** curr = _init_raw_cache(*gridColSize, k);

    for (int i = 0, r = 0; i < *gridColSize; i++) {
        r = (r + grid[0][i]) % k;
        curr[i][r] = 1; // 1 means the only one way, because we go along top row now 
    }

    // lt_reminder is left top reminder
    for (int i = 1, lt_reminder = grid[0][0] % k; i < gridSize; i++) {
        top = curr;
        curr = _init_raw_cache(*gridColSize, k);
        lt_reminder = (lt_reminder + grid[i][0]) % k;
        curr[0][lt_reminder] = 1; // the only path here is from the to

        for (int j = 1; j < *gridColSize; j++) {
            for (int l = 0; l < k; l++) {
                if (top[j][l] != 0) {
                    int r = (l + grid[i][j]) % k;
                    curr[j][r] = (curr[j][r] + top[j][l]) % modulo;
                }
                if (curr[j-1][l] != 0) {
                    int r = (l + grid[i][j]) % k;
                    curr[j][r] = (curr[j][r] + curr[j-1][l]) % modulo;
                }
            }
        }

        free_raw_cache(top, *gridColSize);
    }


    int res = curr[*gridColSize-1][0]; // reminder zero means that path is endede here is divisible on 3
    free_raw_cache(curr, *gridColSize);

    return res;
}