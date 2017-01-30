/*
Properties for numbers which are powers of 2, is that they have one and only one bit set in their binary representation. If the number is neither zero nor a power of two, it will have 1 in more than one place. So if x is a power of 2 then x & (x-1) will be 0.
*/

#include <stdio.h>

typedef int bool;
#define true 1
#define false 0

bool isPowerOfTwo(int x) {
	return (x && !(x &(x -1)));
}

int main() {
	printf("%d", isPowerOfTwo(2));
}

