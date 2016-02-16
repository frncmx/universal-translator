package hu

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"BHD": ut.Currency{Currency: "BHD", DisplayName: "bahreini dinár", Symbol: "BHD"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "bahamai dollár", Symbol: "BSD"}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Francia frank", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Brazil cruzado novo (1989–1990)", Symbol: ""}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR euro", Symbol: ""}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Mali frank", Symbol: ""}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Ukrán karbovanec", Symbol: ""}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Dél-afrikai rand (pénzügyi)", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "gibraltári font", Symbol: "GIP"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Máltai líra", Symbol: ""}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Venezuelai bolivar (1871–2008)", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Spanyol peseta", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Szent Ilona-i font", Symbol: "SHP"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Európai kontó egység (XBD)", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Angolai kwanza reajustado (1995–1999)", Symbol: ""}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Ecuadori Unidad de Valor Constante (UVC)", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "indiai rúpia", Symbol: "INR"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Sierra Leone-i leone", Symbol: "SLL"}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Uruguay-i peso (1975–1993)", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "USA-dollár", Symbol: "USD"}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Andorrai peseta", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "EAE-dirham", Symbol: "AED"}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Holland forint", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "orosz rubel (1991–1998)", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "üzbegisztáni szum", Symbol: "UZS"}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "örmény dram", Symbol: "AMD"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "angolai kwanza", Symbol: "AOA"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "bermudai dollár", Symbol: "BMD"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Mexikói ezüst peso (1861–1992)", Symbol: ""}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "azerbajdzsáni manat", Symbol: "AZN"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "eritreai nakfa", Symbol: "ERN"}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Grúz kupon larit", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "hongkongi dollár", Symbol: "HKD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "lett lats", Symbol: "LVL"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "maldív-szigeteki rufiyaa", Symbol: "MVR"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "ukrán hrivnya", Symbol: "UAH"}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Zairei zaire", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Perui inti", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ausztrál dollár", Symbol: "AUD"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "japán jen", Symbol: "¥"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "mauritániai ouguiya", Symbol: "MRO"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "malajziai ringgit", Symbol: "MYR"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "ruandai frank", Symbol: "RWF"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Ugandai shilling (1966–1987)", Symbol: ""}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Európai kontó egység (XBC)", Symbol: ""}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Belga frank (pénzügyi)", Symbol: ""}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Brazi cruzado (1986–1989)", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "kínai jüan", Symbol: "CNY"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "marokkói dirham", Symbol: "MAD"}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Szlovén tolar", Symbol: ""}, "TND": ut.Currency{Currency: "TND", DisplayName: "tunéziai dínár", Symbol: "TND"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Bolíviai mvdol", Symbol: ""}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Guineai syli", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "mongóliai tugrik", Symbol: "MNT"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Mexikói Unidad de Inversion (UDI)", Symbol: ""}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Európai monetáris egység", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "észak-koreai won", Symbol: "KPW"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "mozambiki metikális", Symbol: "MZN"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ugandai shilling", Symbol: "UGX"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Angolai kwanza (1977–1990)", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "mauritiusi rúpia", Symbol: "MUR"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "nicaraguai córdoba", Symbol: "NIO"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Jugoszláv kemény dínár", Symbol: ""}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "arubai florin", Symbol: "AWG"}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "spanyol peseta (konvertibilis kontó)", Symbol: ""}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Ghánai cedi (1979–2007)", Symbol: ""}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "luxemburgi pénzügyi frank", Symbol: ""}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Szudáni font (1957–1998)", Symbol: ""}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Zimbabwei dollár (2009)", Symbol: ""}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "horvát kuna", Symbol: "HRK"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "lengyel zloty", Symbol: "PLN"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "szudáni font", Symbol: "SDG"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Argentín peso (1983–1985)", Symbol: ""}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "belize-i dollár", Symbol: "BZD"}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Észt korona", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "kazahsztáni tenge", Symbol: "KZT"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "nepáli rúpia", Symbol: "NPR"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Német márka", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "fidzsi dollár", Symbol: "FJD"}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Litvániai talonas", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Lett rubel", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Francia UIC-frank", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Guinea-Bissaui peso", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "hodurasi lempira", Symbol: "HNL"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Lesothoi loti", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "madagaszkári ariary", Symbol: "MGA"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "paraguayi guarani", Symbol: "PYG"}, "STD": ut.Currency{Currency: "STD", DisplayName: "São Tomé és Príncipe-i dobra", Symbol: "STD"}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Zambiai kwacha (1968–2012)", Symbol: "ZMK"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "argentin peso", Symbol: "ARS"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "burundi frank", Symbol: "BIF"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Francia arany frank", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "vietnami dong", Symbol: "VND"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "kanadai dollár", Symbol: "CAD"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Máltai font", Symbol: ""}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "román lej (1952–2006)", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "tajvani új dollár", Symbol: "TWD"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "magyar forint", Symbol: "Ft"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "namíbiai dollár", Symbol: "NAD"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ománi rial", Symbol: "OMR"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "szomáli shilling", Symbol: "SOS"}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Bolíviai peso", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "kubai konvertibilis peso", Symbol: "CUC"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "mianmari kyat", Symbol: "MMK"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "dél-koreai won", Symbol: "KRW"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "szerb dínár", Symbol: "RSD"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Szovjet rubel", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "izlandi korona", Symbol: "ISK"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "salamon-szigeteki dollár", Symbol: "SBD"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "szváziföldi lilangeni", Symbol: "SZL"}, "WST": ut.Currency{Currency: "WST", DisplayName: "nyugat-szamoai tala", Symbol: "WST"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "afgán afghani", Symbol: "AFN"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Cape Verde-i escudo", Symbol: "CVE"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Portugál guinea escudo", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Luxemburgi frank", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "türkmenisztáni manat (1993–2009)", Symbol: ""}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Zimbabwei dollár (1980–2008)", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Belga frank (konvertibilis)", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "türkmenisztáni manat", Symbol: "TMT"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "tongai paanga", Symbol: "TOP"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET tőke", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Zimbabwei dollár (2008)", Symbol: ""}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "barbadosi dollár", Symbol: "BBD"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "guatemalai quetzal", Symbol: "GTQ"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "kuvaiti dínár", Symbol: "KWD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "nigériai naira", Symbol: "NGN"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "pápua új-guineai kina", Symbol: "PGK"}, "COU": ut.Currency{Currency: "COU", DisplayName: "Unidad de Valor Real", Symbol: ""}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "grúz lari", Symbol: "GEL"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Görög drachma", Symbol: ""}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Szudáni dínár (1992–2007)", Symbol: ""}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "albán lek", Symbol: "ALL"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Ezüst", Symbol: ""}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "kelet-karibi dollár", Symbol: "XCD"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "bolgár új leva", Symbol: "BGN"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "botswanai pula", Symbol: "BWP"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "jamaicai dollár", Symbol: "JMD"}, "KES": ut.Currency{Currency: "KES", DisplayName: "kenyai shilling", Symbol: "KES"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "makaói pataca", Symbol: "MOP"}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Lengyel zloty (1950–1995)", Symbol: ""}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "szaúdi riyal", Symbol: "SAR"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Brazil cruzeiro (1993–1994)", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "litvániai litas", Symbol: "LTL"}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "rhodéziai dollár", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "venezuelai bolivar", Symbol: "VEF"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "CFA frank BCEAO", Symbol: "CFA"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Brazi cruzeiro novo (1967–1986)", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "thai baht", Symbol: "THB"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "török líra", Symbol: "TRY"}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Brazil cruzeiro (1990–1993)", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "iraki dínár", Symbol: "IQD"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Mozambik metical", Symbol: ""}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Tádzsikisztáni rubel", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Jugoszláv új dínár", Symbol: ""}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Belga frank", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Burmai kyat", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "új-zélandi dollár", Symbol: "NZD"}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Fehérorosz új rubel (1994–1999)", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "kubai peso", Symbol: "CUP"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "falkland-szigeteki font", Symbol: "FKP"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "luxemburgi konvertibilis frank", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Szlovák korona", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "uruguay-i peso", Symbol: "UYU"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "vanuatui vatu", Symbol: "VUV"}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Arany", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Jemeni dínár", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "dán korona", Symbol: "DKK"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "brit font", Symbol: "GBP"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "kambodzsai riel", Symbol: "KHR"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "panamai balboa", Symbol: "PAB"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "orosz rubel", Symbol: "RUB"}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Madagaszkári frank", Symbol: ""}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "haiti gourde", Symbol: "HTG"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "izraeli új sékel", Symbol: "ILS"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "iráni rial", Symbol: "IRR"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "moldován lei", Symbol: "MDL"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "svéd korona", Symbol: "SEK"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "dél-szudáni font", Symbol: "SSP"}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Tesztelési pénznemkód", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Angolai új kwanza (1990–2000)", Symbol: ""}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Ecuadori sucre", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "perui sol nuevo", Symbol: "PEN"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Osztrák schilling", Symbol: ""}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Horvát dínár", Symbol: ""}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Európai kompozit egység", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "zambiai kwacha", Symbol: "ZMW"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "comorei frank", Symbol: "KMF"}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Marokkói frank", Symbol: ""}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "fülöp-szigeteki peso", Symbol: "PHP"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Zairei új zaire", Symbol: ""}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "holland antilláki forint", Symbol: "ANG"}, "BND": ut.Currency{Currency: "BND", DisplayName: "brunei dollár", Symbol: "BND"}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Egyenlítői-guineai ekwele guineana", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Salvadori colón", Symbol: ""}, "USN": ut.Currency{Currency: "USN", DisplayName: "USA dollár (következő napi)", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Argentín austral", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Ciprusi font", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "egyiptomi font", Symbol: "EGP"}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Finn markka", Symbol: ""}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "guineai frank", Symbol: "GNF"}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "európai pénznemegység", Symbol: ""}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "afgán afghani (1927–2002)", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Ír font", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "palládium", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "spanyol peseta (A–kontó)", Symbol: ""}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Portugál escudo", Symbol: ""}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Timori escudo", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "bhutáni ngultrum", Symbol: "BTN"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "cseh korona", Symbol: "CZK"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Srí Lanka-i rúpia", Symbol: "LKR"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Perui sol", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Suriname-i gulden", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA frank BEAC", Symbol: "FCFA"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "indonéz rúpia", Symbol: "IDR"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "pakisztáni rúpia", Symbol: "PKR"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Trinidad és Tobago-i dollár", Symbol: "TTD"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Uruguayi peso en unidades indexadas", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "ismeretlen pénznem", Symbol: ""}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "bosznia-hercegovinai konvertibilis márka", Symbol: "BAM"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "bolíviai boliviano", Symbol: "BOB"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "dzsibuti frank", Symbol: "DJF"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "líbiai dínár", Symbol: "LYD"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "dominikai peso", Symbol: "DOP"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "mexikói peso", Symbol: "MXN"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "tanzániai shilling", Symbol: "TZS"}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Bolgár kemény leva", Symbol: ""}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Costa Rica-i colon", Symbol: "CRC"}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "szerb dinár", Symbol: ""}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Csehszlovák kemény korona", Symbol: ""}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "guyanai dollár", Symbol: "GYD"}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Izraeli font", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "szíriai font", Symbol: "SYP"}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Special Drawing Rights", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "jemeni rial", Symbol: "YER"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Bosznia-hercegovinai dínár (1992–1994)", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "svájci frank", Symbol: "CHF"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR frank", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Chilei unidades de fomento", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "ghánai cedi", Symbol: "GHS"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "laoszi kip", Symbol: "LAK"}, "USS": ut.Currency{Currency: "USS", DisplayName: "USA dollár (aznapi)", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "brazil real", Symbol: "BRL"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "chilei peso", Symbol: "CLP"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "algériai dínár", Symbol: "DZD"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "kirgizisztáni szom", Symbol: "KGS"}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Nikaraguai cordoba", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "szingapúri dollár", Symbol: "SGD"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "platina", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "dél-afrikai rand", Symbol: "ZAR"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "bangladesi taka", Symbol: "BDT"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Olasz líra", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "jordániai dínár", Symbol: "JOD"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "libanoni font", Symbol: "LBP"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "katari rial", Symbol: "QAR"}, "RON": ut.Currency{Currency: "RON", DisplayName: "román lej", Symbol: "RON"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "tádzsikisztáni somoni", Symbol: "TJS"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "azerbajdzsáni manat (1993–2006)", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "etiópiai birr", Symbol: "ETB"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "kajmán-szigeteki dollár", Symbol: "KYD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "macedon dínár", Symbol: "MKD"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "norvég korona", Symbol: "NOK"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "seychelle-szigeteki rúpia", Symbol: "SCR"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "suriname-i dollár", Symbol: "SRD"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "török líra (1922–2005)", Symbol: ""}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Kelet-Német márka", Symbol: ""}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Jugoszláv konvertibilis dínár", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "kongói frank", Symbol: "CDF"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "gambiai dalasi", Symbol: "GMD"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "csendes-óceáni valutaközösségi frank", Symbol: "CFPF"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "fehérorosz rubel", Symbol: "BYR"}, "COP": ut.Currency{Currency: "COP", DisplayName: "kolumbiai peso", Symbol: "COP"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "euró", Symbol: "EUR"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "libériai dollár", Symbol: "LRD"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "malawi kwacha", Symbol: "MWK"}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Mozambik escudo", Symbol: ""}}
