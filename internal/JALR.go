package internal

import (
	"errors"
	"fmt"
	"strconv"
)

// NewTypeJALR makes a new JALR instruction
func NewTypeJALR(item Item) (*Instruction, error) {
	if item.Type != itemJALR {
		return nil, errors.New("not a JALR instruction")
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
			var rd, rs1, immValue int
			var imm string
			var err error

			rd, err = strconv.Atoi(operands[0].Value[1:])
			if err != nil {
				panic("unable to parse register destination")
			}

			rs1, err = strconv.Atoi(operands[4].Value[1:])
			if err != nil {
				panic("unable to parse register source 1")
			}

			immValue, err = strconv.Atoi(operands[2].Value)
			if err != nil {
				panic("unable to parse immediate value")
			}
			imm = strconv.FormatInt(int64(immValue), 2)
			imm = fmt.Sprintf("%012s", imm)

			return fmt.Sprintf("%012s%05b%s%05b%s", imm, rs1, "000", rd, opcodeJALR)
		},
	}, nil
}
