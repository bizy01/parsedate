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
PLUS
EOF
ERROR
SPACE
MONTH
HANYEAR
HANMONTH
HANDAY
WEEK
ZONE
AM
PM
T
Z

%type <datetime> ansic unixdate rubydate rfc822 rfc822z rfc850 rfc1123 rfc1123z rfc3339
%type <node> date year4 year2 month day time hour minute second millisecond offest
%type <item> DIGIT

%start main

%%
main: ansic
  {
    yylex.(*Lexer).Result = $1
  }
| unixdate
  {
    yylex.(*Lexer).Result = $1
  }
| rubydate
  {
    yylex.(*Lexer).Result = $1
  }
| rfc822
  {
    yylex.(*Lexer).Result = $1
  }
| rfc822z
  {
    yylex.(*Lexer).Result = $1
  }
| rfc850
  {
    yylex.(*Lexer).Result = $1
  }
| rfc1123
  {
    yylex.(*Lexer).Result = $1
  }
rfc1123z
  {
    yylex.(*Lexer).Result = $1
  }
| rfc3339
  {
    yylex.(*Lexer).Result = $1
  }
| error
  {
    yylex.(*Lexer).unexpected("", "")
  }

/* ansic time (Mon Jan _2 15:04:05 2006) */
ansic: WEEK SPACE MONTH SPACE day SPACE time SPACE year4
  {
    var m = ConvToManth($3.Val)
    var tm = $7.(*Time)
    $$ = &DateTime{Year: int($9.(Year)), Month: int(m), Day: int($5.(Day)), Hour: tm.Hour, Minute: tm.Minute, Second: tm.Second, Millisecond: tm.Millisecond}
  }

/* unixdate time (Mon Jan _2 15:04:05 MST 2006) */
unixdate: WEEK SPACE MONTH SPACE day SPACE time SPACE ZONE SPACE year4
  {
    var m = ConvToManth($3.Val)
    var tm = $7.(*Time)
    $$ = &DateTime{Year: int($11.(Year)), Month: int(m), Day: int($5.(Day)), Hour: tm.Hour, Minute: tm.Minute, Second: tm.Second, Millisecond: tm.Millisecond}
  }

/* rubydate time (Mon Jan 02 15:04:05 -0700 2006) */
rubydate: WEEK SPACE MONTH SPACE day SPACE time SPACE offest SPACE year4
  {
    var m = ConvToManth($3.Val)
    var tm = $7.(*Time)
    $$ = &DateTime{Year: int($11.(Year)), Month: int(m), Day: int($5.(Day)), Hour: tm.Hour, Minute: tm.Minute, Second: tm.Second, Millisecond: tm.Millisecond}
  }

/* rfc822 (02 Jan 06 15:04 MST) */
rfc822: day SPACE MONTH SPACE year2 SPACE time SPACE ZONE
  {
    var m = ConvToManth($3.Val)
    var tm = $7.(*Time)
    $$ = &DateTime{Year: int($5.(Year)), Month: int(m), Day: int($1.(Day)), Hour: tm.Hour, Minute: tm.Minute, Second: tm.Second, Millisecond: tm.Millisecond}
  }

/* RFC822Z (02 Jan 06 15:04 -0700) */
rfc822z: day SPACE MONTH SPACE year2 SPACE time SPACE offest
  {
    var m = ConvToManth($3.Val)
    var tm = $7.(*Time)
    $$ = &DateTime{Year: int($5.(Year)), Month: int(m), Day: int($1.(Day)), Hour: tm.Hour, Minute: tm.Minute, Second: tm.Second, Millisecond: tm.Millisecond}
  }

/* rfc850 (Monday, 02-Jan-06 15:04:05 MST) */
rfc850: WEEK COMMA SPACE day DASH MONTH DASH year2 SPACE time SPACE ZONE
  {
    var m = ConvToManth($6.Val)
    var tm = $10.(*Time)
    $$ = &DateTime{Year: int($8.(Year)), Month: int(m), Day: int($4.(Day)), Hour: tm.Hour, Minute: tm.Minute, Second: tm.Second, Millisecond: tm.Millisecond}
  }

/* RFC1123 (Mon, 02 Jan 2006 15:04:05 MST) */
rfc1123: WEEK COMMA SPACE day SPACE MONTH SPACE year4 SPACE time SPACE ZONE
  {
    var m = ConvToManth($6.Val)
    var tm = $10.(*Time)
    $$ = &DateTime{Year: int($8.(Year)), Month: int(m), Day: int($4.(Day)), Hour: tm.Hour, Minute: tm.Minute, Second: tm.Second, Millisecond: tm.Millisecond}
  }

/* RFC1123Z (Mon, 02 Jan 2006 15:04:05 -0700) */
rfc1123z: WEEK COMMA SPACE day SPACE MONTH SPACE year4 SPACE time SPACE offest
  {
    var m = ConvToManth($6.Val)
    var tm = $10.(*Time)
    $$ = &DateTime{Year: int($8.(Year)), Month: int(m), Day: int($4.(Day)), Hour: tm.Hour, Minute: tm.Minute, Second: tm.Second, Millisecond: tm.Millisecond}
  }

/* RFC3339 (2006-01-02T15:04:05Z07:00) */
rfc3339: date T time Z offest
  {
    var dt = $1.(*Date)
    var tm = $3.(*Time)
    $$ = &DateTime{Year: dt.Year, Month: dt.Month, Day: dt.Day, Hour: tm.Hour, Minute: tm.Minute, Second: tm.Second, Millisecond: tm.Millisecond}
  }

date: year4 DASH month DASH day
  {
    $$ = &Date{Year: int($1.(Year)), Month: int($3.(Month)), Day: int($5.(Day))}
  }

year4: DIGIT DIGIT DIGIT DIGIT
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

day: DIGIT DIGIT
  {
  	$$ = JoinDay($1, $2)
  }
| SPACE DIGIT
  {
    $$ = JoinDay($2)
  }
| DIGIT
  {
    $$ = JoinDay($1)
  }

time: hour COLON minute
  {
    $$ = &Time{Hour: int($1.(Hour)), Minute: int($3.(Minute))}
  }
| hour COLON minute COLON second
  {
    $$ = &Time{Hour: int($1.(Hour)), Minute: int($3.(Minute)), Second: int($5.(Second))}
  }
| hour COLON minute COLON second DOT millisecond
  {
    $$ = &Time{Hour: int($1.(Hour)), Minute: int($3.(Minute)), Second: int($5.(Second)), Millisecond: int($7.(Millisecond))}
  }

/* time zone offest */
offest: PLUS DIGIT DIGIT DIGIT DIGIT
  {
  }
| DASH DIGIT DIGIT DIGIT DIGIT
  {
  }
| PLUS DIGIT DIGIT COLON DIGIT DIGIT
  {
  }
| DASH DIGIT DIGIT COLON DIGIT DIGIT
  {
  }
| DIGIT DIGIT COLON DIGIT DIGIT
  {
  }
| PLUS DIGIT DIGIT
  {
  }
| DASH DIGIT DIGIT
  {
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
| DIGIT DIGIT DIGIT DIGIT DIGIT DIGIT DIGIT DIGIT DIGIT
  {
    $$ = JoinMillisecond($1, $2, $3, $4, $5, $6, $7, $8, $9)
  }