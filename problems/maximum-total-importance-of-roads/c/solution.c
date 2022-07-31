#include <stdlib.h>

int cmpfunc (const void * a, const void * b) {
   return ( *(int*)a - *(int*)b );
}

int reverseCMPFunc(const void * a, const void * b) {
   return (-1) * cmpfunc(a, b);
}

long long maximumImportance(int n, int** roads, int roadsSize, int* roadsColSize){
    int *numOfConnectedPoints = (int *)malloc(n * sizeof(int));
    for (int i = 0; i < n; i++) {
        numOfConnectedPoints[i] = 0;
    }

    for (int i = 0; i < roadsSize; i++) {
        numOfConnectedPoints[roads[i][0]]++;
        numOfConnectedPoints[roads[i][1]]++;
    }

    qsort(numOfConnectedPoints, n, sizeof(int), reverseCMPFunc);

    long long res = 0;
    for (int i = 0, m = n; i < n; i++, m--) {
        res += (long long) numOfConnectedPoints[i] * m;
    }

    return res;
}