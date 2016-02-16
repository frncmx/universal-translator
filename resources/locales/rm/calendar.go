package rm

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, 'ils' d 'da' MMMM y", Long: "d 'da' MMMM y", Medium: "dd-MM-y", Short: "dd-MM-yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "schan.", 4: "avr.", 8: "avust", 9: "sett.", 12: "dec.", 10: "oct.", 11: "nov.", 2: "favr.", 3: "mars", 5: "matg", 6: "zercl.", 7: "fan."}, Narrow: ut.CalendarMonthFormatNameValue{9: "S", 11: "N", 12: "D", 2: "F", 3: "M", 4: "A", 7: "F", 8: "A", 10: "O", 1: "S", 5: "M", 6: "Z"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "favrer", 5: "matg", 6: "zercladur", 7: "fanadur", 9: "settember", 10: "october", 12: "december", 1: "schaner", 4: "avrigl", 8: "avust", 11: "november", 3: "mars"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "gli", 2: "ma", 3: "me", 4: "gie", 5: "ve", 6: "so", 0: "du"}, Narrow: ut.CalendarDayFormatNameValue{5: "V", 6: "S", 0: "D", 1: "G", 2: "M", 3: "M", 4: "G"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{2: "mardi", 3: "mesemna", 4: "gievgia", 5: "venderdi", 6: "sonda", 0: "dumengia", 1: "glindesdi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "am", "pm": "sm"}}}}
