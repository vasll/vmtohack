/* eq.asm - implementation of the "eq" vm command in Hack
Pseudocode:
if(stack[sp]-stack[sp-1]==0) {
    stack[sp-1] = -1 // Equal is true
} else {
    stack[sp-1] = 0  // Equal is false
}
*/

@SP
M=M-1   
A=M
D=M      // D = stack[i]
@SP
M=M-1
A=M
M=D-M
D=M      // D = stack[i] - stack[i-1]

@EQ_0_T  // if(D==0) { GOTO EQ_0_T, stack[i]=-1, GOTO EQ_0_END } ELSE { stack[i]=0, GOTO EQ_0_END }
D;JEQ

@SP      // Set stack[i]=-1 and go to EQ_0_END
A=M
M=0
@EQ_0_END
0;JMP

(EQ_0_T)    // Set stack[i]=-1 and go to EQ_0_END
    @SP
    A=M
    M=-1
    @EQ_0_END
    0;JMP

(EQ_0_END)  // End of eq comparison
