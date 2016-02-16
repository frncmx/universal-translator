package am

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE ፣d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "ኦገስ", 9: "ሴፕቴ", 10: "ኦክቶ", 12: "ዲሴም", 3: "ማርች", 4: "ኤፕሪ", 5: "ሜይ", 6: "ጁን", 7: "ጁላይ", 11: "ኖቬም", 1: "ጃንዩ", 2: "ፌብሩ"}, Narrow: ut.CalendarMonthFormatNameValue{9: "ሴ", 11: "ኖ", 2: "ፌ", 4: "ኤ", 6: "ጁ", 7: "ጁ", 10: "ኦ", 12: "ዲ", 1: "ጃ", 3: "ማ", 5: "ሜ", 8: "ኦ"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "ፌብሩወሪ", 4: "ኤፕሪል", 9: "ሴፕቴምበር", 11: "ኖቬምበር", 12: "ዲሴምበር", 1: "ጃንዩወሪ", 3: "ማርች", 5: "ሜይ", 6: "ጁን", 7: "ጁላይ", 8: "ኦገስት", 10: "ኦክቶበር"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "እሑድ", 1: "ሰኞ", 2: "ማክሰ", 3: "ረቡዕ", 4: "ሐሙስ", 5: "ዓርብ", 6: "ቅዳሜ"}, Narrow: ut.CalendarDayFormatNameValue{0: "እ", 1: "ሰ", 2: "ማ", 3: "ረ", 4: "ሐ", 5: "ዓ", 6: "ቅ"}, Short: ut.CalendarDayFormatNameValue{0: "እ", 1: "ሰ", 2: "ማ", 3: "ረ", 4: "ሐ", 5: "ዓ", 6: "ቅ"}, Wide: ut.CalendarDayFormatNameValue{0: "እሑድ", 1: "ሰኞ", 2: "ማክሰኞ", 3: "ረቡዕ", 4: "ሐሙስ", 5: "ዓርብ", 6: "ቅዳሜ"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "እኩለ ሌሊት", "am": "ጥዋት", "noon": "ቀትር", "pm": "ከሰዓት", "morning1": "ጥዋት1", "afternoon1": "ከሰዓት በኋላ", "evening1": "ማታ", "night1": "ሌሊት"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "እኩለ ሌሊት", "am": "ጠ", "noon": "ቀትር", "pm": "ከ", "morning1": "ጥዋት", "afternoon1": "ከሰዓት በኋላ", "evening1": "ማታ", "night1": "ሌሊት"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"night1": "ሌሊት", "midnight": "እኩለ ሌሊት", "am": "ጥዋት", "noon": "ቀትር", "pm": "ከሰዓት", "morning1": "ጥዋት1", "afternoon1": "ከሰዓት በኋላ", "evening1": "ማታ"}}}}
