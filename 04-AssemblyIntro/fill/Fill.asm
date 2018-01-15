// This file is part of www.nand2tetris.org
// and the book "The Elements of Computing Systems"
// by Nisan and Schocken, MIT Press.
// File name: projects/04/Fill.asm

// Runs an infinite loop that listens to the keyboard input.
// When a key is pressed (any key), the program blackens the screen,
// i.e. writes "black" in every pixel;
// the screen should remain fully black as long as the key is pressed. 
// When no key is pressed, the program clears the screen, i.e. writes
// "white" in every pixel;
// the screen should remain fully clear as long as no key is pressed.

// Put your code here.

// Initialize variables to 0
//@0
//D=M
//@color
//M=D

// Infinte loop
(LOOP)

// If the keyboard register is 0 (no key pressed), goto white
@KBD
D=M
@WHITE
D;JEQ

// BLACK code
// Set color variable to 0. When 1 is substracted will be -1 (black)
@0
D=A
@color
M=D
// goto shared code loop
@SET
0;JMP

// WHITE code
(WHITE)
// Set color variable to 1. When 1 is subtracted will be 0 (white)
@1
D=A
@color
M=D
// goto shared code loop
@SET
0;JMP

// Shared code to set all pixel blocks
(SET)
// Initialize pointer pxaddr to SCREEN (first pixel block)
@SCREEN
D=A
@pxaddr
M=D
// Initialize final screen address + 1 to 24576
@24576
D=A
@endaddr
M=D
// Top of pixel setting loop
(PXLOOP)
// If pxaddr equals one greater than the last screen pixel block address, jump to top of infinte loop
@pxaddr
D=M
@endaddr
D=D-M
@LOOP
D;JEQ
// Set pixels at pxaddr to value of variable 'color' - 1
@color
D=M
@pxaddr
A=M
M=D-1
// Add 1 to pxaddr
@pxaddr
M=M+1
// Jump to top of pixel setting loop
@PXLOOP
0;JMP
