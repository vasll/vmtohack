/* pop/pointer.asm - implementation of the "pop pointer i" vm command in Hack

Accessing pointer 0 should result in accessing THIS (@3)
Accessing pointer 1 should result in accessing THAT (@4)

Explanation:
SP--
Take the value from the stack
Put the value into THIS/THAT
*/

@SP
M=M-1
A=M
D=M
@3      // 3 or 4 depending on THIS/THAT
M=D
