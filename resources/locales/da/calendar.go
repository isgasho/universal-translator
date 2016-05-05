package da

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE 'den' d. MMMM y", Long: "d. MMMM y", Medium: "d. MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH.mm.ss zzzz", Long: "HH.mm.ss z", Medium: "HH.mm.ss", Short: "HH.mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} 'kl'. {0}", Long: "{1} 'kl'. {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{12: "dec", 4: "apr", 7: "jul", 8: "aug", 10: "okt", 6: "jun", 9: "sep", 11: "nov", 1: "jan", 2: "feb", 3: "mar", 5: "maj"}, Narrow: ut.CalendarMonthFormatNameValue{11: "N", 1: "J", 4: "A", 7: "J", 8: "A", 9: "S", 10: "O", 2: "F", 3: "M", 5: "M", 6: "J", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{4: "april", 5: "maj", 6: "juni", 9: "september", 1: "januar", 2: "februar", 8: "august", 10: "oktober", 11: "november", 12: "december", 3: "marts", 7: "juli"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{5: "fre", 6: "lør", 0: "søn", 1: "man", 2: "tir", 3: "ons", 4: "tor"}, Narrow: ut.CalendarDayFormatNameValue{5: "F", 6: "L", 0: "S", 1: "M", 2: "T", 3: "O", 4: "T"}, Short: ut.CalendarDayFormatNameValue{5: "fr", 6: "lø", 0: "sø", 1: "ma", 2: "ti", 3: "on", 4: "to"}, Wide: ut.CalendarDayFormatNameValue{5: "fredag", 6: "lørdag", 0: "søndag", 1: "mandag", 2: "tirsdag", 3: "onsdag", 4: "torsdag"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "morgen", "morning2": "formiddag", "afternoon1": "eftermiddag", "evening1": "aften", "night1": "nat", "midnight": "midnat", "am": "AM"}, Narrow: ut.CalendarPeriodFormatNameValue{"morning1": "morgen", "morning2": "formiddag", "afternoon1": "eftermiddag", "evening1": "aften", "night1": "nat", "midnight": "midnat", "am": "AM", "pm": "PM"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "AM", "pm": "PM", "morning1": "morgen", "morning2": "formiddag", "afternoon1": "eftermiddag", "evening1": "aften", "night1": "nat", "midnight": "midnat"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "e.Kr.", Abbrev: "e.Kr.", Narrow: "eKr"}, BC: ut.CalendarEraFormatNames{Full: "f.Kr.", Abbrev: "f.Kr.", Narrow: "fKr"}}}}
