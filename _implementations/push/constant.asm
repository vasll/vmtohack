/* push/constant.asm - implementation of the "push constant i" vm command in Hack
Pseudocode:
stack[sp] = n
stackpointer++
*/

@i
D=A
@SP
A=M
M=D
@SP
M=M+1
