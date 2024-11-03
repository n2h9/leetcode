#include <assert.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

extern int doesAliceWin(char* s);

struct data_entity {
  bool expected;

  char* s;
};

struct data_entity data_array[] = {
    {true, "leetcoder"},
    {true, "ololo"},
    {false, "bbcd"},
};

int main() {
  size_t length = sizeof(data_array) / sizeof(data_array[0]);
  for (size_t i = 0; i < length; i++) {
    bool res = doesAliceWin(data_array[i].s);
    printf("test case #%zu . . .\n", i);
    assert(data_array[i].expected == res);
    printf("test case #%zu is ok\n", i);
  }
  printf("all passed\n");
  return 0;
}