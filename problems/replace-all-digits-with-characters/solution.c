#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char * replaceDigits(char * s);

int main() {
	char * s = "a1c1e1";
    printf("%s\n", replaceDigits(s));

	char * s2 = "a1b2c3d4e";
	printf("%s\n", replaceDigits(s2));

	char * s3 = "a";
    printf("%s\n", replaceDigits(s3));

	char * s4 = "";
    printf("%s\n", replaceDigits(s4));

    return 0;
}

char * replaceDigits(char * s){
	int len = strlen(s);
	char * rs = calloc(len+1, sizeof(char));
	rs[len] = '\0';

	for (int i = 1; i < len; i += 2) {
		rs[i-1] = s[i-1];
		rs[i] = s[i-1]+s[i]-'0';
	}
	if (len & 1) {
		// len of s is odd number => just copy last char to rs
		rs[len-1] = s[len-1];
	}
	return rs;
}