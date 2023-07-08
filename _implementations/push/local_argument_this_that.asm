/* pop/local_argument_this_that.asm - implementation of the "push local/argument/this/that i" vm command in Hack

Pseudocode:
Go to address LCL + i
Store M into D
Go to @SP
Store D into SP
SP++
*/


@LCL
D=M     // Store the value in LCL in D

@%d      // i
A=D+A   // A=LCL+i
D=M     // Store the value at LCL+i in D

@SP
A=M     // Set A to the stack pointer
M=D     // Store the value from D at the stack pointer address

@SP
M=M+1   // Increment the stack pointer