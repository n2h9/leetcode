#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include "./uthash.h"

struct item {
    int key;
    int size;

    UT_hash_handle hh; /* makes this structure hashable */
};

void map_add_val(struct item **map, int key, int pair[2]) {
    struct item *a;
    HASH_FIND_INT(*map, &key, a);
    if (a == NULL) {
        a = (struct item *) malloc(sizeof(struct item));
        a->key = key; 
        a->size = 0;

        HASH_ADD_INT(*map, key, a);   
    }

    a->size++;

}

struct item * map_get_val(struct item **map, int key) {
    struct item *a;

    HASH_FIND_INT(*map, &key, a);

    return a;
}

void map_delete_all(struct item **map) {
  struct item *a, *tmp;

  HASH_ITER(hh, *map, a, tmp) {
    HASH_DEL(*map,a);  /* delete; users advances to next */

    free(a);            /* optional- if you want to free  */
  }
}

struct item *sum_map_1_2 = NULL;
struct item *sum_map_3_4 = NULL;

int fourSumCount(int* nums1, int nums1Size, int* nums2, int nums2Size, int* nums3, int nums3Size, int* nums4, int nums4Size) {
    for (int i = 0; i < nums1Size; i++) {
        for (int j = 0; j < nums1Size; j++) {
            int sum12 = nums1[i] + nums2[j];
            int pair[2] = {i,j};
            map_add_val(&sum_map_1_2, sum12, pair);

            int sum34 = nums3[i] + nums4[j];
             map_add_val(&sum_map_3_4, sum34, pair);
        }
    }

    int result = 0;

    {
        struct item *a, *tmp;
        HASH_ITER(hh, sum_map_1_2, a, tmp) {
            struct item * b = map_get_val(&sum_map_3_4, (-1) * a->key);
            if (b != NULL) {
                result += a->size * b->size; 
            }
        }
    }
    

    map_delete_all(&sum_map_1_2);
    map_delete_all(&sum_map_3_4);
    
    return result;
}