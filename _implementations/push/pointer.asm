/* push/pointer.asm - implementation of the "push pointer i" vm command in Hack

Accessing pointer 0 should result in accessing THIS (@3)
Accessing pointer 1 should result in accessing THAT (@4)

Explanation:
Put the value of THIS/THAT into the stack
Stack++
*/

@3      // 3 or 4 depending on THIS/THAT
D=M
@SP
A=M
M=D
@SP
M=M+1