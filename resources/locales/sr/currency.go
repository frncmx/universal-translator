package sr

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"COU": ut.Currency{Currency: "COU", DisplayName: "Унидад де валоршки реал", Symbol: ""}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Хондурашка лемпира", Symbol: "HNL"}, "PTE": ut.Currency{Currency: "PTE", DisplayName: "Португалски ескудо", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Сиријска фунта", Symbol: "SYP"}, "UYI": ut.Currency{Currency: "UYI", DisplayName: "Уругвајски пезо ен унидадес индексадас", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "CFA франак BEAC", Symbol: "FCFA"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Мјанмарски кјат", Symbol: "MMK"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Панамска балбоа", Symbol: "PAB"}, "SDD": ut.Currency{Currency: "SDD", DisplayName: "Стари судански динар", Symbol: ""}, "USS": ut.Currency{Currency: "USS", DisplayName: "САД долар (исти дан)", Symbol: ""}, "YUR": ut.Currency{Currency: "YUR", DisplayName: "Југословенски реформирани динар", Symbol: ""}, "AOR": ut.Currency{Currency: "AOR", DisplayName: "Анголијска кванза реађустадо (1995–1999)", Symbol: ""}, "ARL": ut.Currency{Currency: "ARL", DisplayName: "Аргентински пезос леј", Symbol: ""}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Боцванска пула", Symbol: "BWP"}, "CHW": ut.Currency{Currency: "CHW", DisplayName: "WIR франак", Symbol: ""}, "EEK": ut.Currency{Currency: "EEK", DisplayName: "Естонска кроон", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Индонежанска рупија", Symbol: "IDR"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Сејшелска рупија", Symbol: "SCR"}, "ANG": ut.Currency{Currency: "ANG", DisplayName: "Холандскоантилски гулден", Symbol: "ANG"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Костарикански колон", Symbol: "CRC"}, "MDC": ut.Currency{Currency: "MDC", DisplayName: "Молдовански купон", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Перуански нови сол", Symbol: "PEN"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Пакистанскa рупиja", Symbol: "PKR"}, "TPE": ut.Currency{Currency: "TPE", DisplayName: "Тиморшки ескудо", Symbol: ""}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "ЦФА франак БЦЕАО", Symbol: "CFA"}, "MKN": ut.Currency{Currency: "MKN", DisplayName: "Стари македонски денар", Symbol: ""}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Зеленортски ескудо", Symbol: "CVE"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Руандски франак", Symbol: "RWF"}, "YUD": ut.Currency{Currency: "YUD", DisplayName: "Југословенски тврди динар", Symbol: ""}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Гамбијски даласи", Symbol: "GMD"}, "MTP": ut.Currency{Currency: "MTP", DisplayName: "Малтешка фунта", Symbol: ""}, "XAU": ut.Currency{Currency: "XAU", DisplayName: "Злато", Symbol: ""}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Тринидад-тобагошки долар", Symbol: "TTD"}, "XRE": ut.Currency{Currency: "XRE", DisplayName: "RINET фонд", Symbol: ""}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Аргентински пезос", Symbol: "ARS"}, "BRC": ut.Currency{Currency: "BRC", DisplayName: "Бразилијски крузадо", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Бутански нгултрум", Symbol: "BTN"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Грузијски лари", Symbol: "GEL"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Руска рубља", Symbol: "RUB"}, "VNN": ut.Currency{Currency: "VNN", DisplayName: "Вијетнамски донг (1978–1985)", Symbol: ""}, "AED": ut.Currency{Currency: "AED", DisplayName: "УАЕ дирхам", Symbol: "AED"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Азербејџански манат (1993–2006)", Symbol: ""}, "ILP": ut.Currency{Currency: "ILP", DisplayName: "Израелска фунта", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Исландска круна", Symbol: "ISK"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Малдивска руфија", Symbol: "MVR"}, "PES": ut.Currency{Currency: "PES", DisplayName: "Перуански сол", Symbol: ""}, "TMM": ut.Currency{Currency: "TMM", DisplayName: "Туркменистански манат (1993–2009)", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Египатска фунта", Symbol: "EGP"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Хрватска куна", Symbol: "HRK"}, "ZWD": ut.Currency{Currency: "ZWD", DisplayName: "Зимбабвеански долар (1980–2008)", Symbol: ""}, "ALK": ut.Currency{Currency: "ALK", DisplayName: "стари албански лек", Symbol: ""}, "GEK": ut.Currency{Currency: "GEK", DisplayName: "Грузијски купон ларит", Symbol: ""}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Гански седи", Symbol: "GHS"}, "ITL": ut.Currency{Currency: "ITL", DisplayName: "Италијанска лира", Symbol: ""}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Македонски денар", Symbol: "MKD"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Парагвајски гварани", Symbol: "PYG"}, "XBB": ut.Currency{Currency: "XBB", DisplayName: "Европска новчана јединица", Symbol: ""}, "YUM": ut.Currency{Currency: "YUM", DisplayName: "Југословенски нови динар", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Авганистански авгани", Symbol: "AFN"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Босанско-херцеговачка конвертибилна марка", Symbol: "КМ"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Белоруска рубља", Symbol: "[BYR]"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Ирански риjал", Symbol: "IRR"}, "MXV": ut.Currency{Currency: "MXV", DisplayName: "Мексички унидад де инверсион (UDI)", Symbol: ""}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Норвешка круна", Symbol: "NOK"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Вијетнамски донг", Symbol: "VND"}, "BOV": ut.Currency{Currency: "BOV", DisplayName: "Боливијски мвдол", Symbol: ""}, "XDR": ut.Currency{Currency: "XDR", DisplayName: "Посебна цртаћа права", Symbol: ""}, "ZMK": ut.Currency{Currency: "ZMK", DisplayName: "Замбијска квача (1968–2012)", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Либеријски долар", Symbol: "LRD"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Румунски леј (1952–2006)", Symbol: "RON"}, "BAN": ut.Currency{Currency: "BAN", DisplayName: "Босанско-херцеговачки нови динар", Symbol: ""}, "BND": ut.Currency{Currency: "BND", DisplayName: "Брунејски долар", Symbol: "BND"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Шриланканскa рупиja", Symbol: "LKR"}, "LUC": ut.Currency{Currency: "LUC", DisplayName: "Луксембуршки конвертибилни франак", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Малезијски рингит", Symbol: "MYR"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Намибијски долар", Symbol: "NAD"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Филипински пезос", Symbol: "PHP"}, "CLE": ut.Currency{Currency: "CLE", DisplayName: "Чилеански ескудо", Symbol: ""}, "CLF": ut.Currency{Currency: "CLF", DisplayName: "Чилеовски унидадес се фоменто", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Чилеански пезос", Symbol: "CLP"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Кенијски шилинг", Symbol: "KES"}, "SDP": ut.Currency{Currency: "SDP", DisplayName: "Стара суданска фунта", Symbol: ""}, "STD": ut.Currency{Currency: "STD", DisplayName: "Саотомска добра", Symbol: "STD"}, "XBD": ut.Currency{Currency: "XBD", DisplayName: "Европска јединица рачуна (XBD)", Symbol: ""}, "YER": ut.Currency{Currency: "YER", DisplayName: "Јеменски риjал", Symbol: "YER"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Хонгконшки долар", Symbol: "HK$"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Кувајтски динар", Symbol: "KWD"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Монголски тугрик", Symbol: "MNT"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Катарски ријал", Symbol: "QAR"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Турска лира", Symbol: "TRY"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Угандски шилинг", Symbol: "UGX"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Источнокарипски долар", Symbol: "EC$"}, "ARP": ut.Currency{Currency: "ARP", DisplayName: "Аргентински пезо (1983–1985)", Symbol: ""}, "BRN": ut.Currency{Currency: "BRN", DisplayName: "Бразилијски нови крузадо", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Бангладешка така", Symbol: "BDT"}, "BEL": ut.Currency{Currency: "BEL", DisplayName: "Белгијски франак (финансијски)", Symbol: ""}, "RHD": ut.Currency{Currency: "RHD", DisplayName: "Родејскидолар", Symbol: ""}, "SRG": ut.Currency{Currency: "SRG", DisplayName: "Суринамски гилдер", Symbol: ""}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Сијералеонски леоне", Symbol: "SLL"}, "BRR": ut.Currency{Currency: "BRR", DisplayName: "Бразилски крузеиро", Symbol: ""}, "GQE": ut.Currency{Currency: "GQE", DisplayName: "Екваторијално-гвинејски еквеле", Symbol: ""}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Марокански дирхам", Symbol: "MAD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Макаоска патака", Symbol: "MOP"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Никарагванска златна кордоба", Symbol: "NIO"}, "ROL": ut.Currency{Currency: "ROL", DisplayName: "Румунски леј", Symbol: ""}, "SKK": ut.Currency{Currency: "SKK", DisplayName: "Словачка круна", Symbol: ""}, "ARA": ut.Currency{Currency: "ARA", DisplayName: "Аргентински аустрал", Symbol: ""}, "BYB": ut.Currency{Currency: "BYB", DisplayName: "Белоруска нова рубља (1994–1999)", Symbol: ""}, "CSD": ut.Currency{Currency: "CSD", DisplayName: "Стари српски динар", Symbol: ""}, "ISJ": ut.Currency{Currency: "ISJ", DisplayName: "Стара исландска круна", Symbol: ""}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Маурицијска рупија", Symbol: "MUR"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Мозамбички метикал", Symbol: "MZN"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Сингапурски долар", Symbol: "SGD"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Сомалијски шилинг", Symbol: "SOS"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Свазилендски лилангени", Symbol: "SZL"}, "XBA": ut.Currency{Currency: "XBA", DisplayName: "Европска композитна јединица", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Tаџихистански сомон", Symbol: "TJS"}, "XPT": ut.Currency{Currency: "XPT", DisplayName: "Платина", Symbol: ""}, "BEC": ut.Currency{Currency: "BEC", DisplayName: "Белгијски франак (конвертибилни)", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Шведска круна", Symbol: "SEK"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Џибутански франак", Symbol: "DJF"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Вануатски вату", Symbol: "VUV"}, "XAG": ut.Currency{Currency: "XAG", DisplayName: "Сребро", Symbol: ""}, "ZRZ": ut.Currency{Currency: "ZRZ", DisplayName: "Заирски заир", Symbol: ""}, "CNX": ut.Currency{Currency: "CNX", DisplayName: "Долар кинеске народне банке", Symbol: ""}, "LVR": ut.Currency{Currency: "LVR", DisplayName: "Латвијска рубља", Symbol: ""}, "MAF": ut.Currency{Currency: "MAF", DisplayName: "Марокански франак", Symbol: ""}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Туркменистански манат", Symbol: "TMT"}, "TRL": ut.Currency{Currency: "TRL", DisplayName: "Турска лира (1922–2005)", Symbol: ""}, "XEU": ut.Currency{Currency: "XEU", DisplayName: "Европска валутна јединица", Symbol: ""}, "BGL": ut.Currency{Currency: "BGL", DisplayName: "Бугарски тврди лев", Symbol: ""}, "CYP": ut.Currency{Currency: "CYP", DisplayName: "Кипарска фунта", Symbol: ""}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Eритрејска накфa", Symbol: "ERN"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Гвинејски франак", Symbol: "GNF"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Јамајчански долар", Symbol: "JMD"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Аустралијски долар", Symbol: "AUD"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Танзанијски шилинг", Symbol: "TZS"}, "YUN": ut.Currency{Currency: "YUN", DisplayName: "Југословенски конвертибилни динар", Symbol: ""}, "BOP": ut.Currency{Currency: "BOP", DisplayName: "Боливијски пезо", Symbol: ""}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Папуанска кина", Symbol: "PGK"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Суданска фунта", Symbol: "SDG"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Узбекистански сом", Symbol: "UZS"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Азербејџански манат", Symbol: "AZN"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Барбадошки долар", Symbol: "BBD"}, "BGO": ut.Currency{Currency: "BGO", DisplayName: "Стари бугарски лев", Symbol: ""}, "BRE": ut.Currency{Currency: "BRE", DisplayName: "Бразилски крузеиро (1990–1993)", Symbol: ""}, "IEP": ut.Currency{Currency: "IEP", DisplayName: "Ирска фунта", Symbol: ""}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Казахстански тенге", Symbol: "KZT"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Гвајански долар", Symbol: "GYD"}, "KRO": ut.Currency{Currency: "KRO", DisplayName: "Стари јужнокорејски вон", Symbol: ""}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Српски динар", Symbol: "RSD"}, "ZWL": ut.Currency{Currency: "ZWL", DisplayName: "Зимбабвеански долар (2009)", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Кинески јуан", Symbol: "CN¥"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Мађарска форинта", Symbol: "HUF"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Нови тајвански долар", Symbol: "NT$"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Канадски долар", Symbol: "CA$"}, "MZM": ut.Currency{Currency: "MZM", DisplayName: "Стари мозамбијски метикал", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Омански ријал", Symbol: "OMR"}, "ZAL": ut.Currency{Currency: "ZAL", DisplayName: "Јужно-афрички ранд (финансијски)", Symbol: ""}, "MZE": ut.Currency{Currency: "MZE", DisplayName: "Мозамбијски ескудо", Symbol: ""}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Тонганска панга", Symbol: "TOP"}, "UAK": ut.Currency{Currency: "UAK", DisplayName: "Украјински карбованети", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Венецуелански боливар", Symbol: "VEF"}, "XFO": ut.Currency{Currency: "XFO", DisplayName: "Француски златни франак", Symbol: ""}, "XFU": ut.Currency{Currency: "XFU", DisplayName: "Француски UIC-франак", Symbol: ""}, "XPD": ut.Currency{Currency: "XPD", DisplayName: "Паладијум", Symbol: ""}, "BUK": ut.Currency{Currency: "BUK", DisplayName: "Бурмански кјат", Symbol: ""}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Израелски нови шекел", Symbol: "₪"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Либанска фунта", Symbol: "LBP"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Тајландски бат", Symbol: "THB"}, "BAD": ut.Currency{Currency: "BAD", DisplayName: "Босанско-Херцеговачки динар", Symbol: ""}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Мадагаскарски ариари", Symbol: "MGA"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Туниски динар", Symbol: "TND"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Бермудски долар", Symbol: "BMD"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Бразилски реал", Symbol: "R$"}, "DEM": ut.Currency{Currency: "DEM", DisplayName: "Немачка марка", Symbol: ""}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Доминикански пезос", Symbol: "DOP"}, "GNS": ut.Currency{Currency: "GNS", DisplayName: "Гвинејски сили", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Непалскa рупиja", Symbol: "NPR"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Јужносуданска фунта", Symbol: "SSP"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Белиски долар", Symbol: "BZD"}, "BRB": ut.Currency{Currency: "BRB", DisplayName: "Бразилски нови крузеиро (1967–1986)", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "Индијска рупија", Symbol: "₹"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Коморски франак", Symbol: "KMF"}, "KRH": ut.Currency{Currency: "KRH", DisplayName: "Јужнокорејски хван", Symbol: ""}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Литвански литас", Symbol: "LTL"}, "AOK": ut.Currency{Currency: "AOK", DisplayName: "Анголијска кванза (1977–1990)", Symbol: ""}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Британска фунта", Symbol: "£"}, "LSL": ut.Currency{Currency: "LSL", DisplayName: "Лесото лоти", Symbol: ""}, "LUF": ut.Currency{Currency: "LUF", DisplayName: "Луксембуршки франак", Symbol: ""}, "NIC": ut.Currency{Currency: "NIC", DisplayName: "Никарагванска кордоба", Symbol: ""}, "XTS": ut.Currency{Currency: "XTS", DisplayName: "Код тестиране валуте", Symbol: ""}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "Руска рубља (1991–1998)", Symbol: ""}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Св. јеленска фунта", Symbol: "SHP"}, "USN": ut.Currency{Currency: "USN", DisplayName: "САД долар (следећи дан)", Symbol: ""}, "ARM": ut.Currency{Currency: "ARM", DisplayName: "Аргентински пезос монедо национал", Symbol: ""}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Бахреински динар", Symbol: "BHD"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Молдавски леј", Symbol: "MDL"}, "MXP": ut.Currency{Currency: "MXP", DisplayName: "Мексички сребрни пезо (1861–1992)", Symbol: ""}, "ADP": ut.Currency{Currency: "ADP", DisplayName: "Андорска пезета", Symbol: ""}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Анголска кванза", Symbol: "AOA"}, "ATS": ut.Currency{Currency: "ATS", DisplayName: "Аустријски шилинг", Symbol: ""}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Чешка круна", Symbol: "CZK"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Фокландска фунта", Symbol: "FKP"}, "LUL": ut.Currency{Currency: "LUL", DisplayName: "Луксембуршки финансијски франак", Symbol: ""}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Малавијска квача", Symbol: "MWK"}, "USD": ut.Currency{Currency: "USD", DisplayName: "Амерички долар", Symbol: "US$"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Албански лек", Symbol: "ALL"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Бахамски долар", Symbol: "BSD"}, "CHE": ut.Currency{Currency: "CHE", DisplayName: "WIR евро", Symbol: ""}, "NLG": ut.Currency{Currency: "NLG", DisplayName: "Холандски гулден", Symbol: ""}, "PEI": ut.Currency{Currency: "PEI", DisplayName: "Перуански инти", Symbol: ""}, "PLZ": ut.Currency{Currency: "PLZ", DisplayName: "Пољски злоти (1950–1995)", Symbol: ""}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Соломонски долар", Symbol: "SBD"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Уругвајски пезос", Symbol: "UYU"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Алжирски динар", Symbol: "DZD"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Јужнокорејски Вон", Symbol: "KRW"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Лаошки кип", Symbol: "LAK"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Мексички пезос", Symbol: "MX$"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Саудијски ријал", Symbol: "SAR"}, "ECV": ut.Currency{Currency: "ECV", DisplayName: "Еквадорски унидад де валор константе", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Ирачки динар", Symbol: "IQD"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Латвијски лати", Symbol: "LVL"}, "ZRN": ut.Currency{Currency: "ZRN", DisplayName: "Заирски нови заир", Symbol: ""}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Боливијски боливијано", Symbol: "BOB"}, "DDM": ut.Currency{Currency: "DDM", DisplayName: "Источно-немачка марка", Symbol: ""}, "ESA": ut.Currency{Currency: "ESA", DisplayName: "Шпанска пезета (рачун)", Symbol: ""}, "ESB": ut.Currency{Currency: "ESB", DisplayName: "Шпанска пезета (конвертибилнирачун)", Symbol: ""}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Етиопски бир", Symbol: "ETB"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Kамбоџански ријел", Symbol: "KHR"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Либијски динар", Symbol: "LYD"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Швајцарски франак", Symbol: "CHF"}, "ECS": ut.Currency{Currency: "ECS", DisplayName: "Еквадорски сакр", Symbol: ""}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Гибралтарска фунта", Symbol: "GIP"}, "GRD": ut.Currency{Currency: "GRD", DisplayName: "Грчка драхма", Symbol: ""}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Јордански динар", Symbol: "JOD"}, "TJR": ut.Currency{Currency: "TJR", DisplayName: "Таџихистанска рубља", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Јерменски драм", Symbol: "AMD"}, "BEF": ut.Currency{Currency: "BEF", DisplayName: "Белгијски франак", Symbol: ""}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Бурундски франак", Symbol: "BIF"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Непозната валута", Symbol: ""}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Фиџијски долар", Symbol: "FJD"}, "MLF": ut.Currency{Currency: "MLF", DisplayName: "Малијански франак", Symbol: ""}, "ZWR": ut.Currency{Currency: "ZWR", DisplayName: "Зимбабвеански долар (2008)", Symbol: ""}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Кубански конвертибилни пезос", Symbol: "CUC"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Хаићански гурд", Symbol: "HTG"}, "MTL": ut.Currency{Currency: "MTL", DisplayName: "Малтешка лира", Symbol: ""}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Суринамски долар", Symbol: "SRD"}, "SUR": ut.Currency{Currency: "SUR", DisplayName: "Совјетска рубља", Symbol: ""}, "SVC": ut.Currency{Currency: "SVC", DisplayName: "Салвадорски колон", Symbol: ""}, "YDD": ut.Currency{Currency: "YDD", DisplayName: "Јеменски динар", Symbol: ""}, "AON": ut.Currency{Currency: "AON", DisplayName: "Анголијска нова кванза (1990–2000)", Symbol: ""}, "BOL": ut.Currency{Currency: "BOL", DisplayName: "Стари боливијски боливијано", Symbol: ""}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Конгоански франак", Symbol: "CDF"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Гватемалски кецал", Symbol: "GTQ"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Самоанска тала", Symbol: "WST"}, "XBC": ut.Currency{Currency: "XBC", DisplayName: "Европска јединица рачуна (XBC)", Symbol: ""}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Мауританијска oгија", Symbol: "MRO"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Арубански флорин", Symbol: "AWG"}, "BRZ": ut.Currency{Currency: "BRZ", DisplayName: "Стари бразилски крузеиро", Symbol: ""}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "Шпанска пезета", Symbol: ""}, "FIM": ut.Currency{Currency: "FIM", DisplayName: "Финска марка", Symbol: ""}, "ILR": ut.Currency{Currency: "ILR", DisplayName: "Стари израелски шекели", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Јапански јен", Symbol: "¥"}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Киргистански сом", Symbol: "KGS"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Новозеландски долар", Symbol: "NZD"}, "UGS": ut.Currency{Currency: "UGS", DisplayName: "Угандски шилинг (1966–1987)", Symbol: ""}, "MGF": ut.Currency{Currency: "MGF", DisplayName: "Малагасијски франак", Symbol: ""}, "COP": ut.Currency{Currency: "COP", DisplayName: "Колумбијски пезос", Symbol: "COP"}, "CSK": ut.Currency{Currency: "CSK", DisplayName: "Чехословачка тврда круна", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Кубански пезос", Symbol: "CUP"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Евро", Symbol: "€"}, "GWE": ut.Currency{Currency: "GWE", DisplayName: "Португалска гвинеја ескудо", Symbol: ""}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Кајмански долар", Symbol: "KYD"}, "MCF": ut.Currency{Currency: "MCF", DisplayName: "Монегаскански франак", Symbol: ""}, "SIT": ut.Currency{Currency: "SIT", DisplayName: "Словеначки толар", Symbol: ""}, "UYP": ut.Currency{Currency: "UYP", DisplayName: "Уругвајски пезо (1975–1993)", Symbol: ""}, "FRF": ut.Currency{Currency: "FRF", DisplayName: "Француски франак", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Пољски злот", Symbol: "PLN"}, "VEB": ut.Currency{Currency: "VEB", DisplayName: "Венецуелански боливар (1871–2008)", Symbol: ""}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Замбијска квача", Symbol: "ZMW"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Бугарски лев", Symbol: "[BGN]"}, "HRD": ut.Currency{Currency: "HRD", DisplayName: "Хрватски динар", Symbol: ""}, "LTT": ut.Currency{Currency: "LTT", DisplayName: "Литвански талонас", Symbol: ""}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Нигеријска наира", Symbol: "NGN"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "CFP франак", Symbol: "CFPF"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Јужноафрички ранд", Symbol: "ZAR"}, "AFA": ut.Currency{Currency: "AFA", DisplayName: "Авганистански авгани (1927–2002)", Symbol: ""}, "BGM": ut.Currency{Currency: "BGM", DisplayName: "Бугарски социјалистички лев", Symbol: ""}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Данска круна", Symbol: "DKK"}, "GHC": ut.Currency{Currency: "GHC", DisplayName: "Гански цеди (1979–2007)", Symbol: ""}, "GWP": ut.Currency{Currency: "GWP", DisplayName: "Гвинеја Бисао Пезо", Symbol: ""}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Севернокорејски вон", Symbol: "KPW"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Украјинска хривња", Symbol: "UAH"}}
