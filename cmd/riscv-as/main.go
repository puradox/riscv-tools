package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/puradox/riscv-tools/internal"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "View the machine code representation of RISC-V assembly"
	app.UsageText = "riscv-as [arguments...] input"
	app.Version = "0.1.0"
	app.Action = assemble
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Sam Balana",
			Email: "sbalana@uci.edu",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "input, i",
			Value: "",
			Usage: "filename for the assembly input",
		},
		cli.StringFlag{
			Name:  "out, o",
			Value: "",
			Usage: "filename for the output file",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func assemble(c *cli.Context) error {
	filename := c.String("input")

	if c.NArg() > 0 {
		filename = c.Args().Get(0)
	}

	if c.NArg() < 1 && filename == "" {
		cli.ShowAppHelpAndExit(c, 1)
	}

	if filename == "" {
		log.Fatalln("expected an assembly filename")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	p := internal.NewParser(file)
	instrs := p.Parse()

	fmt.Println("Hex      - Instruction  -> Expected value")

	for i := 0; i < len(instrs); i++ {
		if instrs[i].Binary == nil {
			log.Panicf("cannot convert '%s' to binary", instrs[i].Item.Value)
		}

		binary := instrs[i].Binary(instrs[i].Operands)

		binaryValue, err := strconv.ParseUint(binary, 2, 32)
		if err != nil {
			log.Panicf("cannot convert '%s' to binary", instrs[i].Item.Value)
		}

		fmt.Printf("%08x - %s\n", binaryValue, instrs[i].Comment())
	}

	return nil
}
