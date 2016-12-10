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

## Substraction
Before subtracting binary numbers, you must know how many bits your computer
uses to add and subtract. Subtraction is performed as addition of a complement.
Technique: 2's complement:

1. Compute the "1's complement" of the number being subtracted (if we're
   subtracting a - b, we need the complement of b).The 1's complement is simply
   the inverse of the number: every 1 is replaced by 0 and every 0 is replaced
   by 1.
2. Insert leading zeroes: this is necessary whenever we are subtracting a
   number which has fewer bits than the number of bits in our adder (enough
   leading zeros must added so that the two numbers have exactly that number
   of bits).
3. Obtain the 2's complement by simply adding 1 to the 1's complement. If the
   left most bit is 1; this tells you that the value is negative.
4. Add the 2's complement to the number which we originally wanted to subtract
   from. Notice that you will have a carry left over; when using 2's complement
   subtraction, this is expected: the final carry is "thrown away" or ignored.


## Unsigned and Signed Integers

An integer is a number with no fractional part; it can be positive, negative or
zero. In ordinary usage, one uses a minus sign to designate a negative integer.
However, a computer can only store information in bits, which can only have the
values zero or one. We might expect, therefore, that the storage of negative
integers in a computer might require some special technique.

### Unsigned integer

As you might imagine, an unsigned integer is either positive or zero. An
unsigned integer containing n bits can have a value between 0 and 2n - 1
(which is 2n different values).

Most modern computers store memory in units of 8 bits, called a "byte" (also
called an "octet"). Arithmetic in such computers can be done in bytes, but is
more often done in larger units called "(short) integers" (16 bits), "long
integers" (32 bits) or "double integers" (64 bits). Short integers can be used
to store numbers between 0 and 216 - 1, or 65,535. Long integers can be used to
store numbers between 0 and 232 - 1, or 4,294,967,295. and double integers can
be used to store numbers between 0 and 264 - 1, or 18,446,744,073,709,551,615.

### Signed integer
Signed integers are stored in a computer using 2's complement. As you recall,
when computing the 2's complement of a number it was necessary to know how many
bits were to be used in the final result; leading zeroes were appended to the
most significant digit in order to make the number the appropriate length.
Since the process of computing the 2's complement involves first computing the
1's complement, these leading zeros become leading ones, and the left most bit
of a negative number is therefore always 1. In computers, the left most bit of
a signed integer is called the "sign bit".

* Note that the most negative number which we can store in an 8 bit signed
  integer is -128, which is - 28 - 1, and that the largest positive signed
  integer we can store in an 8 bit signed integer is 127, which is 28 - 1 - 1.
* The number of integers between -128 and + 127 (inclusive) is 256, which is
  28; this is the same number of values which an unsigned 8 bit integer can
  contain (from 0 to 255).
* If a signed integer has n bits, it can contain a number between - 2n - 1 and +
  (2n - 1 - 1).
* Since both signed and unsigned integers of n bits in length can represent 2n
  different values, there is no inherent way to distinguish signed integers
  from unsigned integers simply by looking at them; the software designer is
  responsible for using them correctly.
* No matter what the length, if a signed integer has a binary value of all 1's,
  it is equal to decimal -1.

There is an interesting consequence to the fact that in 2's complement
arithmetic, one expects to throw away the final carry: in unsigned arithmetic a
carry out of the most significant digit means that there has been an overflow,
but in signed arithmetic an overflow is not so easy to detect. In fact, signed
arithmetic overflows are detected by checking the consistency of the signs of
the operands and the final answer.
A signed overflow has occurred in an addition or subtraction if:

* the sum of two positive numbers is negative;
* the sum of two negative numbers is non-negative;
* subtracting a positive number from a negative one yields a positive result;
  or subtracting a negative number from a non-negative one yields a negative
  result.

Integer arithmetic on computers is often called "fixed point" arithmetic and
the integers themselves are often called fixed point numbers. Real numbers on
computers (which may have fractional parts) are often called "floating point"
numbers.
