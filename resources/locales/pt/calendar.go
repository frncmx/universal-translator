package pt

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "jan", 3: "mar", 6: "jun", 9: "set", 12: "dez", 11: "nov", 2: "fev", 4: "abr", 5: "mai", 7: "jul", 8: "ago", 10: "out"}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 6: "J", 7: "J", 9: "S", 11: "N", 10: "O", 12: "D", 1: "J", 2: "F", 3: "M", 5: "M", 8: "A"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "outubro", 11: "novembro", 12: "dezembro", 2: "fevereiro", 3: "março", 5: "maio", 7: "julho", 8: "agosto", 1: "janeiro", 4: "abril", 6: "junho", 9: "setembro"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "sáb", 0: "dom", 1: "seg", 2: "ter", 3: "qua", 4: "qui", 5: "sex"}, Narrow: ut.CalendarDayFormatNameValue{3: "Q", 4: "Q", 5: "S", 6: "S", 0: "D", 1: "S", 2: "T"}, Short: ut.CalendarDayFormatNameValue{1: "seg", 2: "ter", 3: "qua", 4: "qui", 5: "sex", 6: "sáb", 0: "dom"}, Wide: ut.CalendarDayFormatNameValue{2: "terça-feira", 3: "quarta-feira", 4: "quinta-feira", 5: "sexta-feira", 6: "sábado", 0: "domingo", 1: "segunda-feira"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "manhã", "afternoon1": "tarde", "evening1": "noite", "night1": "madrugada", "midnight": "meia-noite", "am": "AM", "noon": "meio-dia", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "madrugada", "midnight": "meia-noite", "am": "AM", "noon": "meio-dia", "pm": "PM", "morning1": "manhã", "afternoon1": "tarde", "evening1": "noite"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"evening1": "noite", "night1": "madrugada", "midnight": "meia-noite", "am": "AM", "noon": "meio-dia", "pm": "PM", "morning1": "manhã", "afternoon1": "tarde"}}}}
