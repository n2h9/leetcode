#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include "./uthash.h"

struct item {
    int val;
    UT_hash_handle hh; /* makes this structure hashable */
};

void mapAddVal(struct item **map, int n) {
    struct item *a;
    a = malloc(sizeof(struct item));
    a->val = n;

    HASH_ADD_INT(*map, val, a);
}

bool mapExistVal(struct item **map, int n) {
    struct item *a;

    HASH_FIND_INT(*map, &n, a);

    return a != NULL;
}

void mapDeleteAll(struct item **map) {
  struct item *a, *tmp;

  HASH_ITER(hh, *map, a, tmp) {
    HASH_DEL(*map,a);  /* delete; users advances to next */
    free(a);            /* optional- if you want to free  */
  }
}

bool _isHappyCurrent(int n) {
    return n == 1 || n == 10 || n == 100;
}

int sumOfSquares(int n) {
    int sum = 0;
    for (;n > 0;) {
        int a = n % 10;
        n = n / 10;
        sum += a * a;
        // printf("next n = %d a= %d sum = %d\n", n, a, sum);
    }
    return sum;
}

struct item *happyMap = NULL;
struct item *unhappyMap = NULL;

bool isHappy(int n) {
    // check cached results
    if (mapExistVal(&happyMap, n)) {
        return true;
    }
    // check cached results
    if (mapExistVal(&unhappyMap, n)) {
        return false;
    }

    // save all steps in a map
    struct item *traceMap = NULL;
    
    bool result = false;

    for (;;) {
        // printf("isHappy it: n = %d\n", n);
        if (_isHappyCurrent(n)) {
            result = true;
            break;
        }
        if (mapExistVal(&traceMap, n)) {
            break;
        }
        mapAddVal(&traceMap, n);
        n = sumOfSquares(n);
    }
    // copy all nums from trace map to one of the cached maps
    {
        struct item **cached = &happyMap;
        if (result == false) {
            cached = &unhappyMap;
        }
        for(struct item *a=traceMap; a != NULL; a=a->hh.next) {
            mapAddVal(cached, a->val);
        }
    }

    mapDeleteAll(&traceMap);
    return result;
}