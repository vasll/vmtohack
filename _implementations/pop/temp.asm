/* pop/temp.asm - implementation of the "pop temp i" vm command in Hack

Pseudocode:
@R15 = 5 + i
SP--
*@R15 = *@SP
*/

// @R13 = 5+i
@5
D=M
@i
D=D+A
@R13
M=D
// @SP--; D=*@SP
@SP
M=M-1
A=M
D=M
// *@R13 = M
@R13
A=M
M=D
