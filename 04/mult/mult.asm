// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Mult.asm

// Multiplies R0 and R1 and stores the result in R2.
// (R0, R1, R2 refer to RAM[0], RAM[1], and RAM[2], respectively.)

// Put your code here.

// Initialize the product to zero
@0
D=A
@R2
M=D

// Initialize the counter i to zero
@i
M=D

(LOOP)
// If i equals R0, jump to END
@i
D=M
@R0
D=D-M
@END
D;JEQ

// Add value in R1 to value in R2
@R1
D=M
@R2
M=D+M

// Add 1 to the counter i
@i
M=M+1

// Jump to LOOP
@LOOP
0;JMP

// End the loop
(END)
@END
0;JMP
