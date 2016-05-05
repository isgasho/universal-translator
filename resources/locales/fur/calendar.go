package fur

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE d 'di' MMMM 'dal' y", Long: "d 'di' MMMM 'dal' y", Medium: "dd/MM/y", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "Mar", 5: "Mai", 6: "Jug", 9: "Set", 1: "Zen", 4: "Avr", 7: "Lui", 8: "Avo", 10: "Otu", 11: "Nov", 12: "Dic", 2: "Fev"}, Narrow: ut.CalendarMonthFormatNameValue{4: "A", 5: "M", 6: "J", 7: "L", 8: "A", 9: "S", 10: "O", 1: "Z", 2: "F", 3: "M", 11: "N", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{6: "Jugn", 9: "Setembar", 5: "Mai", 2: "Fevrâr", 3: "Març", 4: "Avrîl", 7: "Lui", 8: "Avost", 10: "Otubar", 11: "Novembar", 1: "Zenâr", 12: "Dicembar"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{1: "lun", 2: "mar", 3: "mie", 4: "joi", 5: "vin", 6: "sab", 0: "dom"}, Narrow: ut.CalendarDayFormatNameValue{5: "V", 6: "S", 0: "D", 1: "L", 2: "M", 3: "M", 4: "J"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{2: "martars", 3: "miercus", 4: "joibe", 5: "vinars", 6: "sabide", 0: "domenie", 1: "lunis"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "a.", "pm": "p."}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "a.", "pm": "p."}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "pdC", Narrow: ""}}}}
