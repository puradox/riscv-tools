package internal

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

// NewTypeJAL makes a new JAL instruction
func NewTypeJAL(item Item) (*Instruction, error) {
	if item.Type != itemJAL {
		return nil, errors.New("not a JAL instruction")
	}

	return &Instruction{
		Item: item,
		ValidOperands: []itemType{
			itemRegister,
			itemComma,
			itemInteger,
		},
		Binary: func(operands []Item) string {
			var imm string

			rd, err := strconv.Atoi(operands[0].Value[1:])
			if err != nil {
				panic("unable to parse register destination")
			}

			imm, err = parseIntToBinary(operands[2].Value, 21)
			if err != nil {
				log.Fatal(err)
			}

			return fmt.Sprintf("%01c%010s%01c%08s%05b%s", pick(imm, 20), pluck(imm, 10, 1), pick(imm, 11), pluck(imm, 19, 12), rd, opcodeJAL)
		},
	}, nil
}
