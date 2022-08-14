#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern long long numberOfWays(char * s);

void testCase0() {
    char s[] = "001101";
    long long expected = 6;

    long long res = numberOfWays(s);

    assert(expected == res);
}

void testCase1() {
    char s[] = "0101010101010101010101010101010101010101010101010101010101010111111111111111100000000000000000000010101010101010101010001110001101010101010101";
    long long expected = 116942;

    long long res = numberOfWays(s);

    assert(expected == res);
}

int main() {
    testCase0();
    testCase1();

    printf("ok\n");
    return 0;
}