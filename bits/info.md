## Buses

Running throughout the system is a collection of electrical conduits called buses
that carry bytes of information back and forth between the components. Buses
are typically designed to transfer fixed-sized chunks of bytes known as words. The
number of bytes in a word (the word size) is a fundamental system parameter that
varies across systems. Most machines today have word sizes of either 4 bytes (32
bits) or 8 bytes (64 bits). For the sake of our discussion here, we will assume a word
size of 4 bytes, and we will assume that buses transfer only one word at a time.

## Main Memory

The main memory is a temporary storage device that holds both a program and
the data it manipulates while the processor is executing the program. Physically,
main memory consists of a collection of dynamic random access memory(DRAM)
chips. Logically, memory is organized as a linear array of bytes, each with its own
unique address (array index) starting at zero. In general, each of the machine
instructions that constitute a program can consist of a variable number of bytes.
The sizes of data items that correspond to C program variables vary according to
type. For example, on an IA32 machine running Linux, data of type short requires
two bytes, types int, float, and long four bytes, and type double eight bytes.
chips work, and how they are combined to form main memory.

## Bits

Modern computers store and process information represented as 2-valued signals.
These lowly binary digits, or bits, form the basis of the digital revolution. The
familiar decimal, or base-10, representation has been in use for over 1000 years,
having been developed in India, improved by Arab mathematicians in the 12th
century, and brought to the West in the 13th century by the Italian mathematician
Leonardo Pisano (c. 1170 – c. 1250), better known as Fibonacci. Using decimal
notation is natural for ten-fingered humans, but binary values work better when
building machines that store and process information. Two-valued signals can
readily be represented, stored, and transmitted, for example, as the presence or
absence of a hole in a punched card, as a high or low voltage on a wire, or as a
magnetic domain oriented clockwise or counterclockwise. The electronic circuitry
for storing and performing computations on 2-valued signals is very simple and
reliable, enabling manufacturers to integrate millions, or even billions, of such
circuits on a single silicon chip.

Unsigned encodings are based on traditional binary notation, representing numbers greater
than or equal to 0. Two’s-complement encodings are the most common way to
represent signed integers, that is, numbers that may be either positive or negative.
Floating-point encodings are a base-two version of scientific notation for
representing real numbers.
