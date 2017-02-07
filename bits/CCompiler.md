# How to Use the Plan 9 C Compiler

The language accepted by the compilers is the core ANSI C language with some
modest extensions, a greatly simplified preprocessor, a smaller library that
includes system calls and related facilities, and a completely different
structure for include files.

Official ANSI C accepts the old (K&R) style of declarations for functions; the
Plan 9 compilers are more demanding. Without an explicit run-time flag (`-B`)
whose use is discouraged, the compilers insist on new-style function
declarations, that is, prototypes for function arguments. The function
declarations in the libraries’ include files are all in the new style so the
interfaces are checked at compile time. For C programmers who have not yet
switched to function prototypes the clumsy syntax may seem repellent but the
payoff in stronger typing is substantial. Those who wish to import existing
software to Plan 9 are urged to use the opportunity to update their code.

The compilers include an integrated preprocessor that accepts the familiar
`#include`, `#define` for macros both with and without arguments, `#undef`,
`#line`, `#ifdef`, `#ifndef`, and `#endif`. It supports neither `#if` nor `##`,
although it does honor a few `#pragmas`. The `#if` directive was omitted because
it greatly complicates the preprocessor, is never necessary, and is usually
abused. Conditional compilation in general makes code hard to understand; the
Plan 9 source uses it sparingly. Also, because the compilers remove dead code,
regular if statements with constant conditions are more readable equivalents to
many `#ifs`. To compile imported code ineluctably fouled by `#if` there is a
separate command, `/bin/cpp`, that implements the complete ANSI C preprocessor
specification.

Include files fall into two groups: machine-dependent and machine-independent.
The machine-independent files occupy the directory `/sys/include`; the others
are placed in a directory appropriate to the machine, such as `/mips/include`.
The compiler searches for include files first in the machine-dependent directory
and then in the machine-independent directory. At the time of writing there are
thirty-one machine-independent include files and two (per machine)
machine-dependent ones: `<ureg.h>` and `<u.h>`. The first describes the layout
of registers on the system stack, for use by the debugger. The second defines
some architecture-dependent types such as `jmp_buf` for `setjmp` and the `va_arg`
and `va_list` macros for handling arguments to variadic functions, as well as a
set of typedef abbreviations for unsigned short and so on.
