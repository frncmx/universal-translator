package lt

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "y 'm'. MMMM d 'd'., EEEE", Long: "y 'm'. MMMM d 'd'.", Medium: "y-MM-dd", Short: "y-MM-dd"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "rugs.", 5: "geg.", 7: "liep.", 3: "kov.", 4: "bal.", 6: "birž.", 8: "rugp.", 10: "spal.", 11: "lapkr.", 1: "saus.", 2: "vas.", 12: "gruod."}, Narrow: ut.CalendarMonthFormatNameValue{4: "B", 6: "B", 12: "G", 1: "S", 2: "V", 3: "K", 9: "R", 10: "S", 11: "L", 5: "G", 7: "L", 8: "R"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "vasaris", 3: "kovas", 4: "balandis", 6: "birželis", 8: "rugpjūtis", 9: "rugsėjis", 12: "gruodis", 1: "sausis", 5: "gegužė", 7: "liepa", 10: "spalis", 11: "lapkritis"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "pn", 6: "št", 0: "sk", 1: "pr", 2: "an", 3: "tr", 4: "kt"}, Narrow: ut.CalendarDayFormatNameValue{2: "A", 3: "T", 4: "K", 5: "P", 6: "Š", 0: "S", 1: "P"}, Short: ut.CalendarDayFormatNameValue{4: "Kt", 5: "Pn", 6: "Št", 0: "Sk", 1: "Pr", 2: "An", 3: "Tr"}, Wide: ut.CalendarDayFormatNameValue{0: "sekmadienis", 1: "pirmadienis", 2: "antradienis", 3: "trečiadienis", 4: "ketvirtadienis", 5: "penktadienis", 6: "šeštadienis"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"night1": "naktis", "midnight": "vidurnaktis", "am": "priešpiet", "noon": "vidurdienis", "pm": "popiet", "morning1": "rytas", "afternoon1": "diena", "evening1": "vakaras"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "pr. p.", "noon": "vidurdienis", "pm": "pop.", "morning1": "rytas", "afternoon1": "diena", "evening1": "vakaras", "night1": "naktis", "midnight": "vidurnaktis"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"noon": "vidurdienis", "pm": "popiet", "morning1": "rytas", "afternoon1": "diena", "evening1": "vakaras", "night1": "naktis", "midnight": "vidurnaktis", "am": "priešpiet"}}}}
