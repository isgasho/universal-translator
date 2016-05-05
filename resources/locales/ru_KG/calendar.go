package ru_KG

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y 'г'.", Long: "d MMMM y 'г'.", Medium: "d MMM y 'г'.", Short: "dd.MM.yy"}, Time: ut.CalendarDateFormat{Full: "H:mm:ss zzzz", Long: "H:mm:ss z", Medium: "H:mm:ss", Short: "H:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1}, {0}", Long: "{1}, {0}", Medium: "{1}, {0}", Short: "{1}, {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{9: "сент.", 4: "апр.", 11: "нояб.", 3: "март", 5: "май", 7: "июль", 10: "окт.", 1: "янв.", 6: "июнь", 8: "авг.", 2: "февр.", 12: "дек."}, Narrow: ut.CalendarMonthFormatNameValue{1: "Я", 2: "Ф", 8: "А", 11: "Н", 6: "И", 3: "М", 9: "С", 10: "О", 7: "И", 4: "А", 5: "М", 12: "Д"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{12: "декабрь", 1: "январь", 9: "сентябрь", 11: "ноябрь", 8: "август", 3: "март", 5: "май", 6: "июнь", 7: "июль", 10: "октябрь", 2: "февраль", 4: "апрель"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "сб", 0: "вс", 1: "пн", 2: "вт", 3: "ср", 4: "чт", 5: "пт"}, Narrow: ut.CalendarDayFormatNameValue{3: "С", 4: "Ч", 5: "П", 6: "С", 0: "В", 1: "П", 2: "В"}, Short: ut.CalendarDayFormatNameValue{1: "пн", 2: "вт", 3: "ср", 4: "чт", 5: "пт", 6: "сб", 0: "вс"}, Wide: ut.CalendarDayFormatNameValue{2: "вторник", 3: "среда", 4: "четверг", 5: "пятница", 6: "суббота", 0: "воскресенье", 1: "понедельник"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"morning1": "утро", "afternoon1": "день", "evening1": "веч.", "night1": "ночь", "midnight": "полн.", "am": "ДП", "noon": "полд.", "pm": "ПП"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "полн.", "am": "ДП", "noon": "полд.", "pm": "ПП", "morning1": "утро", "afternoon1": "день", "evening1": "веч.", "night1": "ночь"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"night1": "ночь", "midnight": "полночь", "am": "ДП", "noon": "полдень", "pm": "ПП", "morning1": "утро", "afternoon1": "день", "evening1": "вечер"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
