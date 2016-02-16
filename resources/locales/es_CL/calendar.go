package es_CL

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "dd-MM-y", Short: "dd-MM-yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss (zzzz)", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "ene.", 4: "abr.", 5: "may.", 10: "oct.", 12: "dic.", 2: "feb.", 3: "mar.", 6: "jun.", 7: "jul.", 8: "ago.", 9: "sept.", 11: "nov."}, Narrow: ut.CalendarMonthFormatNameValue{12: "D", 1: "E", 6: "J", 2: "F", 9: "S", 11: "N", 3: "M", 4: "A", 10: "O", 5: "M", 7: "J", 8: "A"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{9: "septiembre", 12: "diciembre", 1: "enero", 7: "julio", 10: "octubre", 11: "noviembre", 2: "febrero", 6: "junio", 4: "abril", 8: "agosto", 5: "mayo", 3: "marzo"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "mar.", 3: "mié.", 4: "jue.", 5: "vie.", 6: "sáb.", 0: "dom.", 1: "lun."}, Narrow: ut.CalendarDayFormatNameValue{3: "X", 4: "J", 5: "V", 6: "S", 0: "D", 1: "L", 2: "M"}, Short: ut.CalendarDayFormatNameValue{4: "ju", 5: "vi", 6: "sá", 0: "do", 1: "lu", 2: "ma", 3: "mi"}, Wide: ut.CalendarDayFormatNameValue{6: "sábado", 0: "domingo", 1: "lunes", 2: "martes", 3: "miércoles", 4: "jueves", 5: "viernes"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m."}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "a. m.", "noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"noon": "mediodía", "pm": "p. m.", "morning1": "madrugada", "morning2": "mañana", "evening1": "tarde", "night1": "noche", "am": "a. m."}}}}
