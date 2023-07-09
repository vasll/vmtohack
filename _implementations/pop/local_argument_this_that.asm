/* pop/local_argument_this_that.asm - implementation of the "pop local/argument/this/that i" vm command in Hack
Explanation:
Pop latest item from stack
@SP--
Put the popped item into *@SEGMENT+i
*/

// R13 = @SEGMENT+i
@SEGMENT
D=M
@i
D=D+A
@R13
M=D
// D = stack[SP-1]
@SP
M=M-1
A=M
D=M
// *@R13 = D
@R13
A=M
M=D
