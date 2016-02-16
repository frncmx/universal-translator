package prg

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, y 'mettas' d. MMMM", Long: "y 'mettas' d. MMMM", Medium: "dd.MM 'st'. y", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "was", 7: "līp", 8: "dag", 9: "sil", 10: "spa", 12: "sal", 1: "rag", 4: "sak", 5: "zal", 6: "sīm", 11: "lap", 3: "pūl"}, Narrow: ut.CalendarMonthFormatNameValue{3: "P", 4: "S", 9: "S", 11: "L", 12: "S", 1: "R", 2: "W", 7: "L", 8: "D", 10: "S", 5: "Z", 6: "S"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "rags", 2: "wassarins", 9: "sillins", 10: "spallins", 7: "līpa", 8: "daggis", 11: "lapkrūtis", 12: "sallaws", 3: "pūlis", 4: "sakkis", 5: "zallaws", 6: "sīmenis"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "nad", 1: "pan", 2: "wis", 3: "pus", 4: "ket", 5: "pēn", 6: "sab"}, Narrow: ut.CalendarDayFormatNameValue{0: "N", 1: "P", 2: "W", 3: "P", 4: "K", 5: "P", 6: "S"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{0: "nadīli", 1: "panadīli", 2: "wisasīdis", 3: "pussisawaiti", 4: "ketwirtiks", 5: "pēntniks", 6: "sabattika"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "ankstāinan", "pm": "pa pussideinan"}}}}
