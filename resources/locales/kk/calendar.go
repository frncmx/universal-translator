package kk

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "y 'ж'. d MMMM, EEEE", Long: "y 'ж'. d MMMM", Medium: "y 'ж'. dd MMM", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "Ақп.", 3: "Нау.", 4: "Сәу.", 5: "Мам.", 6: "Мау.", 7: "Шіл.", 8: "Там.", 1: "Қаң.", 11: "Қар.", 10: "Қаз.", 12: "Жел.", 9: "Қыр."}, Narrow: ut.CalendarMonthFormatNameValue{6: "М", 10: "Қ", 11: "Қ", 1: "Қ", 3: "Н", 5: "М", 8: "Т", 9: "Қ", 12: "Ж", 2: "А", 4: "С", 7: "Ш"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{7: "Шілде", 8: "Тамыз", 9: "Қыркүйек", 11: "Қараша", 12: "Желтоқсан", 3: "Наурыз", 2: "Ақпан", 4: "Сәуір", 5: "Мамыр", 6: "Маусым", 10: "Қазан", 1: "Қаңтар"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Дс", 2: "Сс", 3: "Ср", 4: "Бс", 5: "Жм", 6: "Сб", 0: "Жс"}, Narrow: ut.CalendarDayFormatNameValue{5: "Ж", 6: "С", 0: "Ж", 1: "Д", 2: "С", 3: "С", 4: "Б"}, Short: ut.CalendarDayFormatNameValue{2: "Сс", 3: "Ср", 4: "Бс", 5: "Жм", 6: "Сб", 0: "Жс", 1: "Дс"}, Wide: ut.CalendarDayFormatNameValue{3: "Сәрсенбі", 4: "Бейсенбі", 5: "Жұма", 6: "Сенбі", 0: "Жексенбі", 1: "Дүйсенбі", 2: "Сейсенбі"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "кеш", "night1": "түн", "midnight": "түн жарымы", "am": "таң", "noon": "талтүс", "pm": "түс/кеш", "morning1": "таң", "afternoon1": "түстен кейін"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "түс/кеш", "morning1": "таң", "afternoon1": "түстен кейін", "evening1": "кеш", "night1": "түн", "midnight": "түн жарымы", "am": "таң", "noon": "талтүс"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"evening1": "кеш", "night1": "түн", "midnight": "түн жарымы", "am": "таң", "noon": "талтүс", "pm": "түс/кеш", "morning1": "таң", "afternoon1": "түстен кейін"}}}}
