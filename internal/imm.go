package internal

import (
	"errors"
	"fmt"
	"strconv"
)

// NewTypeI makes a new I-type instruction
func NewTypeI(item Item) (*Instruction, error) {
	var func3 string

	switch item.Type {
	case itemADDI:
		func3 = "000"
	case itemSLLI:
		func3 = "001"
	case itemSLTI:
		func3 = "010"
	case itemSLTIU:
		func3 = "011"
	case itemXORI:
		func3 = "100"
	case itemSRLI:
		func3 = "101"
	case itemSRAI:
		func3 = "101"
	case itemORI:
		func3 = "110"
	case itemANDI:
		func3 = "111"
	default:
		return nil, errors.New("not an I-type instruction")
	}

	return &Instruction{
		Item: item,
		ValidOperands: []itemType{
			itemRegister,
			itemComma,
			itemRegister,
			itemComma,
			itemInteger,
		},
		Binary: func(operands []Item) string {
			var rd, rs1, immValue int
			var imm string
			var err error

			rd, err = strconv.Atoi(operands[0].Value[1:])
			if err != nil {
				panic("unable to parse register destination")
			}

			rs1, err = strconv.Atoi(operands[2].Value[1:])
			if err != nil {
				panic("unable to parse register source 1")
			}

			immValue, err = strconv.Atoi(operands[4].Value)
			if err != nil {
				panic("unable to parse immediate value")
			}

			imm = strconv.FormatUint(uint64(immValue), 2)
			imm = fmt.Sprintf("%012s", imm)
			if immValue < 0 {
				imm = imm[64-12 : 64]
			}

			return fmt.Sprintf("%012s%05b%s%05b%s", imm, rs1, func3, rd, opcodeTypeI)
		},
	}, nil
}
