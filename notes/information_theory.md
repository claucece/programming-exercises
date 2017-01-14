#Information theory

## Entropy and uncertainty

Information thery defines the amount of information in a message as the
minimum number of bits needed to encode all possible meanings of that message,
assuming they are all equally likely.

The amount of information in a message M is measured by the entropy of a
message, denoted by H(m). Formula: `log_2(n)`, n is the number of possible
meanings.

The entropy of a message also measures its uncertainty: the number of plaintext
bits needed to be recovered when the messages is scrambled in ciphertext in
order to learn the plaintext.

## Rate of the language

Formula: `r = H(M)/N`

The absolute rate of a language is the max number of bits that can be coded in
each character.

Formula : `R - log+2 L`

Redundancy of language: D = R - r

## Security of a cryptosystem

Shannon.

No achievable perfect secrecy (a cryptosystem in which the ct yields no possible
information about the pt, expect for possible its length). Ex: one time pad.

The more redundant a language is the easier to crytanalyze. Reduce the size of
text.

The entropy of a cryptosystem is a measure of the size of the keyspace `K`.

Formula: `H(K) = log_2(K)`

## Unicity distance

The number of different keys that will decipher a ciphertext to some
intelligible pt in the same language as the original pt is given by:

`2^H(K)-nD-1`

Unicity distance (U): defined by Shanon. The approximation of the amount of ct
such that the sum of the real information (entropy) in the corresponding pt plus
the entropy of the key equals the number fo ciphertext bits to be used.

`U = H(K)/D`

A system that has infinte unicity distance has *ideal secrecy*.

## Confusion and Diffusion

For obscuring the redundancy in a pt m.

1. Confusion: obscures the relationship between pt and ct. By substitution.
2. Diffusion: dissipates the redundancy of pt by spreading it out over the ct.
By transposition (permutation).
