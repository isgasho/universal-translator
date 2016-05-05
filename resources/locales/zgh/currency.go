package zgh

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MGA": ut.Currency{Currency: "MGA", DisplayName: "ⴼⵔⴰⵏⴽ ⵏ ⵎⴰⴷⴰⵖⴰⵛⵇⴰⵔ", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "ⴰⵛⵉⵍⵉⵏ ⵏ ⴽⵉⵏⵢⴰ", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "ⴼⵔⴰⵏⴽ ⵏ ⵇⵓⵎⵓⵕ", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "ⵍⵉⵍⴰⵏⵊⵉⵏⵉ", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "ⴱⵉⵔ ⵏ ⵉⵜⵢⵓⴱⵢⴰ", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "ⵍⵓⵜⵉ ⵏ ⵍⵉⵚⵓⵟⵓ", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "ⵉⵙⴽⵓⴷⵓ ⵏ ⴽⴰⴱⴱⵉⵔⴷⵉ", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵍⵉⴱⵉⵔⵢⴰ", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵉⵡⵓⵏⴰⴽ ⵉⵎⵓⵏⵏ", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵣⵉⵎⴱⴰⴱⵡⵉ (2009)", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "ⴰⴱⵓⵍⴰ ⵏ ⴱⵓⵜⵙⵡⴰⵏⴰ", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "ⴼⵔⴰⵏⴽ ⵏ ⴽⵓⵏⴳⵓ", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "ⴼⵔⴰⵏⴽ ⵏ ⴷⵊⵉⴱⵓⵜⵉ", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "ⴰⵢⴰⵏ ⵏ ⵍⵢⴰⴱⴰⵏ", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ⵓⵇⵉⵢⵢⴰ ⵏ ⵎⵓⵕⵉⵟⴰⵏⵢⴰ", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵏⴰⵎⵉⴱⵢⴰ", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "ⴰⵔⵢⴰⵍ ⵏ ⵙⵙⴰⵄⵓⴷⵉⵢⴰ", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "ⴰⵔⵓⴱⵉ ⵏ ⵙⵙⵉⵛⵉⵍ", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "ⴰⴷⵉⵏⴰⵔ ⵏ ⵙⵙⵓⴷⴰⵏ", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "ⴽⵡⴰⵏⵣⴰ ⵏ ⴰⵏⴳⵓⵍⴰ", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ⵙⵉⴷⵉ ⵏ ⵖⴰⵏⴰ", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "ⴰⵔⵓⴱⵉ ⵏ ⵍⵀⵉⵏⴷ", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "ⴰⵎⵉⵜⵉⴽⵍ ⵏ ⵎⵓⵣⵏⴱⵉⵇ", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "ⵍⵉⵢⵓⵏ", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "ⴰⴷⵉⵏⴰⵔ ⵏ ⴷⵣⴰⵢⵔ", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "ⵏⴰⴼⴽⴰ ⵏ ⵉⵔⵉⵜⵉⵔⵢⴰ", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "ⴰⴷⵉⵏⴰⵔ ⵏ ⵍⵉⴱⵢⴰ", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "ⴰⴷⵔⵉⵎ ⵏ ⵍⵎⵖⵔⵉⴱ", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "ⵏⴰⵢⵔⴰ ⵏ ⵏⵉⵊⵉⵔⵢⴰ", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "ⴰⴷⵉⵏⴰⵔ ⵏ ⴱⵃⵔⴰⵢⵏ", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "ⴼⵔⴰⵏⴽ ⵏ ⴱⵓⵔⵓⵏⴷⵉ", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "ⴰⴼⵔⴰⵏⴽ ⵏ ⵙⵡⵉⵙⵔⴰ", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "ⴰⵔⵓⴱⵉ ⵏ ⵎⵓⵔⵉⵙ", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "ⴰⴼⵔⴰⵏⴽ ⵏ ⵔⵡⴰⵏⴷⴰ", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "ⴰⵔⴰⵏⴷ ⵏ ⴰⴼⵔⵉⵇⵢⴰ ⵏ ⵉⴼⴼⵓⵙ", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵣⵉⵎⴱⴰⴱⵡⵉ (1980–2008)", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "ⴰⵢⴰⵏ ⵏ ⵛⵛⵉⵏⵡⴰ", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "ⴷⴰⵍⴰⵙⵉ ⵏ ⴳⴰⵎⴱⵢⴰ", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "ⴰⵊⵏⵉⵀ ⵏ ⵙⵙⵓⴷⴰⵏ", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "ⴰⵊⵏⵉⵀ ⵏ ⵙⴰⵏⵜⵉⵍⵉⵏ", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "ⴰⵛⵉⵍⵉⵏ ⵏ ⵟⴰⵏⵥⴰⵏⵢⴰ", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "ⴰⵊⵏⵉⵀ ⴰⵙⵜⵔⵍⵉⵏⵉ ⵏ ⵏⵏⴳⵍⵉⵣ", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "ⴰⴷⵓⴱⵔⴰ ⵏ ⵙⴰⵏⵟⵓⵎⵉ", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵓⵙⵜⵔⴰⵍⵢⴰ", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "ⵓⵔⵓ", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "ⴽⵡⴰⵛⴰ ⵏ ⵎⴰⵍⴰⵡⵉ", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "ⴰⴷⵉⵏⴰⵔ ⵏ ⵜⵓⵏⵙ", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "ⴼⵔⴰⵏⴽ ⵚⵉⴼⴰ", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "ⴰⴽⵡⴰⵛⴰ ⵏ ⵣⴰⵎⴱⵢⴰ (1968–2012)", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "ⴰⴽⵡⴰⵛⴰ ⵏ ⵣⴰⵎⴱⵢⴰ", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⵣⵉⵎⴱⴰⴱⵡⵉ (2008)", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "ⴰⴷⵓⵍⴰⵔ ⵏ ⴽⴰⵏⴰⴷⴰ", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ⴰⵛⵉⵍⵉⵏ ⵏ ⵓⵖⴰⵏⴷⴰ", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "ⴰⵛⵉⵍⵉⵏ ⵏ ⵚⵚⵓⵎⴰⵍ", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "ⴼⵔⴰⵏⴽ ⵚⵉⴼⴰ ⴱⵉⵙⴰⵡ", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "ⴰⴷⵔⵉⵎ ⵏ ⵍⵉⵎⴰⵔⴰⵜ", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "ⴰⵊⵏⵉⵀ ⵏ ⵎⵉⵚⵕ", Symbol: ""}}
