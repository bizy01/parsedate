package parser

import (
    "errors"
    "strings"
    "unicode/utf8"
    "fmt"
)

var errUnexpected = errors.New("unexpected error")
var eof rune = -1

type TokenType int
type Pos int

type Token struct {
    Typ TokenType
    Pos Pos
    Val string
}

var keyWords = map[string]TokenType {
    "january":      JANUARY,
    "february":     FEBRUARY,
    "march":        MARCH,
    "april":        APRIL,
    "may":          MAY,
    "june":         JUNE,
    "july":         JULY,
    "august":       AUGUST,
    "september":    SEPTEMBER,
    "october":      OCTOBER,
    "november":     NOVEMBER,
    "december":     DECEMBER,
    "jan":          JANUARY,
    "feb":          FEBRUARY,
    "mar":          MARCH,
    "apr":          APRIL,
    "jun":          JUNE,
    "jul":          JULY,
    "aug":          AUGUST,
    "sept":         SEPTEMBER,
    "oct":          OCTOBER,
    "nov":          NOVEMBER,
    "dec":          DECEMBER,
}

// Parse parses the input and returs the result.
func ParseDate(input string) (Node, error) {
    l := newLexer(input)
    yyParse(l)
    return l.Result, l.err
}

// 状态机
type stateFn func(*Lexer) stateFn

type Lexer struct {
    input  string
    state  stateFn
    start  Pos
    pos    Pos
    width  Pos
    tmpToken *Token
    scanned bool
    lastPos Pos
    Result Node
    err    error
}

func newLexer(input string) *Lexer {
    l := &Lexer{
        input: input,
        state: lexStatements,
    }
    return l
}

func lexStatements(l *Lexer) stateFn {
    switch r := l.next(); {
        case r == eof:
            l.emit(EOF)
            return nil
        case r == ' ':
            l.emit(SPACE)
        case r == '-':
            l.emit(DASH)
        case r == ':':
            l.emit(COLON)
        case r == '/':
            l.emit(DIV)
        case r == '.':
            l.emit(DOT)
        case r == ',':
            l.emit(COMMA)
        case r == '年':
            l.emit(HANYEAR)
        case r == '月':
            l.emit(HANMONTH)
        case r == '日':
            l.emit(HANDAY)
        case isDigit(r):
            l.emit(DIGIT)
        case isAlpha(r):
            l.back()
            return lexAlphaStatements
        default:
            return l.errorf("unexpected character: %q", r)
    }
    return lexStatements
}

// alpha
func lexAlphaStatements(l *Lexer) stateFn {
Loop:
    for {
        switch r := l.next(); {
        case isAlpha(r):
            // absorb
        default:
            l.back()
            word := l.input[l.start:l.pos]
            if _, ok := keyWords[strings.ToLower(word)]; ok {
                // 英文数字转换(todo)
                l.emit(MONTH)
                break Loop
            } else {
                return l.errorf("unexpected key word: %s", word)
            }
        }
    }
    return lexStatements
}

func (l *Lexer) errorf(format string, args ...interface{}) stateFn {
    *l.tmpToken = Token{ERROR, l.start, fmt.Sprintf(format, args...)}
    l.scanned = true

    return nil
}

func (l *Lexer) next() rune {
    if int(l.pos) >= len(l.input) {
        l.width =0
        return eof
    }
    r, w := utf8.DecodeRuneInString(l.input[l.pos:])
    l.width = Pos(w)
    l.pos += l.width

    return r
}

func (l *Lexer) back() {
    l.pos -= l.width
}

// 预读
func (l *Lexer) peek() rune {
    r := l.next()
    l.back()

    return r
}

func (l *Lexer) ignore() {
    l.start = l.pos
}

func (l *Lexer) emit(t TokenType) {
    tmpToken := Token{
        Typ: t,
        Pos: l.start,
        Val: l.input[l.start:l.pos],
    }

    *l.tmpToken = tmpToken
    l.start = l.pos
    l.scanned = true
}

// Lex satisfies yyLexer.
func (l *Lexer) Lex(lval *yySymType) int {
    var typ TokenType
    for {
        l.NextItem(&lval.item)
        typ = lval.item.Typ
        break
    }

    switch typ {
    case ERROR:
        return 0
    case EOF:
        return 0
    }

    return int(typ)
}

func (l *Lexer) NextItem(itemp *Token) {
    l.scanned = false
    l.tmpToken = itemp

    if l.state != nil {
        for !l.scanned {
            l.state = l.state(l)
        }
    } else {
        l.emit(EOF)
    }

    l.lastPos = l.tmpToken.Pos
}

// Error satisfies yyLexer.
func (l *Lexer) Error(s string) {
    l.err = errors.New(s)
}

func isSpace(r rune) bool {
    return r == ' ' || r == '\t'
}

func isDigit(r rune) bool {
    return '0' <= r && r <= '9'
}

// isAlpha reports whether r is an alphabetic or underscore.
func isAlpha(r rune) bool {
    return r == '_' || ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z')
}
