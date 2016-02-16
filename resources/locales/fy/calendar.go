package fy

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "mrt", 6: "jun", 7: "jul", 10: "okt", 1: "jan", 4: "apr", 5: "mai", 8: "aug", 9: "sep", 11: "nov", 12: "des", 2: "feb"}, Narrow: ut.CalendarMonthFormatNameValue{12: "D", 1: "J", 6: "J", 4: "A", 5: "M", 7: "J", 8: "A", 9: "S", 10: "O", 2: "F", 3: "M", 11: "N"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{8: "augustus", 10: "oktober", 12: "desimber", 2: "febrewaris", 3: "maart", 5: "maaie", 7: "july", 11: "novimber", 1: "jannewaris", 4: "april", 6: "juny", 9: "septimber"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "si", 1: "mo", 2: "ti", 3: "wo", 4: "to", 5: "fr", 6: "so"}, Narrow: ut.CalendarDayFormatNameValue{1: "M", 2: "D", 3: "W", 4: "D", 5: "V", 6: "Z", 0: "Z"}, Short: ut.CalendarDayFormatNameValue{5: "fr", 6: "so", 0: "si", 1: "mo", 2: "ti", 3: "wo", 4: "to"}, Wide: ut.CalendarDayFormatNameValue{5: "freed", 6: "sneon", 0: "snein", 1: "moandei", 2: "tiisdei", 3: "woansdei", 4: "tongersdei"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "a.m.", "pm": "p.m."}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"pm": "p.m.", "am": "foarmiddei"}}}}
