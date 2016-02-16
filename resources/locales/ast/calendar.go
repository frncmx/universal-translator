package ast

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM 'de' y", Long: "d MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'a' 'les' {0}", Long: "{1} 'a' 'les' {0}", Medium: "{1}, {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Xin", 2: "Feb", 6: "Xun", 8: "Ago", 10: "Och", 11: "Pay", 3: "Mar", 4: "Abr", 5: "May", 7: "Xnt", 9: "Set", 12: "Avi"}, Narrow: ut.CalendarMonthFormatNameValue{1: "X", 4: "A", 5: "M", 7: "X", 9: "S", 10: "O", 11: "P", 12: "A", 2: "F", 3: "M", 6: "X", 8: "A"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "marzu", 5: "mayu", 7: "xunetu", 8: "agostu", 9: "setiembre", 10: "ochobre", 11: "payares", 2: "febreru", 4: "abril", 6: "xunu", 12: "avientu", 1: "xineru"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "mié", 4: "xue", 5: "vie", 6: "sáb", 0: "dom", 1: "llu", 2: "mar"}, Narrow: ut.CalendarDayFormatNameValue{1: "L", 2: "M", 3: "M", 4: "X", 5: "V", 6: "S", 0: "D"}, Short: ut.CalendarDayFormatNameValue{2: "ma", 3: "mi", 4: "xu", 5: "vi", 6: "sá", 0: "do", 1: "ll"}, Wide: ut.CalendarDayFormatNameValue{3: "miércoles", 4: "xueves", 5: "vienres", 6: "sábadu", 0: "domingu", 1: "llunes", 2: "martes"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "PM", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "a", "pm": "p"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "mañana", "pm": "tardi"}}}}
