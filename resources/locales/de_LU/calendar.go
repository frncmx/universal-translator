package de_LU

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'um' {0}", Long: "{1} 'um' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 2: "Feb", 6: "Jun", 12: "Dez", 9: "Sep", 5: "Mai", 11: "Nov", 7: "Jul", 8: "Aug", 10: "Okt", 4: "Apr", 3: "Mär"}, Narrow: ut.CalendarMonthFormatNameValue{10: "O", 11: "N", 1: "J", 5: "M", 9: "S", 3: "M", 8: "A", 4: "A", 6: "J", 7: "J", 2: "F", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{10: "Oktober", 4: "April", 2: "Februar", 7: "Juli", 8: "August", 9: "September", 6: "Juni", 5: "Mai", 3: "März", 11: "November", 12: "Dezember", 1: "Januar"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "Sa", 0: "So", 1: "Mo", 2: "Di", 3: "Mi", 4: "Do", 5: "Fr"}, Narrow: ut.CalendarDayFormatNameValue{0: "S", 1: "M", 2: "D", 3: "M", 4: "D", 5: "F", 6: "S"}, Short: ut.CalendarDayFormatNameValue{2: "Di.", 3: "Mi.", 4: "Do.", 5: "Fr.", 6: "Sa.", 0: "So.", 1: "Mo."}, Wide: ut.CalendarDayFormatNameValue{2: "Dienstag", 3: "Mittwoch", 4: "Donnerstag", 5: "Freitag", 6: "Samstag", 0: "Sonntag", 1: "Montag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "Morgen", "midnight": "Mitternacht", "afternoon1": "Mittag", "evening1": "Abend", "morning2": "Vormittag", "am": "vorm.", "pm": "nachm.", "night1": "Nacht", "afternoon2": "Nachmittag"}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon2": "Nachmittag", "am": "vorm.", "afternoon1": "Mittag", "evening1": "Abend", "midnight": "Mitternacht", "pm": "nachm.", "morning2": "Vormittag", "night1": "Nacht", "morning1": "Morgen"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"midnight": "Mitternacht", "am": "vorm.", "morning1": "Morgen", "afternoon1": "Mittag", "night1": "Nacht", "afternoon2": "Nachmittag", "evening1": "Abend", "pm": "nachm.", "morning2": "Vormittag"}}}}
