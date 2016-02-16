package khq

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM, y", Short: "d/M/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "Sek", 10: "Okt", 12: "Dee", 4: "Awi", 5: "Me", 6: "Žuw", 7: "Žuy", 11: "Noo", 1: "Žan", 2: "Fee", 3: "Mar", 8: "Ut"}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 6: "Ž", 7: "Ž", 8: "U", 9: "S", 10: "O", 11: "N", 1: "Ž", 12: "D", 3: "M", 5: "M", 2: "F"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{8: "Ut", 9: "Sektanbur", 4: "Awiril", 6: "Žuweŋ", 3: "Marsi", 5: "Me", 7: "Žuyye", 10: "Oktoobur", 11: "Noowanbur", 12: "Deesanbur", 1: "Žanwiye", 2: "Feewiriye"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Alh", 1: "Ati", 2: "Ata", 3: "Ala", 4: "Alm", 5: "Alj", 6: "Ass"}, Narrow: ut.CalendarDayFormatNameValue{1: "T", 2: "T", 3: "L", 4: "L", 5: "L", 6: "S", 0: "H"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{2: "Atalata", 3: "Alarba", 4: "Alhamiisa", 5: "Aljuma", 6: "Assabdu", 0: "Alhadi", 1: "Atini"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Adduha", "pm": "Aluula"}}}}
