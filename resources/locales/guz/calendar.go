package guz

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "Feb", 3: "Mac", 4: "Apr", 5: "Mei", 6: "Jun", 8: "Agt", 10: "Okt", 12: "Dis", 1: "Can", 7: "Cul", 9: "Sep", 11: "Nob"}, Narrow: ut.CalendarMonthFormatNameValue{12: "D", 3: "M", 4: "A", 8: "A", 9: "S", 10: "O", 11: "N", 1: "C", 2: "F", 5: "M", 6: "J", 7: "C"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{7: "Chulai", 11: "Nobemba", 12: "Disemba", 1: "Chanuari", 2: "Feburari", 3: "Machi", 4: "Apiriri", 6: "Juni", 5: "Mei", 8: "Agosti", 9: "Septemba", 10: "Okitoba"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "Ars", 5: "Icm", 6: "Est", 0: "Cpr", 1: "Ctt", 2: "Cmn", 3: "Cmt"}, Narrow: ut.CalendarDayFormatNameValue{0: "C", 1: "C", 2: "C", 3: "C", 4: "A", 5: "I", 6: "E"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{1: "Chumatato", 2: "Chumaine", 3: "Chumatano", 4: "Aramisi", 5: "Ichuma", 6: "Esabato", 0: "Chumapiri"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Ma/Mo", "pm": "Mambia/Mog"}}}}
