#include <stdlib.h>

struct _qnode {
    int value;
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


qnode * newQNode(int value) {
    qnode* node = (qnode *)malloc(sizeof(qnode));
    node->value = value;
    node->next = NULL;

    return node;
}

int longestCycle(int* edges, int edgesSize){
    int *parentsNumber = (int *)malloc(edgesSize * sizeof(int));
    for (int i = 0; i < edgesSize; i++) {
        parentsNumber[i] = 0;
    }
    for (int i = 0; i < edgesSize; i++) {
        if (edges[i] == -1) {
            continue;
        }
        parentsNumber[edges[i]]++;
    }

    queue * q = newQueue();
    for (int i = 0; i < edgesSize; i++) {
        if (parentsNumber[i] == 0) {
            q_push(q, newQNode(i));
        }    
    }

    while(!q_isEmpty(q)) {
        qnode * node = q_top(q);
        q_pop(q);
        if (edges[node->value] == -1) {
            continue;
        }
        parentsNumber[edges[node->value]]--;
        if (parentsNumber[edges[node->value]] <= 0) {
            q_push(q, newQNode(edges[node->value]));
        }
    }

    int res = -1;
    
    for (int i = 0; i < edgesSize; i++) {
        if (parentsNumber[i] <= 0) {
            continue;
        }
        int len = 0;
        int j = i;
        do {
            len++;
            j = edges[j];
            parentsNumber[j]--;
        } while (j != i);
        if (len > res) {
            res = len;
        }   
    }

    return res;
}