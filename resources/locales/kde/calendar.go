package kde

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{7: "Jul", 8: "Ago", 12: "Des", 1: "Jan", 2: "Feb", 3: "Mac", 4: "Apr", 11: "Nov", 5: "Mei", 6: "Jun", 9: "Sep", 10: "Okt"}, Narrow: ut.CalendarMonthFormatNameValue{7: "J", 8: "A", 1: "J", 2: "F", 3: "M", 4: "A", 5: "M", 6: "J", 11: "N", 12: "D", 9: "S", 10: "O"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{6: "Mwedi wa Nnyano na Umo", 11: "Mwedi wa Nnyano na Nnyano na U", 2: "Mwedi wa Pili", 5: "Mwedi wa Nnyano", 4: "Mwedi wa Nchechi", 7: "Mwedi wa Nnyano na Mivili", 8: "Mwedi wa Nnyano na Mitatu", 9: "Mwedi wa Nnyano na Nchechi", 10: "Mwedi wa Nnyano na Nnyano", 12: "Mwedi wa Nnyano na Nnyano na M", 1: "Mwedi Ntandi", 3: "Mwedi wa Tatu"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "Ll7", 6: "Ll1", 0: "Ll2", 1: "Ll3", 2: "Ll4", 3: "Ll5", 4: "Ll6"}, Narrow: ut.CalendarDayFormatNameValue{2: "4", 3: "5", 4: "6", 5: "7", 6: "1", 0: "2", 1: "3"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{4: "Liduva lyannyano na linji", 5: "Liduva lyannyano na mavili", 6: "Liduva litandi", 0: "Liduva lyapili", 1: "Liduva lyatatu", 2: "Liduva lyanchechi", 3: "Liduva lyannyano"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Muhi", "pm": "Chilo"}}}}
