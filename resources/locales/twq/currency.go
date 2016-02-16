package twq

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"UGX": ut.Currency{Currency: "UGX", DisplayName: "Uganda Šiiliŋ", Symbol: ""}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Alžeeri Dinar", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Jaapoŋ Yen", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Mooritaani Ugiya", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Saudiya Riyal", Symbol: ""}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Somaali Šiiliŋ", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Angoola Kwanza", Symbol: ""}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Rwanda Fraŋ", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Seešel Rupii", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "Ameriki Dollar", Symbol: ""}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Jibuuti Fraŋ", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA Fraŋ (BEAC)", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "Tunizi Dinar", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Burundi Fraŋ", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Kapuver Escudo", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Gaana Šiidi", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Indu Rupii", Symbol: ""}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mozambik Metikal", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Sao Tome nda Prinsipe Dobra", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Zambi Kwaca", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Bahareen Dinar", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Suudaŋ Pund", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Hawasa Afriki Rand", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambi Kwaca (1968–2012)", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Swisu Fraŋ", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Eero", Symbol: ""}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Komoor Fraŋ", Symbol: ""}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Leezoto Loti", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Seŋ Helena Fraŋ", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Sinwa Yuan Renminbi", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Gambi Dalasi", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Malaawi Kwaca", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eritree Nafka", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Mooris Rupii", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA Fraŋ (BCEAO)", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwe Dollar", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leeon", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Kongo Fraŋ", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Ecioopi Birr", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "Keeniya Šiiliŋ", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Malgaaši Fraŋ", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Botswaana Pund", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Liberia Dollar", Symbol: ""}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Naamibi Dollar", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naajiriya Neera", Symbol: ""}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni", Symbol: ""}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Suudaŋ Dinar", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "Laaraw Immaara Margantey Dirham", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Britin Pund", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Ostraali Dollar", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Misra Pund", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Ginee Fraŋ", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Liibi Dinar", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Tanzaani Šiiliŋ", Symbol: ""}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Kanaada Dollar", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Maarok Dirham", Symbol: ""}}
