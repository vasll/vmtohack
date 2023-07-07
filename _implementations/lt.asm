/* lt.asm - implementation of the "lt" vm command in Hack
Pseudocode:
if(stack[sp]-stack[sp-1]>0) {
    stack[sp-1] = -1 // lt is true
} else {
    stack[sp-1] = 0  // lt is false
}
sp++
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

@LT_0_T  // if(D==0) { GOTO LT_0_T, stack[i]=-1, GOTO LT_0_END } ELSE { stack[i]=0, GOTO LT_0_END }
D;JGT

@SP      // Set stack[i]=-1 and go to LT_0_END
A=M
M=0
@LT_0_END
0;JMP

(LT_0_T)    // Set stack[i]=-1 and go to LT_0_END
    @SP
    A=M
    M=-1
    @LT_0_END
    0;JMP

(LT_0_END)  // End of lt comparison
@SP
M=M+1
