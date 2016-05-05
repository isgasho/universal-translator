package zh

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"AOA": ut.Currency{Currency: "AOA", DisplayName: "安哥拉宽扎", Symbol: "AOA"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "刚果法郎", Symbol: "CDF"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "厄瓜多尔 (UVC)", Symbol: "ECV"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "芬兰马克", Symbol: "FIM"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "阿曼里亚尔", Symbol: "OMR"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "乌兹别克斯坦苏姆", Symbol: "UZS"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "波士尼亚-赫塞哥维纳第纳尔 (1992–1994)", Symbol: "BAD"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "希腊德拉克马", Symbol: "GRD"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "洪都拉斯伦皮拉", Symbol: "HNL"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "以色列谢克尔(1980–1985)", Symbol: "ILS"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "越南盾 (1978–1985)", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "巴西克鲁塞罗 (1990–1993)", Symbol: "BRE"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "塞拉利昂利昂", Symbol: "SLL"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "南非兰特 (金融)", Symbol: "ZAL"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "古巴可兑换比索", Symbol: "CUC"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "哈萨克斯坦坚戈", Symbol: "KZT"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "莫桑比克美提卡", Symbol: "MZN"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "扎伊尔 (1971–1993)", Symbol: "ZRZ"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "铂", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "巴西克鲁扎多 (1986–1989)", Symbol: "BRC"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "佛得角埃斯库多", Symbol: "CVE"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "伊朗里亚尔", Symbol: "IRR"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "开曼元", Symbol: "KYD"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "斯里兰卡卢比", Symbol: "LKR"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "沙特里亚尔", Symbol: "SAR"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "欧洲复合单位", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "印度尼西亚盾", Symbol: "IDR"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "科威特第纳尔", Symbol: "KWD"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "立陶宛立特", Symbol: "LTL"}, "USD": ut.Currency{Currency: "USD", DisplayName: "美元", Symbol: "US$"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "澳大利亚元", Symbol: "AU$"}, "BND": ut.Currency{Currency: "BND", DisplayName: "文莱元", Symbol: "BND"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "人民币", Symbol: "￥"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "危地马拉格查尔", Symbol: "GTQ"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "拉脱维亚拉特", Symbol: "LVL"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "马耳他里拉", Symbol: "MTL"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "尼日利亚奈拉", Symbol: "NGN"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "伯利兹元", Symbol: "BZD"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "纳米比亚元", Symbol: "NAD"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "缅元", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "克罗地亚库纳", Symbol: "HRK"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "以色列镑", Symbol: "ILP"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "厄瓜多尔苏克雷", Symbol: "ECS"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "朝鲜元", Symbol: "KPW"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "马其顿第纳尔 (1992–1993)", Symbol: "MKN"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "帝汶埃斯库多", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "土耳其里拉 (1922–2005)", Symbol: "TRL"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "乌克兰币", Symbol: "UAK"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "黄金", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "东德奥斯特马克", Symbol: "DDM"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "秘鲁印第", Symbol: "PEI"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "巴西新克鲁赛罗 (1967–1986)", Symbol: "BRB"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "卢森堡法郎", Symbol: "LUF"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "瑞典克朗", Symbol: "SEK"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "西班牙比塞塔（帐户 A）", Symbol: "ESA"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "俄罗斯卢布", Symbol: "RUB"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "博茨瓦纳普拉", Symbol: "BWP"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "厄立特里亚纳克法", Symbol: "ERN"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "埃塞俄比亚比尔", Symbol: "ETB"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "葡萄牙埃斯库多", Symbol: "PTE"}, "RON": ut.Currency{Currency: "RON", DisplayName: "罗马尼亚列伊", Symbol: "RON"}, "THB": ut.Currency{Currency: "THB", DisplayName: "泰铢", Symbol: "THB"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "塔吉克斯坦卢布", Symbol: "TJR"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "拉脱维亚卢布", Symbol: "LVR"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "瓦努阿图瓦图", Symbol: "VUV"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "港元", Symbol: "HK$"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "委内瑞拉玻利瓦尔 (1871–2008)", Symbol: "VEB"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "阿根廷比索 (1983–1985)", Symbol: "ARP"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "比利时法郎（可兑换）", Symbol: "BEC"}, "COU": ut.Currency{Currency: "COU", DisplayName: "哥伦比亚币", Symbol: "COU"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "韩元", Symbol: "￦"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "利比亚第纳尔", Symbol: "LYD"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "波兰兹罗提 (1950–1995)", Symbol: "PLZ"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "斯洛伐克克朗", Symbol: "SKK"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "巴巴多斯元", Symbol: "BBD"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "玻利维亚 Mvdol（资金）", Symbol: "BOV"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "哥斯达黎加科朗", Symbol: "CRC"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "格鲁吉亚拉里", Symbol: "GEL"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "黎巴嫩镑", Symbol: "LBP"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "墨西哥（资金）", Symbol: "MXV"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "巴基斯坦卢比", Symbol: "PKR"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "阿根廷比索", Symbol: "ARS"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "波斯尼亚-黑塞哥维那可兑换马克", Symbol: "BAM"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "马拉维克瓦查", Symbol: "MWK"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "坦桑尼亚先令", Symbol: "TZS"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "阿根廷法定比索 (1970–1983)", Symbol: "ARL"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "捷克硬克郎", Symbol: "CSK"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "法国法郎", Symbol: "FRF"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "圭亚那元", Symbol: "GYD"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "约旦第纳尔", Symbol: "JOD"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "布隆迪法郎", Symbol: "BIF"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "伊拉克第纳尔", Symbol: "IQD"}, "TND": ut.Currency{Currency: "TND", DisplayName: "突尼斯第纳尔", Symbol: "TND"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "特别提款权", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "亚美尼亚德拉姆", Symbol: "AMD"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "巴西新克鲁扎多 (1989–1990)", Symbol: "BRN"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "乔治亚库蓬拉瑞特", Symbol: "GEK"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "尼泊尔卢比", Symbol: "NPR"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "罗得西亚元", Symbol: "RHD"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "圣赫勒拿群岛磅", Symbol: "SHP"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "赤道几内亚埃奎勒", Symbol: "GQE"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "利比里亚元", Symbol: "LRD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "马其顿第纳尔", Symbol: "MKD"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "所罗门群岛元", Symbol: "SBD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "越南盾", Symbol: "₫"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "欧洲计算单位 (XBC)", Symbol: ""}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "安哥拉宽扎 (1977–1990)", Symbol: "AOK"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "直布罗陀镑", Symbol: "GIP"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "法国法郎 (UIC)", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "也门第纳尔", Symbol: "YDD"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "阿鲁巴基尔德元", Symbol: "AWG"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "苏丹镑", Symbol: "SDG"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "汤加潘加", Symbol: "TOP"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "乌拉圭比索 (1975–1993)", Symbol: "UYP"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "欧洲货币单位", Symbol: "XEU"}, "AED": ut.Currency{Currency: "AED", DisplayName: "阿联酋迪拉姆", Symbol: "AED"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "比利时法郎", Symbol: "BEF"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "莫桑比克埃斯库多", Symbol: "MZE"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "波兰兹罗提", Symbol: "PLN"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "乌克兰格里夫纳", Symbol: "UAH"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "保加利亚列弗 (1879–1952)", Symbol: "BGO"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "智利（资金）", Symbol: "CLF"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "加纳塞第", Symbol: "GHC"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "斯威士兰里兰吉尼", Symbol: "SZL"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "阿塞拜疆马纳特 (1993–2006)", Symbol: "AZM"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "加拿大元", Symbol: "CA$"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "冈比亚达拉西", Symbol: "GMD"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "克罗地亚第纳尔", Symbol: "HRD"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "莱索托洛蒂", Symbol: "LSL"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "秘鲁新索尔", Symbol: "PEN"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "南斯拉夫可兑换第纳尔 (1990–1992)", Symbol: "YUN"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "法郎 (WIR)", Symbol: "CHW"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "德国马克", Symbol: "DEM"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "韩元 (1945–1953)", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "卡塔尔里亚尔", Symbol: "QAR"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "苏里南盾", Symbol: "SRG"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "古巴比索", Symbol: "CUP"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "马达加斯加阿里亚里", Symbol: "MGA"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "澳门元", Symbol: "MOP"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "马来西亚林吉特", Symbol: "MYR"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "巴拿马巴波亚", Symbol: "PAB"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "乌干达先令", Symbol: "UGX"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "荷属安的列斯盾", Symbol: "ANG"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "牙买加元", Symbol: "JMD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "巴拉圭瓜拉尼", Symbol: "PYG"}, "YER": ut.Currency{Currency: "YER", DisplayName: "也门里亚尔", Symbol: "YER"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "安哥拉重新调整宽扎 (1995–1999)", Symbol: "AOR"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "捷克克郎", Symbol: "CZK"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "旧莫桑比克美提卡", Symbol: "MZM"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "叙利亚镑", Symbol: "SYP"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "乌干达先令 (1966–1987)", Symbol: "UGS"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "欧元 (WIR)", Symbol: "CHE"}, "COP": ut.Currency{Currency: "COP", DisplayName: "哥伦比亚比索", Symbol: "COP"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "爱沙尼亚克朗", Symbol: "EEK"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "几内亚西里", Symbol: "GNS"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "韩元 (1953–1962)", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "中非金融合作法郎", Symbol: "FCFA"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "津巴布韦元 (1980–2008)", Symbol: "ZWD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "肯尼亚先令", Symbol: "KES"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "马里法郎", Symbol: "MLF"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "土库曼斯坦马纳特", Symbol: "TMT"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "银", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "缅甸元", Symbol: "MMK"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "蒙古图格里克", Symbol: "MNT"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "苏里南元", Symbol: "SRD"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "不丹努尔特鲁姆", Symbol: "BTN"}, "INR": ut.Currency{Currency: "INR", DisplayName: "印度卢比", Symbol: "₹"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "意大利里拉", Symbol: "ITL"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "摩纳哥法郎", Symbol: "MCF"}, "PES": ut.Currency{Currency: "PES", DisplayName: "秘鲁索尔 (1863–1965)", Symbol: "PES"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "安道尔比塞塔", Symbol: "ADP"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "玻利维亚比索", Symbol: "BOP"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "丹麦克朗", Symbol: "DKK"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "科摩罗法郎", Symbol: "KMF"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "尼加拉瓜科多巴 (1988–1991)", Symbol: "NIC"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "索马里先令", Symbol: "SOS"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "乌拉圭比索（索引单位）", Symbol: "UYI"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "欧洲计算单位 (XBD)", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "阿尔巴尼亚列克", Symbol: "ALL"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "吉布提法郎", Symbol: "DJF"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "摩尔多瓦列伊", Symbol: "MDL"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "毛里求斯卢比", Symbol: "MUR"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "尼加拉瓜金科多巴", Symbol: "NIO"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "卢旺达法郎", Symbol: "RWF"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "土耳其里拉", Symbol: "TRY"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "吉尔吉斯斯坦索姆", Symbol: "KGS"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "卢森堡金融法郎", Symbol: "LUL"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "新西兰元", Symbol: "NZ$"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "非洲金融共同体法郎", Symbol: "CFA"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "未知货币", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "欧元", Symbol: "€"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "葡萄牙几内亚埃斯库多", Symbol: "GWE"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "南苏丹镑", Symbol: "SSP"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET 基金", Symbol: ""}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "巴西克鲁塞罗 (1942–1967)", Symbol: "BRZ"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "土库曼斯坦马纳特 (1993–2009)", Symbol: "TMM"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "乌拉圭比索", Symbol: "UYU"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "测试货币代码", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "阿富汗尼", Symbol: "AFN"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "法国金法郎", Symbol: ""}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "南斯拉夫硬第纳尔 (1966–1990)", Symbol: "YUD"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "旧塞尔维亚第纳尔", Symbol: "CSD"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "埃及镑", Symbol: "EGP"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "柬埔寨瑞尔", Symbol: "KHR"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "摩洛哥迪拉姆", Symbol: "MAD"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "旧苏丹镑", Symbol: "SDP"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "冰岛克朗(1918–1981)", Symbol: "ISJ"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "摩尔多瓦库邦", Symbol: "MDC"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "塞尔维亚第纳尔", Symbol: "RSD"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "玻利维亚诺 (1863–1963)", Symbol: "BOL"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "立陶宛塔咯呐司", Symbol: "LTT"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "俄国卢布 (1991–1998)", Symbol: "RUR"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "塞舌尔卢比", Symbol: "SCR"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "阿根廷比索 (1881–1970)", Symbol: "ARM"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "保加利亚硬列弗", Symbol: "BGL"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "福克兰群岛镑", Symbol: "FKP"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "马耳他镑", Symbol: "MTP"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "阿根廷奥斯特拉尔", Symbol: "ARA"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "阿塞拜疆马纳特", Symbol: "AZN"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "保加利亚社会党列弗", Symbol: "BGM"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "斐济元", Symbol: "FJD"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "海地古德", Symbol: "HTG"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "萨尔瓦多科朗", Symbol: "SVC"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "墨西哥银比索 (1861–1992)", Symbol: "MXP"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "巴布亚新几内亚基那", Symbol: "PGK"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "特立尼达和多巴哥元", Symbol: "TTD"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "太平洋法郎", Symbol: "CFPF"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "欧洲货币联盟", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "阿富汗尼 (1927–2002)", Symbol: ""}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "奥地利先令", Symbol: "ATS"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "西班牙比塞塔", Symbol: "ESP"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "老挝基普", Symbol: "LAK"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "毛里塔尼亚乌吉亚", Symbol: "MRO"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "苏丹第纳尔 (1992–2007)", Symbol: "SDD"}, "WST": ut.Currency{Currency: "WST", DisplayName: "萨摩亚塔拉", Symbol: "WST"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "巴林第纳尔", Symbol: "BHD"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "爱尔兰镑", Symbol: "IEP"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "塔吉克斯坦索莫尼", Symbol: "TJS"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "保加利亚新列弗", Symbol: "BGN"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "阿尔及利亚第纳尔", Symbol: "DZD"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "卢森堡可兑换法郎", Symbol: "LUC"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "津巴布韦元 (2009)", Symbol: "ZWL"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "委内瑞拉玻利瓦尔", Symbol: "VEF"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "南非兰特", Symbol: "ZAR"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "阿尔巴尼亚列克(1946–1965)", Symbol: "ALK"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "比利时法郎（金融）", Symbol: "BEL"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "智利比索", Symbol: "CLP"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "以色列新谢克尔", Symbol: "₪"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "日元", Symbol: "JP¥"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "旧罗马尼亚列伊", Symbol: "ROL"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "东加勒比元", Symbol: "EC$"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "挪威克朗", Symbol: "NOK"}, "USS": ut.Currency{Currency: "USS", DisplayName: "美元（当日）", Symbol: "USS"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "钯", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "赞比亚克瓦查", Symbol: "ZMW"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "波士尼亚-赫塞哥维纳新第纳尔 (1994–1997)", Symbol: "BAN"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "巴西克鲁塞罗 (1993–1994)", Symbol: "BRR"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "瑞士法郎", Symbol: "CHF"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "摩洛哥法郎", Symbol: "MAF"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "马达加斯加法郎", Symbol: "MGF"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "马尔代夫卢菲亚", Symbol: "MVR"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "斯洛文尼亚托拉尔", Symbol: "SIT"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "孟加拉塔卡", Symbol: "BDT"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "玻利维亚诺", Symbol: "BOB"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "巴西雷亚尔", Symbol: "R$"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "白俄罗斯卢布", Symbol: "BYR"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "荷兰盾", Symbol: "NLG"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "菲律宾比索", Symbol: "PHP"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "新加坡元", Symbol: "SGD"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "南斯拉夫改良第纳尔 (1992–1993)", Symbol: "YUR"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "赞比亚克瓦查 (1968–2012)", Symbol: "ZMK"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "巴哈马元", Symbol: "BSD"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "西班牙比塞塔（兑换帐户）", Symbol: "ESB"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "百慕大元", Symbol: "BMD"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "多米尼加比索", Symbol: "DOP"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "英镑", Symbol: "£"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "几内亚比绍比索", Symbol: "GWP"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "冰岛克朗", Symbol: "ISK"}, "USN": ut.Currency{Currency: "USN", DisplayName: "美元（次日）", Symbol: "USN"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "新扎伊尔 (1993–1998)", Symbol: "ZRN"}, "AON": ut.Currency{Currency: "AON", DisplayName: "安哥拉新宽扎 (1990–2000)", Symbol: "AON"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "智利埃斯库多", Symbol: "CLE"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "塞浦路斯镑", Symbol: "CYP"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "几内亚法郎", Symbol: "GNF"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "匈牙利福林", Symbol: "HUF"}, "STD": ut.Currency{Currency: "STD", DisplayName: "圣多美和普林西比多布拉", Symbol: "STD"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "苏联卢布", Symbol: "SUR"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "白俄罗斯新卢布 (1994–1999)", Symbol: "BYB"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "墨西哥比索", Symbol: "MX$"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "南斯拉夫新第纳尔 (1994–2002)", Symbol: "YUM"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "津巴布韦元 (2008)", Symbol: "ZWR"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "加纳塞地", Symbol: "GHS"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "新台币", Symbol: "NT$"}}
