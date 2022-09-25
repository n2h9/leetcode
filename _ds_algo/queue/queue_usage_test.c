#include <stdio.h>
#include <assert.h>
#include "./queue.c"

struct _value_c {
    int i, j, t;
};
typedef struct _value_c value_c;

value_c * new_value_c(int i, int j, int t) {
    value_c* v = (value_c *) malloc(sizeof(value_c));
    v->i = i;
    v->j = j;
    v->t = t;

    return v;
}

void testCase0() {
    int i = 2, j = 4, t = 8;

    value_c * vc = new_value_c(i, j, t);
    qnode * qn = new_q_node((qnvalue) vc);
    
    queue * q = new_queue();
    q_push(q, qn);

    qnode * qn0 = q_top(q);
    value_c * vc0 = (value_c *) qn0->value;
    q_pop(q);

    assert(i == vc0->i);
    assert(j == vc0->j);
    assert(t == vc0->t);

    free(vc0);
    free(qn0);
    free(q);
}


int main() {
    testCase0();
    printf("ok\n");

    return 0;
}