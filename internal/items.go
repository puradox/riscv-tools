package internal

import "fmt"

// Item represents a token or text string returned from the scanner.
type Item struct {
	Type  itemType // The type of this item, such as itemNumber.
	Value string   // The value of this item, such as "0x32".
}

func (i Item) String() string {
	switch {
	case i.Type == itemEOF:
		return "EOF"
	case i.Type == itemError:
		return i.Value
	case i.Type > itemKeyword:
		return fmt.Sprintf("<%s>", i.Value)
	case len(i.Value) > 10:
		return fmt.Sprintf("%.10q...", i.Value)
	}
	return fmt.Sprintf("%q", i.Value)
}

// itemType identifies the type of lex items.
type itemType int

const (
	itemError   itemType = iota // error occured; value is text of error
	itemComma                   // comma symbol separating arguments
	itemComment                 // comment starting with "#" and ending with a newline
	itemEOF
	itemExpect     // "->" token indicating to assert the value
	itemHex        // hexidecimal number
	itemInteger    // positive or negative integer number
	itemLeftParen  // "(" preceeding register offset
	itemNumber     // simple number, including hexidecimal
	itemRegister   // register address, denoted with "x"
	itemRightParen // ")" following register offset
	itemSpace      // run of spaces separating arguments
	// Keywords appear after all the rest.
	itemKeyword // used only to delimit the keywords
	itemADD
	itemADDI
	itemAND
	itemANDI
	itemAUIPC
	itemBEQ
	itemBGE
	itemBGEU
	itemBLT
	itemBLTU
	itemBNE
	itemJAL
	itemJALR
	itemLB
	itemLBU
	itemLH
	itemLHU
	itemLUI
	itemLW
	itemOR
	itemORI
	itemSB
	itemSH
	itemSLL
	itemSLLI
	itemSLT
	itemSLTI
	itemSLTIU
	itemSLTU
	itemSRA
	itemSRAI
	itemSRL
	itemSRLI
	itemSUB
	itemSW
	itemXOR
	itemXORI
)

var itemKey = map[string]itemType{
	"ADD":   itemADD,
	"ADDI":  itemADDI,
	"AND":   itemAND,
	"ANDI":  itemANDI,
	"AUIPC": itemAUIPC,
	"BEQ":   itemBEQ,
	"BGE":   itemBGE,
	"BGEU":  itemBGEU,
	"BLT":   itemBLT,
	"BLTU":  itemBLTU,
	"BNE":   itemBNE,
	"JAL":   itemJAL,
	"JALR":  itemJALR,
	"LB":    itemLB,
	"LBU":   itemLBU,
	"LH":    itemLH,
	"LHU":   itemLHU,
	"LUI":   itemLUI,
	"LW":    itemLW,
	"OR":    itemOR,
	"ORI":   itemORI,
	"SB":    itemSB,
	"SH":    itemSH,
	"SLL":   itemSLL,
	"SLLI":  itemSLLI,
	"SLT":   itemSLT,
	"SLTI":  itemSLTI,
	"SLTIU": itemSLTIU,
	"SLTU":  itemSLTU,
	"SRA":   itemSRA,
	"SRAI":  itemSRAI,
	"SRL":   itemSRL,
	"SRLI":  itemSRLI,
	"SUB":   itemSUB,
	"SW":    itemSW,
	"XOR":   itemXOR,
	"XORI":  itemXORI,
}
