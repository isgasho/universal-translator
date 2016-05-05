package fr_DZ

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "h:mm:ss a zzzz", Long: "h:mm:ss a z", Medium: "h:mm:ss a", Short: "h:mm a"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'à' {0}", Long: "{1} 'à' {0}", Medium: "{1} 'à' {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "déc.", 4: "avr.", 9: "sept.", 1: "janv.", 5: "mai", 11: "nov.", 3: "mars", 8: "août", 10: "oct.", 7: "juil.", 2: "févr.", 6: "juin"}, Narrow: ut.CalendarMonthFormatNameValue{9: "S", 2: "F", 5: "M", 7: "J", 4: "A", 3: "M", 6: "J", 8: "A", 11: "N", 12: "D", 10: "O", 1: "J"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{7: "juillet", 10: "octobre", 4: "avril", 12: "décembre", 3: "mars", 6: "juin", 11: "novembre", 1: "janvier", 2: "février", 5: "mai", 8: "août", 9: "septembre"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "ven.", 6: "sam.", 0: "dim.", 1: "lun.", 2: "mar.", 3: "mer.", 4: "jeu."}, Narrow: ut.CalendarDayFormatNameValue{5: "V", 6: "S", 0: "D", 1: "L", 2: "M", 3: "M", 4: "J"}, Short: ut.CalendarDayFormatNameValue{4: "je", 5: "ve", 6: "sa", 0: "di", 1: "lu", 2: "ma", 3: "me"}, Wide: ut.CalendarDayFormatNameValue{2: "mardi", 3: "mercredi", 4: "jeudi", 5: "vendredi", 6: "samedi", 0: "dimanche", 1: "lundi"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "mat.", "afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi"}, Narrow: ut.CalendarPeriodFormatNameValue{"afternoon1": "ap.m.", "evening1": "soir", "night1": "nuit", "midnight": "min.", "am": "AM", "noon": "midi", "pm": "PM", "morning1": "mat."}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"morning1": "matin", "afternoon1": "après-midi", "evening1": "soir", "night1": "nuit", "midnight": "minuit", "am": "AM", "noon": "midi", "pm": "PM"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
