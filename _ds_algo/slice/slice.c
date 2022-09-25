#include <stdlib.h>

// pointer to the value which is stored in a slice
typedef void* slice_value_ptr;
// ----

struct _slice {
    slice_value_ptr* storage;
    int length;
    int capacity;
};
typedef struct _slice slice;

int slice_len(slice * s) {
    return s->length;
}

int slice_capacity(slice * s) {
    return s->capacity;
}

slice * slice_new(int capacity) {
    if (capacity <= 0) {
        // default capacity is 32 
        capacity = 1 << 5;
    }
    slice * s = (slice *) malloc(sizeof(slice));
    s->capacity = capacity;
    s->length = 0;
    s->storage = (slice_value_ptr *) malloc(capacity * sizeof(slice_value_ptr));

    return s;
}

// free only slice memory allocated: slice and slice->storage
void _slice_free(slice * s) {
    free(s->storage);
    free(s);
}

// free also the memory of the slice items
void slice_free_all(slice * s) {
    for (int i = 0, l = slice_len(s); i < l; i++) {
        free(s->storage[i]);
    }
    _slice_free(s);
}


// i must be < s->length
slice_value_ptr slice_get(slice * s, int i) {
    return s->storage[i];
}

// i must be < s->length
void slice_set(slice * s, int i, slice_value_ptr value) {
    if (s->storage[i] != NULL) {
        free(s->storage[i]);
    }
    s->storage[i] = value;
}

slice * _slice_increase_size(slice *s);

// returns new link to slice
// if slice was increased in size
// old memory is freed
slice * slice_append(slice *s, slice_value_ptr value) {
    if (s->length >= s->capacity) { // when s->length == s.capacity allocate new memory
       s = _slice_increase_size(s);
    }

    s->storage[s->length] = value;
    s->length++;

    return s;
}

slice * _slice_increase_size(slice *s) {
    slice * s_new = slice_new(
        slice_capacity(s) << 1 // double the capacity
    );

    for (int i = 0, l = slice_len(s); i < l; i++) {
        s_new = slice_append(s_new, slice_get(s, i));
    }

    _slice_free(s);

    return s_new;
}

