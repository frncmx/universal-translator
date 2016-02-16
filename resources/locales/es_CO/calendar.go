package es_CO

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d/MM/y", Short: "d/MM/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "feb.", 3: "mar.", 5: "may.", 8: "ago.", 10: "oct.", 11: "nov.", 1: "ene.", 4: "abr.", 6: "jun.", 7: "jul.", 9: "sept.", 12: "dic."}, Narrow: ut.CalendarMonthFormatNameValue{11: "N", 1: "E", 3: "M", 9: "S", 8: "A", 12: "D", 4: "A", 6: "J", 2: "F", 5: "M", 7: "J", 10: "O"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{11: "noviembre", 7: "julio", 10: "octubre", 8: "agosto", 9: "septiembre", 12: "diciembre", 4: "abril", 6: "junio", 2: "febrero", 5: "mayo", 1: "enero", 3: "marzo"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "dom.", 1: "lun.", 2: "mar.", 3: "mié.", 4: "jue.", 5: "vie.", 6: "sáb."}, Narrow: ut.CalendarDayFormatNameValue{4: "j", 5: "v", 6: "s", 0: "d", 1: "l", 2: "m", 3: "m"}, Short: ut.CalendarDayFormatNameValue{4: "JU", 5: "VI", 6: "SA", 0: "DO", 1: "LU", 2: "MA", 3: "MI"}, Wide: ut.CalendarDayFormatNameValue{1: "lunes", 2: "martes", 3: "miércoles", 4: "jueves", 5: "viernes", 6: "sábado", 0: "domingo"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "tarde", "night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning2": "a.\u00a0m.", "evening1": "p.\u00a0m.", "night1": "p.\u00a0m.", "morning1": "madrugada", "am": "a. m.", "noon": "m.", "pm": "p. m."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "noche", "am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde"}}}}
