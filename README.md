<big><h1 align="center">riscv-tools</h1></big>

<p align="center">
  <a href="https://github.com/puradox/riscv-tools/issues">
    <img src="https://img.shields.io/github/issues/puradox/riscv-tools.svg" alt="Github Issues">
  </a>
  <a href="https://goreportcard.com/report/github.com/puradox/riscv-tools">
    <img src="https://goreportcard.com/badge/github.com/puradox/riscv-tools" alt="Go Report Card">
  </a>
  <a href="https://godoc.org/github.com/puradox/riscv-tools">
    <img src="https://godoc.org/github.com/puradox/riscv-tools?status.svg" alt="GoDoc">
  </a>
</p>

<p align="center"><big>
  Tools for translating RISC-V assembly into machine code, testbenches, and more!
</big></p>


## Features
  - Parse RISC-V assembly files
  - Output machine code in binary and hexidecimal
  - Generate SystemVerilog workbenches using "expect" syntax
    - Support for pipelined processors
  - Generate SystemVerilog instruction memories preloaded with machine code

## Install

It is strongly recommended that you use a released version. Release binaries are available on the [releases](https://github.com/puradox/riscv-tools/releases) page.

If you're interested in hacking on `riscv-tools`, you can install via `go get`:
```bash
go get -u github.com/puradox/riscv-tools/...
```

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can be easily used:
```bash
export PATH=$PATH:$GOPATH/bin
```

## Usage

```bash
$ riscv-as examples/simple.asm
Hex      - Instruction  -> Expected value
00007033 - and x0, x0, x0
00100093 - addi x1, x0, 1
00100113 - addi x2, x0, 1 -> 0x1
004001e7 - jalr x3, 4(x0) -> 0x4

$ cd examples
$ riscv-tb -o tb.sv simple.asm tb.sv.tmpl
```

See `examples/tb.sv` for the generated testbench.

## License

- **MIT** : http://opensource.org/licenses/MIT

## Contributing

Contributions are highly welcome and even encouraged! This tool was made to make RISC-V development easier. Filing bug reports and feature requests will help further improve the tool for all.

## Future work
 - [ ] Add program counter
 - [ ] Support labels
 - [ ] Add macros for customizing testbench translation
 - [ ] Expose API for developers wishing to extend RISC-V instructions themselves
 - [ ] Simulate RISC-V for _automagicially_ generating expected values for testbenches
