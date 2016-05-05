package ug

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE، MMMM d، y", Long: "MMMM d، y", Medium: "MMM d، y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1}، {0}", Short: "{1}، {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{11: "نويابىر", 12: "دېكابىر", 3: "مارت", 9: "سېنتەبىر", 4: "ئاپرېل", 5: "ماي", 6: "ئىيۇن", 7: "ئىيۇل", 8: "ئاۋغۇست", 10: "ئۆكتەبىر", 1: "يانۋار", 2: "فېۋرال"}, Narrow: ut.CalendarMonthFormatNameValue{4: "4", 7: "7", 8: "8", 9: "9", 2: "2", 3: "3", 6: "6", 10: "10", 11: "11", 12: "12", 1: "1", 5: "5"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{8: "ئاۋغۇست", 10: "ئۆكتەبىر", 1: "يانۋار", 2: "فېۋرال", 4: "ئاپرېل", 5: "ماي", 7: "ئىيۇل", 3: "مارت", 6: "ئىيۇن", 9: "سېنتەبىر", 11: "نويابىر", 12: "دېكابىر"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "پە", 5: "جۈ", 6: "شە", 0: "يە", 1: "دۈ", 2: "سە", 3: "چا"}, Narrow: ut.CalendarDayFormatNameValue{0: "ي", 1: "د", 2: "س", 3: "چ", 4: "پ", 5: "ج", 6: "ش"}, Short: ut.CalendarDayFormatNameValue{6: "ش", 0: "ي", 1: "د", 2: "س", 3: "چ", 4: "پ", 5: "ج"}, Wide: ut.CalendarDayFormatNameValue{4: "پەيشەنبە", 5: "جۈمە", 6: "شەنبە", 0: "يەكشەنبە", 1: "دۈشەنبە", 2: "سەيشەنبە", 3: "چارشەنبە"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "چۈشتىن كېيىن", "am": "چۈشتىن بۇرۇن"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "چۈشتىن بۇرۇن", "pm": "چۈشتىن كېيىن"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "مىلادىيەدىن بۇرۇن", Abbrev: "BCE", Narrow: "BCE"}}}}
