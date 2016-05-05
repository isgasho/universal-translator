package en_GG

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "Apr", 6: "Jun", 11: "Nov", 5: "May", 9: "Sep", 10: "Oct", 12: "Dec", 2: "Feb", 7: "Jul", 8: "Aug", 1: "Jan", 3: "Mar"}, Narrow: ut.CalendarMonthFormatNameValue{11: "N", 5: "M", 6: "J", 9: "S", 1: "J", 3: "M", 4: "A", 7: "J", 8: "A", 10: "O", 12: "D", 2: "F"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{2: "February", 3: "March", 5: "May", 8: "August", 1: "January", 4: "April", 11: "November", 12: "December", 10: "October", 9: "September", 6: "June", 7: "July"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{0: "Sun", 1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat"}, Narrow: ut.CalendarDayFormatNameValue{0: "S", 1: "M", 2: "T", 3: "W", 4: "T", 5: "F", 6: "S"}, Short: ut.CalendarDayFormatNameValue{0: "Su", 1: "Mo", 2: "Tu", 3: "We", 4: "Th", 5: "Fr", 6: "Sa"}, Wide: ut.CalendarDayFormatNameValue{6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n", "pm": "p"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
