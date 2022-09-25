#include <stdlib.h>

// pointer to the value which is stored in a queue
typedef void * qnvalue;
// ----


// ----
// queue node struct
// encapsulate node

struct _qnode {
    qnvalue value;
    struct _qnode * next;
};
typedef struct _qnode qnode;

qnode * new_q_node(qnvalue value) {
    qnode* node = (qnode *)malloc(sizeof(qnode));
    node->value = value;

    return node;
}

// end of queue node struct
// ----

// ----
// queue
// all bellow relate to queue

typedef struct {
    qnode * head;
    qnode * tail;
} queue;

queue * new_queue() {
    queue * q = (queue *)malloc(sizeof(queue));
    q->head = NULL;
    q->tail = NULL;
    return q;
}

unsigned char q_is_empty(queue * q) {
    if (q->head == NULL) {
        return 1;
    }
    return 0;
}

void q_push(queue * q, qnode * node) {
    node->next = NULL;
    if (q_is_empty(q)) {
        q->head = node;
        q->tail = node;
        return ;
    }

    q->tail->next = node;
    q->tail = q->tail->next;
}

void q_pop(queue * q) {
    qnode * old_head = q->head;
    q->head = q->head->next;
    old_head->next = NULL;

    if (q->head == NULL) {
        q->tail = NULL;
    }
}

// do not check on empty queue
qnode * q_top(queue * q) {
    return q->head;
}