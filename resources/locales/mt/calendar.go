package mt

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'ta'’ MMMM y", Long: "d 'ta'’ MMMM y", Medium: "dd MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "Aww", 9: "Set", 10: "Ott", 11: "Nov", 6: "Ġun", 7: "Lul", 3: "Mar", 4: "Apr", 5: "Mej", 12: "Diċ", 1: "Jan", 2: "Fra"}, Narrow: ut.CalendarMonthFormatNameValue{7: "Lj", 8: "Aw", 11: "Nv", 12: "Dċ", 1: "Jn", 2: "Fr", 4: "Ap", 9: "St", 10: "Ob", 3: "Mz", 5: "Mj", 6: "Ġn"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "Marzu", 4: "April", 5: "Mejju", 6: "Ġunju", 9: "Settembru", 10: "Ottubru", 11: "Novembru", 1: "Jannar", 2: "Frar", 7: "Lulju", 8: "Awwissu", 12: "Diċembru"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Tli", 3: "Erb", 4: "Ħam", 5: "Ġim", 6: "Sib", 0: "Ħad", 1: "Tne"}, Narrow: ut.CalendarDayFormatNameValue{4: "Ħm", 5: "Ġm", 6: "Sb", 0: "Ħd", 1: "Tn", 2: "Tl", 3: "Er"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "Il-Ħadd", 1: "It-Tnejn", 2: "It-Tlieta", 3: "L-Erbgħa", 4: "Il-Ħamis", 5: "Il-Ġimgħa", 6: "Is-Sibt"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "am": "AM"}}}}
