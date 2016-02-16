package fr_HT

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "déc.", 3: "mars", 5: "mai", 6: "juin", 7: "juil.", 10: "oct.", 1: "janv.", 2: "févr.", 4: "avr.", 8: "août", 11: "nov.", 9: "sept."}, Narrow: ut.CalendarMonthFormatNameValue{6: "J", 12: "D", 2: "F", 8: "A", 3: "M", 4: "A", 11: "N", 9: "S", 10: "O", 1: "J", 5: "M", 7: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "mai", 12: "décembre", 11: "novembre", 6: "juin", 7: "juillet", 2: "février", 8: "août", 10: "octobre", 1: "janvier", 3: "mars", 4: "avril", 9: "septembre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "ven.", 6: "sam.", 0: "dim.", 1: "lun.", 2: "mar.", 3: "mer.", 4: "jeu."}, Narrow: ut.CalendarDayFormatNameValue{2: "M", 3: "M", 4: "J", 5: "V", 6: "S", 0: "D", 1: "L"}, Short: ut.CalendarDayFormatNameValue{2: "ma", 3: "me", 4: "je", 5: "ve", 6: "sa", 0: "di", 1: "lu"}, Wide: ut.CalendarDayFormatNameValue{4: "jeudi", 5: "vendredi", 6: "samedi", 0: "dimanche", 1: "lundi", 2: "mardi", 3: "mercredi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"evening1": "soir", "night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "matin", "afternoon1": "après-midi"}}}}
