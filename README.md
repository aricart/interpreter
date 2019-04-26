# interpreter
simple interpreter based on some on a primitive assembler alejandro gave me.

## Install Go

https://golang.org/dl/

## Write a program

This machine has a:
 - limited number of register (0-9)
 - limited number of memory addresses (0-9)


The instruction set for the interpreter

- 1RA - Load Register R with memory value at address A
- 2Rn - Load Register R with a number
- 3RA - Store Register R at memory address
- 4RS - Move Register R to register S
- 5RST - Add and store into register R the values of register S and T
- 6RST - Not Implemented
- 7RST - OR and store into register R the values in register S and T
- 8RST - AND and store into register R the values in register S and T
- 9RST - XOR and store into register R the values in register S and T
- ARn  - ROT value in R on bit to the right n times
- BRA  - JUMP to instruction at address A if register R and register 0 equal
- C    - HALT
- P    - Print the current machine

Enter all the instructions one per line.
After the last instruction enter <ctrl+D> to execute the program

```bash
> go get github.com/aricart/interpreter

> interpreter
2110
2210
5312
P
C
<ctrl + d>

╭────────────╮
│  Machine   │
├───┬────┬───┤
│   │ R  │ M │
├───┼────┼───┤
│ 0 │ 0  │ 0 │
│ 1 │ 10 │ 0 │
│ 2 │ 20 │ 0 │
│ 3 │ 30 │ 0 │
│ 4 │ 0  │ 0 │
│ 5 │ 0  │ 0 │
│ 6 │ 0  │ 0 │
│ 7 │ 0  │ 0 │
│ 8 │ 0  │ 0 │
│ 9 │ 0  │ 0 │
╰───┴────┴───╯


```

