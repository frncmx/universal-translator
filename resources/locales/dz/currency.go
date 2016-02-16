package dz

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"CAD": ut.Currency{Currency: "CAD", DisplayName: "ཀེ་ན་ཌ་གི་དངུལ་ ཌོ་ལར", Symbol: "CA$"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "ཌེན་མཱཀ་གི་དངུལ་ ཀྲོན", Symbol: ""}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "བལ་པོའི་དངུལ་ རུ་པི", Symbol: ""}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "ནིའུ་ཛི་ལེནཌ་གི་དངུལ་ ཌོ་ལར", Symbol: "NZ$"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "བར་མུ་ཌ་གི་དངུལ་ ཌོ་ལར", Symbol: ""}, "MVR": ut.Currency{Currency: "MVR", DisplayName: "མཱལ་དིབས་ཀྱི་དངུལ་ རུ་ཕི་ཡ", Symbol: ""}, "PAB": ut.Currency{Currency: "PAB", DisplayName: "པ་ན་མ་གི་དངུལ་ བཱལ་བོ་ཝ", Symbol: ""}, "TJS": ut.Currency{Currency: "TJS", DisplayName: "ཏ་ཇི་ཀིས་ཏཱན་གྱི་དངུལ་ སོ་མོ་ནི", Symbol: ""}, "UGX": ut.Currency{Currency: "UGX", DisplayName: "ཡུ་གྷེན་ཌ་གི་དངུལ་ ཤི་ལིང", Symbol: ""}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "ཀིའུ་བྷ་གི་དངུལ་ པེ་སོ", Symbol: ""}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "ཧོང་ཀོང་གི་དངུལ་ ཌོ་ལར", Symbol: "HK$"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "ཀེམ་བྷོ་ཌི་ཡ་གི་དངུལ་ རི་ཨཱལ", Symbol: ""}, "LYD": ut.Currency{Currency: "LYD", DisplayName: "ལི་བི་ཡ་གི་དངུལ་ ཌི་ན", Symbol: ""}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "མ་ལེ་ཤི་ཡ་གི་དངུལ་ རིང་གིཊ", Symbol: ""}, "PEN": ut.Currency{Currency: "PEN", DisplayName: "པ་རུ་གི་དངུལ་ ནུ་བོ་ སཱོལ", Symbol: ""}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "པ་ཀིས་ཏཱན་གྱི་དངུལ་ རུ་པི", Symbol: ""}, "XXX": ut.Currency{Currency: "XXX", DisplayName: "མ་ཤེས་པའི་དངུལ", Symbol: ""}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "བྷང་ལ་དེཤ་གི་དངུལ་ ཏ་ཀ", Symbol: ""}, "CHF": ut.Currency{Currency: "CHF", DisplayName: "སུ་ཡིས་ཀྱི་དངུལ་ ཕྲངཀ", Symbol: ""}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "ཅི་ལི་གི་དངུལ་ པེ་སོ", Symbol: ""}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "རྒྱ་ནག་གི་དངུལ་ ཡུ་ཝཱན", Symbol: "CN¥"}, "DZD": ut.Currency{Currency: "DZD", DisplayName: "ཨཱལ་ཇི་རི་ཡ་གི་དངུལ་ ཌའི་ནར", Symbol: ""}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "ཨི་ཇིབཊ་གི་དངུལ་ པ་འུནཌ", Symbol: ""}, "IQD": ut.Currency{Currency: "IQD", DisplayName: "ཨི་རཱཀ་གི་དངུལ་ ཌི་ན", Symbol: ""}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "ལཱའི་བེ་རི་ཡ་གི་དངུལ་ ཌོ་ལར", Symbol: ""}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "མེཀ་སི་ཀོ་གི་དངུལ་ པེ་སོ", Symbol: "MX$"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "ནོར་ཝེ་གི་དངུལ་ ཀྲོ་ན", Symbol: ""}, "QAR": ut.Currency{Currency: "QAR", DisplayName: "ཀ་ཊར་གྱི་དངུལ་ རི་ཨཱལ", Symbol: ""}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "སཱའུཐ་ ཨཕ་རི་ཀ་གི་དངུལ་ རཱནད", Symbol: ""}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "ཨཱོས་ཊྲེ་ལི་ཡ་གི་དངུལ་ ཌོ་ལར", Symbol: "AU$"}, "COP": ut.Currency{Currency: "COP", DisplayName: "ཀོ་ལོམ་བྷི་ཡ་གི་དངུལ་ པེ་སོ", Symbol: ""}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "ཨིན་ཌོ་ནེ་ཤི་ཡ་གི་དངུལ་ རུ་པི་ཡ", Symbol: ""}, "INR": ut.Currency{Currency: "INR", DisplayName: "རྒྱ་གར་གྱི་དངུལ་ རུ་པི", Symbol: "₹"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "ལཱ་ཝོས་ཀྱི་དངུལ་ ཀིཔ", Symbol: ""}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "མི་ཡཱན་མར་གྱི་དངུལ་ ཅཱཏ", Symbol: ""}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "བེ་ནི་ཛུ་ཝེ་ལ་གི་དངུལ་ བོ་ལི་བར", Symbol: ""}, "EUR": ut.Currency{Currency: "EUR", DisplayName: "ཡུ་རོ༌དངུལ་", Symbol: "€"}, "SAR": ut.Currency{Currency: "SAR", DisplayName: "སཱཝ་དིའི་དངུལ་ རི་ཡཱལ", Symbol: ""}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "སུའི་ཌེན་གྱི་དངུལ་ ཀྲོ་ན", Symbol: ""}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "སིང་ག་པོར་གྱི་དངུལ་ ཌོ་ལར", Symbol: ""}, "USD": ut.Currency{Currency: "USD", DisplayName: "ཡུ་ཨེས་ ཌོ་ལར", Symbol: "US$"}, "AED": ut.Currency{Currency: "AED", DisplayName: "ཡུ་ནཱའི་ཊེཌ་ ཨ་རབ་ ཨེ་མེ་རེཊས་ཀྱི་དངུལ་ ཌིར་ཧཱམ", Symbol: ""}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "བྲ་ཛིལ་གྱི་དངུལ་ རེ་ཡལ", Symbol: "R$"}, "JOD": ut.Currency{Currency: "JOD", DisplayName: "ཇོར་ཌན་གྱི་དངུལ་ ཌི་ན", Symbol: ""}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "ཊཱའི་ཝཱན་གི་དངུལ ཌོ་ལར", Symbol: "NT$"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "བྲི་ཊིཤ་ པ་འུནཌ་ ཨིས་ཊར་ལིང", Symbol: "£"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "ཀ་ཛགས་ཏཱན་གྱི་དངུལ་ ཏེང་གེ", Symbol: ""}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "པོ་ལེནཌ་ཀྱི་དངུལ ཛ྄ལོ་ཊི", Symbol: ""}, "SCR": ut.Currency{Currency: "SCR", DisplayName: "སེ་ཤཱལས་ཀྱི་དངུལ་ རུ་པི", Symbol: ""}, "AFN": ut.Currency{Currency: "AFN", DisplayName: "ཨཕ་གཱན་གྱི་དངུལ་ ཨཕ་ག་ནི", Symbol: ""}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "ཇཱ་མཻ་ཀ་གི་དངུལ་ ཌོ་ལར", Symbol: ""}, "KES": ut.Currency{Currency: "KES", DisplayName: "ཀེན་ཡ་གི་དངུལ་ ཤི་ལིང", Symbol: ""}, "UZS": ut.Currency{Currency: "UZS", DisplayName: "ཨུས་བེ་ཀིས་ཏཱན་གྱི་དངུལ་ སོམ", Symbol: ""}, "VND": ut.Currency{Currency: "VND", DisplayName: "བེཊ་ནཱམ་གྱི་དངུལ་ ཌོང", Symbol: ""}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "", Symbol: "XAF"}, "IRR": ut.Currency{Currency: "IRR", DisplayName: "ཨི་རཱན་གྱི་དངུལ་ རི་ཨཱལ", Symbol: ""}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "སཱའུཐ་ ཀོ་རི་ཡ་གི་དངུལ་ ཝོན", Symbol: "KR₩"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "ཨིས་རེལ་གྱི་དངུལ་གསརཔ་ ཤེ་ཀེལ", Symbol: "ILS"}, "MAD": ut.Currency{Currency: "MAD", DisplayName: "མོ་རོ་ཀོ་གི་དངུལ་ ཌིར་ཧཱམ", Symbol: ""}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "སི་རི་ཡ་གི་དངུལ་ པ་འུནཌ", Symbol: ""}, "THB": ut.Currency{Currency: "THB", DisplayName: "ཐཱའི་ལེནཌ་གི་དངུལ་ བཱཏ", Symbol: "TH฿"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "ཊཱར་ཀི་གི་དངུལ་ ལི་ར", Symbol: ""}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "ཤྲི་ ལང་ཀ་གི་དངུལ་ རུ་པི", Symbol: ""}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "ཡུ་རུ་གུ་ཝའི་གི་དངུལ་ པེ་སོ", Symbol: ""}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "ཨཱཡིས་ལེནཌ་གི་དངུལ་ ཀྲོ་ན", Symbol: ""}, "KWD": ut.Currency{Currency: "KWD", DisplayName: "ཀུ་ཝེཊ་གི་དངུལ་ ཌི་ན", Symbol: ""}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "སོག་པོའི་དངུལ་ ཏུ་གྲིཀ", Symbol: ""}, "OMR": ut.Currency{Currency: "OMR", DisplayName: "ཨོ་མཱན་གྱི་དངུལ་ རི་ཨཱལ", Symbol: ""}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "ཇཱ་པཱན་གྱི་དངུལ་ ཡེན", Symbol: "JP¥"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "ཕི་ལི་པིནས་གྱི་དངུལ་ པེ་སོ", Symbol: ""}, "TZS": ut.Currency{Currency: "TZS", DisplayName: "ཊཱན་ཛཱ་ནི་ཡ་གི་དངུལ་ ཤི་ལིང", Symbol: ""}, "BTN": ut.Currency{Currency: "BTN", DisplayName: "དངུལ་ཀྲམ", Symbol: "Nu."}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "ནོརཐ་ ཀོ་རི་ཡ་གི་དངུལ་ ཝོན", Symbol: ""}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "ལེ་བ་ནོན་གྱི་དངུལ་ པ་འུནཌ", Symbol: ""}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "ཨུ་རུ་སུ་གི་དངུལ་ རུ་བཱལ", Symbol: ""}}
