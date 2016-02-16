package eu

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BZD": ut.Currency{Currency: "BZD", DisplayName: "Belizeko dolarra", Symbol: "BZD"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Kolonbiako pesoa", Symbol: "COP"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Salomon uharteetako dolarra", Symbol: "SBD"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Surinameko dolarra", Symbol: "SRD"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Armeniako drama", Symbol: "AMD"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Euroa", Symbol: "€"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somaliako txelina", Symbol: "SOS"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Ricako colona", Symbol: "CRC"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Kubako pesoa", Symbol: "CUP"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Moldaviako leua", Symbol: "MDL"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Myanmarreko kyata", Symbol: "MMK"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Peruko sol berria", Symbol: "PEN"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Uzbekistango soma", Symbol: "UZS"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Afrika erdialdeko frankoa BEAC", Symbol: "FCFA"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambiako kwacha", Symbol: "ZMW"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Arubeko florina", Symbol: "AWG"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Bruneiko dolarra", Symbol: "BND"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malawiko kwacha", Symbol: "MWK"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vanuatuko vatua", Symbol: "VUV"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Dominikar Errepublikako pesoa", Symbol: "DOP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Ganbiako dalasia", Symbol: "GMD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Laoseko kipa", Symbol: "LAK"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Letoniako latsa", Symbol: "LVL"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Omango riala", Symbol: "OMR"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzaniako txelina", Symbol: "TZS"}, "USD": ut.Currency{Currency: "USD", DisplayName: "AEBetako dolarra", Symbol: "US$"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Samoko tala", Symbol: "WST"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Hungariako florina", Symbol: "HUF"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Sri Lankako errupia", Symbol: "LKR"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunisiako dinarra", Symbol: "TND"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Moneta ezezaguna", Symbol: ""}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Bielorrusiako errubloa", Symbol: "BYR"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Marokoko dirhama", Symbol: "MAD"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tajikistango somonia", Symbol: "TJS"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Tongako Paʻanga", Symbol: "TOP"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Bhutango ngultruma", Symbol: "BTN"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Ghanako cedia", Symbol: "GHS"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Kuwaiteko dinarra", Symbol: "KWD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Mongoliako tugrika", Symbol: "MNT"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Mexikoko pesoa", Symbol: "MX$"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Sudango libera", Symbol: "SDG"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Karibe ekialdeko dolarra", Symbol: "EC$"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seychelleetako errupia", Symbol: "SCR"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Afrika ekialdeko frankoa BCEAO", Symbol: "CFA"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afganistango afghania", Symbol: "AFN"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Bermudetako dolarra", Symbol: "BMD"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Suitzako frankoa", Symbol: "CHF"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Algeriako dinarra", Symbol: "DZD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kanbodiako riela", Symbol: "KHR"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Paraguaiko guarania", Symbol: "PYG"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Ruandako frankoa", Symbol: "RWF"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Bosnia-Herzegovinako marko trukakorra", Symbol: "BAM"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Etiopiako birra", Symbol: "ETB"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Gineako frankoa", Symbol: "GNF"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Hong Kongeko dolarra", Symbol: "HK$"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Libiako dinarra", Symbol: "LYD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Maurizio uharteetako errupia", Symbol: "MUR"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Errumaniako leua", Symbol: "RON"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Singapurreko dolarra", Symbol: "SGD"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Arabiar Emirerri Batuetako dirhama", Symbol: "AED"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Egiptoko libera", Symbol: "EGP"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "", Symbol: "₧"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Malaysiako ringgita", Symbol: "MYR"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Azerbaijango manata", Symbol: "AZN"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Zeelanda Berriko dolarra", Symbol: "NZ$"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Ugandako txelina", Symbol: "UGX"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angolako kwanza", Symbol: "AOA"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongoko frankoa", Symbol: "CDF"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Gibraltarreko libera", Symbol: "GIP"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Israeleko sheqel berria", Symbol: "₪"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Ipar Koreako wona", Symbol: "KPW"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Hego Koreako wona", Symbol: "₩"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Mozambikeko metikala", Symbol: "MZN"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Norvegiako koroa", Symbol: "NOK"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Thailandiako bahta", Symbol: "฿"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Yemengo riala", Symbol: "YER"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Irango riala", Symbol: "IRR"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Islandiako koroa", Symbol: "ISK"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Macanako pataca", Symbol: "MOP"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Maldivetako rufiyaa", Symbol: "MVR"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Swaziko lilangenia", Symbol: "SZL"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Djibouteko frankoa", Symbol: "DJF"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Kirgizistango soma", Symbol: "KGS"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Errusiako errubloa", Symbol: "RUB"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Suediako koroa", Symbol: "SEK"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP frankoa", Symbol: "CFPF"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Albaniako leka", Symbol: "ALL"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Cabo Verdeko ezkutua", Symbol: "CVE"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Libanoko libera", Symbol: "LBP"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Serbiako dinarra", Symbol: "RSD"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Txekiar Errepublikako koroa", Symbol: "CZK"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritreako nakfa", Symbol: "ERN"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Jamaikako dolarra", Symbol: "JMD"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Kaiman uharteetako dolarra", Symbol: "KYD"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Namibiako dolarra", Symbol: "NAD"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leonako leona", Symbol: "SLL"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Sao Tomeko eta Principeko dobra", Symbol: "STD"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Taiwango dolar berria", Symbol: "NT$"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Txinako yuana", Symbol: "CN¥"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Fijiko dolarra", Symbol: "FJD"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Papua Ginea Berriko kina", Symbol: "PGK"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Hegoafrikako randa", Symbol: "ZAR"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Holandarren Antilletako guilderra", Symbol: "ANG"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Barbadosetako dolarra", Symbol: "BBD"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mauritaniako ouguiya", Symbol: "MRO"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Argentinako pesoa", Symbol: "ARS"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indiako errupia", Symbol: "₹"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Nepalgo errupia", Symbol: "NPR"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Bosniako pula", Symbol: "BWP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Falkland uharteetako libera", Symbol: "FKP"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Georgiako laria", Symbol: "GEL"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Japoniako yena", Symbol: "JP¥"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Pakistango errupia", Symbol: "PKR"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Hego Sudango libera", Symbol: "SSP"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Vietnameko donga", Symbol: "₫"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Kubako peso trukakorra", Symbol: "CUC"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Danimarkako koroa", Symbol: "DKK"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Indonesiako errupia", Symbol: "IDR"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesothoko lotia", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Nigeriako naira", Symbol: "NGN"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Uruguaiko pesoa", Symbol: "UYU"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komoreetako frankoa", Symbol: "KMF"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberiako dolarra", Symbol: "LRD"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Nikaraguako cordoba", Symbol: "NIO"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Venezuelako bolivarra", Symbol: "VEF"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanadako dolarra", Symbol: "CA$"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Libera esterlina", Symbol: "£"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Lituaniako litasa", Symbol: "LTL"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Filipinetako pesoa", Symbol: "PHP"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Irakeko dinarra", Symbol: "IQD"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Ukrainako hryvnia", Symbol: "UAH"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambiako kwacha (1968–2012)", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Guyanako dolarra", Symbol: "GYD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Kenyako txelina", Symbol: "KES"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Kazakhstango tengea", Symbol: "KZT"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Poloniako zlotya", Symbol: "PLN"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudiko riala", Symbol: "SAR"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Turkmenistango manata", Symbol: "TMT"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviako bolivianoa", Symbol: "BOB"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Guatemalako quetzala", Symbol: "GTQ"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kroaziako kuna", Symbol: "HRK"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Jordaniako dinarra", Symbol: "JOD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Mazedoniako denara", Symbol: "MKD"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Santa Helenako libera", Symbol: "SHP"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Bahametako dolarra", Symbol: "BSD"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Bangladesheko taka", Symbol: "BDT"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundiko frankoa", Symbol: "BIF"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Txileko pesoa", Symbol: "CLP"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Madagaskarreko ariarya", Symbol: "MGA"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Qatarreko riala", Symbol: "QAR"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Turkiako lira", Symbol: "TRY"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Australiako dolarra", Symbol: "A$"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Bulgariako leva", Symbol: "BGN"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahraingo dinarra", Symbol: "BHD"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Brasilgo erreala", Symbol: "R$"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Hondurasko lempira", Symbol: "HNL"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Haitiko gourdea", Symbol: "HTG"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Panamako balboa", Symbol: "PAB"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Siriako libera", Symbol: "SYP"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidadeko eta Tobagoko dolarra", Symbol: "TTD"}}
