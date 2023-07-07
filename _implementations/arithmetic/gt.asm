/* gt.asm - implementation of the "gt" vm command in Hack
Pseudocode:
if(stack[sp]-stack[sp-1]<0) {
    stack[sp-1] = -1 // gt is true
} else {
    stack[sp-1] = 0  // gt is false
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

@GT_0_T  // if(D==0) { GOTO GT_0_T, stack[i]=-1, GOTO GT_0_END } ELSE { stack[i]=0, GOTO GT_0_END }
D;JLT

@SP      // Set stack[i]=-1 and go to GT_0_END
A=M
M=0
@GT_0_END
0;JMP

(GT_0_T)    // Set stack[i]=-1 and go to GT_0_END
    @SP
    A=M
    M=-1
    @GT_0_END
    0;JMP

(GT_0_END)  // End of gt comparison
@SP
M=M+1
