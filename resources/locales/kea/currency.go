package kea

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilanjeni", Symbol: "SZL"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "", Symbol: "PKR"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kuanza", Symbol: "AOA"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula di Botsuana", Symbol: "BWP"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Sedi di Gana", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Sili", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kuaxa di Malaui", Symbol: "MWK"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "", Symbol: "OMR"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupia di Seixelis", Symbol: "SCR"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Xelin somalianu", Symbol: "SOS"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "", Symbol: "BGN"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Skudu Kabuverdianu", Symbol: "\u200b"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinar tunizianu", Symbol: "TND"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kuaxa zambianu", Symbol: "ZMW"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Zloty polaku", Symbol: "PLN"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "", Symbol: "CFPF"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "", Symbol: "ALL"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "", Symbol: "CZK"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Xelin kenianu", Symbol: "KES"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metikal", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "", Symbol: "NIO"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dola merkanu", Symbol: "US$"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "", Symbol: "AFN"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dola australianu", Symbol: "AU$"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Franku di Djibuti", Symbol: "DJF"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nafka di Eritreia", Symbol: "ERN"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "", Symbol: "FKP"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinar libiu", Symbol: "LYD"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "", Symbol: "MZN"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dola namibianu", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Franku ruandes", Symbol: "RWF"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "", Symbol: "SYP"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Diren di Emiradus Arabi Unidu", Symbol: "AED"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Kuroa norueges", Symbol: "NOK"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "", Symbol: "PEN"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Libra sudanes antigu", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "", Symbol: "VEF"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Franku suisu", Symbol: "CHF"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "", Symbol: "GNF"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "", Symbol: "HTG"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "", Symbol: "MDL"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "", Symbol: "MOP"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "", Symbol: "KWD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "", Symbol: "TMT"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Rial brazileru", Symbol: "R$"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Ieni japones", Symbol: "JP¥"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "", Symbol: "PAB"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dola di Zimbabue", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Franku kongoles", Symbol: "CDF"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "", Symbol: "TJS"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "", Symbol: "MYR"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Pezu mexikanu", Symbol: "MX$"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "", Symbol: "NZ$"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "", Symbol: "ANG"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar di Barain", Symbol: "BHD"}, "BND": ut.Currency{Currency: "BND", DisplayName: "", Symbol: "$"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "", Symbol: "HNL"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "", Symbol: "LBP"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Kuroa dinamarkeza", Symbol: "DKK"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "", Symbol: "ISK"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "", Symbol: "VUV"}, "WST": ut.Currency{Currency: "WST", DisplayName: "", Symbol: "WST"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "", Symbol: "KPW"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "", Symbol: "BOB"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "", Symbol: "QAR"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Franku CFA BCEAO", Symbol: "CFA"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Bir etiopi", Symbol: "ETB"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "", Symbol: "$"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupia indianu", Symbol: "₹"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dola liberianu", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Rublu rusu", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "", Symbol: "BTN"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "", Symbol: "JOD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupia di Maurisias", Symbol: "MUR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "", Symbol: "RSD"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "", Symbol: "AMD"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Kuroa sueku", Symbol: "SEK"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra di Sãu Tume i Prinsipi", Symbol: "STD"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "", Symbol: "UZS"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "", Symbol: "AZN"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinar arjelinu", Symbol: "DZD"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "", Symbol: "HRK"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "", Symbol: "HUF"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "", Symbol: "₮"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "", Symbol: "MVR"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Rial saudita", Symbol: "SAR"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Xelin di Tanzania", Symbol: "TZS"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi", Symbol: "GMD"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "", Symbol: "IQD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "", Symbol: "LTL"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Franku CFA BEAC", Symbol: "FCFA"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "", Symbol: "EC$"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Rupia indoneziu", Symbol: "IDR"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "", Symbol: "₸"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "", Symbol: "LKR"}, "VND": ut.Currency{Currency: "VND", DisplayName: "", Symbol: "₫"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Lira turku", Symbol: "₺"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Dola Novu di Taiwan", Symbol: "NT$"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dola kanadianu", Symbol: "CA$"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Libra britaniku", Symbol: "£"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Diren marokinu", Symbol: "MAD"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "", Symbol: "MMK"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "", Symbol: "₱"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand sulafrikanu", Symbol: "ZAR"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kuaxa zambianu (1968–2012)", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won sul-koreanu", Symbol: "₩"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ougia", Symbol: "MRO"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "", Symbol: "$"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "", Symbol: "TOP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Xelin ugandensi", Symbol: "UGX"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "", Symbol: "PGK"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "", Symbol: "$"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Mueda diskonxedu", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "", Symbol: "AWG"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Libra ejipsiu", Symbol: "EGP"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "", Symbol: "៛"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "", Symbol: "₭"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "", Symbol: "NPR"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "", Symbol: "GEL"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dola di Ong Kong", Symbol: "HK$"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht tailandes", Symbol: "฿"}, "YER": ut.Currency{Currency: "YER", DisplayName: "", Symbol: "YER"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Franku borundes", Symbol: "BIF"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Iuan xines", Symbol: "CN¥"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti di Lezotu", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariari di Madagaskar", Symbol: "MGA"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "", Symbol: "MKD"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "", Symbol: "₪"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "", Symbol: "KGS"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Libra di Santa Ilena", Symbol: "SHP"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone di Sera Leoa", Symbol: "SLL"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "", Symbol: "BAM"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "", Symbol: "GTQ"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Franku di Komoris", Symbol: "KMF"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Libra sudanes", Symbol: "SDG"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "", Symbol: "৳"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "", Symbol: "IRR"}}
