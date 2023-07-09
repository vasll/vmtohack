/* push/local_argument_this_that.asm - implementation of the "push local/argument/this/that i" vm command in Hack
Explanation:
Get item at address @SEGMENT+i
Push item onto the stack
@SP++
*/

// D = *@SEGMENT+i
@SEGMENT
D=M
@i
A=D+A
D=M
// *@SP = M 
@SP
A=M
M=D
// @SP++
@SP
M=M+1



