### 时间介绍

时间是一个物理的计量概念，与生产生活息息相关。在程序的世界中，通常会看到各种眼花缭乱的时间格式，本文将对这一概念进行展开。

经历了15世纪的大航海时代，18世纪的工业革命，世界各地的交往更加的密切。各地区的生活劳作习惯都以太阳为参考，各地区都有一套自己的时区划分。为了统一时区标准，为提高沟通效率避免混乱，各国的代表1884年在美国华盛顿召开了国际大会，选出英国伦敦的格林威治作为全球时间的中心点，并由它负责维护和计算，从1924年开始，格林威治天文台每小时就会向全世界播报时间（截止到1979年）

以本初子午线为中心，按照地球自转方向，每隔经度15°划分一个时区的方法，全球共分为24个时区：东1区至东12区，西1区至西12区，其中东西12区跨度都是7.5°也叫半时区。

### GMT时间

格林威治时间，也叫世界标准时间。是指位于英国伦敦郊区的【皇家格林尼治天文台】的标准时间，是本初子午线上的地方时，是0时区的区时

GMT：老的时间计量标准，根据地球的自转和公转来计算时间的，自转一圈是一天，公转一圈是一年。但是呢，地球公转的轨道是椭圆形的, 存在一定的时间偏差

'Mon, 30 Nov 2020 01:07:30 GMT'


```
HTTP-date    = rfc1123-date | rfc850-date | asctime-date
rfc1123-date = wkday "," SP date1 SP time SP "GMT"
rfc850-date  = weekday "," SP date2 SP time SP "GMT"
asctime-date = wkday SP date3 SP time SP 4DIGIT
date1        = 2DIGIT SP month SP 4DIGIT
                    ; day month year (e.g., 02 Jun 1982)
date2        = 2DIGIT "-" month "-" 2DIGIT
                    ; day-month-year (e.g., 02-Jun-82)
date3        = month SP ( 2DIGIT | ( SP 1DIGIT ))
                    ; month day (e.g., Jun  2)
time         = 2DIGIT ":" 2DIGIT ":" 2DIGIT
                    ; 00:00:00 - 23:59:59
wkday        = "Mon" | "Tue" | "Wed"
                    | "Thu" | "Fri" | "Sat" | "Sun"
weekday      = "Monday" | "Tuesday" | "Wednesday"
                    | "Thursday" | "Friday" | "Saturday" | "Sunday"
month        = "Jan" | "Feb" | "Mar" | "Apr"
                    | "May" | "Jun" | "Jul" | "Aug"
                    | "Sep" | "Oct" | "Nov" | "Dec"
```

```
date-time   =  [ day "," ] date time        ; dd mm yy
                                            ;  hh:mm:ss zzz

day         =  "Mon"  / "Tue" /  "Wed"  / "Thu"
                /  "Fri"  / "Sat" /  "Sun"

date        =  1*2DIGIT month 2DIGIT        ; day month year
                                            ;  e.g. 20 Jun 82

month       =  "Jan"  /  "Feb" /  "Mar"  /  "Apr"
                /  "May"  /  "Jun" /  "Jul"  /  "Aug"
                 /  "Sep"  /  "Oct" /  "Nov"  /  "Dec"

time        =  hour zone                    ; ANSI and Military

hour        =  2DIGIT ":" 2DIGIT [":" 2DIGIT]
                                            ; 00:00:00 - 23:59:59

zone        =  "UT"  / "GMT"                ; Universal Time
                                            ; North American : UT
                /  "EST" / "EDT"            ;  Eastern:  - 5/ - 4
                /  "CST" / "CDT"            ;  Central:  - 6/ - 5
                /  "MST" / "MDT"            ;  Mountain: - 7/ - 6
                /  "PST" / "PDT"            ;  Pacific:  - 8/ - 7
                /  1ALPHA                   ; Military: Z = UT;
                                            ;  A:-1; (J not used)
                                            ;  M:-12; N:+1; Y:+12
                / ( ("+" / "-") 4DIGIT )    ; Local differential
                                            ;  hours+min. (HHMM)
```

### utc时间(时间标准时间)

世界协调时间。它是以原子时作为计量单位的时间，计算结果极其严谨和精密

原子时：物质的原子内部发射的电磁振荡频率为基准的时间计量系统。美国的物理实验市在2014年造出了人类历史上最精确的原子钟，50亿年误差1s，可谓相当靠谱了。中国的铯原子钟也能确保2000万年误差不超过1s。

大事记：1979年12月初内瓦举行的世界无线电行政大会通过决议，确定用“世界协调时间（UTC时间）”取代“格林威治时间（GMT时间）”，作为无线电通信领域内的国际标准时间

### iso时间

### 时间戳



