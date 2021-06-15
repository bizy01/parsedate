%{
package parser
%}

%union {
  node      Node
  item      Token
}

%token <item>
DASH
DIV
DOT
DIGIT
BLANK
COLON
COMMA
EOF
ERROR
SPACE
MONTH
HANYEAR
HANMONTH
HANDAY

// Keywords
%token	keywordsStart
%token <item>
JANUARY
JAN
FEBRUARY
FEB
MARCH
MAR
APRIL
APR
MAY
JUNE
JUN
JULY
JUL
AUGUST
AUG
SEPTEMBER
SEPT
OCTOBER
OCT
NOVEMBER
NOV
DECEMBER
DEC
%token keywordsEnd

%type <node> date year month day year2
%type <item> DIGIT

%start main

%%
main: date
  {
    yylex.(*Lexer).Result = $1
  }

date: year SPACE month SPACE day
  {
    $$ = &Date{Year: int($1.(Year)), Month: int($3.(Month)), Day: int($5.(Day))}
  }
| year DIV month DIV day
  {
    $$ = &Date{Year: int($1.(Year)), Month: int($3.(Month)), Day: int($5.(Day))}
  }
| year HANYEAR month HANMONTH day HANDAY
  {
    $$ = &Date{Year: int($1.(Year)), Month: int($3.(Month)), Day: int($5.(Day))}
  }
| month DIV day DIV year
  {
    $$ = &Date{Year: int($5.(Year)), Month: int($1.(Month)), Day: int($3.(Day))}
  }
| year DASH month DASH day
  {
    $$ = &Date{Year: int($1.(Year)), Month: int($3.(Month)), Day: int($5.(Day))}
  }
| day SPACE month SPACE year
  {
    $$ = &Date{Year: int($5.(Year)), Month: int($3.(Month)), Day: int($1.(Day))}
  }
| day SPACE month SPACE year2
  {
    $$ = &Date{Year: int($5.(Year)), Month: int($3.(Month)), Day: int($1.(Day))}
  }
| MONTH SPACE day COMMA SPACE year
  {
    var m = ConvToManth($1.Val)
    $$ = &Date{Year: int($6.(Year)), Month: int(m), Day: int($3.(Day))}
  }
| MONTH DOT SPACE day COMMA SPACE year
  {
    var m = ConvToManth($1.Val)
    $$ = &Date{Year: int($7.(Year)), Month: int(m), Day: int($4.(Day))}
  }

year: DIGIT DIGIT DIGIT DIGIT
  {
    $$ = JoinYear($1, $2, $3, $4)
  }

year2: DIGIT DIGIT
  {
    $$ = JoinYear($1, $2)
  }

month: DIGIT DIGIT
  {
  	$$ = JoinMonth($1, $2)
  }
| DIGIT
  {
    $$ = JoinMonth($1)
  }
| MONTH
  {
    $$ = ConvToManth($1.Val)
  }

day: DIGIT DIGIT
  {
  	$$ = JoinDay($1, $2)
  }
| DIGIT
  {
    $$ = JoinDay($1)
  }

