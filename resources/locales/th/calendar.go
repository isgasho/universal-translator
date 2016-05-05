package th

import "github.com/go-playground/universal-translator"

var calendar = ut.Calendar{Formats: ut.CalendarFormats{Date: ut.CalendarDateFormat{Full: "EEEEที่ d MMMM G y", Long: "d MMMM G y", Medium: "d MMM y", Short: "d/M/yy"}, Time: ut.CalendarDateFormat{Full: "H นาฬิกา mm นาที ss วินาที zzzz", Long: "H นาฬิกา mm นาที ss วินาที z", Medium: "HH:mm:ss", Short: "HH:mm"}, DateTime: ut.CalendarDateFormat{Full: "{1} {0}", Long: "{1} {0}", Medium: "{1} {0}", Short: "{1} {0}"}}, FormatNames: ut.CalendarFormatNames{Months: ut.CalendarMonthFormatNames{Abbreviated: ut.CalendarMonthFormatNameValue{5: "พ.ค.", 7: "ก.ค.", 8: "ส.ค.", 10: "ต.ค.", 11: "พ.ย.", 12: "ธ.ค.", 1: "ม.ค.", 2: "ก.พ.", 3: "มี.ค.", 4: "เม.ย.", 6: "มิ.ย.", 9: "ก.ย."}, Narrow: ut.CalendarMonthFormatNameValue{2: "ก.พ.", 3: "มี.ค.", 9: "ก.ย.", 11: "พ.ย.", 7: "ก.ค.", 8: "ส.ค.", 10: "ต.ค.", 12: "ธ.ค.", 1: "ม.ค.", 4: "เม.ย.", 5: "พ.ค.", 6: "มิ.ย."}, Short: ut.CalendarMonthFormatNameValue(nil), Wide: ut.CalendarMonthFormatNameValue{3: "มีนาคม", 5: "พฤษภาคม", 8: "สิงหาคม", 9: "กันยายน", 10: "ตุลาคม", 12: "ธันวาคม", 1: "มกราคม", 2: "กุมภาพันธ์", 4: "เมษายน", 6: "มิถุนายน", 7: "กรกฎาคม", 11: "พฤศจิกายน"}}, Days: ut.CalendarDayFormatNames{Abbreviated: ut.CalendarDayFormatNameValue{6: "ส.", 0: "อา.", 1: "จ.", 2: "อ.", 3: "พ.", 4: "พฤ.", 5: "ศ."}, Narrow: ut.CalendarDayFormatNameValue{3: "พ", 4: "พฤ", 5: "ศ", 6: "ส", 0: "อา", 1: "จ", 2: "อ"}, Short: ut.CalendarDayFormatNameValue{6: "ส.", 0: "อา.", 1: "จ.", 2: "อ.", 3: "พ.", 4: "พฤ.", 5: "ศ."}, Wide: ut.CalendarDayFormatNameValue{4: "วันพฤหัสบดี", 5: "วันศุกร์", 6: "วันเสาร์", 0: "วันอาทิตย์", 1: "วันจันทร์", 2: "วันอังคาร", 3: "วันพุธ"}}, Periods: ut.CalendarPeriodFormatNames{Abbreviated: ut.CalendarPeriodFormatNameValue{"am": "ก่อนเที่ยง", "pm": "หลังเที่ยง"}, Narrow: ut.CalendarPeriodFormatNameValue{"midnight": "เที่ยงคืน", "noon": "เที่ยง", "pm": "หลังเที่ยง", "afternoon2": "บ่าย", "evening2": "ค่ำ", "night1": "กลางคืน", "am": "ก่อนเที่ยง", "morning1": "เช้า", "afternoon1": "ช่วงเที่ยง", "evening1": "เย็น"}, Short: ut.CalendarPeriodFormatNameValue(nil), Wide: ut.CalendarPeriodFormatNameValue{"midnight": "เที่ยงคืน", "pm": "หลังเที่ยง", "morning1": "เช้า", "afternoon2": "บ่าย", "evening2": "ค่ำ", "night1": "กลางคืน", "am": "ก่อนเที่ยง", "noon": "เที่ยง", "afternoon1": "ช่วงเที่ยง", "evening1": "เย็น"}}, Eras: ut.Eras{AD: ut.CalendarEraFormatNames{Full: "คริสต์ศักราช", Abbrev: "ค.ศ.", Narrow: "ค.ศ."}, BC: ut.CalendarEraFormatNames{Full: "ปีก่อนคริสต์ศักราช", Abbrev: "ปีก่อน ค.ศ.", Narrow: "ก่อน ค.ศ."}}}}
