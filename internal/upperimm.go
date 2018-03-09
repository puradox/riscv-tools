package internal

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

// NewTypeU makes a new U-type instruction
func NewTypeU(item Item) (*Instruction, error) {
	var opcode string

	switch item.Type {
	case itemLUI:
		opcode = opcodeLUI
	case itemAUIPC:
		opcode = opcodeAUIPC
	default:
		return nil, errors.New("not a U-type instruction")
	}

	return &Instruction{
		Item: item,
		ValidOperands: []itemType{
			itemRegister,
			itemComma,
			itemHex,
		},
		Binary: func(operands []Item) string {
			var rd int
			var immValue uint64
			var imm string
			var err error

			rd, err = strconv.Atoi(operands[0].Value[1:])
			if err != nil {
				panic("unable to parse register destination")
			}

			immValue, err = strconv.ParseUint(operands[2].Value[2:], 16, 32)
			if err != nil {
				log.Panic(err)
			}

			imm = strconv.FormatUint(immValue, 2)
			imm = fmt.Sprintf("%032s", imm)

			return fmt.Sprintf("%020s%05b%s", imm[12:31], rd, opcode)
		},
	}, nil
}
