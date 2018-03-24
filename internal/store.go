package internal

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

// NewTypeS makes a new S-type instruction: SB SH SW
func NewTypeS(item Item) (*Instruction, error) {
	var func3 string

	switch item.Type {
	case itemSB:
		func3 = "000"
	case itemSH:
		func3 = "001"
	case itemSW:
		func3 = "010"
	default:
		return nil, errors.New("not a S-type instruction")
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
			var rs1, rs2 int
			var imm string
			var err error

			rs1, err = strconv.Atoi(operands[4].Value[1:])
			if err != nil {
				panic("unable to parse register source 1")
			}

			rs2, err = strconv.Atoi(operands[0].Value[1:])
			if err != nil {
				panic("unable to parse register source 2")
			}

			imm, err = parseIntToBinary(operands[2].Value, 12)
			if err != nil {
				log.Fatal(err)
			}

			return fmt.Sprintf("%07s%05b%05b%s%05s%s", pluck(imm, 11, 5), rs2, rs1, func3, pluck(imm, 4, 0), opcodeStore)
		},
	}, nil
}
