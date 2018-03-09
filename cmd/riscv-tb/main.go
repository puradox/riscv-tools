package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/puradox/riscv-tools/internal"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Generate a SystemVerilog testbench from RISC-V assembly"
	app.UsageText = "riscv-tb [arguments...] input template"
	app.Version = "0.1.0"
	app.Action = testbench
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
			Name:  "template, t",
			Value: "",
			Usage: "filename for the template",
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

func testbench(c *cli.Context) error {
	filename := c.String("input")
	templateFilename := c.String("template")
	outFilename := c.String("out")

	if c.NArg() > 1 {
		filename = c.Args().Get(0)
		templateFilename = c.Args().Get(1)
	}

	if c.NArg() < 2 && filename == "" && templateFilename == "" {
		cli.ShowAppHelpAndExit(c, 1)
	}

	if filename == "" {
		log.Fatalln("expected an assembly filename")
	}
	if templateFilename == "" || !strings.HasSuffix(templateFilename, ".tmpl") {
		log.Fatalln("expected a template filename ending in .tmpl")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	p := internal.NewParser(file)
	instrs := p.Parse()

	tmpl, errTmpl := template.ParseFiles(templateFilename)
	if errTmpl != nil {
		log.Fatalf("failed to parse template file '%s'\n", templateFilename)
	}

	testbenchCode := bytes.NewBufferString("")
	for i := 0; i < len(instrs); i++ {
		internal.WritePipelinedComment(testbenchCode, &instrs[i])
	}
	internal.WriteLastTests(testbenchCode)

	var output io.Writer
	if outFilename == "" {
		output = os.Stdout
	} else {
		var errOut error
		output, errOut = os.Create(outFilename)
		if errOut != nil {
			log.Fatalf("failed to create file '%s'\n", outFilename)
		}
	}

	tmpl.Execute(output, map[string]interface{}{
		"testbench": testbenchCode.String(),
	})

	return nil
}
