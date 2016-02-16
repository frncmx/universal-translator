package brx

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue(nil), Narrow: ut.CalendarMonthFormatNameValue{8: "आ", 9: "से", 11: "न", 12: "दि", 1: "ज", 4: "ए", 6: "जु", 7: "जु", 10: "अ", 2: "फे", 3: "मा", 5: "मे"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{5: "मे", 7: "जुलाइ", 8: "आगस्थ", 9: "सेबथेज्ब़र", 10: "अखथबर", 4: "एफ्रिल", 2: "फेब्रुवारी", 3: "मार्स", 6: "जुन", 11: "नबेज्ब़र", 12: "दिसेज्ब़र", 1: "जानुवारी"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "सम", 2: "मंगल", 3: "बुद", 4: "बिसथि", 5: "सुखुर", 6: "सुनि", 0: "रबि"}, Narrow: ut.CalendarDayFormatNameValue{2: "मं", 3: "बु", 4: "बि", 5: "सु", 6: "सु", 0: "र", 1: "स"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{4: "बिसथिबार", 5: "सुखुरबार", 6: "सुनिबार", 0: "रबिबार", 1: "समबार", 2: "मंगलबार", 3: "बुदबार"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "फुं", "pm": "बेलासे"}}}}
