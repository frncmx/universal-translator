package es_US

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "dic.", 9: "sept.", 3: "mar.", 5: "may.", 8: "ago.", 10: "oct.", 11: "nov.", 2: "feb.", 6: "jun.", 1: "ene.", 4: "abr.", 7: "jul."}, Narrow: ut.CalendarMonthFormatNameValue{5: "M", 7: "J", 8: "A", 10: "O", 4: "A", 9: "S", 11: "N", 12: "D", 2: "F", 1: "E", 3: "M", 6: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{10: "octubre", 11: "noviembre", 1: "enero", 3: "marzo", 4: "abril", 2: "febrero", 7: "julio", 6: "junio", 8: "agosto", 9: "septiembre", 12: "diciembre", 5: "mayo"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "jue.", 5: "vie.", 6: "sáb.", 0: "dom.", 1: "lun.", 2: "mar.", 3: "mié."}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "D", 1: "L", 2: "M", 3: "X", 4: "J", 5: "V"}, Short: ut.CalendarDayFormatNameValue{0: "DO", 1: "LU", 2: "MA", 3: "MI", 4: "JU", 5: "VI", 6: "SA"}, Wide: ut.CalendarDayFormatNameValue{2: "martes", 3: "miércoles", 4: "jueves", 5: "viernes", 6: "sábado", 0: "domingo", 1: "lunes"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "p", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a", "noon": "m"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning2": "mañana", "evening1": "tarde", "night1": "noche", "noon": "mediodía", "am": "AM", "pm": "PM", "morning1": "madrugada"}}}}
