package internal

import (
	"bytes"
	"strconv"
)

// Instruction represents a single RISC-V instruction.
type Instruction struct {
	Item      Item   // item of the instruction
	Operands  []Item // operands used with the instruction
	expect    Item   // value that is expected after this instruction runs
	hasExpect bool   // does this instruction have an expected value?

	ValidOperands []itemType          // types of items allowed as operands
	Binary        func([]Item) string // generate a string of hex representing the instruction
}

// Comment generates a comment for this instruction
func (instr *Instruction) Comment() string {
	buf := bytes.NewBufferString("")

	buf.WriteString(instr.Item.Value)
	buf.WriteString(" ")

	for _, op := range instr.Operands {
		buf.WriteString(op.Value)
		if op.Type == itemComma {
			buf.WriteString(" ")
		}
	}

	if instr.hasExpect {
		buf.WriteString(" -> ")
		buf.WriteString(instr.expect.Value)
	}

	return buf.String()
}

// Testbench generates a test for this instruction
func (instr *Instruction) Testbench() string {
	buf := bytes.NewBufferString("")

	if instr.hasExpect {
		buf.WriteString("test(32'h")
		buf.WriteString(instr.expect.Value[2:])
		buf.WriteString(");")
	} else {
		buf.WriteString("skip();")
	}

	return buf.String()
}

const warmupTime = 20
const nsPerCycle = 20

var cycleCount = 1
var lastInstructions []Instruction
var pipelineStages = []string{"FE", "DE", "ME", "EX", "WB"}

func writeCycleInfo(buf *bytes.Buffer) {
	buf.WriteString("    //\n")
	buf.WriteString("    // Cycle ")
	buf.WriteString(strconv.Itoa(cycleCount - 1))
	buf.WriteString("\n")
	buf.WriteString("    // Time ")
	buf.WriteString(strconv.Itoa((cycleCount-1)*nsPerCycle + warmupTime))
	buf.WriteString("ns\n")
	buf.WriteString("    //\n")
}

func writePipelineComment(buf *bytes.Buffer, instrIndex, stageIndex int) {
	buf.WriteString("    ")
	buf.WriteString("// ")
	buf.WriteString(pipelineStages[stageIndex])
	buf.WriteString(" - ")
	buf.WriteString(lastInstructions[instrIndex].Comment())
	buf.WriteString("\n")
}

func writeTest(buf *bytes.Buffer, instrIndex int) {
	buf.WriteString("    ")
	if instrIndex >= 0 {
		buf.WriteString(lastInstructions[instrIndex].Testbench())
	} else {
		buf.WriteString("skip();")
	}
	buf.WriteString("\n\n")
}

// WritePipelinedComment writes comments for pipelined processors.
func WritePipelinedComment(buf *bytes.Buffer, instr *Instruction) {
	instrCount := len(lastInstructions)
	lastInstructions = append(lastInstructions, *instr)
	stageCount := len(pipelineStages)

	writeCycleInfo(buf)

	for i := 0; i < stageCount && i <= instrCount; i++ {
		writePipelineComment(buf, instrCount-i, i)
	}

	writeTest(buf, instrCount-stageCount+1)
	cycleCount++
}

// WriteLastTests writes the last tests
func WriteLastTests(buf *bytes.Buffer) {
	instrCount := len(lastInstructions)
	stageCount := len(pipelineStages)

	for i := 1; i < stageCount; i++ {
		writeCycleInfo(buf)

		for j := i; j < stageCount && j <= instrCount; j++ {
			writePipelineComment(buf, instrCount-j+i-1, j)
		}

		writeTest(buf, instrCount-stageCount+i)
		cycleCount++
	}
}

type newTypeFn func(item Item) (*Instruction, error)

var newTypeFns = map[itemType]newTypeFn{
	itemADD:   NewTypeR,
	itemADDI:  NewTypeI,
	itemAND:   NewTypeR,
	itemANDI:  NewTypeI,
	itemAUIPC: NewTypeU,
	itemBEQ:   NewTypeB,
	itemBGE:   NewTypeB,
	itemBGEU:  NewTypeB,
	itemBLT:   NewTypeB,
	itemBLTU:  NewTypeB,
	itemBNE:   NewTypeB,
	itemJAL:   NewTypeJAL,
	itemJALR:  NewTypeJALR,
	itemLB:    NewTypeLoad,
	itemLBU:   NewTypeLoad,
	itemLH:    NewTypeLoad,
	itemLHU:   NewTypeLoad,
	itemLUI:   NewTypeU,
	itemLW:    NewTypeLoad,
	itemOR:    NewTypeR,
	itemORI:   NewTypeI,
	itemSB:    NewTypeS,
	itemSH:    NewTypeS,
	itemSLL:   NewTypeR,
	itemSLLI:  NewTypeI,
	itemSLT:   NewTypeR,
	itemSLTI:  NewTypeI,
	itemSLTIU: NewTypeI,
	itemSLTU:  NewTypeR,
	itemSRA:   NewTypeR,
	itemSRAI:  NewTypeI,
	itemSRL:   NewTypeR,
	itemSRLI:  NewTypeI,
	itemSUB:   NewTypeR,
	itemSW:    NewTypeS,
	itemXOR:   NewTypeR,
	itemXORI:  NewTypeI,
}
