# Simple RISC-V assembly for starting off with

and  x0, x0, x0         # Register addresses begin with x
addi x1, x0, 1          # Immediates are integers
addi x2, x0, 1  -> 0x1  # Expected values are hex
jalr x3, 4(x0)  -> 0x14 # Register offset
