%{
package parser
%}

%union {
  node      Node
  item      Token
  datetime  *DateTime
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

%type <datetime> datetime
%type <node>  date year month day year2 time hour minute second millisecond
%type <item> DIGIT

%start main

%%
main: datetime
  {
    yylex.(*Lexer).Result = $1
  }
| error
  {
    yylex.(*Lexer).unexpected("parse error...", "")
  }

datetime: date
  {
    y := $1.(*Date)
    $$ = &DateTime{Year: y.Year , Month: y.Month, Day: y.Day}
  }
| time
  {
    t := $1.(*Time)
    $$ = &DateTime{Hour: t.Hour, Minute: t.Minute, Second: t.Second}
  }
| date SPACE time
  {
    y := $1.(*Date)
    t := $3.(*Time)
    $$ = &DateTime{Year: y.Year, Month: y.Month, Day: y.Day, Hour: t.Hour, Minute: t.Minute, Second: t.Second}
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

time: hour COLON minute COLON second
  {
    $$ = &Time{Hour: int($1.(Hour)), Minute: int($3.(Minute)), Second: int($5.(Second))}
  }
| hour COLON minute COLON second COLON millisecond
  {
    $$ = &Time{Hour: int($1.(Hour)), Minute: int($3.(Minute)), Second: int($5.(Second)), Millisecond: int($7.(Millisecond))}
  }


hour: DIGIT DIGIT
  {
    $$ = JoinHour($1, $2)
  }
| DIGIT
  {
    $$ = JoinHour($1)
  }

minute: DIGIT DIGIT
  {
    $$ = JoinMinute($1, $2)
  }
| DIGIT
  {
    $$ = JoinMinute($1)
  }

second: DIGIT DIGIT
  {
    $$ = JoinSecond($1, $2)
  }
| DIGIT
  {
    $$ = JoinSecond($1)
  }

millisecond: DIGIT
  {
    $$ = JoinMillisecond($1)
  }
| DIGIT DIGIT
  {
    $$ = JoinMillisecond($1, $2)
  }
| DIGIT DIGIT DIGIT
  {
    $$ = JoinMillisecond($1, $2, $3)
  }
| DIGIT DIGIT DIGIT DIGIT
  {
    $$ = JoinMillisecond($1, $2, $3, $4)
  }
| DIGIT DIGIT DIGIT DIGIT DIGIT
  {
    $$ = JoinMillisecond($1, $2, $3, $4, $5)
  }
| DIGIT DIGIT DIGIT DIGIT DIGIT DIGIT
  {
    $$ = JoinMillisecond($1, $2, $3, $4, $5, $6)
  }