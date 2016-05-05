package ki

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEE, d MMMM y", Long: "d MMMM y", Medium: "d MMM y", Short: "dd/MM/y"}, Time: ut.CalendarDateFormat{Full: "HH:mm:ss zzzz", Long: "HH:mm:ss z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "", Long: "", Medium: "", Short: ""}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{3: "WGT", 4: "WKN", 6: "WTD", 9: "WKD", 11: "WMW", 1: "JEN", 2: "WKR", 5: "WTN", 7: "WMJ", 8: "WNN", 10: "WIK", 12: "DIT"}, Narrow: ut.CalendarMonthFormatNameValue{10: "I", 11: "I", 1: "J", 2: "K", 3: "G", 5: "G", 6: "G", 9: "K", 4: "K", 7: "M", 8: "K", 12: "D"}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{10: "Mwere wa ikũmi", 12: "Ndithemba", 1: "Njenuarĩ", 3: "Mwere wa gatatũ", 5: "Mwere wa gatano", 6: "Mwere wa gatandatũ", 8: "Mwere wa kanana", 2: "Mwere wa kerĩ", 4: "Mwere wa kana", 7: "Mwere wa mũgwanja", 9: "Mwere wa kenda", 11: "Mwere wa ikũmi na ũmwe"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{4: "ART", 5: "NMA", 6: "NMM", 0: "KMA", 1: "NTT", 2: "NMN", 3: "NMT"}, Narrow: ut.CalendarDayFormatNameValue{0: "K", 1: "N", 2: "N", 3: "N", 4: "A", 5: "N", 6: "N"}, Short: ut.CalendarDayFormatNameValue(nil), Wide: ut.CalendarDayFormatNameValue{1: "Njumatatũ", 2: "Njumaine", 3: "Njumatana", 4: "Aramithi", 5: "Njumaa", 6: "Njumamothi", 0: "Kiumia"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "Kiroko", "pm": "Hwaĩ-inĩ"}, Narrow: ut.CalendarPeriodFormatNameValue(nil), Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"am": "Kiroko", "pm": "Hwaĩ-inĩ"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "", Abbrev: "", Narrow: ""}, BC: ut.CalendarEraFormatNames{Full: "Mbere ya Kristo", Abbrev: "MK", Narrow: ""}}}}
