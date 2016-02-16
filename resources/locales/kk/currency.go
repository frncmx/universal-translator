package kk

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"ANG": ut.Currency{Currency: "ANG", DisplayName: "Нидерланд антиль гульдені", Symbol: "ANG"}, "THB": ut.Currency{Currency: "THB", DisplayName: "Тай баты", Symbol: "฿"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "Қытай юані", Symbol: "CN¥"}, "ERN": ut.Currency{Currency: "ERN", DisplayName: "Эритрея накфасы", Symbol: "ERN"}, "MDL": ut.Currency{Currency: "MDL", DisplayName: "Молдова лейі", Symbol: "MDL"}, "SLL": ut.Currency{Currency: "SLL", DisplayName: "Сьерра-Леоне леонесы", Symbol: "SLL"}, "ALL": ut.Currency{Currency: "ALL", DisplayName: "Албания легі", Symbol: "ALL"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "Литва литы", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "Мексика песосы", Symbol: "MX$"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "Пәкістан рупиясы", Symbol: "PKR"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "Доминикан песосы", Symbol: "DOP"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "Фиджи доллары", Symbol: "FJD"}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "Мальдив руфиясы", Symbol: "MVR"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "Аргентина песосы", Symbol: "ARS"}, "AWG": ut.Currency{Currency: "AWG", DisplayName: "Аруба флорині", Symbol: "AWG"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "Иран риалы", Symbol: "IRR"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "Шығыс Кариб доллары", Symbol: "EC$"}, "BIF": ut.Currency{Currency: "BIF", DisplayName: "Бурунди франкы", Symbol: "BIF"}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "Кувейт динары", Symbol: "KWD"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "Лаос кипі", Symbol: "LAK"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "Либерия доллары", Symbol: "LRD"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "Ангола кванзасы", Symbol: "AOA"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "Чехия кронасы", Symbol: "CZK"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "Моңғолия тугригі", Symbol: "MNT"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "Соломон аралдары доллары", Symbol: "SBD"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "Маврикий рупиясы", Symbol: "MUR"}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "Ливия динары", Symbol: "LYD"}, "MOP": ut.Currency{Currency: "MOP", DisplayName: "Макао патакасы", Symbol: "MOP"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "Швеция кроны", Symbol: "SEK"}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "Ауғанстан афганиі", Symbol: "AFN"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "Алжир динары", Symbol: "DZD"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "Венгрия форинты", Symbol: "HUF"}, "INR": ut.Currency{Currency: "INR", DisplayName: "Үндістан рупиясы", Symbol: "₹"}, "AZN": ut.Currency{Currency: "AZN", DisplayName: "Әзірбайжан манаты", Symbol: "AZN"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "Гвинея франкы", Symbol: "GNF"}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "Панама бальбоасы", Symbol: "PAB"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "Уругвай песосы", Symbol: "UYU"}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "Катар риалы", Symbol: "QAR"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "Грузия лариі", Symbol: "GEL"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "Иордания динары", Symbol: "JOD"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "Малайзия ринггиті", Symbol: "MYR"}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "Оман риалы", Symbol: "OMR"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "Оңтүстік Судан фунты", Symbol: "SSP"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "КФП франкы", Symbol: "CFPF"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "Замбия квачасы", Symbol: "ZMW"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "Ботсвана пуласы", Symbol: "BWP"}, "HTG": ut.Currency{Currency: "HTG", DisplayName: "Гаити гурды", Symbol: "HTG"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "Индонезия рупиясы", Symbol: "IDR"}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "Солтүстік Корея воны", Symbol: "KPW"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "Тонга паангасы", Symbol: "TOP"}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "Өзбекстан сомы", Symbol: "UZS"}, "AED": ut.Currency{Currency: "AED", DisplayName: "Біріккен Араб Әмірліктерінің дирхамы", Symbol: "AED"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "Куба айырбасталмалы песосы", Symbol: "CUC"}, "MZN": ut.Currency{Currency: "MZN", DisplayName: "Мозамбик метикалы", Symbol: "MZN"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "Сауд Арабиясының риалы", Symbol: "SAR"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "Шри-Ланка рупиясы", Symbol: "LKR"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "Жаңа Тайван доллары", Symbol: "NT$"}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "Танзания шиллингі", Symbol: "TZS"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "Белиз доллары", Symbol: "BZD"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "Дат кроны", Symbol: "DKK"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "Гонконг доллары", Symbol: "HK$"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "Гондурас лемпираcы", Symbol: "HNL"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "Қазақстан теңгесі", Symbol: "₸"}, "SOS": ut.Currency{Currency: "SOS", DisplayName: "Сомали шиллингі", Symbol: "SOS"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "КФА ВЕАС франкы", Symbol: "FCFA"}, "TND": ut.Currency{Currency: "TND", DisplayName: "Тунис динары", Symbol: "TND"}, "DJF": ut.Currency{Currency: "DJF", DisplayName: "Джибути франкы", Symbol: "DJF"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "Британия фунты", Symbol: "£"}, "MWK": ut.Currency{Currency: "MWK", DisplayName: "Малави квачасы", Symbol: "MWK"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "Норвегия кроны", Symbol: "NOK"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "Австралия доллары", Symbol: "A$"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "Барбадос доллары", Symbol: "BBD"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "Боливия боливианосы", Symbol: "BOB"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "Мысыр фунты", Symbol: "EGP"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "Филиппин песосы", Symbol: "PHP"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "Парагвай гуараниі", Symbol: "PYG"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "Сирия фунты", Symbol: "SYP"}, "USD": ut.Currency{Currency: "USD", DisplayName: "АҚШ доллары", Symbol: "$"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "Канада доллары", Symbol: "КД$"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "Фолкленд аралдарының фунты", Symbol: "FKP"}, "MRO": ut.Currency{Currency: "MRO", DisplayName: "Мавритания угиясы", Symbol: "MRO"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "Жаңа Зеландия доллары", Symbol: "NZ$"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "Оңтүстік Африка рэнді", Symbol: "ZAR"}, "BHD": ut.Currency{Currency: "BHD", DisplayName: "Бахрейн динары", Symbol: "BHD"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "Бермуд доллары", Symbol: "БД"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "Чили песосы", Symbol: "CLP"}, "CVE": ut.Currency{Currency: "CVE", DisplayName: "Кабо-Верде эскудосы", Symbol: "CVE"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "Исландия кронасы", Symbol: "ISK"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "Жапония иенасы", Symbol: "JP¥"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "Малагаси ариариы", Symbol: "MGA"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "Никарагуа кордобасы", Symbol: "NIO"}, "BND": ut.Currency{Currency: "BND", DisplayName: "Бруней доллары", Symbol: "BND"}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "Бутан нгултрумы", Symbol: "BTN"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "Куба песосы", Symbol: "CUP"}, "GMD": ut.Currency{Currency: "GMD", DisplayName: "Гамбия даласиі", Symbol: "GMD"}, "RSD": ut.Currency{Currency: "RSD", DisplayName: "Сербия динары", Symbol: "RSD"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "Ресей рублі", Symbol: "₽"}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "Сейшель рупиясы", Symbol: "SCR"}, "SDG": ut.Currency{Currency: "SDG", DisplayName: "Судан фунты", Symbol: "SDG"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "Багам доллары", Symbol: "BSD"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "Ямайка доллары", Symbol: "JMD"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "Кайман аралдары доллары", Symbol: "KYD"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "Нигерия найрасы", Symbol: "NGN"}, "WST": ut.Currency{Currency: "WST", DisplayName: "Самоа таласы", Symbol: "WST"}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "Уганда шиллингі", Symbol: "UGX"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "Гайана доллары", Symbol: "GYD"}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "Ирак динары", Symbol: "IQD"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "Камбоджа риелі", Symbol: "KHR"}, "PGK": ut.Currency{Currency: "PGK", DisplayName: "Папуа - Жаңа Гвинея кинасы", Symbol: "PGK"}, "RON": ut.Currency{Currency: "RON", DisplayName: "Румыния лейі", Symbol: "RON"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "КФА ВСЕАО франкы", Symbol: "CFA"}, "BGN": ut.Currency{Currency: "BGN", DisplayName: "Болгария леві", Symbol: "BGN"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "Беларусь рублі", Symbol: "BYR"}, "CDF": ut.Currency{Currency: "CDF", DisplayName: "Конго франкы", Symbol: "CDF"}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "Перу жаңа солі", Symbol: "PEN"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "Бразилия реалы", Symbol: "R$"}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "Швейцария франкы", Symbol: "CHF"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "Марокко дирхамы", Symbol: "MAD"}, "MKD": ut.Currency{Currency: "MKD", DisplayName: "Македония динары", Symbol: "MKD"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "Әулие Елена аралы фунты", Symbol: "SHP"}, "SZL": ut.Currency{Currency: "SZL", DisplayName: "Свазиленд лилангениі", Symbol: "SZL"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "Украина гривнасы", Symbol: "UAH"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "Венесуэла боливары", Symbol: "VEF"}, "COP": ut.Currency{Currency: "COP", DisplayName: "Колумбия песосы", Symbol: "COP"}, "ETB": ut.Currency{Currency: "ETB", DisplayName: "Эфиопия быры", Symbol: "ETB"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "Латвия латы", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "Польша злотасы", Symbol: "PLN"}, "VUV": ut.Currency{Currency: "VUV", DisplayName: "Вануату ватуы", Symbol: "VUV"}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "Белгісіз валюта", Symbol: ""}, "AMD": ut.Currency{Currency: "AMD", DisplayName: "Армения драмы", Symbol: "AMD"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "Коста-Рика колоны", Symbol: "CRC"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "Израиль жаңа шекелі", Symbol: "₪"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "Ливан фунты", Symbol: "LBP"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "Суринам доллары", Symbol: "SRD"}, "STD": ut.Currency{Currency: "STD", DisplayName: "Сант-Томе мен Принсипи добрасы", Symbol: "STD"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "Түрік лирасы", Symbol: "TRY"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "Бангладеш такасы", Symbol: "BDT"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "Гибралтар фунты", Symbol: "GIP"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "Гватемала кетсалі", Symbol: "GTQ"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "Непал рупиясы", Symbol: "NPR"}, "AZM": ut.Currency{Currency: "AZM", DisplayName: "Азербайджан манаты (1993–2006)", Symbol: ""}, "KGS": ut.Currency{Currency: "KGS", DisplayName: "Қырғызстан сомы", Symbol: "KGS"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "Мьянма кьяты", Symbol: "MMK"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "Руанда франкы", Symbol: "RWF"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "Тринидад пен Тобаго доллары", Symbol: "TTD"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "Сингапур доллары", Symbol: "SGD"}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "Тәжікстан сомониі", Symbol: "TJS"}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "Еуро", Symbol: "€"}, "GHS": ut.Currency{Currency: "GHS", DisplayName: "Гана седиі", Symbol: "GHS"}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "Хорватия кунасы", Symbol: "HRK"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "Намибия доллары", Symbol: "NAD"}, "VND": ut.Currency{Currency: "VND", DisplayName: "Вьетнам донгы", Symbol: "₫"}, "TMT": ut.Currency{Currency: "TMT", DisplayName: "Түрікменстан манаты", Symbol: "TMT"}, "YER": ut.Currency{Currency: "YER", DisplayName: "Йемен риалы", Symbol: "YER"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "Босния және Герцеговина айырбасталмалы маркасы", Symbol: "BAM"}, "KES": ut.Currency{Currency: "KES", DisplayName: "Кения шиллингі", Symbol: "KES"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "Комор аралдары франкы", Symbol: "KMF"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "Оңтүстік Корея воны", Symbol: "₩"}}
