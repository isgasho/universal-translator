package ce

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, Time: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "окт", 11: "ноя", 12: "дек", 6: "июн", 7: "июл", 9: "сен", 4: "апр", 5: "май", 8: "авг", 1: "янв", 2: "фев", 3: "мар"}, Narrow: ut.CalendarMonthFormatNameValue(nil), Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "октябрь", 1: "январь", 2: "февраль", 3: "март", 4: "апрель", 5: "май", 7: "июль", 9: "сентябрь", 6: "июнь", 8: "август", 11: "ноябрь", 12: "декабрь"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue(nil), Narrow: ut.CalendarDayFormatNameValue(nil), Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{3: "кхаарин де", 4: "еарин де", 5: "пӀераскан де", 6: "шот де", 0: "кӀиранан де", 1: "оршотан де", 2: "шинарин де"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue(nil), Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue(nil)}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
