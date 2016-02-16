package dsb

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "d.M.y", Short: "d.M.yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "jan", 5: "maj", 8: "awg", 9: "sep", 12: "dec", 2: "feb", 3: "měr", 4: "apr", 6: "jun", 7: "jul", 10: "okt", 11: "now"}, Narrow: ut.CalendarMonthFormatNameValue{5: "m", 6: "j", 9: "s", 10: "o", 12: "d", 1: "j", 2: "f", 3: "m", 4: "a", 7: "j", 8: "a", 11: "n"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "februar", 3: "měrc", 5: "maj", 6: "junij", 10: "oktober", 11: "nowember", 12: "december", 1: "januar", 7: "julij", 8: "awgust", 9: "september", 4: "apryl"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "sob", 0: "nje", 1: "pón", 2: "wał", 3: "srj", 4: "stw", 5: "pět"}, Narrow: ut.CalendarDayFormatNameValue{4: "s", 5: "p", 6: "s", 0: "n", 1: "p", 2: "w", 3: "s"}, Short: ut.CalendarDayFormatNameValue{0: "nj", 1: "pó", 2: "wa", 3: "sr", 4: "st", 5: "pě", 6: "so"}, Wide: ut.CalendarDayFormatNameValue{0: "njeźela", 1: "pónjeźele", 2: "wałtora", 3: "srjoda", 4: "stwórtk", 5: "pětk", 6: "sobota"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue{"am": "dop.", "pm": "wótp."}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "dopołdnja", "pm": "wótpołdnja"}}}}
