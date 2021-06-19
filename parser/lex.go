package parser

import (
    "errors"
    "strings"
    "unicode/utf8"
    "fmt"
)

var errUnexpected = errors.New("unexpected error")
var eof rune = -1

type ParseErr struct {
    PositionRange PositionRange
    Err           error
    Input         string
    LineOffset    int
}

func (e *ParseErr) Error() string {
    pos := int(e.PositionRange.Start)
    lastLineBreak := -1
    line := e.LineOffset + 1

    var positionStr string

    if pos < 0 || pos > len(e.Input) {
        positionStr = "invalid position:"
    } else {

        for i, c := range e.Input[:pos] {
            if c == '\n' {
                lastLineBreak = i
                line++
            }
        }

        col := pos - lastLineBreak
        positionStr = fmt.Sprintf("%d:%d:", line, col)
    }
    return fmt.Sprintf("%s parse error: %s", positionStr, e.Err)
}

type ParseErrors []ParseErr

func (errs ParseErrors) Error() string {
    if len(errs) != 0 {
        return errs[0].Error()
    }
    return "error contains no error message"
}

type TokenType int
type Pos int

type Token struct {
    Typ TokenType
    Pos Pos
    Val string
}

var monthWords = map[string]int {
    "january":      1,
    "february":     2,
    "march":        3,
    "april":        4,
    "may":          5,
    "june":         6,
    "july":         7,
    "august":       8,
    "september":    9,
    "october":      10,
    "november":     11,
    "december":     12,
    "jan":          1,
    "feb":          2,
    "mar":          3,
    "apr":          4,
    "jun":          6,
    "jul":          7,
    "aug":          8,
    "sept":         9,
    "oct":          10,
    "nov":          11,
    "dec":          12,
}

var weekWords = map[string]int {
    mon: 1,
    tue: 2,
    wed: 3,
    thur: 4,
    fri: 5,
    sat: 6,
    sun: 0,
    monday: 1,
    tuesday: 2,
    wednesday: 3,
    thursday: 4,
    friday: 5,
    saturday: 6,
    sunday: 0,
}

var zoneWords = map[string]int {
    gmt: 1,
}

// Parse parses the input and returs the result.
func ParseDate(input string) (*DateTime, error) {
    var err error
    l := newLexer(input)
    yyParse(l)

    if len(l.parseErrors) != 0 {
        err = l.parseErrors
    }

    return l.Result, err
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
    Result *DateTime
    parseErrors ParseErrors
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
            if _, ok := monthWords[strings.ToLower(word)]; ok {
                l.emit(MONTH)
                break Loop
            } else if _, ok := weekWords[strings.ToLower(word)]; ok {
                l.emit(WEEK)
                break Loop
            } else if _, ok := zoneWords[strings.ToLower(word)]; ok {
                l.emit(ZONE)
                break Loop
            } else {
                return l.errorf("unexpected key word: %s", word)
            }
        }
    }
    return lexStatements
}

func (l *Lexer) unexpected(context string, expected string) {
    var errMsg strings.Builder

    // Do not report lexer errors twice
    if l.tmpToken.Typ == ERROR {
        return
    }

    errMsg.WriteString("unexpected ")

    if context != "" {
        errMsg.WriteString(" in ")
        errMsg.WriteString(context)
    }

    if expected != "" {
        errMsg.WriteString(", expected ")
        errMsg.WriteString(expected)
    }

    l.addParseErr(PositionRange{}, errors.New(errMsg.String()))
}

func (l *Lexer) addParseErr(positionRange PositionRange, err error) {
    perr := ParseErr{
        PositionRange: positionRange,
        Err:           err,
        Input:         l.input,
    }

    l.parseErrors = append(l.parseErrors, perr)
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
        pos := PositionRange{
            Start: l.start,
            End:   Pos(len(l.input)),
        }

        l.addParseErr(pos, errors.New(lval.item.Val))

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
