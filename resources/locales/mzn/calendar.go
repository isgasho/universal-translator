package mzn

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, Time: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "آوریل", 6: "ژوئن", 11: "نوامبر", 9: "سپتامبر", 10: "اکتبر", 1: "ژانویه", 2: "فوریه", 3: "مارس", 5: "مه", 7: "ژوئیه", 8: "اوت", 12: "دسامبر"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "آوریل", 5: "مه", 6: "ژوئن", 7: "ژوئیه", 8: "اوت", 10: "اکتبر", 1: "ژانویه", 2: "فوریه", 11: "نوامبر", 12: "دسامبر", 3: "مارس", 9: "سپتامبر"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue(nil), Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue(nil)}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "بعد میلاد", Abbrev: "م.", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "قبل میلاد", Abbrev: "پ.م", Narrow: ""}}}}
