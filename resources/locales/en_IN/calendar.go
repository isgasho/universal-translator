package en_IN

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "MMMM d, y", Medium: "dd-MMM-y", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{2: "Feb", 9: "Sep", 12: "Dec", 1: "Jan", 5: "May", 10: "Oct", 4: "Apr", 6: "Jun", 7: "Jul", 8: "Aug", 11: "Nov", 3: "Mar"}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 7: "J", 1: "J", 5: "M", 9: "S", 3: "M", 2: "F", 11: "N", 10: "O", 6: "J", 8: "A", 12: "D"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{12: "December", 2: "February", 10: "October", 8: "August", 7: "July", 11: "November", 3: "March", 5: "May", 9: "September", 1: "January", 4: "April", 6: "June"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat"}, Narrow: ut.CalendarDayFormatNameValue{6: "S", 0: "S", 1: "M", 2: "T", 3: "W", 4: "T", 5: "F"}, Short: ut.CalendarDayFormatNameValue{4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo", 2: "Tu", 3: "We"}, Wide: ut.CalendarDayFormatNameValue{0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n", "pm": "p"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
