package en_NF

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Jan", 4: "Apr", 6: "Jun", 10: "Oct", 12: "Dec", 9: "Sep", 2: "Feb", 3: "Mar", 5: "May", 8: "Aug", 11: "Nov", 7: "Jul"}, Narrow: ut.CalendarMonthFormatNameValue{10: "O", 6: "J", 1: "J", 3: "M", 8: "A", 2: "F", 11: "N", 4: "A", 5: "M", 7: "J", 9: "S", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{8: "August", 12: "December", 11: "November", 3: "March", 4: "April", 7: "July", 9: "September", 6: "June", 10: "October", 1: "January", 2: "February", 5: "May"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed"}, Narrow: ut.CalendarDayFormatNameValue{0: "S", 1: "M", 2: "T", 3: "W", 4: "T", 5: "F", 6: "S"}, Short: ut.CalendarDayFormatNameValue{1: "Mo", 2: "Tu", 3: "We", 4: "Th", 5: "Fr", 6: "Sa", 0: "Su"}, Wide: ut.CalendarDayFormatNameValue{6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "a", "noon": "n", "pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight", "am": "AM"}}}}
