#include <stdlib.h>

struct _qnvalue {
    int i, j;
    int t;
};
typedef struct _qnvalue qnvalue;

struct _qnode {
    qnvalue * value;
    struct _qnode * next;
};
typedef struct _qnode qnode;

typedef struct {
    qnode * head;
    qnode * tail;
} queue;

queue * newQueue() {
    queue * q = (queue *)malloc(sizeof(queue));
    q->head = NULL;
    q->tail = NULL;
    return q;
}

unsigned char q_isEmpty(queue * q) {
    if (q->head == NULL) {
        return 1;
    }
    return 0;
}

void q_push(queue * q, qnode * node) {
    node->next = NULL;
    if (q_isEmpty(q)) {
        q->head = node;
        q->tail = node;
        return ;
    }

    q->tail->next = node;
    q->tail = q->tail->next;
}

void q_pop(queue * q) {
    qnode * oldHead = q->head;
    q->head = q->head->next;
    oldHead->next = NULL;

    if (q->head == NULL) {
        q->tail = NULL;
    }
}

// do not check on empty queue
qnode * q_top(queue * q) {
    return q->head;
}


qnode * newQNode(qnvalue* value) {
    qnode* node = (qnode *)malloc(sizeof(qnode));
    node->value = value;

    return node;
}

qnvalue * newQNValue(int i, int j, int t) {
    qnvalue* value = (qnvalue *) malloc(sizeof(qnvalue));
    value->i = i;
    value->j = j;
    value->t = t;

    return value;
}