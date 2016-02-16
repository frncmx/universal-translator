package ca

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"KWD": ut.Currency{Currency: "KWD", DisplayName: "dinar kuwaitià", Symbol: "KWD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "rupia de les Seychelles", Symbol: "SCR"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "dòlar de Trinitat i Tobago", Symbol: "TTD"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "franc congolès", Symbol: "CDF"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "lliura xipriota", Symbol: "CYP"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "corona estoniana", Symbol: "EEK"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "lari georgià", Symbol: "GEL"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "dòlar de Barbados", Symbol: "BBD"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "dòlar del Banc Popular Xinès", Symbol: ""}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "marc alemany", Symbol: "DEM"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "ruble letó", Symbol: "LVR"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ouguiya maurità", Symbol: "MRO"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "florí neerlandès", Symbol: "NLG"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "dinar serbi", Symbol: "RSD"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "cruzeiro brasiler", Symbol: "BRR"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "zloty polonès (1950–1995)", Symbol: "PLZ"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "lliura de Santa Helena", Symbol: "SHP"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "unitat de moneda europea", Symbol: "XEU"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "franc CFA BCEAO", Symbol: "CFA"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "lev fort búlgar", Symbol: "BGL"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "escut xilè", Symbol: "CLE"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "franc guineà", Symbol: "GNF"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "antic leu romanès", Symbol: "ROL"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "franc UIC francès", Symbol: "XFU"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "lek albanès", Symbol: "ALL"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "peso convertible cubà", Symbol: "CUC"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "kwacha malawià", Symbol: "MWK"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "dòlar namibià", Symbol: "NAD"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "plata", Symbol: "XAG"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "ruble bielorús", Symbol: "BYR"}, "COP": ut.Currency{Currency: "COP", DisplayName: "peso colombià", Symbol: "COP"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "escut portuguès", Symbol: "PTE"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "escut de Timor", Symbol: "TPE"}, "USS": ut.Currency{Currency: "USS", DisplayName: "dòlar dels Estats Units (mateix dia)", Symbol: "USS"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "bolívar veneçolà (1871–2008)", Symbol: "VEB"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "cruzeiro brasiler (1990–1993)", Symbol: "BRE"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "dinar serbi antic", Symbol: "CSD"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "corona islandesa", Symbol: "ISK"}, "KES": ut.Currency{Currency: "KES", DisplayName: "xíling kenyà", Symbol: "KES"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ringgit de Malàisia", Symbol: "MYR"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "rupia nepalesa", Symbol: "NPR"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "kwanza angolès", Symbol: "AOA"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "peso bolivià", Symbol: "BOP"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "real brasiler", Symbol: "BRL"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "kyat birmà", Symbol: "BUK"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "nakfa eritreu", Symbol: "ERN"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "dracma grega", Symbol: "GRD"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "dòlar de Guyana", Symbol: "GYD"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platí", Symbol: "XPT"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "kwanza angolès reajustat (1995–1999)", Symbol: "AOR"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "dòlar de Belize", Symbol: "BZD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "tögrög mongol", Symbol: "MNT"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "dòlar rhodesià", Symbol: "RHD"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "corona eslovaca", Symbol: "SKK"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "nou dinar iugoslau", Symbol: "YUM"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "fons RINET", Symbol: "XRE"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "lev socialista búlgar", Symbol: "BGM"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "lliura egípcia", Symbol: "EGP"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euro", Symbol: "€"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "dòlar fijià", Symbol: "FJD"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "lempira hondurenya", Symbol: "HNL"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "lira turca", Symbol: "TRY"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "xíling tanzà", Symbol: "TZS"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "manat azerbaidjanès (1993–2006)", Symbol: "AZM"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "sucre equatorià", Symbol: "ECS"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "syli guineà", Symbol: "GNS"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "dinar iraquià", Symbol: "IQD"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "dinar jordà", Symbol: "JOD"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "rial de Qatar", Symbol: "QAR"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "florí de les Antilles Neerlandeses", Symbol: "ANG"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "peso argentí", Symbol: "ARS"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "lev búlgar (1879–1952)", Symbol: "BGO"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "peso mexicà", Symbol: "MXN"}, "WST": ut.Currency{Currency: "WST", DisplayName: "tala samoà", Symbol: "WST"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "hrívnia ucraïnesa", Symbol: "UAH"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "florí d’Aruba", Symbol: "AWG"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "birr etíop", Symbol: "ETB"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "rial iranià", Symbol: "IRR"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "franc convertible luxemburguès", Symbol: "LUC"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "dírham marroquí", Symbol: "MAD"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "escut moçambiquès", Symbol: "MZE"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "lilangeni swazi", Symbol: "SZL"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "peso uruguaià (1975–1993)", Symbol: "UYP"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "cupó de lari georgià", Symbol: "GEK"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "kuna croata", Symbol: "HRK"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "rupia de les Maldives", Symbol: "MVR"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "iuan xinès", Symbol: "¥"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "rupia de Sri Lanka", Symbol: "LKR"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "manat turcman", Symbol: "TMT"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "franc CFA BEAC", Symbol: "FCFA"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "xíling austríac", Symbol: "ATS"}, "BND": ut.Currency{Currency: "BND", DisplayName: "dòlar de Brunei", Symbol: "BND"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "franc de Djibouti", Symbol: "DJF"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "dòlar jamaicà", Symbol: "JMD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "lliura libanesa", Symbol: "LBP"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "denar macedoni", Symbol: "MKD"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "rial saudita", Symbol: "SAR"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "peso ley argentí (1970–1983)", Symbol: "ARL"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "peso argentí (1981–1970)", Symbol: "ARM"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "MVDOL bolivià", Symbol: "BOV"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "córdoba nicaragüenca", Symbol: "NIC"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "dram armeni", Symbol: "AMD"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "dinar algerià", Symbol: "DZD"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "lliura de les illes Malvines", Symbol: "FKP"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "franc monegasc", Symbol: "MCF"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "pal·ladi", Symbol: "XPD"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "colón costa-riqueny", Symbol: "CRC"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "riel cambodjà", Symbol: "KHR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "zloty polonès", Symbol: "PLN"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "som uzbek", Symbol: "UZS"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "unitat de compte europea (XBD)", Symbol: "XBD"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "nou xéquel israelià", Symbol: "₪"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "lliura maltesa", Symbol: "MTP"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "inti peruà", Symbol: "PEI"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "xíling somali", Symbol: "SOS"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "lliura del Sudan del Sud", Symbol: "SSP"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "pesseta espanyola (compte convertible)", Symbol: "ESB"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "dòlar neozelandès", Symbol: "NZ$"}, "THB": ut.Currency{Currency: "THB", DisplayName: "baht tailandès", Symbol: "฿"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "franc CFP", Symbol: "CFPF"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgani afganès", Symbol: "AFN"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "lek albanès (1946–1965)", Symbol: ""}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "peso argentí (1983–1985)", Symbol: "ARP"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "corona forta txecoslovaca", Symbol: "CSK"}, "STD": ut.Currency{Currency: "STD", DisplayName: "dobra de São Tomé i Príncipe", Symbol: "STD"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "kwacha zambià (1968–2012)", Symbol: "ZMK"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "franc belga (convertible)", Symbol: "BEC"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "euro WIR", Symbol: "CHE"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "franc WIR", Symbol: "CHW"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "escut de la Guinea Portuguesa", Symbol: "GWE"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "dinar croat", Symbol: "HRD"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "rupia pakistanesa", Symbol: "PKR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "rial omanita", Symbol: "OMR"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "dinar iemenita", Symbol: "YDD"}, "YER": ut.Currency{Currency: "YER", DisplayName: "rial iemenita", Symbol: "YER"}, "AON": ut.Currency{Currency: "AON", DisplayName: "nou kwanza angolès (1990–2000)", Symbol: "AON"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "manat azerbaidjanès", Symbol: "AZN"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "dòlar de les Bermudes", Symbol: "BMD"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "peso filipí", Symbol: "PHP"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "zaire zairès", Symbol: "ZRZ"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "ruble rus (1991–1998)", Symbol: "RUR"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "or", Symbol: "XAU"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "unidad de fomento xilena", Symbol: "CLF"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "pa‘anga tongà", Symbol: "TOP"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "xíling ugandès (1966–1987)", Symbol: "UGS"}, "AED": ut.Currency{Currency: "AED", DisplayName: "dírham de la UEA", Symbol: "AED"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "dinar de Bòsnia i Hercegovina (1992–1994)", Symbol: "BAD"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "antic metical moçambiquès", Symbol: "MZM"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "ruble soviètic", Symbol: "SUR"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "nou dinar de Bòsnia i Hercegovina (1994–1997)", Symbol: "BAN"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "lliura israeliana", Symbol: "ILP"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "franc luxemburguès", Symbol: "LUF"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "nou dòlar de Taiwan", Symbol: "NT$"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "codi reservat per a proves", Symbol: "XTS"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "peso de plata mexicà (1861–1992)", Symbol: "MXP"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "moneda desconeguda", Symbol: "XXX"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "austral argentí", Symbol: "ARA"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "córdoba oro nicaragüenca", Symbol: "NIO"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "drets especials de gir", Symbol: "XDR"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "dinar fort iugoslau", Symbol: "YUD"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "dòlar australià", Symbol: "AU$"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ngultrum de Bhutan", Symbol: "BTN"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "dòlar de Surinam", Symbol: "SRD"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "nou zaire zairès", Symbol: "ZRN"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "ien japonès", Symbol: "JP¥"}, "USN": ut.Currency{Currency: "USN", DisplayName: "dòlar dels Estats Units (dia següent)", Symbol: "USN"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vatu de Vanuatu", Symbol: "VUV"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "antic cruzeiro brasiler", Symbol: "BRZ"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "corona txeca", Symbol: "CZK"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "pesseta espanyola (compte A)", Symbol: "ESA"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "lliura sudanesa", Symbol: "SDG"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "dinar iugoslau reformat (1992–1993)", Symbol: "YUR"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "ekwele de Guinea Equatorial", Symbol: "GQE"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "hwan sud-coreà (1953–1962)", Symbol: "KRH"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "dòlar de les illes Caiman", Symbol: "KYD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "naira nigerià", Symbol: "NGN"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "dinar sudanès", Symbol: "SDD"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "marc convertible de Bòsnia i Hercegovina", Symbol: "BAM"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "nou ruble bielorús (1994–1999)", Symbol: "BYB"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "peso dominicà", Symbol: "DOP"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "lliura esterlina britànica", Symbol: "£"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "lira italiana", Symbol: "ITL"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "tolar eslovè", Symbol: "SIT"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "leone de Sierra Leone", Symbol: "SLL"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "franc belga (financer)", Symbol: "BEL"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "dòlar de Hong Kong", Symbol: "HK$"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "unitat de compte europea (XBC)", Symbol: "XBC"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "franc belga", Symbol: "BEF"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "bolivià", Symbol: "BOB"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "gourde haitià", Symbol: "HTG"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "tenge kazakh", Symbol: "KZT"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lats letó", Symbol: "LVL"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "lliura síria", Symbol: "SYP"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "dinar de Bahrain", Symbol: "BHD"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "antic won sud-coreà", Symbol: "KRO"}, "PES": ut.Currency{Currency: "PES", DisplayName: "sol peruà", Symbol: "PES"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "taka de Bangla Desh", Symbol: "BDT"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "cruzado brasiler", Symbol: "BRC"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "karbóvanets ucraïnès", Symbol: "UAK"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "cruzado novo brasiler", Symbol: "BRN"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "franc malià", Symbol: "MLF"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "rand sud-africà", Symbol: "ZAR"}, "INR": ut.Currency{Currency: "INR", DisplayName: "rupia índia", Symbol: "₹"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "marc finlandès", Symbol: "FIM"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "cedi ghanès", Symbol: "GHS"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "won sud-coreà", Symbol: "₩"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "franc marroquí", Symbol: "MAF"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "pataca de Macau", Symbol: "MOP"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "corona sueca", Symbol: "SEK"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "rand sud-africà (financer)", Symbol: "ZAL"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "boliviano bolivià (1863–1963)", Symbol: "BOL"}, "COU": ut.Currency{Currency: "COU", DisplayName: "unidad de valor real colombiana", Symbol: "COU"}, "RON": ut.Currency{Currency: "RON", DisplayName: "leu romanès", Symbol: "RON"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "kwanza angolès (1977–1991)", Symbol: "AOK"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "franc suís", Symbol: "CHF"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "cupó moldau", Symbol: "MDC"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "leu moldau", Symbol: "MDL"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "franc de Burundi", Symbol: "BIF"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ariary malgaix", Symbol: "MGA"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "dòlar de les illes Salomó", Symbol: "SBD"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "florí de Surinam", Symbol: "SRG"}, "USD": ut.Currency{Currency: "USD", DisplayName: "dòlar dels Estats Units", Symbol: "USD"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "kwacha zambià", Symbol: "ZMW"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "dòlar canadenc", Symbol: "CAD"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "unidad de valor constante (UVC) equatoriana", Symbol: "ECV"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "peso de Guinea Bissau", Symbol: "GWP"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "dong vietnamita (1978–1985)", Symbol: "VNN"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "dòlar zimbabuès (1980–2008)", Symbol: "ZWD"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afgani afganès (1927–2002)", Symbol: "AFA"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "corona danesa", Symbol: "DKK"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "dalasi gambià", Symbol: "GMD"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "fòrint hongarès", Symbol: "HUF"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "franc financer luxemburguès", Symbol: "LUL"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "ruble rus", Symbol: "RUB"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "unitat compensatòria europea", Symbol: "XBA"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "pesseta espanyola", Symbol: "₧"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "kip laosià", Symbol: "LAK"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "dòlar liberià", Symbol: "LRD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "dòlar de Singapur", Symbol: "SGD"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "xíling ugandès", Symbol: "UGX"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "dinar convertible iugoslau", Symbol: "YUN"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "lev búlgar", Symbol: "BGN"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "lliura irlandesa", Symbol: "IEP"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "loti de Lesotho", Symbol: "LSL"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "franc de Ruanda", Symbol: "RWF"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "antiga lliura sudanesa", Symbol: "SDP"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "colón salvadorenc", Symbol: "SVC"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "bolívar veneçolà", Symbol: "VEF"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "peso xilè", Symbol: "CLP"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "franc francès", Symbol: "FRF"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "quetzal guatemalenc", Symbol: "GTQ"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "rupia indonèsia", Symbol: "IDR"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litas lituà", Symbol: "LTL"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "kina de Papua Nova Guinea", Symbol: "PGK"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "lira turca (1922–2005)", Symbol: "TRL"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "pesseta andorrana", Symbol: "ADP"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "dòlar de les Bahames", Symbol: "BSD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "pula de Botswana", Symbol: "BWP"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "cedi ghanès (1979–2007)", Symbol: "GHC"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "lliura de Gibraltar", Symbol: "GIP"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "won nord-coreà", Symbol: "KPW"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "kyat de Myanmar", Symbol: "MMK"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "dòlar zimbabuès (2009)", Symbol: "ZWL"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "dòlar zimbabuès (2008)", Symbol: "ZWR"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "xéquel israelià", Symbol: ""}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "corona islandesa antiga", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "franc malgaix", Symbol: "MGF"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "unidad de inversión (UDI) mexicana", Symbol: "MXV"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "metical moçambiquès", Symbol: "MZN"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "corona noruega", Symbol: "NOK"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "unitat monetària europea", Symbol: "XBB"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "franc de les Comores", Symbol: "KMF"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "talonas lituà", Symbol: "LTT"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "rupia mauriciana", Symbol: "MUR"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "balboa panameny", Symbol: "PAB"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "peso cubà", Symbol: "CUP"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "dinar libi", Symbol: "LYD"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "nou sol peruà", Symbol: "PEN"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "peso uruguaià en unitats indexades", Symbol: "UYI"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "escut de Cap Verd", Symbol: "CVE"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "guaraní paraguaià", Symbol: "PYG"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "manat turcman (1993–2009)", Symbol: "TMM"}, "VND": ut.Currency{Currency: "VND", DisplayName: "dong vietnamita", Symbol: "₫"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "ruble tadjik", Symbol: "TJR"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "somoni tadjik", Symbol: "TJS"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "peso uruguaià", Symbol: "UYU"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "franc or francès", Symbol: "XFO"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "cruzeiro novo brasiler (1967–1986)", Symbol: "BRB"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "marc de l’Alemanya Oriental", Symbol: "DDM"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "som kirguís", Symbol: "KGS"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "denar macedoni (1992–1993)", Symbol: "MKN"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "lira maltesa", Symbol: "MTL"}, "TND": ut.Currency{Currency: "TND", DisplayName: "dinar tunisià", Symbol: "TND"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "dòlar del Carib Oriental", Symbol: "XCD"}}
