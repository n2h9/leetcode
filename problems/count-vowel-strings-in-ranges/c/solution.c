#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>
#include <string.h>

const char vowel[] = {'a', 'e', 'i', 'o', 'u'};

bool isVowel(char *str) {
    // assume str is not null and contains at least one char (do not check it on null)
    // last char index
    int li = strlen(str) - 1;
    bool first = false, last = false;
    for (int i = sizeof(vowel) - 1; i >= 0; i--) {
        if (str[0] == vowel[i]) {
            first = true;
        }
        if (str[li] == vowel[i]) {
            last = true;
        }
    }
    return first && last;
}

/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* vowelStrings(char ** words, int wordsSize, int** queries, int queriesSize, int* queriesColSize, int* returnSize){
    // create an array which [i] element contains 
    // number of words in words array starting with vowel letter
    // from index 0 to index i (including)
    int *count = (int *) malloc(sizeof(int) * wordsSize);
    
    count[0] = isVowel(words[0]) ? 1 : 0;
    for (int i = 1; i < wordsSize; i++) {
        count[i] = count[i-1];
        if (isVowel(words[i])) {
            count[i]++;
        }
    }

    // the result array has same size as a query array
    *returnSize = queriesSize;
    int *returnArray = (int *) malloc(sizeof(int) * *returnSize);
    for (int i = 0; i < queriesSize; i++) {
        if (queries[i][0] > 0) {
            returnArray[i] = count[queries[i][1]] - count[queries[i][0] - 1];
        } else {
            returnArray[i] = count[queries[i][1]];
        }
    }
    
    return returnArray;
}