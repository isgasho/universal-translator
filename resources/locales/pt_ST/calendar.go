package pt_ST

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d 'de' MMMM 'de' y", Long: "d 'de' MMMM 'de' y", Medium: "d 'de' MMM 'de' y", Short: "dd/MM/yy"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{4: "abr", 12: "dez", 11: "nov", 1: "jan", 2: "fev", 6: "jun", 8: "ago", 9: "set", 3: "mar", 10: "out", 5: "mai", 7: "jul"}, Narrow: ut.CalendarMonthFormatNameValue{9: "S", 11: "N", 3: "M", 12: "D", 5: "M", 6: "J", 7: "J", 8: "A", 1: "J", 2: "F", 4: "A", 10: "O"}, Short: ut.CalendarMonthFormatNameValue{}, Wide: ut.CalendarMonthFormatNameValue{3: "março", 5: "maio", 6: "junho", 1: "janeiro", 8: "agosto", 11: "novembro", 12: "dezembro", 2: "fevereiro", 4: "abril", 10: "outubro", 7: "julho", 9: "setembro"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "sáb", 0: "dom", 1: "seg", 2: "ter", 3: "qua", 4: "qui", 5: "sex"}, Narrow: ut.CalendarDayFormatNameValue{4: "Q", 5: "S", 6: "S", 0: "D", 1: "S", 2: "T", 3: "Q"}, Short: ut.CalendarDayFormatNameValue{6: "sáb", 0: "dom", 1: "seg", 2: "ter", 3: "qua", 4: "qui", 5: "sex"}, Wide: ut.CalendarDayFormatNameValue{3: "quarta-feira", 4: "quinta-feira", 5: "sexta-feira", 6: "sábado", 0: "domingo", 1: "segunda-feira", 2: "terça-feira"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "AM", "noon": "meio-dia", "pm": "PM", "morning1": "manhã", "afternoon1": "tarde", "evening1": "noite", "night1": "madrugada", "midnight": "meia-noite"}, Narrow: ut.CalendarPeriodFormatNameValue{"noon": "meio-dia", "pm": "PM", "morning1": "manhã", "afternoon1": "tarde", "evening1": "noite", "night1": "madrugada", "midnight": "meia-noite", "am": "AM"}, Short: ut.CalendarPeriodFormatNameValue{}, Wide: ut.CalendarPeriodFormatNameValue{"pm": "PM", "morning1": "manhã", "afternoon1": "tarde", "evening1": "noite", "night1": "madrugada", "midnight": "meia-noite", "am": "AM", "noon": "meio-dia"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}}}}
