/*

the relationship between the bits of x and x-1. So as in x-1, the rightmost 1 and bits right to it are flipped, then by performing x&(x-1), and storing it in x, will reduce x to a number containing number of ones(in its binary form) less than the previous state of x, thus increasing the value of count in each iteration.

*/
#include <stdio.h>

int count_one (int n) {
	while( n )
	{
		n = n&(n-1);
		count++;
	}
	return count;
}
