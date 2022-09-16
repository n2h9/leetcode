#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <assert.h>

extern bool isHappy(int n);

void testCase0() {
    int n = 19;
    bool expected = true;

    int res = isHappy(n);

    assert(expected == res);
}

void testCase1() {
    int n = 2;
    bool expected = false;

    int res = isHappy(n);

    assert(expected == res);
}

void testCase2() {
    int n = 7;
    bool expected = true;

    int res = isHappy(n);

    assert(expected == res);
}

void testCase3() {
    int n = 23;
    bool expected = true;

    int res = isHappy(n);

    assert(expected == res);
}


int main() {
    testCase0();
    testCase1();
    testCase2();
    testCase3();

    printf("ok\n");
    return 0;
}