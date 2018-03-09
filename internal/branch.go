package internal

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
)

// NewTypeB makes a new B-type instruction
func NewTypeB(item Item) (*Instruction, error) {
	var func3 string

	switch item.Type {
	case itemBEQ:
		func3 = "000"
	case itemBNE:
		func3 = "001"
	case itemBLT:
		func3 = "100"
	case itemBGE:
		func3 = "101"
	case itemBLTU:
		func3 = "110"
	case itemBGEU:
		func3 = "111"
	default:
		return nil, errors.New("not a B-type instruction")
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
			var rs1, rs2 int
			var imm string
			var err error

			rs1, err = strconv.Atoi(operands[0].Value[1:])
			if err != nil {
				panic("unable to parse register source 1")
			}

			rs2, err = strconv.Atoi(operands[2].Value[1:])
			if err != nil {
				panic("unable to parse register source 2")
			}

			imm, err = parseIntToBinary(operands[4].Value, 13)
			if err != nil {
				log.Fatal(err)
			}

			return fmt.Sprintf("%01c%05s%05b%05b%s%04s%01c%s", pick(imm, 12), pluck(imm, 10, 5), rs2, rs1, func3, pluck(imm, 4, 1), pick(imm, 11), opcodeBranch)
		},
	}, nil
}

func pick(imm string, pos int) byte {
	return imm[len(imm)-pos-1]
}

func pluck(imm string, from, to int) string {
	return imm[len(imm)-from-1 : len(imm)-to]
}

// parseIntToBinary parses a string of an integer into a string of sign-extended binary.
func parseIntToBinary(s string, bitSize int) (string, error) {
	value, err := strconv.ParseInt(s, 10, bitSize)
	if err != nil {
		return "", err
	}

	result := strconv.FormatUint(uint64(value), 2)

	// Clip sign extend
	if value < 0 {
		result = fmt.Sprintf("%s", result[64-bitSize:])
	} else {
		var buf bytes.Buffer
		for i := len(result); i < bitSize; i++ {
			buf.WriteRune('0')
		}
		buf.WriteString(result)
		result = buf.String()
	}

	// fmt.Printf("%d => %s (%c)\n", value, result, result[0])

	return result, nil
}
