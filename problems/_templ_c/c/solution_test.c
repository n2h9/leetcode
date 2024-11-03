#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

extern int someFunc(int* array, int size, char* s, bool b);

struct data_entity {
  int expected;

  int* array;
  int size;
  char* s;
  bool b;
};

struct data_entity data_array[] = {
    {0, (int[]){1, 2, 3, 4}, 4, "yellow", true},
    {0, (int[]){5, 6, 7}, 3, "bus", false},
};

int main() {
  size_t length = sizeof(data_array) / sizeof(data_array[0]);

  for (size_t i = 0; i < length; i++) {
    int res = someFunc(data_array[i].array, data_array[i].size, data_array[i].s,
                       data_array[i].b);
    printf("test case #%zu . . .\n", i);
    assert(data_array[i].expected == res);
    printf("test case #%zu is ok\n", i);
  }

  printf("all passed\n");

  return 0;
}