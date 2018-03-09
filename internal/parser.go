package internal

import (
	"bufio"
	"io"
	"log"
)

// Parser takes token from the lexer and creates instructions.
type Parser struct {
	lex     *lexer // lexer to parser the input
	last    Item   // last item read
	hasLast bool   // has a last read item
}

// NewParser creates a new parser.
func NewParser(reader io.Reader) *Parser {
	return &Parser{
		lex: &lexer{
			reader: bufio.NewReader(reader),
		},
	}
}

func (p *Parser) next() Item {
	if p.hasLast {
		p.hasLast = false
	} else {
		p.last = p.lex.nextItem()

		// Skip whitespace and comments
		for p.last.Type == itemSpace || p.last.Type == itemComment {
			p.last = p.next()
		}
	}

	return p.last
}

func (p *Parser) backup() {
	if p.hasLast {
		log.Fatal("there is already something buffered")
	}
	p.hasLast = true
}

// Parse parses a RISC-V assembly file into instructions for pragmatic use.
func (p *Parser) Parse() []Instruction {
	var result []Instruction

	for currItem := p.next(); currItem.Type != itemEOF; currItem = p.next() {
		if currItem.Type == itemSpace || currItem.Type == itemComment {
			continue
		}
		if currItem.Type <= itemKeyword {
			log.Fatalf("instruction expected, got (%d) %s", currItem.Type, currItem.Value)
		}
		// fmt.Printf("(%d) %s\n", currItem.Type, currItem)
		result = append(result, *p.parseInstruction(currItem))
	}

	return result
}

func (p *Parser) parseInstruction(item Item) *Instruction {
	getTypeFn, ok := newTypeFns[item.Type]
	if !ok {
		log.Fatalf("instruction type not found for %s", item.Value)
	}

	instr, err := getTypeFn(item)
	if err != nil {
		log.Panic(err)
	}

	validOps := instr.ValidOperands
	currItem := p.next()
	operandCount := 0

	// Get operands of instruction
	for currItem.Type != itemEOF && operandCount < len(validOps) {
		if currItem.Type == itemSpace {
			currItem = p.next()
			continue
		}

		if currItem.Type != validOps[operandCount] {
			log.Fatalf("invalid instruction operands (%d) %s", currItem.Type, currItem.Value)
		}

		// fmt.Printf("   (%d) %s\n", currItem.Type, currItem)
		instr.Operands = append(instr.Operands, currItem)

		currItem = p.next()
		operandCount++
	}

	// Look for expected value
	if currItem.Type == itemExpect {
		instr.hasExpect = true
		instr.expect = p.next()
		// fmt.Println("   ->", instr.expect.Value)
	} else {
		// fmt.Println(currItem.Type)
		p.backup()
	}

	return instr
}
