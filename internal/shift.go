package internal

import (
	"errors"
	"fmt"
	"strconv"
)

// NewTypeShiftI makes a new I-type shift instruction
func NewTypeShiftI(item Item) (*Instruction, error) {
	var func3 string
	opcode := opcodeTypeI
	func7 := "0000000"

	switch item.Type {
	case itemSLLI:
		func3 = "001"
	case itemSRLI:
		func3 = "101"
	case itemSRAI:
		func3 = "101"
		func7 = "0100000"
	default:
		return nil, errors.New("not a R-type instruction")
	}

	return &Instruction{
		Item: item,
		ValidOperands: []itemType{
			itemRegister,
			itemComma,
			itemInteger,
			itemLeftParen,
			itemRegister,
			itemRightParen,
		},
		Binary: func(operands []Item) string {
			var rd, rs1, shamtBase10 int
			var shamt string
			var err error

			rd, err = strconv.Atoi(operands[0].Value[1:])
			if err != nil {
				panic("unable to parse register destination")
			}

			rs1, err = strconv.Atoi(operands[4].Value[1:])
			if err != nil {
				panic("unable to parse register source 1")
			}

			shamtBase10, err = strconv.Atoi(operands[2].Value[1:])
			if err != nil {
				panic("unable to parse register source 2")
			}

			shamt = strconv.FormatInt(int64(shamtBase10), 2)

			return fmt.Sprintf("%s%05s%05b%s%05b%s", func7, shamt, rs1, func3, rd, opcode)
		},
	}, nil
}
