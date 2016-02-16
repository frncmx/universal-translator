package es_DO

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "sept.", 1: "ene.", 7: "jul.", 10: "oct.", 3: "mar.", 12: "dic.", 4: "abr.", 5: "may.", 6: "jun.", 8: "ago.", 11: "nov.", 2: "feb."}, Narrow: ut.CalendarMonthFormatNameValue{10: "O", 5: "M", 11: "N", 12: "D", 3: "M", 4: "A", 6: "J", 2: "F", 7: "J", 8: "A", 9: "S", 1: "E"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{9: "septiembre", 12: "diciembre", 2: "febrero", 7: "julio", 11: "noviembre", 4: "abril", 8: "agosto", 5: "mayo", 10: "octubre", 1: "enero", 3: "marzo", 6: "junio"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "lun.", 2: "mar.", 3: "mié.", 4: "jue.", 5: "vie.", 6: "sáb.", 0: "dom."}, Narrow: ut.CalendarDayFormatNameValue{5: "V", 6: "S", 0: "D", 1: "L", 2: "M", 3: "X", 4: "J"}, Short: ut.CalendarDayFormatNameValue{6: "SA", 0: "DO", 1: "LU", 2: "MA", 3: "MI", 4: "JU", 5: "VI"}, Wide: ut.CalendarDayFormatNameValue{3: "miércoles", 4: "jueves", 5: "viernes", 6: "sábado", 0: "domingo", 1: "lunes", 2: "martes"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning2": "mañana", "evening1": "tarde", "night1": "noche", "noon": "m.", "am": "a. m.", "pm": "p. m.", "morning1": "madrugada"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche"}}}}
