package ug

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"NIO": ut.Currency{Currency: "NIO", DisplayName: "نىگېرىيە كوردوباسى", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "چىلى ھېسابات بىرلىكى (UF)", Symbol: ""}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "پورتۇگالىيە گىۋىنېيە ئېسكۇدوسى", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "موزامبىك مېتىكالى (1980–2006)", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "فىنلاندىيە مارككاسى", Symbol: ""}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "ۋىيېتنام دوڭى (1978–1985)", Symbol: ""}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "ئەزەربەيجان ماناتى (1993–2006)", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "بوسنىيە-خېرتسېگوۋىنا ئالماشتۇرۇشچان ماركى", Symbol: ""}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "بۇلغارىيە لېۋاسى (1879–1952)", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "ساينىت-ھېلېنا فوندستېرلىڭى", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "يۇگوسلاۋىيە ئالماشتۇرۇشچان دىنارى (1990–1992)", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "پاناما بالبوئاسى", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "پېرۇ يېڭى سولى", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "سىنگاپور دوللىرى", Symbol: ""}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "پاراگۋاي گۇئارانىسى", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "ئافرىقا قىتئەسى پۇل-مۇئامىلە ئىتتىپاقى فرانكى (BCEAO)", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "پۇل سىناش بىرلىكى", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "ئالجىرىيە دىنارى", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "كايمان ئاراللىرى دوللىرى", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "پاكىستان رۇپىسى", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "چېخ جۇمھۇرىيىتى كورۇناسى", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "ئىسلاندىيە كروناسى", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "باربادوس دوللىرى", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "زايىر زايىرى (1971–1993)", Symbol: ""}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "مېكسىكا مەبلەغ بىرلىكى", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "ئامېرىكا دوللىرى", Symbol: "$"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "بىرازىلىيە يېڭى كرۇزېروسى (1990–1993)", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "ئېكۋادور سۇكرېسى", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "ئىيوردانىيە دىنارى", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "رۋاندا فرانكى", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "تۈركىيە لىراسى", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "ئانگولا كۇۋانزاسى", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "گىۋىنېيە فرانكى", Symbol: ""}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "جەنۇبىي كورېيە ۋونى (1945–1953)", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "ئاۋسترىيە شىللىڭى", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "ھوندۇراس لېمپىراسى", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "تىمور ئېسكۇدوسى", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "ۋېنېزۇئېلا بولىۋارى (1871–2008)", Symbol: ""}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "كۈمۈش", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "گامبىيە دالاسى", Symbol: ""}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "ئىسلاندىيە كروناسى (1918–1981)", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "شىۋېتسىيە كروناسى", Symbol: ""}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "مالدىۋى رۇپىسى", Symbol: ""}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "پىلاتىنا", Symbol: ""}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "سوۋىت رۇبلىسى", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "ياۋروپا مۇرەككەپ بىرلىكى", Symbol: ""}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "كامبودژا رىئېلى", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "پېرۇ ئىنتىسى", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "پولشا زىلوتى (1950–1995)", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "بۇتان نگۇلترۇمى", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "فىلىپپىن پېسوسى", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "مالايشىيا رىڭگىتى", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ماداغاسقار ئارىئارىسى", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "رودېزىيە دوللىرى", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "بولىۋىيە پىسوسى", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "ئېكۋادور تۇراقلىق قىممەت بىرلىكى", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "ئىرېلاندىيە فوندستېرلىڭى", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "ئارگېنتىنا ئاۋسترالى", Symbol: ""}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "ليۇكسېمبۇرگ ئالماشتۇرۇشچان پېسوسى", Symbol: ""}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "ماكېدونىيە دىنارى (1992–1993)", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "قاتار رىيالى", Symbol: ""}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "ئاسىيا تەرەققىيات بانكىسى ھېسابات بىرلىكى", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "ئىسپانىيە پېسېتاسى (A ھېسابات)", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "ئىسپانىيە پېسېتاسى (ئالماشتۇرۇش ھېساباتى)", Symbol: ""}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "موزامبىك مېتىكالى", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "جەبىلتارىق فوند سىتېرلىڭى", Symbol: ""}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "گىۋاتېمالا كۇۋېتزالى", Symbol: ""}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "موناكو فرانكى", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "رۇسىيە رۇبلىسى", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "ئافغان ئافغانى (1927–2002)", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "ئارۇبان فىلورۇنى", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "كونگو فرانكى", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "يوچۇن پۇل", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "ئەرەب بىرلەشمە خەلىپىلىكى دەرھەمى", Symbol: ""}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "بوسنىيە-خېرتسېگوۋىنا دىنارى (1992–1994)", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "سېررالېئون لېئونېسى", Symbol: ""}, "PES": ut.Currency{Currency: "PES", DisplayName: "پېرۇ سولى (1863–1965)", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "پورتۇگالىيە ئېسكۇدوسى", Symbol: ""}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "لائوس كىپى", Symbol: ""}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "لاتۋىيە لاتى", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "گوللاندىيە گۈلدىنى", Symbol: ""}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "ۋېنگىرىيە فورېنتى", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "ياپونىيە يېنى", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "لاتۋىيە رۇبلىسى", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ئۇگاندا شىللىڭى", Symbol: ""}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "ئۇرۇگۋاي پېسوسى (ئىندېكىسلاش بىرلىكى)", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "پاللادىي", Symbol: ""}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "سۇكرې", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR ياۋرو", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "چېخسىلوۋاكىيە قاتتىق كورۇناسى", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "تۈركمەنىستان ماناتى", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ماۋرىتانىيە ئۇگىيەسى", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "سۇدان فوندستېرلىڭى (1957–1998)", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "بولىۋىيە بولىۋىيانوسى", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "سېربىيە دىنارى (2002–2006)", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "كىرودىيە كۇناسى", Symbol: ""}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "مالتا لىراسى", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "ئارگېنتىنا پېسو لېيى (1970–1983)", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "بېلىز دوللىرى", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "دانىيە كرونى", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "يېڭى زېلاندىيە دوللىرى", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "بوتسۋانا پۇلاسى", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "قازاقىستان تەڭگىسى", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "لىۋىيە دىنارى", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "بېلارۇسىيە رۇبلىسى", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "ئافرىقا قىتئەسى پۇل-مۇئامىلە ئىتتىپاقى فرانكى", Symbol: ""}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "باھاما دوللىرى", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "گىۋىنېيە سىلىسى", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "ئالاھىدە پۇل ئېلىش ھوقۇقى", Symbol: ""}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "جەنۇبىي كورېيە خۋانى (1953–1962)", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "موڭغۇلىيە تۈگرىكى", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "ئامېرىكا دوللىرى (كېيىنكى كۈن)", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "ئېستونىيە كرۇنى", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "لىيۇكسېمبۇرگ پۇل-مۇئامىلە فرانكى", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "ماكېدونىيە دىنارى", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "نېپال رۇپىسى", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "سېيشېل رۇپىسى", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "ئالبانىيە لېكى", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "ئەرمېنىيە دىرامى", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "سىپرۇس فوند ستېرلىڭى", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "تۈركمەنىستان ماناتى (1993–2009)", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "يېڭى تەيۋەن دوللىرى", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "مىسىر فوند سىتېرلىڭى", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "تايلاند باختى", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "جەنۇبىي ئافرىقا راندى", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "ئارگېنتىنا پېسوسى (1983–1985)", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "بۇرۇندى فرانكى", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "بىرما كىياتى", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "سالۋادور كولونى", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "ئارگېنتىنا پېسوسى (1881–1970)", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "بۇلغارىيە ئىجتىمائىي لېۋاسى", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "بىرازىلىيە رىيالى", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "زىمبابۋې دوللىرى (2009)", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "گوللاندىيەگە قاراشلىق ئانتىللېن گۇلدېنى", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "بەھرەين دىنارى", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "پاپۇئا يېڭى گىۋىنېيە كىناسى", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "ماراكەش فرانكى", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "پولشا زىلوتى", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "تونگا پائانگاسى", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "جەنۇبىي ئافرىقا راندى (پۇل–مۇئامىلە)", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "زامبىيە كۋاچاسى (1968–2012)", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "ئانگولا كۇۋانزاسى (1977–1991)", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "كانادا دوللىرى", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "ئىسرائىل يېڭى شېكېلى", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "جۇڭگو خەلق بانكىسى دوللىرى", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "كۇبا ئالماشتۇرۇشچان پېسوسى", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "گانا سېدىسى", Symbol: ""}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "ئىتالىيە لىراسى", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "مالدىۋى رۇفىياسى", Symbol: ""}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "بوسنىيە-خېرتسېگوۋىنا يېڭى دىنارى (1994–1997)", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "بۇلغارىيە قاتتىق لېۋاسى", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "بولىۋىيە بولىۋىيانوسى (1863–1963)", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ئومان رىيالى", Symbol: ""}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "ئۇگاندا شىللىڭى (1966–1987)", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "زىمبابۋې دوللىرى (1980–2008)", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "ئۇرۇگۋاي پېسوسى", Symbol: ""}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "ياۋروپا ھېسابات بىرلىكى (XBD)", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "ياۋروپا پۇل بىرلىكى", Symbol: ""}, "COU": ut.Currency{Currency: "COU", DisplayName: "كولومبىيە ھەقىقىي قىممەت بىرلىكى", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "فىجى دوللىرى", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "شىمالىي كورېيە ۋونى", Symbol: ""}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ئۇكرائىنا خرىۋناسى", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "مالاۋى كۋاچاسى", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "نىگېرىيە نايراسى", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "رۇسىيە رۇبلىسى (1991–1998)", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "نىگېرىيە كوردوباسى (1988–1991)", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "سەئۇدى رىيالى", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "سىۋېزىلاند لىلانگېنى", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "زىمبابۋې دوللىرى (2008)", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "بېلگىيە فرانكى", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "چىلى پېسوسى", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "لېسوتو لوتىسى", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "لىبېرىيە دوللىرى", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "ئانگولا قايتا تەڭشەلگەن كۇۋانزاسى (1995–1999)", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "جۇڭگو يۈەنى", Symbol: "￥"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "سىرىلانكا رۇپىسى", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "گىرۇزىيە كۇپون لارىتى", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "مىيانمار كىياتى", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "تۇنىس دىنارى", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "بىرازىلىيە يېڭى كرۇزادوسى (1989–1990)", Symbol: ""}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "بىرازىلىيە كرۇزېروسى (1993–1994)", Symbol: ""}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "فالكلاند ئاراللىرى فوند سىتېرلىڭى", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "فىرانسىيە فرانكى", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "شياڭگاڭ دوللىرى", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "سىلوۋاكىيە كورۇناسى", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "ۋىيېتنام دوڭى", Symbol: ""}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "ۋانۇئاتۇ ۋاتۇسى", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "بىرۇنېي دوللىرى", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "بىرازىلىيە كرۇزادوسى (1986–1989)", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "شەرقىي گېرمانىيە ماركى", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "فىرانسىيە UIC فرانكى", Symbol: ""}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "يۇگوسلاۋىيە ئىسلاھات دىنارى (1992–1993)", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "زامبىيە كۋاچاسى", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "مېكسىكا پېسوسى", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "رۇمىنىيە لېيى (1952–2006)", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "تاجىكىستان سومونىسى", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "ياۋروپا پۇل بىرلىكى (XBB)", Symbol: ""}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "چىلى ئېسكۇدوسى", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "كۇبا پېسوسى", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "گىۋىنېيە-بىسسائۇ پېسوسى", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "موزامبىك ئېسكۇدوسى", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "ئۇرۇگۋاي پېسوسى (1975–1993)", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "سۇدان دىنارى (1992–2007)", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "ئالتۇن", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "بېلگىيە فرانكى (پۇل–مۇئامىلە)", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "گىۋىئانا دوللىرى", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "ماۋرىتىئۇس رۇپىسى", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "ئالبانىيە لېكى (1946–1965)", Symbol: ""}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "بۇلغارىيە لېۋاسى", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "گېرمانىيە ماركى", Symbol: ""}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "مولدوۋا لېۋى", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "يەمەن دىنارى", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "كوستارىكا كولونى", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "ئېفىيوپىيە بىررى", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "ياۋرو", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "ئىسرائىلىيە فوندستېرلىڭى", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "ئانگولا يېڭى كۇۋانزاسى (1990–2000)", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "بېلارۇسىيە يېڭى رۇبلىسى (1994–1999)", Symbol: ""}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "گىرېتسىيە دراخماسى", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "ئامېرىكا دوللىرى (ئوخشاش كۈن)", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "ئارگېنتىنا پېسوسى", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "كولومبىيە پېسوسى", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "كېنىيە شىللىڭى", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "سۇرىنام گۈلدىنى", Symbol: ""}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "جەنۇبىي سۇدان فوندستېرلىڭى", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "ئەنگلىيە فوند سىتېرلىڭى", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "جەنۇبىي كورېيە ۋونى", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "سۇرىنام دوللىرى", Symbol: ""}, "WST": ut.Currency{Currency: "WST", DisplayName: "ساموئا تالاسى", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "لىتۋا تالوناسى", Symbol: ""}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "مولدوۋا كۇپونى", Symbol: ""}, "RON": ut.Currency{Currency: "RON", DisplayName: "رۇمىنىيە لېيى", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "سومالى شىللىڭى", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "ئۇكرائىنا كاربوۋانېتسى", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "شىۋېتسىيە فرانكى", Symbol: ""}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "ئىسرائىل شېكېلى (1980–1985)", Symbol: ""}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "مېكسىكا كۈمۈش پېسوسى (1861–1992)", Symbol: ""}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "بولىۋىيە مۇدولى", Symbol: ""}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "بىرازىلىيە يېڭى كرۇزېروسى (1967–1986)", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "بىرازىلىيە كرۇزېروسى (1942–1967)", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ئېكۋاتور گىۋىنېيە ئېكۋېلېسى", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "سېربىيە دىنارى", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "يەمەن رىيالى", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "ھىندونېزىيە رۇپىيەسى", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "كۇۋەيت دىنارى", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "نورۋېگىيە كرونى", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "تاجىكىستان رۇبلىسى", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "ئاندورران پېسېتاسى", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "ئەزەربەيجان ماناتى", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "ئىسپانىيە پېسېتاسى", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "نامىبىيە دوللىرى", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ئاۋسترالىيە دوللىرى", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "ئىراق دىنارى", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "مالى فرانكى", Symbol: ""}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "زايىر يېڭى زايىرى (1993–1998)", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "گىرۇزىيە لارىسى", Symbol: ""}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "ئىران رىيالى", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "لىۋان فوند سىتېرلىڭى", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "يېشىل تۇمشۇق ئېسكۇدوسى", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "ھايتى گۇردېسى", Symbol: ""}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "مالتا فوندستېرلىڭى", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "سىلوۋېنىيە تولارى", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "تۈركىيە لىراسى (1922–2005)", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "تىرىنىداد ۋە توباگو دوللىرى", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "ۋېنېزۇئېلا بولىۋارى", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "ئېرىترېيە ناكفاسى", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "ليۇكسېمبۇرگ فرانكى", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "ماراكەش دىرھەمى", Symbol: ""}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET فوندى", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "كىرودىيە دىنارى", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "قىرغىزىستان سومى", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "تانزانىيە شىللىڭى", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "شەرقىي كارىب دوللىرى", Symbol: ""}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "فىرانسىيە ئالتۇن فرانكى", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "ئافغان ئافغانى", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "بېلگىيە فرانكى (ئالماشتۇرۇشچان)", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "گانا سېدىسى (1979–2007)", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "يۇگوسلاۋىيە قاتتىق دىنارى (1966–1990)", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "ئاۋمېن پاتاكاسى", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "سان-تومې ۋە پىرىنسىپى دوبراسى", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "سۈرىيە فوندستېرلىڭى", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "يامايكا دوللىرى", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "لىتۋا لىتاسى", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "ماداغاسقار فرانكى", Symbol: ""}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "تىنچ ئوكيان پۇل-مۇئامىلە ئورتاق گەۋدىسى فرانكى", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "يۇگوسلاۋىيە يېڭى دىنارى (1994–2002)", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "باڭلادىش تاكاسى", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "سۇدان فوندستېرلىڭى", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "ياۋروپا ھېسابات بىرلىكى (XBC)", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "كومورو فرانكى", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "ئۆزبېكىستان سومى", Symbol: ""}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR فرانكى", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "سولومون ئاراللىرى دوللىرى", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "ھىندىستان رۇپىسى", Symbol: ""}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "بېرمۇدا دوللىرى", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "جىبۇتى فرانكى", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "دومىنىكا پېسوسى", Symbol: ""}}
