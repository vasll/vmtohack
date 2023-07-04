/* add.asm - implementation of the "add" vm command in Hack
Pseudocode:
stack[sp-1] = stack[sp] + stack[sp-1]
*/

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=D+M