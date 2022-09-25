#include <stdio.h>
#include <assert.h>
#include "./slice.c"

void test_case_0() {
    // initial slice capacity is 2
    slice * s = slice_new(2);

    // insert 16  values in slice, it will increase the slice capacity
    for (int i = 0; i < 16; i++) {
        int * ptr = (int *) malloc(sizeof(int));
        *ptr = i;
        s = slice_append(s, (slice_value_ptr) ptr);
    }

    // check the value we inserted exist
    assert(15 == *(int *)slice_get(s, 15));

    int * ptr = (int *) malloc(sizeof(int));
    *ptr = 128;
    slice_set(s, 15, (slice_value_ptr) ptr);
    assert(128 == *(int *)slice_get(s, 15));

    slice_free_all(s);
}

int main() {
    test_case_0();
    printf("ok\n");

    return 0;
}