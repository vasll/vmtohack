/* not.asm - implementation of the "not" vm command in Hack
Pseudocode:
stack[sp] = !stack[sp]
sp++
*/

@SP
M=M-1
A=M
D=!M
@SP
M=M+1
