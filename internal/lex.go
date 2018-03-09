package internal

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"strings"
)

type stateFn func(*lexer) stateFn

type lexer struct {
	reader *bufio.Reader // reader for the file
	buf    bytes.Buffer  // buffer holding the current item value
	width  int           // size of the last rune read
}

const eof = rune(0)

// next reads the next rune from the buffered reader.
func (l *lexer) next() rune {
	r, w, err := l.reader.ReadRune()
	if err != nil {
		return eof
	}
	l.buf.WriteRune(r)
	l.width = w
	return r
}

// backup places the previously read rune back on the reader.
func (l *lexer) backup() {
	err := l.reader.UnreadRune()
	if err != nil {
		log.Fatal(err)
		panic("unable to unread rune")
	}
	l.buf.Truncate(l.buf.Len() - l.width)
	l.width = 0
}

// peek returns but does not consume the next run in the input.
func (l *lexer) peek() rune {
	r, _, err := l.reader.ReadRune()
	if err != nil {
		return eof
	}

	err = l.reader.UnreadRune()
	if err != nil {
		panic("unable to unread rune")
	}

	return r
}

func (l *lexer) acceptAll(validFn func(rune) bool) {
	for {
		if r := l.next(); r == eof {
			break
		} else if !validFn(r) {
			l.backup()
			break
		}
	}
}

// emit passes an item back to the client
func (l *lexer) emit(t itemType) Item {
	i := Item{t, l.buf.String()}
	l.buf.Reset()
	return i
}

// emit passes an item back to the client
func (l *lexer) emitError(err string) Item {
	i := Item{itemError, fmt.Sprintf("%s: %s", err, l.buf.String())}
	l.buf.Reset()
	return i
}

// nextItem returns the next token and literal value.
func (l *lexer) nextItem() Item {
	switch r := l.next(); {
	case r == eof:
		return l.emit(itemEOF)
	case isWhitespace(r):
		return l.scanWhitespace()
	case r == '#':
		return l.scanComment()
	case r == 'x' || r == 'X':
		return l.scanRegister()
	case r == ',':
		return l.emit(itemComma)
	case r == '-':
		return l.scanNegative()
	case r == '0':
		return l.scanZero()
	case r == '(':
		return l.emit(itemLeftParen)
	case r == ')':
		return l.emit(itemRightParen)
	case isLetter(r):
		return l.scanIdentifier()
	case isInteger(r):
		return l.scanInteger()
	default:
		return l.emitError("unrecognized character")
	}
}

// scanComments reads comments. # is known to be present already
func (l *lexer) scanComment() Item {
	l.acceptAll(isComment)
	return l.emit(itemComment)
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (l *lexer) scanWhitespace() Item {
	l.acceptAll(isWhitespace)
	return l.emit(itemSpace)
}

func (l *lexer) scanRegister() Item {
	if r := l.next(); !isInteger(r) {
		return l.scanIdentifier()
	}

	l.acceptAll(isInteger)
	return l.emit(itemRegister)
}

func (l *lexer) scanNegative() Item {
	switch r := l.next(); {
	case r == '>':
		return l.emit(itemExpect)
	case isInteger(r):
		return l.scanInteger()
	default:
		return l.emitError("unidentified token after negative sign")
	}
}

func (l *lexer) scanInteger() Item {
	l.acceptAll(isInteger)
	return l.emit(itemInteger)
}

// scanZero decides what to do after reading a zero. 0 is already known to be present.
func (l *lexer) scanZero() Item {
	// Check for "x"
	switch r := l.next(); {
	case r == 'x':
		return l.scanHex()
	case r == '(':
		l.backup()
		return l.emit(itemInteger)
	case isWhitespace(r) || r == eof:
		l.backup()
		return l.emit(itemInteger)
	default:
		return l.emitError("unidentified token after 0")
	}
}

// scanHex looks for hexidecimal token. 0x is already known to be present.
func (l *lexer) scanHex() Item {
	l.acceptAll(isHex)
	return l.emit(itemHex)
}

// scanIndentifier consumes the current rune and all contiguous identifier runes.
func (l *lexer) scanIdentifier() Item {
	l.acceptAll(isLetter)

	// If the string matches a keyword then return that keyword.
	upper := strings.ToUpper(l.buf.String())
	typ, ok := itemKey[upper]

	// Check if it's a valid identifier
	if !ok {
		return l.emitError("unrecognized identifier")
	}

	return l.emit(typ)
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}

func isComment(r rune) bool {
	return r != '\n'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isInteger(r rune) bool {
	return r >= '0' && r <= '9'
}

func isHex(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')
}
