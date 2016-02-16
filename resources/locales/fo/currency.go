package fo

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"ALL": ut.Currency{Currency: "ALL", DisplayName: "Albania lek", Symbol: "ALL"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Niðurlonds Karibia gyllin", Symbol: "ANG"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Avstralskur dollari", Symbol: "A$"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Aruba florin", Symbol: "AWG"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Kuba peso (sum kann vekslast)", Symbol: "CUC"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatu vatu", Symbol: "VUV"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad & Tobago dollari", Symbol: "TTD"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venesuela bolívar", Symbol: "VEF"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angola kwanza", Symbol: "AOA"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentina peso", Symbol: "ARS"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Brunei dollari", Symbol: "BND"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Gujana dollari", Symbol: "GYD"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Honduras lempira", Symbol: "HNL"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Pólland zloty", Symbol: "PLN"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "kinesiskur yuan", Symbol: "CN¥"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Meksiko peso", Symbol: "MX$"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibia dollari", Symbol: "NAD"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP frankur", Symbol: "CFPF"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudan pund", Symbol: "SDG"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somalia skillingur", Symbol: "SOS"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongo frankur", Symbol: "CDF"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kroatia kuna", Symbol: "HRK"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "iranskir rials", Symbol: "IRR"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "japanskur yen", Symbol: "JP¥"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Norðurkorea won", Symbol: "KPW"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomonoyggjar dollari", Symbol: "SBD"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Sao Tome & Prinsipi dobra", Symbol: "STD"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Jemen rial", Symbol: "YER"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanada dollari", Symbol: "CA$"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Kosta Rika colón", Symbol: "CRC"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malavi kwacha", Symbol: "MWK"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tonga paʻanga", Symbol: "TOP"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Móritius rupi", Symbol: "MUR"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "St. Helena pund", Symbol: "SHP"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Usbekistan som", Symbol: "UZS"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokko dirham", Symbol: "MAD"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbia dinar", Symbol: "RSD"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Hvítarussland ruble", Symbol: "BYR"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algeria dinar", Symbol: "DZD"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Guinea frankur", Symbol: "GNF"}, "INR": ut.Currency{Currency: "INR", DisplayName: "indiskir rupis", Symbol: "₹"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgisia som", Symbol: "KGS"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laos kip", Symbol: "LAK"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Svasiland lilangeni", Symbol: "SZL"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Miðafrika CFA frankur", Symbol: "FCFA"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "ókent gjaldoyra", Symbol: ""}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkaland liri", Symbol: "TRY"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tansania skillingur", Symbol: "TZS"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armenia dram", Symbol: "AMD"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladesj taka", Symbol: "BDT"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Kekkia koruna", Symbol: "CZK"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgia lari", Symbol: "GEL"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Rumenia leu", Symbol: "RON"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tadsjikistan somoni", Symbol: "TJS"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoa tala", Symbol: "WST"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritrea nakfa", Symbol: "ERN"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lanka rupi", Symbol: "LKR"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Tailand baht", Symbol: "THB"}, "USD": ut.Currency{Currency: "USD", DisplayName: "US dollari", Symbol: "US$"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "unse guld", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Kolombia peso", Symbol: "COP"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominika peso", Symbol: "DOP"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemala quetzal", Symbol: "GTQ"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Madagaskar ariary", Symbol: "MGA"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivoyggjar rufiyaa", Symbol: "MVR"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panama balboa", Symbol: "PAB"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botsvana pula", Symbol: "BWP"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "sveisiskur frankur", Symbol: "CHF"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Irak dinar", Symbol: "IQD"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "íslendsk króna", Symbol: "ISK"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kambodja riel", Symbol: "KHR"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Makedónia denar", Symbol: "MKD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmenistan manat", Symbol: "TMT"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vjetnam dong", Symbol: "₫"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Ísrael new sheqel", Symbol: "₪"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Móritania ouguiya", Symbol: "MRO"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukraina hryvnia", Symbol: "UAH"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afganistan afghani", Symbol: "AFN"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "bretsk pund", Symbol: "£"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepal rupi", Symbol: "NPR"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peru nuevo sol", Symbol: "PEN"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapor dollari", Symbol: "SGD"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haiti gourde", Symbol: "HTG"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komoroyggjar frankur", Symbol: "KMF"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libya dinar", Symbol: "LYD"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nikaragua córdoba", Symbol: "NIO"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Uganda skillingur", Symbol: "UGX"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbados dollari", Symbol: "BBD"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberia dollari", Symbol: "LRD"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "norsk króna", Symbol: "NOK"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinam dollari", Symbol: "SRD"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Suðursudan pund", Symbol: "SSP"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Suðurafrika rand", Symbol: "ZAR"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brasilianskur real", Symbol: "R$"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudiarabia riyal", Symbol: "SAR"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seyskelloyggjar rupi", Symbol: "SCR"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Butan ngultrum", Symbol: "BTN"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Kili peso", Symbol: "CLP"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kasakstan tenge", Symbol: "KZT"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambia dalasi", Symbol: "GMD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguai peso", Symbol: "UYU"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnia-Hersegovina mark (kann vekslast)", Symbol: "BAM"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Caymanoyggjar dollari", Symbol: "KYD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongolia tugrik", Symbol: "MNT"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigeria naira", Symbol: "NGN"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Russland ruble", Symbol: "RUB"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Aserbadjan manat", Symbol: "AZN"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldova leu", Symbol: "MDL"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Oman rial", Symbol: "OMR"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "svensk króna", Symbol: "SEK"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesia rupiah", Symbol: "IDR"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaika dollari", Symbol: "JMD"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Taivan new dollari", Symbol: "NT$"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Eystur Karibia dollari", Symbol: "EC$"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltar pund", Symbol: "GIP"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papua Nýguinea kina", Symbol: "PGK"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "unse palladium", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong Kong dollari", Symbol: "HK$"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filipsoyggjar peso", Symbol: "PHP"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leona leone", Symbol: "SLL"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Sýria pund", Symbol: "SYP"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundi frankur", Symbol: "BIF"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fiji dollari", Symbol: "FJD"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falklandsoyggjar pund", Symbol: "FKP"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Suðurkorea won", Symbol: "₩"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermuda dollari", Symbol: "BMD"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egyptaland pund", Symbol: "EGP"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Katar rial", Symbol: "QAR"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "unse sølv", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Sambia kwacha", Symbol: "ZMW"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Nýsæland dollari", Symbol: "NZ$"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunesia dinar", Symbol: "TND"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgaria lev", Symbol: "BGN"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Belis dollari", Symbol: "BZD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "donsk króna", Symbol: "kr"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Etiopia birr", Symbol: "ETB"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libanon pund", Symbol: "LBP"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Makao pataca", Symbol: "MOP"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Sameindu Emirríkini dirham", Symbol: "AED"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Kuba peso", Symbol: "CUP"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Evra", Symbol: "€"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuvait dinar", Symbol: "KWD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malaisia ringgit", Symbol: "MYR"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Vesturafrika CFA frankur", Symbol: "CFA"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Grønhøvdaoyggjar escudo", Symbol: "CVE"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Gana cedi", Symbol: "GHS"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistan rupi", Symbol: "PKR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguai guarani", Symbol: "PYG"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Ungarn forint", Symbol: "HUF"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mosambik metical", Symbol: "MZN"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Ruanda frankur", Symbol: "RWF"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahamaoyggjar dollari", Symbol: "BSD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "kenjanskur skillingur", Symbol: "KES"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Barein dinar", Symbol: "BHD"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Bolivia boliviano", Symbol: "BOB"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djibuti frankur", Symbol: "DJF"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordan dinar", Symbol: "JOD"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmar (Burma) kyat", Symbol: "MMK"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "unse platin", Symbol: ""}}
