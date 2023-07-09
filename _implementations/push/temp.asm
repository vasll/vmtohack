/* push/local.asm - implementation of the "push temp i" vm command in Hack
Explanation:
Go to address 5 + i
Get item from said address, add it to *@SP
@SP++
*/

// D = *(5+i)
@5
D=M
@i
A=D+A
D=M
// *@SP = D; @SP++
@SP
A=M
M=D
@SP
M=M+1
