/* pop/temp.asm - implementation of the "pop temp i" vm command in Hack

Pseudocode:
R13 = 5+i       // get address of 5+i
D = pop()       // get last value from stack
*R13 =  D  
*/

@5
D=A
@%d       // i
D=D+A     
@R13
M=D       

@SP
AM=M-1    // Decrement SP and move A to the address pointed by SP
D=M       // D = *SP (last value on the stack)

@R13
A=M       
M=D       

