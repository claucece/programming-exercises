# Introduction

Data processed by computer:
1. Simple: one value in memory: int, real, char, bool, enum, subrange
2. Structured: : with a name they refer to various values in memory.

## Arrays

A finite, homogenous and ordered colection of elements.

* Finite: has a limit (maximal number)
* Homogeneous: all of its elements are from the same type
* Ordered: the first and subsequent elements can be determined

It has two parts:

* Index: number of components and how to acess them. They refer to components.
* Component: the elements

To refer to one component, you need:

* name of array
* index of the element

### Definition

indentifier = ARRAY[min_limit..max_limit] OF type

NTC: max_limit - min_limit + 1

### Operations

1. Reading/Writing
2. Assigning
3. Updating: inserting, deleting and modifying
4. Ordering
5. Searching


## Bidimensional Arrays

It is a finite, homogenous and ordered group where every element is refered by 
two indexes.
The first one is for row and the second one for column.

### Definitions

array_identifier = ARRAY[min_limita..max_limita, min_limitb..mac_limitb] OF type

min_limita..max_limita: row
min_limitb..mac_limitb: columns

TNC = (min_limita - max_limita + 1) * (max_limita, min_limitb - mac_limitb + 1)

### Order
It can be ordered by row or colum: 

By row:
A[1, 1]; A[1,2]

By column:
A[1, 1]; A[2,1]

### Equation for finding an element

m : row
n : columns
POSINI: position for the array 

By rows:
LOC(A[i, j]) = POSINI + n * (i - 1) + (j - 1)

By columns: 

LOC(A[i, j]) = POSINI + m * (j - 1) + (i - 1)

## Multidimensional Arrays

Length of dimension: Li
Efective index: IEi
Index: Ki

### Equation for finding an element

By rows:

LOC(A[Ki]) = POSINI + ((((IE1 * L2 + IE2) * L3 + EI2) * L4 + ... + IEN-1) * LN + IEN)

Example: 

Array of 2 x 3 x 2

Find A[2, 3, 1]

L1 := 2 - 1 + 1 = 2
L2 := 3 - 1 + 1 = 3
L3 := 2 - 1 + 1 = 2

IE1 := K1 - Liminf1 = 2 - 1 = 1
IE2 := K2 - Liminf2 = 3 - 1 = 2
IE3 := K3 - Liminf3 = 1 - 1 = 0

By colums:

LOC(A[Ki]) = POSINI + (((((IEN * LN-1 + IEN-1) * LN-2) + ... + IE3) * L2 + IE2) * L1 + IE1)

Example: 

Array of 2 x 3 x 2

Find A[2, 3, 1]

L1 := 2 - 1 + 1 = 2
L2 := 3 - 1 + 1 = 3
L3 := 2 - 1 + 1 = 2

IE1 := K1 - Liminf1 = 2 - 1 = 1
IE2 := K2 - Liminf2 = 3 - 1 = 2
IE3 := K3 - Liminf3 = 1 - 1 = 0
