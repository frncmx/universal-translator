package es_EC

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d MMM y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{6: "jun.", 1: "ene.", 4: "abr.", 9: "sept.", 5: "may.", 12: "dic.", 7: "jul.", 8: "ago.", 10: "oct.", 11: "nov.", 2: "feb.", 3: "mar."}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 6: "J", 9: "S", 12: "D", 7: "J", 8: "A", 1: "E", 3: "M", 5: "M", 10: "O", 11: "N", 2: "F"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{10: "octubre", 11: "noviembre", 3: "marzo", 8: "agosto", 9: "septiembre", 4: "abril", 6: "junio", 2: "febrero", 5: "mayo", 7: "julio", 12: "diciembre", 1: "enero"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "jue.", 5: "vie.", 6: "sáb.", 0: "dom.", 1: "lun.", 2: "mar.", 3: "mié."}, Narrow: ut.CalendarDayFormatNameValue{4: "J", 5: "V", 6: "S", 0: "D", 1: "L", 2: "M", 3: "X"}, Short: ut.CalendarDayFormatNameValue{6: "SA", 0: "DO", 1: "LU", 2: "MA", 3: "MI", 4: "JU", 5: "VI"}, Wide: ut.CalendarDayFormatNameValue{4: "jueves", 5: "viernes", 6: "sábado", 0: "domingo", 1: "lunes", 2: "martes", 3: "miércoles"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada"}}}}
