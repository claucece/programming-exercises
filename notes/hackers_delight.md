# Computer math

## Binary and Hexadecimal Numbers

Information is represented electronically in a computer using voltages of different
levels. For instance, a 0 could be represented by a 0 voltage and a 1 could be
represented by a voltage of 3 volts. In principle, one could design a computer
to use a large set of voltage levels; in practice, only two are used. This is
because the closer the voltage levels are to each other, the greater the chance
of distorting data due to voltage drift and noise.

In an arbitrary base "n" (where n is a positive integer), the digits may take
values from 0 to "n-1"; hence in base 2 (the binary number system) the digits
may take the values 0 or 1. For this reason, information in a computer is
represented in binary, and each piece of information is called a bit (binary digit).

### Converting Decimal to Binary

It is a matter of determining which places have 1's (since the remaining places
will have 0's), and this can be accomplished by subtracting consecutive powers
of 2, beginning at the left (the most significant digit). The following is an
algorithm (a sequential set of operations designed to solve a problem) to
convert a decimal number into binary:

1. Identify the largest power of 2 which is less than or equal to the number.
2. Write down all the binary place values from that power of 2 to 1.
3. Put a 1 in the place corresponding to that power of 2.
4. Subtract that power of 2 from the decimal number.
5. Identify the largest power of 2 which is less than or equal to the new number.
6. Repeat the last three steps until the new number is 0.
7. Put 0's in all of the remaining places.

Check this: http://kias.dyndns.org/comath/11.html

Computer arithmetic is reduced mod 2^n, being n the word size of the machine.
Usually the word size is 32 bits.

## Sum
In an arbitrary base "n", if the sum of the digits in any place value is greater
than n - 1, the 1's portion of the sum is written down and the "n's" portion is
carried to the place value on the left. In the binary number system, therefore,
if the sum of the digits in any place value is greater than 1, the 1's portion
of the sum is written down and the 2's portion is carried to the place value on
the left.
