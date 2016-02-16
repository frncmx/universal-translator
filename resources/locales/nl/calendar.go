package nl

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd-MM-yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "mrt.", 4: "apr.", 7: "jul.", 8: "aug.", 10: "okt.", 1: "jan.", 2: "feb.", 5: "mei", 6: "jun.", 9: "sep.", 11: "nov.", 12: "dec."}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 5: "M", 6: "J", 11: "N", 1: "J", 2: "F", 3: "M", 7: "J", 8: "A", 9: "S", 10: "O", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "januari", 6: "juni", 10: "oktober", 9: "september", 11: "november", 2: "februari", 3: "maart", 4: "april", 5: "mei", 7: "juli", 8: "augustus", 12: "december"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "vr", 6: "za", 0: "zo", 1: "ma", 2: "di", 3: "wo", 4: "do"}, Narrow: ut.CalendarDayFormatNameValue{5: "V", 6: "Z", 0: "Z", 1: "M", 2: "D", 3: "W", 4: "D"}, Short: ut.CalendarDayFormatNameValue{2: "di", 3: "wo", 4: "do", 5: "vr", 6: "za", 0: "zo", 1: "ma"}, Wide: ut.CalendarDayFormatNameValue{4: "donderdag", 5: "vrijdag", 6: "zaterdag", 0: "zondag", 1: "maandag", 2: "dinsdag", 3: "woensdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "middernacht", "am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "a.m.", "pm": "p.m.", "morning1": "ochtend", "afternoon1": "middag", "evening1": "avond", "night1": "nacht", "midnight": "middernacht"}}}}
