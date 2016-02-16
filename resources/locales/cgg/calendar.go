package cgg

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "KST", 4: "KKN", 6: "KMK", 8: "KMN", 11: "KNK", 12: "KNB", 1: "KBZ", 2: "KBR", 5: "KTN", 7: "KMS", 9: "KMW", 10: "KKM"}, Narrow: ut.CalendarMonthFormatNameValue{6: "J", 8: "A", 11: "N", 12: "D", 2: "F", 4: "A", 5: "M", 9: "S", 10: "O", 1: "J", 3: "M", 7: "J"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{2: "Okwakabiri", 3: "Okwakashatu", 4: "Okwakana", 5: "Okwakataana", 6: "Okwamukaaga", 1: "Okwokubanza", 7: "Okwamushanju", 8: "Okwamunaana", 9: "Okwamwenda", 10: "Okwaikumi", 11: "Okwaikumi na kumwe", 12: "Okwaikumi na ibiri"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "OKN", 5: "OKT", 6: "OMK", 0: "SAN", 1: "ORK", 2: "OKB", 3: "OKS"}, Narrow: ut.CalendarDayFormatNameValue{0: "S", 1: "K", 2: "R", 3: "S", 4: "N", 5: "T", 6: "M"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{6: "Orwamukaaga", 0: "Sande", 1: "Orwokubanza", 2: "Orwakabiri", 3: "Orwakashatu", 4: "Orwakana", 5: "Orwakataano"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}}}
