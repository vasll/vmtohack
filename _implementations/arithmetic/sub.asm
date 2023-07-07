/* sub.asm - implementation of the "sub" vm command in Hack
Pseudocode:
stack[sp-1] = stack[sp] - stack[sp-1]
sp++
*/

@SP
M=M-1
A=M
D=M
@SP
M=M-1
A=M
M=M-D
@SP
M=M+1
