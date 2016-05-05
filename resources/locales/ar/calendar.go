package ar

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE، d MMMM، y", Long: "d MMMM، y", Medium: "dd\u200f/MM\u200f/y", Short: "d\u200f/M\u200f/y"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{1: "يناير", 3: "مارس", 4: "أبريل", 7: "يوليو", 8: "أغسطس", 12: "ديسمبر", 2: "فبراير", 5: "مايو", 6: "يونيو", 9: "سبتمبر", 10: "أكتوبر", 11: "نوفمبر"}, Narrow: ut.CalendarMonthFormatNameValue{6: "ن", 7: "ل", 10: "ك", 11: "ب", 12: "د", 1: "ي", 2: "ف", 3: "م", 4: "أ", 5: "و", 8: "غ", 9: "س"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{1: "يناير", 5: "مايو", 6: "يونيو", 8: "أغسطس", 10: "أكتوبر", 12: "ديسمبر", 2: "فبراير", 3: "مارس", 4: "أبريل", 7: "يوليو", 9: "سبتمبر", 11: "نوفمبر"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين"}, Narrow: ut.CalendarDayFormatNameValue{3: "ر", 4: "خ", 5: "ج", 6: "س", 0: "ح", 1: "ن", 2: "ث"}, Short: ut.CalendarDayFormatNameValue{6: "السبت", 0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء", 4: "الخميس", 5: "الجمعة"}, Wide: ut.CalendarDayFormatNameValue{4: "الخميس", 5: "الجمعة", 6: "السبت", 0: "الأحد", 1: "الاثنين", 2: "الثلاثاء", 3: "الأربعاء"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"night1": "منتصف الليل", "night2": "ليلاً", "pm": "م", "afternoon2": "بعد الظهر", "evening1": "مساءً", "afternoon1": "ظهرًا", "am": "ص", "morning1": "فجرا", "morning2": "صباحًا"}, Narrow: ut.CalendarPeriodFormatNameValue{"am": "ص", "morning2": "صباحًا", "afternoon2": "بعد الظهر", "evening1": "مساءً", "night1": "منتصف الليل", "night2": "ليلاً", "pm": "م", "morning1": "فجرا", "afternoon1": "ظهرًا"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "صباحًا", "pm": "مساءً", "morning1": "فجرا", "morning2": "صباحًا", "afternoon1": "ظهرًا", "afternoon2": "بعد الظهر", "evening1": "مساءً", "night1": "منتصف الليل", "night2": "ليلاً"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "بعد الميلاد", Abbrev: "ب.م", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "قبل الميلاد", Abbrev: "ق.م", Narrow: ""}}}}
