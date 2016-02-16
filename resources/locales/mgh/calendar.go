package mgh

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Kwa", 2: "Una", 3: "Rar", 4: "Che", 6: "Moc", 7: "Sab", 8: "Nan", 10: "Kum", 11: "Moj", 5: "Tha", 9: "Tis", 12: "Yel"}, Narrow: ut.CalendarMonthFormatNameValue{3: "R", 4: "C", 5: "T", 9: "T", 10: "K", 1: "K", 2: "U", 6: "M", 7: "S", 8: "N", 11: "M", 12: "Y"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{6: "Mweri wo thanu na mocha", 10: "Mweri wo kumi", 11: "Mweri wo kumi na moja", 12: "Mweri wo kumi na yel’li", 1: "Mweri wo kwanza", 2: "Mweri wo unayeli", 3: "Mweri wo uneraru", 4: "Mweri wo unecheshe", 5: "Mweri wo unethanu", 7: "Mweri wo saba", 8: "Mweri wo nane", 9: "Mweri wo tisa"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Sab", 1: "Jtt", 2: "Jnn", 3: "Jtn", 4: "Ara", 5: "Iju", 6: "Jmo"}, Narrow: ut.CalendarDayFormatNameValue{6: "J", 0: "S", 1: "J", 2: "J", 3: "J", 4: "A", 5: "I"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{4: "Arahamisi", 5: "Ijumaa", 6: "Jumamosi", 0: "Sabato", 1: "Jumatatu", 2: "Jumanne", 3: "Jumatano"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "wichishu", "pm": "mchochil’l"}}}}
