package gd

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"IEP": ut.Currency{Currency: "IEP", DisplayName: "Punnd Èireannach", Symbol: "IEP"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Sol Pearùthach (1863–1965)", Symbol: "PES"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Złoty Pòlainneach", Symbol: "PLN"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Dinar Sèirbeach", Symbol: "RSD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Som Cìorgasach", Symbol: "KGS"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Franc ionmhasail Lugsamburgach", Symbol: "LUL"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Franc Monacach", Symbol: "MCF"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Dolar Singeapòrach", Symbol: "SGD"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Rùbal Sovietach", Symbol: "SUR"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Lari Cairtbheileach", Symbol: "GEL"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Rupee Neapàlach", Symbol: "NPR"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Rial Eamanach", Symbol: "YER"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Dram Airmeineach", Symbol: "AMD"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Dinar Bachraineach", Symbol: "BHD"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Peso Boilibhiach", Symbol: "BOP"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Gourde Haidhteach", Symbol: "HTG"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Dinar Libitheach", Symbol: "LYD"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Dolar Barbadach", Symbol: "BBD"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Peso Uruguaidheach (1975–1993)", Symbol: "UYP"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Rand Afraga a Deas", Symbol: "ZAR"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Rùbal Bealaruiseach", Symbol: "BYR"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Punnd Diobraltarach", Symbol: "GIP"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Talonas Liotuaineach", Symbol: "LTT"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Peseta Spàinnteach (cunntas iompachail)", Symbol: "ESB"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Córdoba Niocaragach (1988–1991)", Symbol: "NIC"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Dolar Shealainn Nuaidh", Symbol: "NZ$"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Dinar Iòrdanach", Symbol: "JOD"}, "MVP": ut.Currency{Currency: "MVP", DisplayName: "Rupee Maladaibheach", Symbol: "MVP"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Naira Nigèiriach", Symbol: "NGN"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Dinar iompachail Iùgoslabhach (1990–1992)", Symbol: "YUN"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Manat Asarbaideànach (1993–2006)", Symbol: "AZM"}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Lev sòisealach Bulgarach", Symbol: "BGM"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Peseta Spàinnteach", Symbol: "ESP"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Rufiyaa Maladaibheach", Symbol: "MVR"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Rùbal Taidigeach", Symbol: "TJR"}, "AON": ut.Currency{Currency: "AON", DisplayName: "Kwanza ùr Angòlach (1990–2000)", Symbol: "AON"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Eòro", Symbol: "€"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Dinar Cròthaiseach", Symbol: "HRD"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Somoni Taidigeach", Symbol: "TJS"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Òr", Symbol: "XAU"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Dolar Astràilianach", Symbol: "A$"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Franc Comorosach", Symbol: "KMF"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Lats Laitbheach", Symbol: "LVL"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zaïre ùr Zaïreach (1993–1998)", Symbol: "ZRN"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Aonad co-dhèanta Eòrpach", Symbol: "XBA"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Còd airgeadra fo dheuchainn", Symbol: "XTS"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Leu Moldobhach", Symbol: "MDL"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Dolar Sìombabuthach (2009)", Symbol: "ZWL"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Florin Arùbach", Symbol: "AWG"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Taka Bangladaiseach", Symbol: "BDT"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Franc Burundaidheach", Symbol: "BIF"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Escudo Mòsaimbiceach", Symbol: "MZE"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Guaraní Paraguaidheach", Symbol: "PYG"}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Peso Gini-Biosothach", Symbol: "GWP"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Punnd Maltach", Symbol: "MTP"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Leu Romàineach (1952–2006)", Symbol: "ROL"}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Cruzado ùr Braisileach (1989–1990)", Symbol: "BRN"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Krone Danmhairgeach", Symbol: "DKK"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Franc Ginitheach", Symbol: "GNF"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Punnd Eilean Naomh Eilidh", Symbol: "SHP"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Paʻanga Tongach", Symbol: "TOP"}, "USN": ut.Currency{Currency: "USN", DisplayName: "Dolar nan SA (an ath–latha)", Symbol: "USN"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Lev cruaidh Bulgarach", Symbol: "BGL"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Nakfa Eartrach", Symbol: "ERN"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Quetzal Guatamalach", Symbol: "GTQ"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Denar Masadonach (1992–1993)", Symbol: "MKN"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Dinar Sudànach (1992–2007)", Symbol: "SDD"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Manat Turcmanach", Symbol: "TMT"}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Pallaideam", Symbol: "XPD"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Dirham nan Iomaratan Arabach Aonaichte", Symbol: "AED"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Cruzeiro Braisileach (1993–1994)", Symbol: "BRR"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Dolar Guidheànach", Symbol: "GYD"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Ngultrum Butànach", Symbol: "BTN"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Colón Costa Rìceach", Symbol: "CRC"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Escudo Tìomorach", Symbol: "TPE"}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Kyat Burmach", Symbol: "BUK"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Peso Cùbach", Symbol: "CUP"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "Franc CFA Meadhan-Afragach", Symbol: "FCFA"}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Dinar Eamanach", Symbol: "YDD"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Dinar cruaidh Iùgoslabhach (1966–1990)", Symbol: "YUD"}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Dinar ùr Iùgoslabhach (1994–2002)", Symbol: "YUM"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Dolar Brùnaigheach", Symbol: "BND"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Cruzeiro Braisileach (1990–1993)", Symbol: "BRE"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Forint Ungaireach", Symbol: "HUF"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Punnd Iosraeleach", Symbol: "ILP"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Kip Làthosach", Symbol: "LAK"}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Franc Lugsamburgach", Symbol: "LUF"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Gulden Duitseach", Symbol: "NLG"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Franc Rubhandach", Symbol: "RWF"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Airgead", Symbol: "XAG"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Kroon Eastoineach", Symbol: "EEK"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Punnd Èipheiteach", Symbol: "EGP"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Tugrik Mongolach", Symbol: "MNT"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Ringgit Malaidheach", Symbol: "MYR"}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Koruna Slòbhacach", Symbol: "SKK"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Dolar Bearmùdach", Symbol: "BMD"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Peso iompachail Cùbach", Symbol: "CUC"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Kupon larit Cairtbheileach", Symbol: "GEK"}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Króna Innis Tìleach (1918–1981)", Symbol: "ISJ"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Kwacha Malabhaidheach", Symbol: "MWK"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Metical Mòsaimbiceach", Symbol: "MZN"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Còir tarraing shònraichte", Symbol: "XDR"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Schilling Ostaireach", Symbol: "ATS"}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Boliviano Boilibhiach (1863–1963)", Symbol: "BOL"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Koruna Seiceach", Symbol: "CZK"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Dinar Cuibhèiteach", Symbol: "KWD"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Dinar Sèirbeach (2002–2006)", Symbol: "CSD"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Pula Botsuanach", Symbol: "BWP"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Rupee Moiriseasach", Symbol: "MUR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Rial Omànach", Symbol: "OMR"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Airgeadra neo-aithnichte", Symbol: "XXX"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Dolar Caimeanach", Symbol: "KYD"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Dolar Fìditheach", Symbol: "FJD"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Kwanza Angòlach (1977–1991)", Symbol: "AOK"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Ouguiya Moratàineach", Symbol: "MRO"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Aonad inbheistidh Meagsagach", Symbol: "MXV"}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Gulden Suranamach", Symbol: "SRG"}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Colón Salbhadorach", Symbol: "SVC"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Baht Tàidheach", Symbol: "฿"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Dinar Tuiniseach", Symbol: "TND"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Dong Bhiet-Namach", Symbol: "₫"}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "Lek Albàineach (1946–1965)", Symbol: "ALK"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Escudo a’ Chip Uaine", Symbol: "CVE"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Cedi Gànach (1979–2007)", Symbol: "GHC"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Lira Eadailteach", Symbol: "ITL"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Rupee Seiseallach", Symbol: "SCR"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Lilangeni Suasaidheach", Symbol: "SZL"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Cruzeiro ùr Braisileach (1967–1986)", Symbol: "BRB"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Aonad fìor-luach Coloimbeach", Symbol: "COU"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Punnd Sasannach", Symbol: "£"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Franc Madagasgarach", Symbol: "MGF"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Lira Turcach (1922–2005)", Symbol: "TRL"}, "XUA": ut.Currency{Currency: "XUA", DisplayName: "Aonad cunntasachd ADB", Symbol: "XUA"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Dolar Sìombabuthach (1980–2008)", Symbol: "ZWD"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Dolar Beilìseach", Symbol: "BZD"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Dirham Morocach", Symbol: "MAD"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Peso Uruguaidheach (aonadan inneacsaichte)", Symbol: "UYI"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Afghani Afghanach (1927–2002)", Symbol: "AFA"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Afghani Afghanach", Symbol: "AFN"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Mark na Gearmailte an Ear", Symbol: "DDM"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Aonad luach chunbhalaich Eacuadorach", Symbol: "ECV"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Litas Liotuaineach", Symbol: "LTL"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Córdoba Niocaragach", Symbol: "NIO"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Leone Siarra Leòmhannach", Symbol: "SLL"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Punnd Sudàin a Deas", Symbol: "SSP"}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Peso Argantaineach (1881–1970)", Symbol: "ARM"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Franc Beilgeach (ionmhasail)", Symbol: "BEL"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "Franc WIR", Symbol: "CHW"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Peso Coloimbeach", Symbol: "COP"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Won Choirèa a Tuath", Symbol: "KPW"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "Franc CFA Afraga an Iar", Symbol: "CFA"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "Aonad maoine RINET", Symbol: "XRE"}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Peso ley Argantaineach (1970–1983)", Symbol: "ARL"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Dalasi Gaimbitheach", Symbol: "GMD"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Dolar Rhodesiach", Symbol: "RHD"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Dobra São Tomé agus Príncipe", Symbol: "STD"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Franc òir Frangach", Symbol: "XFO"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Mvdol Boilibhiach", Symbol: "BOV"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Sucre Eacuadorach", Symbol: "ECS"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Kuna Cròthaiseach", Symbol: "HRK"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Rial Iorànach", Symbol: "IRR"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Won Choirèa a Deas (1945–1953)", Symbol: "KRO"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Krona Suaineach", Symbol: "SEK"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Dolar nan Stàitean Aonaichte", Symbol: "$"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Dong Bhiet-Namach (1978–1985)", Symbol: "VNN"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Kwanza Angòlach", Symbol: "AOA"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Punnd Fàclannach", Symbol: "FKP"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Rupiah Innd-Innseach", Symbol: "IDR"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Dolar Caraibeach earach", Symbol: "EC$"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Kwacha Sàimbitheach (1968–2012)", Symbol: "ZMK"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Rùbal ùr Bealaruiseach (1994–1999)", Symbol: "BYB"}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Inti Pearùthach", Symbol: "PEI"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Escudo Portagaileach", Symbol: "PTE"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Aonad cunntasachd Eòrpach (XBD)", Symbol: "XBD"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Peso Sileach", Symbol: "CLP"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Mark Gearmailteach", Symbol: "DEM"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Dolar Libèireach", Symbol: "LRD"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Peso Meagsagach", Symbol: "MX$"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Peso Filipineach", Symbol: "PHP"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Leu Romàineach", Symbol: "RON"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Peso Uruguaidheach", Symbol: "UYU"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Rand Afraga a Deas (ionmhasail)", Symbol: "ZAL"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Kwacha Sàimbitheach", Symbol: "ZMW"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Escudo Sileach", Symbol: "CLE"}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Aonad cunntasachd Sileach (UF)", Symbol: "CLF"}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Dolar an t-sluagh-bhanca Shìnich", Symbol: "CNX"}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Peseta Spàinnteach (cunntas A)", Symbol: "ESA"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Franc Màilitheach", Symbol: "MLF"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Lira Maltach", Symbol: "MTL"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Rial Catarach", Symbol: "QAR"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Punnd Sudànach", Symbol: "SDG"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Shilling Ugandach (1966–1987)", Symbol: "UGS"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Manat Asarbaideànach", Symbol: "AZN"}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Punnd Cìoprasach", Symbol: "CYP"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Dinar Aildireach", Symbol: "DZD"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Tenge Casachach", Symbol: "KZT"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Cupon Moldobhach", Symbol: "MDC"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Tolar Slòbhaineach", Symbol: "SIT"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Dolar Thrianaid agus Thobago", Symbol: "TTD"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Tala Samothach", Symbol: "WST"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Lev Bulgarach (1879–1952)", Symbol: "BGO"}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Sheqel Iosraeleach (1980–1985)", Symbol: "ILR"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Króna Innis Tìleach", Symbol: "ISK"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Dolar Naimibitheach", Symbol: "NAD"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Sol ùr Pearùthach", Symbol: "PEN"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Rupee Pagastànach", Symbol: "PKR"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Aonad cunntasachd Eòrpach (XBC)", Symbol: "XBC"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "Franc CFP", Symbol: "CFPF"}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Dolar Sìombabuthach (2008)", Symbol: "ZWR"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Dinar ùr Bhosna agus Hearsagobhana (1994–1997)", Symbol: "BAN"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Dolar Canadach", Symbol: "CA$"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Yuan Sìneach", Symbol: "CN¥"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Punnd Leabanach", Symbol: "LBP"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Rupee Sri Lancach", Symbol: "LKR"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Peseta Andorrach", Symbol: "ADP"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Won Choirèa a Deas", Symbol: "₩"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Rùbal Ruiseach", Symbol: "RUB"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Peso Argantaineach", Symbol: "ARS"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Franc Beilgeach", Symbol: "BEF"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Dolar Bathamach", Symbol: "BSD"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Franc iompachail Lugsamburgach", Symbol: "LUC"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Punnd Siridheach", Symbol: "SYP"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Birr Itiopach", Symbol: "ETB"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Markka Fionnlannach", Symbol: "FIM"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Som Usbagach", Symbol: "UZS"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Cruzeiro Braisileach (1942–1967)", Symbol: "BRZ"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Dolar Diameugach", Symbol: "JMD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Pataca Macàthuach", Symbol: "MOP"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Krone Nirribheach", Symbol: "NOK"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Kina Ghini Nuaidh Paputhaiche", Symbol: "PGK"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Dolar Eileanan Sholaimh", Symbol: "SBD"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Shilling Somàilitheach", Symbol: "SOS"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Hryvnia Ucràineach", Symbol: "UAH"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Bolívar Bheinisealach", Symbol: "VEF"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Cedi Gànach", Symbol: "GHS"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Franc Eilbheiseach", Symbol: "CHF"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Franc Diobùtaidheach", Symbol: "DJF"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Dinar Ioràcach", Symbol: "IQD"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Lek Albàineach", Symbol: "ALL"}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Franc Beilgeach (iompachail)", Symbol: "BEC"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Boliviano Boilibhiach", Symbol: "BOB"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Drachma Greugach", Symbol: "GRD"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Platanam", Symbol: "XPT"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Yen Seapanach", Symbol: "JP¥"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Rùbal Ruiseach (1991–1998)", Symbol: "RUR"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Lira Turcach", Symbol: "TRY"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zaïre Zaïreach (1971–1993)", Symbol: "ZRZ"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "Eòro WIR", Symbol: "CHE"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Riel Cambuideach", Symbol: "KHR"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Punnd Sudànach (1957–1998)", Symbol: "SDP"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Bolívar Bheinisealach (1871–2008)", Symbol: "VEB"}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Franc UIC Frangach", Symbol: "XFU"}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Dinar ath-leasaichte Iùgoslabhach (1992–1993)", Symbol: "YUR"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Cruzado Braisileach (1986–1989)", Symbol: "BRC"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Balboa Panamach", Symbol: "PAB"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Peso Argantaineach (1983–1985)", Symbol: "ARP"}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Rùbal Laitbheach", Symbol: "LVR"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Shilling Ugandach", Symbol: "UGX"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Vatu Vanuatuthach", Symbol: "VUV"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Aonad airgid Eòrpach", Symbol: "XBB"}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Kwanza ath-ghleusaichte Angòlach (1995–1999)", Symbol: "AOR"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Dinar Bhosna agus Hearsagobhana (1992–1994)", Symbol: "BAD"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Riyal Sabhdach", Symbol: "SAR"}, "USS": ut.Currency{Currency: "USS", DisplayName: "Dolar nan SA (an aon latha)", Symbol: "USS"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Aonad airgeadra Eòrpach", Symbol: "XEU"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Mark iompachail Bhosna agus Hearsagobhana", Symbol: "BAM"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Real Braisileach", Symbol: "R$"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Franc Congothach", Symbol: "CDF"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Rupee Innseanach", Symbol: "₹"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Shilling Tansanaidheach", Symbol: "TZS"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Karbovanets Ucràineach", Symbol: "UAK"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Gulden Eileanan Aintilia nan Tìrean Ìsle", Symbol: "ANG"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Denar Masadonach", Symbol: "MKD"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Złoty Pòlainneach (1950–1995)", Symbol: "PLZ"}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Austral Argantaineach", Symbol: "ARA"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Dolar Hong Kong", Symbol: "HK$"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Lempira Hondùrach", Symbol: "HNL"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Dolar ùr Taidh-Bhànach", Symbol: "NT$"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Peso Doiminiceach", Symbol: "DOP"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Ekwele Gini Meadhan-Chriosaich", Symbol: "GQE"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Franc Morocach", Symbol: "MAF"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Dolar Suranamach", Symbol: "SRD"}, "XSU": ut.Currency{Currency: "XSU", DisplayName: "Sucre", Symbol: "XSU"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Lev Bulgarach", Symbol: "BGN"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Koruna cruaidh Seic-Slòbhacach", Symbol: "CSK"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Escudo Gini na Portagaile", Symbol: "GWE"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Sheqel ùr Iosraeleach", Symbol: "₪"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Shilling Ceineach", Symbol: "KES"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Peso airgid Meagsagach (1861–1992)", Symbol: "MXP"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Metical Mòsaimbiceach (1980–2006)", Symbol: "MZM"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Franc Frangach", Symbol: "FRF"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Syli Ginitheach", Symbol: "GNS"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Hwan Choirèa a Deas (1953–1962)", Symbol: "KRH"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Loti Leasotach", Symbol: "LSL"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Ariary Madagasgarach", Symbol: "MGA"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Kyat Miànmarach", Symbol: "MMK"}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Manat Turcmanach (1993–2009)", Symbol: "TMM"}}
