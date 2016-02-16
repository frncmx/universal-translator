package ta_SG

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM, y", Long: "d MMMM, y", Medium: "d MMM, y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "a h:mm:ss zzzz", Long: "a h:mm:ss z", Medium: "a h:mm:ss", Short: "a h:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} ’அன்று’ {0}", Long: "{1} ’அன்று’ {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "ஜன.", 5: "மே", 11: "நவ.", 2: "பிப்.", 6: "ஜூன்", 8: "ஆக.", 9: "செப்.", 7: "ஜூலை", 10: "அக்.", 12: "டிச.", 3: "மார்.", 4: "ஏப்."}, Narrow: ut.CalendarMonthFormatNameValue{3: "மா", 8: "ஆ", 11: "ந", 1: "ஜ", 4: "ஏ", 5: "மே", 6: "ஜூ", 7: "ஜூ", 12: "டி", 10: "அ", 2: "பி", 9: "செ"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{6: "ஜூன்", 8: "ஆகஸ்டு", 11: "நவம்பர்", 5: "மே", 7: "ஜூலை", 3: "மார்ச்", 10: "அக்டோபர்", 1: "ஜனவரி", 2: "பிப்ரவரி", 9: "செப்டம்பர்", 12: "டிசம்பர்", 4: "ஏப்ரல்"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "ஞாயி.", 1: "திங்.", 2: "செவ்.", 3: "புத.", 4: "வியா.", 5: "வெள்.", 6: "சனி"}, Narrow: ut.CalendarDayFormatNameValue{5: "வெ", 6: "ச", 0: "ஞா", 1: "தி", 2: "செ", 3: "பு", 4: "வி"}, Short: ut.CalendarDayFormatNameValue{3: "பு", 4: "வி", 5: "வெ", 6: "ச", 0: "ஞா", 1: "தி", 2: "செ"}, Wide: ut.CalendarDayFormatNameValue{3: "புதன்", 4: "வியாழன்", 5: "வெள்ளி", 6: "சனி", 0: "ஞாயிறு", 1: "திங்கள்", 2: "செவ்வாய்"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "முற்பகல்", "pm": "பிற்பகல்", "midnight": "நள்ளிரவு", "morning2": "காலை", "afternoon2": "பிற்பகல்", "night1": "இரவு", "evening2": "அந்தி மாலை", "noon": "நண்பகல்", "morning1": "அதிகாலை", "afternoon1": "மதியம்", "evening1": "மாலை"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "பி.ப", "morning2": "கா.", "evening1": "மா.", "noon": "நண்.", "evening2": "அந்தி மா.", "am": "மு.ப", "morning1": "அதி.", "afternoon1": "மதி.", "afternoon2": "பிற்.", "midnight": "நள்.", "night1": "இ."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon2": "பிற்பகல்", "night1": "இரவு", "midnight": "நள்ளிரவு", "afternoon1": "மதியம்", "morning1": "அதிகாலை", "am": "முற்பகல்", "morning2": "காலை", "evening1": "மாலை", "evening2": "அந்தி மாலை", "noon": "நண்பகல்", "pm": "பிற்பகல்"}}}}
