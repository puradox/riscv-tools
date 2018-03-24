package internal

import (
	"errors"
	"fmt"
	"strconv"
)

// NewTypeR makes a new R-type instruction
func NewTypeR(item Item) (*Instruction, error) {
	var func3 string
	opcode := opcodeTypeR
	func7 := "0000000"

	switch item.Type {
	case itemADD:
		func3 = "000"
	case itemSUB:
		func3 = "000"
		func7 = "0100000"
	case itemSLL:
		func3 = "001"
	case itemSLT:
		func3 = "010"
	case itemSLTU:
		func3 = "011"
	case itemXOR:
		func3 = "100"
	case itemSRL:
		func3 = "101"
	case itemSRA:
		func3 = "101"
		func7 = "0100000"
	case itemOR:
		func3 = "110"
	case itemAND:
		func3 = "111"
	default:
		return nil, errors.New("not a R-type instruction")
	}

	return &Instruction{
		Item: item,
		ValidOperands: []itemType{
			itemRegister,
			itemComma,
			itemRegister,
			itemComma,
			itemRegister,
		},
		Binary: func(operands []Item) string {
			var rd, rs1, rs2 int
			var err error

			rd, err = strconv.Atoi(operands[0].Value[1:])
			if err != nil {
				panic("unable to parse register destination")
			}

			rs1, err = strconv.Atoi(operands[2].Value[1:])
			if err != nil {
				panic("unable to parse register source 1")
			}

			rs2, err = strconv.Atoi(operands[4].Value[1:])
			if err != nil {
				panic("unable to parse register source 2")
			}

			return fmt.Sprintf("%s%05b%05b%s%05b%s", func7, rs2, rs1, func3, rd, opcode)
		},
	}, nil
}
