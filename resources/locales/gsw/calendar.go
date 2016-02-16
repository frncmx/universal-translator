package gsw

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d. MMMM y", Long: "d. MMMM y", Medium: "dd.MM.y", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 5: "Mai", 11: "Nov", 8: "Aug", 9: "Sep", 10: "Okt", 2: "Feb", 3: "Mär", 4: "Apr", 6: "Jun", 7: "Jul", 12: "Dez"}, Narrow: ut.CalendarMonthFormatNameValue{10: "O", 11: "N", 1: "J", 3: "M", 4: "A", 6: "J", 7: "J", 2: "F", 5: "M", 8: "A", 9: "S", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{9: "Septämber", 11: "Novämber", 2: "Februar", 3: "März", 6: "Juni", 7: "Juli", 8: "Auguscht", 10: "Oktoober", 12: "Dezämber", 1: "Januar", 4: "April", 5: "Mai"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "Mi.", 4: "Du.", 5: "Fr.", 6: "Sa.", 0: "Su.", 1: "Mä.", 2: "Zi."}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "S", 1: "M", 2: "D", 3: "M", 4: "D", 5: "F"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{3: "Mittwuch", 4: "Dunschtig", 5: "Friitig", 6: "Samschtig", 0: "Sunntig", 1: "Määntig", 2: "Ziischtig"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "n.m.", "am": "v.m."}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Vormittag", "pm": "Namittag"}}}}
