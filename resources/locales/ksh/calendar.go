package ksh

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, 'dä' d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM. y", Short: "d. M. y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan.", 2: "Fäb.", 4: "Apr.", 5: "Mäi", 8: "Ouj.", 9: "Säp.", 10: "Okt.", 12: "Dez.", 3: "Mäz.", 6: "Jun.", 7: "Jul.", 11: "Nov."}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 6: "J", 7: "J", 11: "N", 12: "D", 10: "O", 1: "J", 2: "F", 3: "M", 5: "M", 8: "A", 9: "S"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{9: "Septämber", 12: "Dezämber", 2: "Fäbrowa", 3: "Määz", 4: "Aprell", 5: "Mäi", 7: "Juuli", 1: "Jannewa", 6: "Juuni", 8: "Oujoß", 10: "Oktoober", 11: "Novämber"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "Me.", 4: "Du.", 5: "Fr.", 6: "Sa.", 0: "Su.", 1: "Mo.", 2: "Di."}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "S", 1: "M", 2: "D", 3: "M", 4: "D", 5: "F"}, Short: ut.CalendarDayFormatNameValue{3: "Me", 4: "Du", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo", 2: "Di"}, Wide: ut.CalendarDayFormatNameValue{6: "Samsdaach", 0: "Sunndaach", 1: "Moondaach", 2: "Dinnsdaach", 3: "Metwoch", 4: "Dunnersdaach", 5: "Friidaach"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "v.m.", "pm": "n.m."}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"pm": "Nachmittag", "am": "Vormittag"}}}}
