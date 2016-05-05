package uz_Cyrl

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, y MMMM dd", Long: "y MMMM d", Medium: "y MMM d", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "Янв", 3: "Мар", 4: "Апр", 6: "Июн", 10: "Окт", 12: "Дек", 2: "Фев", 5: "Май", 7: "Июл", 8: "Авг", 9: "Сен", 11: "Ноя"}, Narrow: ut.CalendarMonthFormatNameValue{12: "Д", 1: "Я", 4: "А", 5: "М", 6: "И", 9: "С", 11: "Н", 2: "Ф", 3: "М", 7: "И", 8: "А", 10: "О"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{3: "Март", 4: "Апрел", 9: "Сентябр", 1: "Январ", 2: "Феврал", 7: "Июл", 8: "Август", 10: "Октябр", 11: "Ноябр", 12: "Декабр", 5: "Май", 6: "Июн"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "Сеш", 3: "Чор", 4: "Пай", 5: "Жум", 6: "Шан", 0: "Якш", 1: "Душ"}, Narrow: ut.CalendarDayFormatNameValue{4: "П", 5: "Ж", 6: "Ш", 0: "Я", 1: "Д", 2: "С", 3: "Ч"}, Short: ut.CalendarDayFormatNameValue{0: "Якш", 1: "Душ", 2: "Сеш", 3: "Чор", 4: "Пай", 5: "Жум", 6: "Шан"}, Wide: ut.CalendarDayFormatNameValue{6: "шанба", 0: "якшанба", 1: "душанба", 2: "сешанба", 3: "чоршанба", 4: "пайшанба", 5: "жума"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"noon": "tush payti", "pm": "TK", "morning1": "ertalab", "afternoon1": "kunduzi", "evening1": "kechqurun", "night1": "kechasi", "midnight": "yarim kechasi", "am": "TO"}, Narrow: ut.CalendarPeriodFormatNameValue{"night1": "kechasi", "midnight": "yarim kechasi", "am": "TO", "noon": "tush payti", "pm": "TK", "morning1": "ertalab", "afternoon1": "kunduzi", "evening1": "kechqurun"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"am": "TO", "noon": "tush payti", "pm": "TK", "morning1": "ertalab", "afternoon1": "kunduzi", "evening1": "kechqurun", "night1": "kechasi", "midnight": "yarim kechasi"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "М.А.", Narrow: ""}}}}
