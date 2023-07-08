/* pop/local_argument_this_that.asm - implementation of the "pop local/argument/this/that i" vm command in Hack

Pseudocode:
R13 = SEG+i     // get address of SEG+i where SEG is either local, argument, this or that
D = pop()       // get last value from stack
*R13 =  D  
*/

@LCL
D=M
@%d
D=D+A
@R13
M=D   

@SP
AM=M-1
D=M

@R13
A=M
M=D

