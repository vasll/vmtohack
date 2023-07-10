/* pop/static.asm - implementation of the "pop static i" vm command in Hack

Pseudocode:
D = stack.pop()
@<FILENAME>.<INDEX>
M = D
*/

@SP
M=M-1
A=M
D=M
@%s.%d      // @<FILENAME>.<INDEX>
M=D
