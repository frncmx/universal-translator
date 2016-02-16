package pa

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"IRR": ut.Currency{Currency: "IRR", DisplayName: "ਈਰਾਨੀ ਰਿਆਲ", Symbol: "IRR"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "ਮੋਰੱਕਨ ਦਿਰਹਾਮ", Symbol: "MAD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "ਮਾਲਦੀਵੀ ਰੁਫੀਆ", Symbol: "MVR"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "ਸਰਬੀਆਈ ਦਿਨਾਰ", Symbol: "RSD"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "ਅਰੂਬਨ ਫਲੋਰਿਨ", Symbol: "AWG"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "ਯੂਰੋ", Symbol: "€"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "ਬ੍ਰਿਟਿਸ਼ ਪੌਂਡ", Symbol: "£"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "ਹਾਨਡੂਰਨ ਲੇਮਪਿਰਾ", Symbol: "HNL"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "ਆਈਸਲੈਂਡਿਕ ਕਰੌਨ", Symbol: "ISK"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "ਸੋਵੀਅਤ ਰੂਬਲ", Symbol: "SUR"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "ਟੌਂਗਨ ਪੈਂਗਾ", Symbol: "TOP"}, "USD": ut.Currency{Currency: "USD", DisplayName: "ਯੂ.ਐਸ. ਡਾਲਰ", Symbol: "US$"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "ਅਰਜਨਟੀਨੀ ਅਸਟਰਾਲ", Symbol: "ARA"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਕਰੁਜ਼ਿਰੋਸ (1990–1993)", Symbol: "BRE"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "ਬੇਲੀਜ਼ ਡਾਲਰ", Symbol: "BZD"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "ਹੈਤੀ ਗੌਰਡੇ", Symbol: "HTG"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "ਵੇਨੇਜ਼ੂਏਲਨ ਬੋਲੀਵਰ", Symbol: "VEF"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "", Symbol: "XBD"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "ਅਰਜਨਟੀਨੀ ਪੇਸੋ ਲੇ (1970–1983)", Symbol: "ARL"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "ਲਾਟਵਿਆਈ ਲਾਟਸ", Symbol: "LVL"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "ਸਾਊਦੀ ਰਿਆਲ", Symbol: "SAR"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ਯੂਕਰੇਨੀਆਈ ਰਿਵਨਿਆ", Symbol: "UAH"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "ਨਿਕਾਰਾਗੁਆਈ ਕੋਰਡੋਬਾ", Symbol: "NIO"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ਯੂਗਾਂਡੀਆਈ ਸ਼ਿਲਿੰਗ", Symbol: "UGX"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "ਉਰੂਗੁਵਾਇਨ ਪੇਸੋ", Symbol: "UYU"}, "YER": ut.Currency{Currency: "YER", DisplayName: "ਯਮਨੀ ਰਿਆਲ", Symbol: "YER"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "ਬਹਿਰੀਨੀ ਦਿਨਾਰ", Symbol: "BHD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "ਕਿਊਬਨ ਬਦਲਣਯੋਗ ਪੇਸੋ", Symbol: "CUC"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "ਇਥੋਪੀਆਈ ਬਿਰ", Symbol: "ETB"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "ਤਨਜ਼ਾਨੀਆਈ ਸ਼ਿਲਿੰਗ", Symbol: "TZS"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "ਸਿਏਰਾ ਲਿਓਨੀਅਨ ਲਿਓਨ", Symbol: "SLL"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "ਗੁਆਨਾਆਈ ਡਾਲਰ", Symbol: "GYD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "ਕਿਰਗਿਸਤਾਨੀ ਸੋਮ", Symbol: "KGS"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "ਸੋਲੋਮਨ ਆਈਲੈਂਡਸ ਡਾਲਰ", Symbol: "SBD"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "ਸੂਡਾਨੀ ਪੌਂਡ", Symbol: "SDG"}, "RON": ut.Currency{Currency: "RON", DisplayName: "ਰੋਮਾਨੀਆਈ ਲੇਉ", Symbol: "RON"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "ਸਿੰਗਾਪੁਰ ਡਾਲਰ", Symbol: "SGD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "ਤੁਰਕਮੇਨਿਸਤਾਨੀ ਮਾਨਤ", Symbol: "TMT"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "ਅਰਜਨਟੀਨੀ ਪੇਸੋ (1881–1970)", Symbol: "ARM"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "ਬਾਹਾਮੀਅਨ ਡਾਲਰ", Symbol: "BSD"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "ਚੀਨੀ ਯੁਆਨ", Symbol: "CN¥"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "ਨਾਰਵੇਜੀਆਈ ਕਰੌਨ", Symbol: "NOK"}, "INR": ut.Currency{Currency: "INR", DisplayName: "ਭਾਰਤੀ ਰੁਪਇਆ", Symbol: "₹"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "ਜਾਰਡਨ ਦਿਨਾਰ", Symbol: "JOD"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "ਪਨਾਮੇਨੀਅਨ ਬਾਲਬੋਆ", Symbol: "PAB"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "ਸੇਂਟ ਹੇਲੇਨਾ ਪੌਂਡ", Symbol: "SHP"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "ਅਰਜਨਟੀਨੀ ਪੇਸੋ (1983–1985)", Symbol: "ARP"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "ਸਵਿਸ ਫ੍ਰੈਂਕ", Symbol: "CHF"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "ਹਾਂਗ ਕਾਂਗ ਡਾਲਰ", Symbol: "HK$"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "ਇਜ਼ਰਾਈਲੀ ਨਵੀਂ ਸ਼ੇਕੇਲ", Symbol: "₪"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "ਉਰੂਗੁਵਾਇਨ ਪੇਸੋ (1975–1993)", Symbol: "UYP"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "ਫ੍ਰੈਂਕ (CFP)", Symbol: "CFPF"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "ਪਾਕਿਸਤਾਨੀ ਰੁਪਇਆ", Symbol: "PKR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "ਪੋਲੈਂਡੀ ਜ਼ਲੌਟੀ", Symbol: "PLN"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "ਸੋਮਾਲੀ ਸ਼ਿਲਿੰਗ", Symbol: "SOS"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "ਅਫ਼ਗਾਨ ਅਫ਼ਗਾਨੀ", Symbol: "AFN"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "ਬੋਲੀਵੀਆਈ ਬੋਲੀਵੀਅਨੋ (1863–1963)", Symbol: "BOL"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "ਕਿਊਬਨ ਪੇਸੋ", Symbol: "CUP"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "ਇੰਡੋਨੇਸ਼ੀਆਈ ਰੁਪਿਆਹ", Symbol: "IDR"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "ਹੰਗਰੀ ਫੋਰਿੰਟ", Symbol: "HUF"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "ਜਮਾਇਕਨ ਡਾਲਰ", Symbol: "JMD"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "ਉੱਤਰੀ ਕੋਰੀਆਈ ਵੋਨ", Symbol: "KPW"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "ਲਾਓਟਿਆਈ ਕਿਪ", Symbol: "LAK"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "ਬੁਲਗਾਰੀਆਈ ਲੇਵ", Symbol: "BGN"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ਭੂਟਾਨੀ ਐਂਗਲਟ੍ਰਮ", Symbol: "BTN"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "ਚਿਲੀ ਪੇਸੋ", Symbol: "CLP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "ਫ਼ਾਕਲੈਂਡ ਆਈਲੈਂਡਸ ਪੌਂਡ", Symbol: "FKP"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "ਲਾਈਬੀਰੀਆਈ ਡਾਲਰ", Symbol: "LRD"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ਓਮਾਨੀ ਰਿਆਲ", Symbol: "OMR"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "ਦੱਖਣੀ ਅਫਰੀਕੀ ਰੈਂਡ", Symbol: "ZAR"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "ਇਰੀਟ੍ਰਿਆਈ ਨਾਫ਼ਾ", Symbol: "ERN"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "ਸਵਾਜ਼ੀ ਲਾਇਲੈਂਗਨੀ", Symbol: "SZL"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "ਵੇਨੇਜ਼ੂਏਲਨ ਬੋਲੀਵਰ (1871–2008)", Symbol: "VEB"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "ਵਾਨੂਆਟੂ ਵਾਟੂ", Symbol: "VUV"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "ਚੈਕ ਗਣਰਾਜ ਕੋਰੁਨਾ", Symbol: "CZK"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "ਜਰਮਨ ਮਾਰਕ", Symbol: "DEM"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "ਦੱਖਣੀ ਕੋਰੀਆਈ ਵੋਨ", Symbol: "₩"}, "VND": ut.Currency{Currency: "VND", DisplayName: "ਵੀਅਤਨਾਮੀ ਡੋਂਗ", Symbol: "₫"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "ਸੀਰੀਆਈ ਪੌਂਡ", Symbol: "SYP"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "ਉਜ਼ਬੇਕਿਸਤਾਨ ਸੋਮ", Symbol: "UZS"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "ਯੂਰਪੀ ਮੁਦਰਾ ਇਕਾਈ", Symbol: "XEU"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "ਜ਼ਾਮਬੀਆਈ ਕਵਾਚਾ", Symbol: "ZMW"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "ਮੈਕਸੀਕਨ ਪੇਸੋ", Symbol: "MX$"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "ਮੋਜ਼ਾਮਬੀਕਨ ਮੈਟੀਕਲ", Symbol: "MZN"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "ਪਾਪੂਆ ਨਿਊ ਗਿਨੀਆਈ ਕੀਨਾ", Symbol: "PGK"}, "STD": ut.Currency{Currency: "STD", DisplayName: "ਸਾਉ ਟੋਮੀ ਐਂਡ ਪ੍ਰਿੰਸਪੀ ਡੋਬਰਾ", Symbol: "STD"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "ਮੋਰਿਟਾਨੀਆਈ ਊਗੀਆ", Symbol: "MRO"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "ਰੂਸੀ ਰੂਬਲ", Symbol: "RUB"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "", Symbol: "р."}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "ਬੋਲੀਵੀਆਈ ਪੇਸੋ", Symbol: "BOP"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "ਡੈਨਿਸ਼ ਕਰੌਨ", Symbol: "DKK"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "ਗਿਨੀ ਫ੍ਰੈਂਕ", Symbol: "GNF"}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "ਆਇਰਿਸ਼ ਪੌਂਡ", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "ਰਵਾਂਡਨ ਫ੍ਰੈਂਕ", Symbol: "RWF"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "ਸੂਰੀਨਾਮੀ ਡਾਲਰ", Symbol: "SRD"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "ਬੋਲੀਵੀਆਈ ਮਵਡੋਲ", Symbol: "BOV"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "ਜ਼ੀਬੂਤੀਅਨ ਫ੍ਰੈਂਕ", Symbol: "DJF"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "ਸ੍ਰੀਲੰਕਾਈ ਰੁਪਇਆ", Symbol: "LKR"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "ਲਿਥੁਆਨੀਆਈ ਲਿਤਾਸ", Symbol: "LTL"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਕਰੁਜ਼ਿਰੋਸ (1942–1967)", Symbol: "BRZ"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "ਬੋਟਸਵਾਨਾ ਪੁਲਾ", Symbol: "BWP"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "ਕੇਮੈਨ ਆਈਲੈਂਡਸ ਡਾਲਰ", Symbol: "KYD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "ਮੌਰਿਸ਼ੀਆਈ ਰੁਪਇਆ", Symbol: "MUR"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "ਪੇਰੂਵੀਅਨ ਨਿਊਵੋ ਸੋਲ", Symbol: "PEN"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "", Symbol: "XBC"}, "AED": ut.Currency{Currency: "AED", DisplayName: "ਸੰਯੁਕਤ ਅਰਬ ਅਮੀਰਾਤ ਦਿਰਹਾਮ", Symbol: "AED"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "ਕੇਨੇਡਿਆਈ ਡਾਲਰ", Symbol: "CA$"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "ਜਿਬਰਾਲਟਰ ਪੌਂਡ", Symbol: "GIP"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "ਲੈਬਨਾਨੀ ਪੌਂਡ", Symbol: "LBP"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "ਨਾਮੀਬੀਆਈ ਡਾਲਰ", Symbol: "NAD"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "ਫਿਲਿਪੀਨੀ ਪੇਸੋ", Symbol: "PHP"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "ਨੀਦਰਲੈਂਡਸ ਐਂਟੀਲੀਅਨ ਗਿਲਡਰ", Symbol: "ANG"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਕਰੁਜ਼ਿਰੋਸ (1993–1994)", Symbol: "BRR"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "ਕੋਸਟਾ ਰੀਕਨ ਕੋਲਨ", Symbol: "CRC"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "ਜਪਾਨੀ ਯੇਨ", Symbol: "JP¥"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "ਤੁਰਕੀ ਲੀਰਾ", Symbol: "TRY"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "ਨਵਾਂ ਤਾਇਵਾਨ ਡਾਲਰ", Symbol: "NT$"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "ਅਜ਼ਰਬਾਈਜਾਨ ਮਾਨਤ", Symbol: "AZN"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਰੀਅਲ", Symbol: "R$"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "ਫ਼ਿਜ਼ੀ ਡਾਲਰ", Symbol: "FJD"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "ਮੋਲਡੋਵਨ ਲੇਉ", Symbol: "MDL"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "ਪੱਛਮੀ ਅਫ਼ਰੀਕੀ (CFA) ਫ੍ਰੈਂਕ", Symbol: "CFA"}, "BND": ut.Currency{Currency: "BND", DisplayName: "ਬਰੂਨੇਈ ਡਾਲਰ", Symbol: "BND"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "ਗੈਂਬੀਆਈ ਦਲਾਸੀ", Symbol: "GMD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "ਕੀਨੀਆਈ ਸ਼ਿਲਿੰਗ", Symbol: "KES"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "ਪੂਰਬੀ ਕੈਰੇਬੀਅਨ ਡਾਲਰ", Symbol: "EC$"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "ਮਿਸਰੀ ਪੌਂਡ", Symbol: "EGP"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "", Symbol: "XBA"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "ਯੂਰਪੀ ਵਿੱਤੀ ਇਕਾਈ", Symbol: "XBB"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "ਮਾਲਾਗਾਸੀ ਅਰਾਇਰੀ", Symbol: "MGA"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "ਨਿਊਜ਼ੀਲੈਂਡ ਡਾਲਰ", Symbol: "NZ$"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "ਫ੍ਰੈਂਕ (CFA BEAC)", Symbol: "FCFA"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "ਬੋਸਨੀਆ-ਹਰਜ਼ੇਗੋਵੀਨਾ ਬਦਲਣਯੋਗ ਮਾਰਕ", Symbol: "BAM"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "ਬੋਲੀਵੀਅਨ ਬੋਲੀਵੀਅਨੋ", Symbol: "BOB"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "ਕਾਂਗੋਲੀਜ਼ ਫ੍ਰੈਂਕ", Symbol: "CDF"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "ਅਲਜੀਰਿਆਈ ਦਿਨਾਰ", Symbol: "DZD"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "ਨੇਪਾਲੀ ਰੁਪਇਆ", Symbol: "NPR"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ਘਾਨਾਈ ਸੇਡੀ", Symbol: "GHS"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "ਇਜ਼ਰਾਈਲੀ ਪੌਂਡ", Symbol: "ILP"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "ਲੀਬੀਅਨ ਦਿਨਾਰ", Symbol: "LYD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "ਮੇਕਾਨੀ ਪਟਾਕਾ", Symbol: "MOP"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "", Symbol: "UYI"}, "WST": ut.Currency{Currency: "WST", DisplayName: "ਸਾਮੋਆਈ ਤਾਲਾ", Symbol: "WST"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "ਬਰਮੂਡਾ ਡਾਲਰ", Symbol: "BMD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "ਮਾਲਾਵੀਆਈ ਕਵਾਚਾ", Symbol: "MWK"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "ਕਤਰੀ ਰਿਆਲ", Symbol: "QAR"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "ਸਵੀਡਿਸ਼ ਕਰੋਨਾ", Symbol: "SEK"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "ਮਲੇਸ਼ੀਆਈ ਰਿੰਗਿਟ", Symbol: "MYR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "ਪੈਰਾਗੁਵਾਇਨ ਗੁਆਰਾਨੀ", Symbol: "PYG"}, "THB": ut.Currency{Currency: "THB", DisplayName: "ਥਾਈ ਬਾਹਤ", Symbol: "฿"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "ਵੀਅਤਨਾਮੀ ਡੋਂਗ (1978–1985)", Symbol: "VNN"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਨਿਊ ਕਰੁਜ਼ਾਡੂ (1989–1990)", Symbol: "BRN"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "ਬੇਲਾਰੂਸੀ ਰੂਬਲ", Symbol: "BYR"}, "COP": ut.Currency{Currency: "COP", DisplayName: "ਕੋਲੰਬਿਆਈ ਪੇਸੋ", Symbol: "COP"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "ਇਰਾਕੀ ਦਿਨਾਰ", Symbol: "IQD"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "ਚਾਂਦੀ", Symbol: "XAG"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "ਮੰਗੋਲੀਆਈ ਤੁਗਰਿਕ", Symbol: "MNT"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "ਨਾਇਜੀਰੀਆਈ ਨਾਇਰਾ", Symbol: "NGN"}, "TND": ut.Currency{Currency: "TND", DisplayName: "ਟਿਉਨੀਸ਼ੀਆਈ ਦਿਨਾਰ", Symbol: "TND"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "ਸੋਨਾ", Symbol: "XAU"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "ਅਰਜਨਟੀਨੀ ਪੇਸੋ", Symbol: "ARS"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਨਿਊ ਕਰੁਜ਼ਿਰੋਸ (1967–1986)", Symbol: "BRB"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "ਕਜ਼ਾਖਸਤਾਨੀ ਤੇਂਗੇ", Symbol: "KZT"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "ਮਿਆਂਮਾਰ ਕਿਆਤ", Symbol: "MMK"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "ਡੌਮਿਨਿਕਨ ਪੇਸੋ", Symbol: "DOP"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "ਗੁਆਟੇਮਾਲਾ ਕੁਏਟਜ਼ਲ", Symbol: "GTQ"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "ਮੈਕਡੋਨੀਆਈ ਡੇਨਾਰ", Symbol: "MKD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "ਕੰਬੋਡੀਆਈ ਰੀਅਲ", Symbol: "KHR"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "ਤਾਜਿਕਿਸਤਾਨੀ ਸੋਮੋਨੀ", Symbol: "TJS"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "ਕਰੋਏਸ਼ੀਆਈ ਕੁਨਾ", Symbol: "HRK"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "ਕੋਮੋਰੀਅਨ ਫ੍ਰੈਂਕ", Symbol: "KMF"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "ਟ੍ਰਿਨੀਡਾਡ ਅਤੇ ਟੋਬਾਗੋ ਡਾਲਰ", Symbol: "TTD"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "ਅਲਬਾਨੀਆਈ ਲੇਕ", Symbol: "ALL"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "ਅਰਮੀਨੀਆਈ ਦਰਮ", Symbol: "AMD"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "ਅੰਗੋਲਾ ਕਵਾਂਜਾ", Symbol: "AOA"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "ਅਗਿਆਤ ਮੁਦਰਾ", Symbol: "XXX"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "ਬੰਗਲਾਦੇਸ਼ੀ ਟਕਾ", Symbol: "BDT"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "ਕੁਵੈਤੀ ਦਿਨਾਰ", Symbol: "KWD"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "ਸੇਸ਼ਲਸ ਰੁਪਇਆ", Symbol: "SCR"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "ਕੇਪ ਵਰਡੀਅਨ ਸਕੂਡੋ", Symbol: "CVE"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "ਜਾਰਜੀਆਈ ਲਾਰੀ", Symbol: "GEL"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "ਦੱਖਣੀ ਸੂਡਾਨੀ ਪੌਂਡ", Symbol: "SSP"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ਆਸਟ੍ਰੇਲੀਆਈ ਡਾਲਰ", Symbol: "A$"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "ਬਾਰਬਾਡੀਅਨ ਡਾਲਰ", Symbol: "BBD"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "ਬੁਰੁੰਡੀਆਈ ਫ੍ਰੈਂਕ", Symbol: "BIF"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "ਬ੍ਰਾਜ਼ੀਲੀਆਈ ਕਰੁਜ਼ਾਡੂ (1986–1989)", Symbol: "BRC"}}
