/* push/static.asm - implementation of the "push static i" vm command in Hack

Pseudocode:
D = @<FILENAME>.<INDEX>
*@SP = D
@SP++ 
*/

@%s.%d      // @<FILENAME>.<INDEX>
D=M
@SP
A=M
M=D
@SP
M=M+1
