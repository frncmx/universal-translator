package sw_UG

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "Nov", 6: "Jun", 10: "Okt", 1: "Jan", 2: "Feb", 8: "Ago", 12: "Des", 4: "Apr", 7: "Jul", 5: "Mei", 3: "Mac", 9: "Sep"}, Narrow: ut.CalendarMonthFormatNameValue{8: "A", 5: "M", 10: "O", 11: "N", 4: "A", 6: "J", 7: "J", 12: "D", 1: "J", 2: "F", 3: "M", 9: "S"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{11: "Novemba", 1: "Januari", 10: "Oktoba", 2: "Februari", 4: "Aprili", 6: "Juni", 7: "Julai", 9: "Septemba", 3: "Machi", 5: "Mei", 8: "Agosti", 12: "Desemba"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Jumapili", 1: "Jumatatu", 2: "Jumanne", 3: "Jumatano", 4: "Alhamisi", 5: "Ijumaa", 6: "Jumamosi"}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "S", 1: "M", 2: "T", 3: "W", 4: "T", 5: "F"}, Short: ut.CalendarDayFormatNameValue{3: "Jumatano", 4: "Alhamisi", 5: "Ijumaa", 6: "Jumamosi", 0: "Jumapili", 1: "Jumatatu", 2: "Jumanne"}, Wide: ut.CalendarDayFormatNameValue{0: "Jumapili", 1: "Jumatatu", 2: "Jumanne", 3: "Jumatano", 4: "Alhamisi", 5: "Ijumaa", 6: "Jumamosi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "morning1": "alfajiri", "afternoon1": "mchana", "evening1": "jioni", "night1": "usiku", "midnight": "saa sita za usiku", "noon": "saa sita za mchana", "pm": "PM", "morning2": "asubuhi"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "am", "morning1": "alfajiri", "noon": "saa sita za mchana", "pm": "pm", "morning2": "asubuhi", "midnight": "saa sita za usiku", "afternoon1": "mchana", "evening1": "jioni", "night1": "usiku"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "alfajiri", "morning2": "asubuhi", "afternoon1": "mchana", "midnight": "saa sita za usiku", "night1": "usiku", "am": "AM", "pm": "PM", "evening1": "jioni", "noon": "saa sita za mchana"}}}}
