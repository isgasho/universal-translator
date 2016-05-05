package fr_KM

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{10: "oct.", 1: "janv.", 2: "févr.", 5: "mai", 6: "juin", 7: "juil.", 4: "avr.", 8: "août", 12: "déc.", 9: "sept.", 3: "mars", 11: "nov."}, Narrow: ut.CalendarMonthFormatNameValue{5: "M", 6: "J", 8: "A", 9: "S", 11: "N", 1: "J", 3: "M", 2: "F", 7: "J", 12: "D", 4: "A", 10: "O"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{5: "mai", 8: "août", 1: "janvier", 3: "mars", 6: "juin", 10: "octobre", 11: "novembre", 12: "décembre", 7: "juillet", 2: "février", 4: "avril", 9: "septembre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "lun.", 2: "mar.", 3: "mer.", 4: "jeu.", 5: "ven.", 6: "sam.", 0: "dim."}, Narrow: ut.CalendarDayFormatNameValue{1: "L", 2: "M", 3: "M", 4: "J", 5: "V", 6: "S", 0: "D"}, Short: ut.CalendarDayFormatNameValue{0: "di", 1: "lu", 2: "ma", 3: "me", 4: "je", 5: "ve", 6: "sa"}, Wide: ut.CalendarDayFormatNameValue{1: "lundi", 2: "mardi", 3: "mercredi", 4: "jeudi", 5: "vendredi", 6: "samedi", 0: "dimanche"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat."}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "matin", "afternoon1": "après-midi", "evening1": "soir", "night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
