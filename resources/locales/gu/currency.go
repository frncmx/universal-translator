package gu

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"INR": ut.Currency{Currency: "INR", DisplayName: "ભારતીય રૂપિયા", Symbol: "₹"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "મેક્સિકન પેસો", Symbol: "MX$"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "સોલોમન આઇલેંડ્સ ડોલર", Symbol: "SBD"}, "COP": ut.Currency{Currency: "COP", DisplayName: "કોલમ્બિયન પેસો", Symbol: "COP"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "શ્રી લંકન રૂપી", Symbol: "LKR"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "પનામેનિયન બાલ્બોઆ", Symbol: "PAB"}, "THB": ut.Currency{Currency: "THB", DisplayName: "થાઇ બાહ્ત", Symbol: "฿"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "ગયાનિઝ ડોલર", Symbol: "GYD"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "હંગેરીયન ફોરિન્ત", Symbol: "HUF"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "મોલડોવેન લિયુ", Symbol: "MDL"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "માલાવિયન ક્વાચા", Symbol: "MWK"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "ફિલીપાઇન પેસો", Symbol: "PHP"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "યુગાંડન શિલિંગ", Symbol: "UGX"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "ઇસ્ટ કેરિબિયન ડોલર", Symbol: "EC$"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "બ્રાઝિલીયન રિઆલ", Symbol: "R$"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "નાઇજીરિયન નૈરા", Symbol: "NGN"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "સુદાનિઝ પાઉન્ડ", Symbol: "SDG"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "ત્રિનિદાદ અને ટોબેગો ડોલર", Symbol: "TTD"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "તાન્ઝાનિયન શિલિંગ", Symbol: "TZS"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "અજ્ઞાત ચલણ", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "ડોમિનિકન પેસો", Symbol: "DOP"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "કોંગોલિઝ ફ્રેંક", Symbol: "CDF"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "ચેક રીપબ્લિક કોરુના", Symbol: "CZK"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "માલ્દિવિયન રુફિયા", Symbol: "MVR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ઓમાની રિયાલ", Symbol: "OMR"}, "USD": ut.Currency{Currency: "USD", DisplayName: "યુઍસ ડોલર", Symbol: "US$"}, "WST": ut.Currency{Currency: "WST", DisplayName: "સમોઅન તાલા", Symbol: "WST"}, "YER": ut.Currency{Currency: "YER", DisplayName: "યેમેની રિઆલ", Symbol: "YER"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "અફ્ગાન અફ્ગાની", Symbol: "AFN"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "રશિયન રૂબલ", Symbol: "RUB"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "તુર્કિશ લિરા", Symbol: "TRY"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "ડેનિશ ક્રોન", Symbol: "DKK"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "જમૈકિયન ડોલર", Symbol: "JMD"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "મ્યાંમાર ક્યાત", Symbol: "MMK"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "આઇસલેન્ડિક ક્રોના", Symbol: "ISK"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "સ્વીડિશ ક્રોના", Symbol: "SEK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "રોમાનિયન લેઉ", Symbol: "RON"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "સૂરીનામિઝ ડોલર", Symbol: "SRD"}, "STD": ut.Currency{Currency: "STD", DisplayName: "સાઓ ટૉમ એન્ડ પ્રિંસાઇપ ડોબ્રા", Symbol: "STD"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "ટોંગન પ’અંગા", Symbol: "TOP"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "લિથુએનિયન લિતાસ", Symbol: "LTL"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "કેનેડિયન ડૉલર", Symbol: "CA$"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "હાઇટિઇન ગોર્ડ", Symbol: "HTG"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "ઇઝરાયેલી ન્યુ શેકલ", Symbol: "₪"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "કેયમેન આઇલેંડ્સ ડોલર", Symbol: "KYD"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "અઝરબૈજાની મનાત", Symbol: "AZN"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "ક્રોએશિયન ક્યુના", Symbol: "HRK"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "લિબ્યન દિનાર", Symbol: "LYD"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "સાઇરિયન પાઉન્ડ", Symbol: "SYP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "ગેમ્બિયન દલાસી", Symbol: "GMD"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "ઇજિપ્તિયન પાઉન્ડ", Symbol: "EGP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "ફૉકલેન્ડ આઇલેંડ્સ પાઉન્ડ", Symbol: "FKP"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "જાપાનીઝ યેન", Symbol: "JP¥"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "મોઝામ્બિકન મેટિકલ", Symbol: "MZN"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "રવાંડન ફ્રેંક", Symbol: "RWF"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "[CFP] ફ્રેંક", Symbol: "CFPF"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "અરુબન ફ્લોરિન", Symbol: "AWG"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "ઇથિયોપીયન બિર", Symbol: "ETB"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "ફિજિઅન ડોલર", Symbol: "FJD"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "જીબ્રાલ્ટર પાઉન્ડ", Symbol: "GIP"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "લિબેરિયન ડોલર", Symbol: "LRD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "મોરેશીઅન રૂપી", Symbol: "MUR"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "પોલિસ ઝ્લોટી", Symbol: "PLN"}, "BND": ut.Currency{Currency: "BND", DisplayName: "બ્રુનેઇ ડોલર", Symbol: "BND"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "સર્બિયન દિનાર", Symbol: "RSD"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "મધ્ય આફ્રિકન [CFA] ફ્રેંક", Symbol: "FCFA"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "બોસ્નિયા અને હર્ઝેગોવિના રૂપાંતર યોગ્ય માર્ક", Symbol: "BAM"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "ન્યૂઝિલેંડ ડૉલર", Symbol: "NZ$"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "પેરુવિયન ન્યુવો સોલ", Symbol: "PEN"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "પાકિસ્તાની રૂપી", Symbol: "PKR"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "બ્રિટિશ પાઉન્ડ", Symbol: "£"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "કેપ વર્દિયન એસ્કુડો", Symbol: "CVE"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "હોંગ કોંગ ડૉલર", Symbol: "HK$"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "દક્ષિણ સુદાનિઝ પાઉન્ડ", Symbol: "SSP"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "બર્મુડન ડોલર", Symbol: "BMD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "કેન્યેન શિલિંગ", Symbol: "KES"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "ઉત્તર કોરિયન વન", Symbol: "KPW"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "લેબેનિઝ પાઉન્ડ", Symbol: "LBP"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "સિંગાપુર ડૉલર", Symbol: "SGD"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "તાજિકિસ્તાની સોમોની", Symbol: "TJS"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "ન્યુ તાઇવાન ડૉલર", Symbol: "NT$"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ઘાનાઇયન સેડી", Symbol: "GHS"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "હોન્ડ્યુરન લેમ્પિરા", Symbol: "HNL"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "જ્યોર્જિઅન લારી", Symbol: "GEL"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "બેલિઝ ડોલર", Symbol: "BZD"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "કઝાકિસ્તાની ટેંગ", Symbol: "KZT"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "મોરીશેનિયન ઓગુયા", Symbol: "MRO"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "નેપાલિઝ રૂપી", Symbol: "NPR"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "વેનેઝુએલન બોલિવર", Symbol: "VEF"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "દક્ષિણ આફ્રિકી રેંડ", Symbol: "ZAR"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "અલ્બેનિયન લેક", Symbol: "ALL"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "કતારી રિયાલ", Symbol: "QAR"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "સોમાલી શિલિંગ", Symbol: "SOS"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "સ્વાઝી લિલાન્ગેની", Symbol: "SZL"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "ઉરુગ્વેયન પેસો", Symbol: "UYU"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "પશ્ચિમી આફ્રિકન [CFA] ફ્રેંક", Symbol: "CFA"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "ઝામ્બિયન ક્વાચા (1968–2012)", Symbol: ""}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "નિકારાગુઅન કોર્ડોબા", Symbol: "NIO"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "કોસ્ટા રિકન કોલોન", Symbol: "CRC"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "એરિટ્રેયન નક્ફા", Symbol: "ERN"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "ગિનીયન ફ્રેંક", Symbol: "GNF"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "લાતવિયન લેત્સ", Symbol: "LVL"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "વનૌતુ વાતુ", Symbol: "VUV"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "ઝામ્બિયન ક્વાચા", Symbol: "ZMW"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "ચિલિઅન પેસો", Symbol: "CLP"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "જોર્ડનિયન દિનાર", Symbol: "JOD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "લાઓશિયન કિપ", Symbol: "LAK"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "નૉર્વેજિયન ક્રોન", Symbol: "NOK"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "સિએરા લિઓનિઅન લિઓન", Symbol: "SLL"}, "VND": ut.Currency{Currency: "VND", DisplayName: "વિયેતનામીસ ડોંગ", Symbol: "₫"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "ગ્વાટેમાલા કુઇટ્ઝલ", Symbol: "GTQ"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "ઇન્ડોનેશિયન રૂપિયા", Symbol: "IDR"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "પાપુઆ ન્યૂ ગિનીયન કિના", Symbol: "PGK"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "અંગોલિયન ક્વાન્ઝા", Symbol: "AOA"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "સાઉદી રિયાલ", Symbol: "SAR"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "ક્યુબન પેસો", Symbol: "CUP"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "ભુતાનિઝ એંગુલ્ત્રમ", Symbol: "BTN"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "ચાઇનિઝ યુઆન", Symbol: "CN¥"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "યુક્રેનિયન હ્રિવિનિયા", Symbol: "UAH"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "બોલિવિયન બોલિવિયાનો", Symbol: "BOB"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "બેલારુશિયન રૂબલ", Symbol: "BYR"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "દક્ષિણ કોરિયન વન", Symbol: "₩"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "મેસેડોનિયન દિનાર", Symbol: "MKD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "મલેશિયન રિંગ્ગેટ", Symbol: "MYR"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "બાર્બાડિયન ડોલર", Symbol: "BBD"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "અર્મેનિયન ડ્રેમ", Symbol: "AMD"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ઑસ્ટ્રેલિયન ડૉલર", Symbol: "A$"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "બલ્ગેરીયન લેવ", Symbol: "BGN"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "ઇરાકી દિનાર", Symbol: "IQD"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "માલાગેસી અરીઆરી", Symbol: "MGA"}, "AED": ut.Currency{Currency: "AED", DisplayName: "યુનાઈટેડ આરબ અમિરાત દિરહામ", Symbol: "AED"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "યુરો", Symbol: "€"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "ઇરાનિયન રિયાલ", Symbol: "IRR"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "લેસોથો લોતી", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "ટ્યુનિશિયન દિનાર", Symbol: "TND"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "જિબુટિયન ફ્રેંક", Symbol: "DJF"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "બોત્સવાનન પુલા", Symbol: "BWP"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "કિર્ગિસ્તાની સોમ", Symbol: "KGS"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "કેમ્બોડિયન રીઅલ", Symbol: "KHR"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "મોરોક્કન દિરહામ", Symbol: "MAD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "મોંગોલિયન ટગરિક", Symbol: "MNT"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "પરાગ્વેયન ગુઆરાની", Symbol: "PYG"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "સેશેલોઈ રૂપી", Symbol: "SCR"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "બેહરિની દિનાર", Symbol: "BHD"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "ઉઝ્બેકિસ્તાન સોમ", Symbol: "UZS"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "સેંટ હેલેના પાઉન્ડ", Symbol: "SHP"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "બાંગલાદેશી ટાકા", Symbol: "BDT"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "બહામિયન ડોલર", Symbol: "BSD"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "સ્વિસ ફ્રેંક", Symbol: "CHF"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "ક્યુબન રૂપાંતર યોગ્ય પેસો", Symbol: "CUC"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "અર્જેન્ટીના પેસો", Symbol: "ARS"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "માકાનિઝ પતાકા", Symbol: "MOP"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "કુવૈતી દિનાર", Symbol: "KWD"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "બુરુન્ડિયન ફ્રેંક", Symbol: "BIF"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "અલ્જિરિયન દિનાર", Symbol: "DZD"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "કોમોરિઅન ફ્રેંક", Symbol: "KMF"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "નામિબિયન ડોલર", Symbol: "NAD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "તુર્કમેનિસ્તાની મનત", Symbol: "TMT"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "નેધરલેંડ એંટિલિન ગિલ્ડર", Symbol: "ANG"}}
