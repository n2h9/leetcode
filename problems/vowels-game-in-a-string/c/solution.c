#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

#include "uthash.h"

// array of vowels
char vowels[] = {'a', 'e', 'i', 'o', 'u'};
#define NUM_OF_VOWELS sizeof(vowels) / sizeof(vowels[0])

// hash map item to have vowels
struct hashmap_item {
  char id;           /* we'll use this field as the key */
  UT_hash_handle hh; /* makes this structure hashable */
};

// hasmap to check if char is a vowel
struct hashmap_item* hashmap = NULL;

void initHashMap() {
  for (size_t i = 0; i < NUM_OF_VOWELS; i++) {
    struct hashmap_item* item =
        (struct hashmap_item*)malloc(sizeof(struct hashmap_item));
    item->id = vowels[i];
    HASH_ADD(hh, hashmap, id, sizeof(char), item);
  }
}

void freeHashMap() {
  struct hashmap_item *item, *tmp;
  // Free all entries in the hashmap
  HASH_ITER(hh, hashmap, item, tmp) {
    HASH_DEL(hashmap, item);
    free(item);
  }
}

bool isVowel(char c) {
  struct hashmap_item* item;
  HASH_FIND(hh, hashmap, &c, sizeof(char), item);
  return item != NULL;
}

bool doesAliceWin(char* s) {
  initHashMap();
  int countVowels = 0;
  for (char* ch = s; *ch != '\0'; ch++) {
    if (isVowel(*ch)) {
      countVowels++;
    }
  }

  freeHashMap();
  return countVowels > 0;
}