package af

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"MGA": ut.Currency{Currency: "MGA", DisplayName: "Malgassiese ariary", Symbol: "MGA"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongoolse toegrik", Symbol: "MNT"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papoea-Nieu-Guinese kina", Symbol: "PGK"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Poolse zloty", Symbol: "PLN"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katarrese rial", Symbol: "QAR"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzaniese sjieling", Symbol: "TZS"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundiese frank", Symbol: "BIF"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongolese frank", Symbol: "CDF"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Keniaanse sjieling", Symbol: "KES"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldowiese leu", Symbol: "MDL"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Mianmese kyat", Symbol: "MMK"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Oekraïnse hriwna", Symbol: "UAH"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falkland-eilande pond", Symbol: "FKP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambiese dalasi", Symbol: "GMD"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filippynse peso", Symbol: "PHP"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapoer dollar", Symbol: "SGD"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Nuwe Taiwanese dollar", Symbol: "NT$"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Iranse rial", Symbol: "IRR"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Meksikaanse peso", Symbol: "MXN"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guinese syli", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Macedoniese denar", Symbol: "MKD"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Noorse kroon", Symbol: "NOK"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Soedannese pond", Symbol: "SDG"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambiese kwacha", Symbol: "ZMW"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kaap Verdiese escudo", Symbol: "CVE"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algeriese dinar", Symbol: "DZD"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalese roepee", Symbol: "NPR"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Nieu-Seeland dollar", Symbol: "NZ$"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistanse roepee", Symbol: "PKR"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP-frank", Symbol: "CFPF"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Onbekende geldeenheid", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australiese dollar", Symbol: "A$"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egiptiese pond", Symbol: "EGP"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tongaanse pa’anga", Symbol: "TOP"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanese cedi", Symbol: "GHS"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Suid-Soedanese pond", Symbol: "SSP"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Noord-Koreaanse won", Symbol: "KPW"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritaniese ouguiya", Symbol: "MRO"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Amerikaanse dollar", Symbol: "US$"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Chileense peso", Symbol: "CLP"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyanese dollar", Symbol: "GYD"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Hongaarse florint", Symbol: "HUF"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kambodjaanse riel", Symbol: "KHR"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lettiese lats", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Oezbekiese som", Symbol: "UZS"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambiese kwacha (1968–2012)", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armeense dram", Symbol: "AMD"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Hondurese lempira", Symbol: "HNL"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbados-dollar", Symbol: "BBD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Deense kroon", Symbol: "DKK"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaikaanse dollar", Symbol: "JMD"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberiese dollar", Symbol: "LRD"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libiese dinar", Symbol: "LYD"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokkaanse dirham", Symbol: "MAD"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Siriese pond", Symbol: "SYP"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda-dollar", Symbol: "BMD"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Italiaanse lier", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kroatiese kuna", Symbol: "HRK"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgisiese som", Symbol: "KGS"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Suid-Koreaanse won", Symbol: "₩"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesotho loti", Symbol: ""}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Turkse lier (1922–2005)", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviaanse boliviano", Symbol: "BOB"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemalaanse quetzal", Symbol: "GTQ"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghanese cedi (1979–2007)", Symbol: ""}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macaose pataca", Symbol: "MOP"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbeidjaanse manat", Symbol: "AZN"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominikaanse peso", Symbol: "DOP"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Tsjeggiese kroon", Symbol: "CZK"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Lebanese pond", Symbol: "LBP"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Sjinese joean renminbi", Symbol: "CN¥"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Kubaanse omskakelbare peso", Symbol: "CUC"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leoniese leone", Symbol: "SLL"}, "STD": ut.Currency{Currency: "STD", DisplayName: "São Tomé en Príncipe dobra", Symbol: "STD"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Verenigde Arabiese Emirate dirham", Symbol: "AED"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haïtiaanse gourde", Symbol: "HTG"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesiese roepia", Symbol: "IDR"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lankaanse roepee", Symbol: "LKR"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruaanse nuwe sol", Symbol: "PEN"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguaanse guarani", Symbol: "PYG"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saoedi-Arabiese riyal", Symbol: "SAR"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordaniese dinar", Symbol: "JOD"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibiese dollar", Symbol: "NAD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seychellese rupee", Symbol: "SCR"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Ugandese sjieling", Symbol: "UGX"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentynse peso", Symbol: "ARS"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Koeweitse dinar", Symbol: "KWD"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoaanse tala", Symbol: "WST"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afgaanse afgani", Symbol: "AFN"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Swazilandse lilangeni", Symbol: "SZL"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Colombiaanse peso", Symbol: "COP"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Ricaanse colón", Symbol: "CRC"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltarese pond", Symbol: "GIP"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Irakse dinar", Symbol: "IQD"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Yslandse kroon", Symbol: "ISK"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinaamse dollar", Symbol: "SRD"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladesjiese taka", Symbol: "BDT"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahreinse dinar", Symbol: "BHD"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatuse vatu", Symbol: "VUV"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmeense manat", Symbol: "TMT"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunisiese dinar", Symbol: "TND"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nicaraguaanse córdoba", Symbol: "NIO"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Roemeense leu", Symbol: "RON"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Etiopiese birr", Symbol: "ETB"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgiese lari", Symbol: "GEL"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Israeliese nuwe sikkel", Symbol: "₪"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mauritiaanse rupee", Symbol: "MUR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omaanse rial", Symbol: "OMR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbiese dinar", Symbol: "RSD"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadjikse roebel", Symbol: "TJS"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhoetanese ngoeltroem", Symbol: "BTN"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djiboeti frank", Symbol: "DJF"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Sweedse kroon", Symbol: "SEK"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Sint Helena pond", Symbol: "SHP"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazakse tenge", Symbol: "KZT"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigeriese naira", Symbol: "NGN"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrese nakfa", Symbol: "ERN"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fidjiaanse dollar", Symbol: "FJD"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong Kong dollar", Symbol: "HK$"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laosiaanse kip", Symbol: "LAK"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Maleisiese ringgit", Symbol: "MYR"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Viëtnamese dong", Symbol: "₫"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albanese lek", Symbol: "ALL"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnies-Herzegowiniese omskakelbare marka", Symbol: "BAM"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Suid-Afrikaanse rand", Symbol: "R"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkse lier", Symbol: "TRY"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jemenitiese rial", Symbol: "YER"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolese kwanza", Symbol: "AOA"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanadese dollar", Symbol: "CA$"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Comoraanse frank", Symbol: "KMF"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA frank BEAC", Symbol: "FCFA"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Arubaanse floryn", Symbol: "AWG"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Belo-Russiese roebel", Symbol: "BYR"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA frank BCEAO", Symbol: "CFA"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwiese dollar", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russiese roebel", Symbol: "RUB"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwandiese frank", Symbol: "RWF"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Switserse frank", Symbol: "CHF"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japannese jen", Symbol: "JP¥"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litause litas", Symbol: "LTL"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panamese balboa", Symbol: "PAB"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgaarse lev", Symbol: "BGN"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brasilliaanse reaal", Symbol: "R$"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Kubaanse peso", Symbol: "CUP"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Britse pond", Symbol: "£"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indiese roepee", Symbol: "₹"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thaise baht", Symbol: "฿"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Oos-Karibbiese dollar", Symbol: "EC$"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamiaanse dollar", Symbol: "BSD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Beliziese dollar", Symbol: "BZD"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mosambiekse metical", Symbol: "MZN"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomonseilande dollar", Symbol: "SBD"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad en Tobago dollar", Symbol: "TTD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguaanse peso", Symbol: "UYU"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Broeneise dollar", Symbol: "BND"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Cayman-eilande dollar", Symbol: "KYD"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinese frank", Symbol: "GNF"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Malediviese rufia", Symbol: "MVR"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawiese kwacha", Symbol: "MWK"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mosambiekse metical (1980–2006)", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Soedannese pond (1957–1998)", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somaliese sjieling", Symbol: "SOS"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswana pula", Symbol: "BWP"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euro", Symbol: "€"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Nederlands-Antilliaanse gulde", Symbol: "ANG"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venezolaanse bolivar", Symbol: "VEF"}}
