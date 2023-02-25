#include <stdio.h>
#include <stdlib.h>
#include <assert.h>

extern int* vowelStrings(char **words, int wordsSize, int **queries, int queriesSize, int *queriesColSize, int *returnSize);

void testCase0() {
    char wordsArray[5][40] = {"aba","bcb","ece","aa","e"};
    int wordsSize = sizeof(wordsArray) / (sizeof(wordsArray[0]));
    char **words = (char **) malloc(sizeof(char *) * wordsSize);
    for (int i = 0; i < wordsSize; i++) {
        words[i] = wordsArray[i];
    }

    int queriesArray[3][2] = {{0,2},{1,4},{1,1}};
    int queriesColSize = sizeof(queriesArray[0]);
    int queriesSize = sizeof(queriesArray) / queriesColSize;
    int **queries = (int **) malloc(sizeof(int *) * queriesSize);
    for (int i = 0; i < queriesSize; i++) {
        queries[i] = queriesArray[i];
    }

    int expected[] = {2,3,0}; 

    int resSize;
    int* res = vowelStrings(words, wordsSize, queries, queriesSize, &queriesColSize, &resSize);

    for (int i = 0; i < resSize; i++) {
        assert(expected[i] == res[i]);    
    }

    free(res);
    free(queries);
    free(words);  
}

void testCase1() {
    char wordsArray[5][40] = {"a","e","i",};
    int wordsSize = sizeof(wordsArray) / (sizeof(wordsArray[0]));
    char **words = (char **) malloc(sizeof(char *) * wordsSize);
    for (int i = 0; i < wordsSize; i++) {
        words[i] = wordsArray[i];
    }

    int queriesArray[3][2] = {{0,2},{0,1},{2,2}};
    int queriesColSize = sizeof(queriesArray[0]);
    int queriesSize = sizeof(queriesArray) / queriesColSize;
    int **queries = (int **) malloc(sizeof(int *) * queriesSize);
    for (int i = 0; i < queriesSize; i++) {
        queries[i] = queriesArray[i];
    }

    int expected[] = {3,2,1}; 

    int resSize;
    int* res = vowelStrings(words, wordsSize, queries, queriesSize, &queriesColSize, &resSize);

    for (int i = 0; i < resSize; i++) {
        assert(expected[i] == res[i]);    
    }

    free(res);
    free(queries);
    free(words);  
}

int main() {
    testCase0();
    testCase1();
    printf("ok\n");
    return 0;
}