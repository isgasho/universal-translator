package en_SC

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "Oct", 2: "Feb", 4: "Apr", 6: "Jun", 7: "Jul", 3: "Mar", 5: "May", 9: "Sep", 12: "Dec", 1: "Jan", 8: "Aug", 11: "Nov"}, Narrow: ut.CalendarMonthFormatNameValue{7: "J", 10: "O", 1: "J", 5: "M", 6: "J", 2: "F", 4: "A", 9: "S", 8: "A", 11: "N", 12: "D", 3: "M"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{12: "December", 4: "April", 6: "June", 7: "July", 2: "February", 5: "May", 10: "October", 11: "November", 3: "March", 1: "January", 8: "August", 9: "September"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun", 1: "Mon"}, Narrow: ut.CalendarDayFormatNameValue{4: "T", 5: "F", 6: "S", 0: "S", 1: "M", 2: "T", 3: "W"}, Short: ut.CalendarDayFormatNameValue{3: "We", 4: "Th", 5: "Fr", 6: "Sa", 0: "Su", 1: "Mo", 2: "Tu"}, Wide: ut.CalendarDayFormatNameValue{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM"}, Narrow: ut.CalendarPeriodFormatNameValue{"pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"afternoon1": "afternoon", "evening1": "evening", "night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
