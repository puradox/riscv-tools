package internal

import (
	"errors"
	"fmt"
	"strconv"
)

// NewTypeLoad makes a new Load instruction
func NewTypeLoad(item Item) (*Instruction, error) {
	var func3 string
	opcode := opcodeLoad

	switch item.Type {
	case itemLB:
		func3 = "000"
	case itemLH:
		func3 = "001"
	case itemLW:
		func3 = "010"
	case itemLBU:
		func3 = "100"
	case itemLHU:
		func3 = "101"
	default:
		return nil, errors.New("not an Load instruction")
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
				panic("unable to parse register source 2")
			}
			imm = strconv.FormatInt(int64(immValue), 2)

			return fmt.Sprintf("%012s%05b%s%05b%s", imm, rs1, func3, rd, opcode)
		},
	}, nil
}
