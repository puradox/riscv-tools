addi  x0,x0,0    -> 0x0        # r0=0     (r0=0)
addi  x1,x0,8    -> 0x8        # r1=r0+8  (r1=8)
addi  x2,x0,4    -> 0x4        # r2=r0+4  (r2=4)
or    x3,x1,x2   -> 0xc        # r3=r1|r2 (r3=12) // Forwarding Test
or    x4,x2,x0   -> 0x4        # r4=r2|r0 (r4=4)  // Forwarding Test
add   x6,x4,x2   -> 0x8        # r6=r4+r2 (r6=8)  // Hazard test

addi  x7,x0,8    -> 0x8
sw    x7,48(x0)  -> 0x30       # Mem[0x30]=8
lw    x8,48(x0)  -> 0x8        # (r8=8)
addi  x9,x8,0    -> 0x8        # (r9=8)

addi  x4,x0,2    -> 0x2        # r4=r0+2  (r4=2)
addi  x5,x0,-2   -> 0xfffffffe # r5=r0-2  (r5=-2)

sll   x18,x1,x4  -> 0x20       # r18=r1<<r4  (r18 = 8<<2)  r18=0000020
srl   x19,x5,x4  -> 0x3fffffff # r19=r5>>r4  (r19 = -2>>2) r19=3FFFFFF // Forwarding Test
sra   x20,x5,x4  -> 0xffffffff # r20=r5>>>r4 (r20 = -2>>2) r20=FFFFFFF

slt   x21,x1,x2  -> 0x0        # if r1<r2, r21=1 (r21=0)
slt   x22,x2,x1  -> 0x1        # if r2<r1, r22=1 (r22=1)
sltu  x23,x5,x1  -> 0x0        # if r5<r1, r23=1 (r23=0)
sltu  x24,x1,x5  -> 0x1        # if r1<r5, r24=1 (r24=1)
slti  x25,x1,8   -> 0x0        # if r1<8,  r25=1 (r25=0)
slti  x26,x1,16  -> 0x1        # if r1<16, r26=1 (r26=1)

addi  x5,x0,-4   -> 0xfffffffc # r5 =r0-4 (r5=-4)

sltiu x27,x1,-2  -> 0x1        # if r1<u(-2),  r27=1 (r27=1)
sltiu x28,x5,-2  -> 0x1        # if r5<u(-2),  r28=1 (r28=1)

slli  x29,x5,1   -> 0xfffffff8 # r29=r5<<1 (r29=FFFFFFF8)
srli  x30,x5,1   -> 0x7ffffffe # r30=r5>>1 (r30=7FFFFFFE)
srai  x31,x5,1   -> 0xfffffffe # r31=r5>>1 (r31=FFFFFFFE)

xori  x6,x1,10   -> 0x2        # r6=r1 xor A (r6=00000002)
ori   x7,x1,2    -> 0xa        # r7=r1 or 2  (r7=0000000A)
andi  x8,x1,10   -> 0x8        # r8=r1 and A (r8=00000008)
xor   x9,x1,x2   -> 0xc        # r9=r1 or r2 (r9=0000000C)
