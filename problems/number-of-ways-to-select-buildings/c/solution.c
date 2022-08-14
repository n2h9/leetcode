#include <stdlib.h>
#include <stdio.h>

long long numberOfWays(char * s){
    char one = '1', zero = '0';
    // contains the counts for the sequences of chars
    // count[0] -> number of "0" in a string
    // count[1] -> number of "1" in a string
    // count[2] -> number of "01" in a string
    // count[3] -> number of "10" in a string
    long long count [4] = {0,0,0,0};
    long long res = 0;

    for (int i = 0; s[i] != '\0'; i++) {
        if (s[i] == zero) {
            res += count[2];
            count[3] += count[1];
            count[0]++;
        } else if (s[i] == one) {
            res += count[3];
            count[2] += count[0];
            count[1]++;
        } else {
            printf("incorrect input string\n");
            exit(1);
        }
        
    }

    return res;
}