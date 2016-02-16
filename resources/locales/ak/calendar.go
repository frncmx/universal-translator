package ak

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "yy/MM/dd"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "D-Ɔ", 9: "F-Ɛ", 12: "M-Ɔ", 1: "S-Ɔ", 3: "E-Ɔ", 4: "E-O", 6: "O-A", 7: "A-K", 2: "K-Ɔ", 5: "E-K", 10: "Ɔ-A", 11: "Ɔ-O"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{11: "Ɔberɛfɛw-Obubuo", 1: "Sanda-Ɔpɛpɔn", 7: "Ayɛwoho-Kitawonsa", 8: "Difuu-Ɔsandaa", 5: "Esusow Aketseaba-Kɔtɔnimba", 6: "Obirade-Ayɛwohomumu", 9: "Fankwa-Ɛbɔ", 10: "Ɔbɛsɛ-Ahinime", 12: "Mumu-Ɔpɛnimba", 2: "Kwakwar-Ɔgyefuo", 3: "Ebɔw-Ɔbenem", 4: "Ebɔbira-Oforisuo"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{3: "Wuk", 4: "Yaw", 5: "Fia", 6: "Mem", 0: "Kwe", 1: "Dwo", 2: "Ben"}, Narrow: ut.CalendarDayFormatNameValue{1: "D", 2: "B", 3: "W", 4: "Y", 5: "F", 6: "M", 0: "K"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{5: "Fida", 6: "Memeneda", 0: "Kwesida", 1: "Dwowda", 2: "Benada", 3: "Wukuda", 4: "Yawda"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "AN", "pm": "EW"}}}}
