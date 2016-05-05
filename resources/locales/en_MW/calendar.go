package en_MW

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, MMMM d, y", Long: "MMMM d, y", Medium: "MMM d, y", Short: "M/d/yy"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'at' {0}", Long: "{1} 'at' {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{8: "Aug", 3: "Mar", 5: "May", 12: "Dec", 4: "Apr", 6: "Jun", 7: "Jul", 11: "Nov", 1: "Jan", 9: "Sep", 10: "Oct", 2: "Feb"}, Narrow: ut.CalendarMonthFormatNameValue{8: "A", 9: "S", 10: "O", 12: "D", 6: "J", 4: "A", 7: "J", 11: "N", 1: "J", 5: "M", 2: "F", 3: "M"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{7: "July", 11: "November", 5: "May", 8: "August", 1: "January", 4: "April", 6: "June", 10: "October", 9: "September", 12: "December", 2: "February", 3: "March"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "Mon", 2: "Tue", 3: "Wed", 4: "Thu", 5: "Fri", 6: "Sat", 0: "Sun"}, Narrow: ut.CalendarDayFormatNameValue{3: "W", 4: "T", 5: "F", 6: "S", 0: "S", 1: "M", 2: "T"}, Short: ut.CalendarDayFormatNameValue{6: "Sa", 0: "Su", 1: "Mo", 2: "Tu", 3: "We", 4: "Th", 5: "Fr"}, Wide: ut.CalendarDayFormatNameValue{4: "Thursday", 5: "Friday", 6: "Saturday", 0: "Sunday", 1: "Monday", 2: "Tuesday", 3: "Wednesday"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "in the morning", "afternoon1": "in the afternoon", "evening1": "in the evening", "night1": "at night"}, Narrow: ut.CalendarPeriodFormatNameValue{"evening1": "in the evening", "night1": "at night", "midnight": "mi", "am": "a", "noon": "n", "pm": "p", "morning1": "in the morning", "afternoon1": "in the afternoon"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "night", "midnight": "midnight", "am": "AM", "noon": "noon", "pm": "PM", "morning1": "morning", "afternoon1": "afternoon", "evening1": "evening"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
