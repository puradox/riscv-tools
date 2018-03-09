# Testbench provided by Ronak

and  x0,  x0, x0      -> 0x0        # x0
addi x1,  x0, 1       -> 0x1        # x1
addi x2,  x0, 2       -> 0x2        # x2
addi x3,  x1, 3       -> 0x4        # x3
addi x4,  x1, 4       -> 0x5        # x4
addi x5,  x2, 5       -> 0x7        # x5
addi x6,  x2, 6       -> 0x8        # x6
addi x7,  x3, 7       -> 0xB        # x7
add  x8,  x1, x2      -> 0x3        # x8
sub  x9,  x8, x4      -> 0xfffffffe # x9 = -2
and  x10, x2, x3      -> 0x0        # x10
or   x11, x3, x4      -> 0x5        # x11
beq  x4,  x11, 36     -> 0x0        # branch taken to inst_mem[21]

addi x8,  x1,  0      -> 0x1        # x8
bgeu x9,  x7,  -36    -> 0x0        # branch taken to inst_mem[13]

addi x8,  x1,  1      -> 0x2        # x8
bne  x3,  x4,  20     -> 0xffffffff # branch taken to inst_mem[19]

addi x8,  x1,  2      -> 0x3        # x8
bltu x2,  x5,  -24    -> 0x1        # branch taken to inst_mem[15]

addi x8,  x1,  3      -> 0x4        # x8
blt  x9,  x1,  4      -> 0x1        # branch taken to inst_mem[17]

addi x8,  x1,  4      -> 0x5        # x8
bge  x7,  x11, 20     -> 0x0        # branch taken to inst_mem[23]

or   x13, x7,  x8     -> 0xf        # x13
jal x11, 24           -> 0x64       # jump to inst[30]

xor x15, x5, x7       -> 0xc        # x15
srl x16, x6, x2       -> 0x2        # x16
sra x17, x9, x3       -> 0xffffffff # x17
jalr x13, 0(x11)      -> 0x88       # branch taken to inst_mem[25]

sw x10, 48(x0)        -> 0x30
sw x8,  352(x0)       -> 0x160
lw x12, 48(x0)        -> 0x0
sll x14, x2, x3       -> 0x20       # x14
beq x12, x10, 20      -> 0x0        # branch taken to inst_mem[34]

xori x10, x2,  22     -> 0x14       # x10
ori  x11, x5,  46     -> 0x2f       # x11
andi x12, x6,  111    -> 0x8        # x12
slli x13, x9,  3      -> 0xfffffff0 # x13
srli x14, x6,  3      -> 0x1        # x14
srai x15, x13, 2      -> 0xfffffffc # x15
slt   x16, x17,x10    -> 0x1        # x16
sltu  x16, x9, x10    -> 0x0        # x16
slti  x16, x9, 2      -> 0x1        # x16
sltiu x16, x9, 2      -> 0x0        # x16
lui   x16, 0xccccc000 -> 0xccccc000 # x16
auipc x16, 0xccccc000 -> 0xccccc0b4 # x16
sw  x9, 20(x0)        -> 0x14
lw  x2, 20(x0)        -> 0xfffffffe # x2
lb  x3, 20(x0)        -> 0xfffffffe # x3
lh  x4, 20(x0)        -> 0xfffffffe # x4
lbu x5, 20(x0)        -> 0x000000fe # x5
lhu x6, 20(x0)        -> 0x0000fffe # x6
slli x13, x11, 4      -> 0x2f0      # x13
sb   x13, 40(x0)      -> 0x28
lw   x14, 40(x0)      -> 0xfffffff0 # x14
sh x13, 40(x0)        -> 0x28
lw x13, 40(x0)        -> 0x000002f0 # x13

