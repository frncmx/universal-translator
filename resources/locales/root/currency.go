package root

import "github.com/go-playground/universal-translator"

var currencies = ut.CurrencyFormatValue{"EUR": ut.Currency{Currency: "EUR", DisplayName: "", Symbol: "€"}, "GNF": ut.Currency{Currency: "GNF", DisplayName: "", Symbol: "FG"}, "JPY": ut.Currency{Currency: "JPY", DisplayName: "", Symbol: "JP¥"}, "NGN": ut.Currency{Currency: "NGN", DisplayName: "", Symbol: "₦"}, "XOF": ut.Currency{Currency: "XOF", DisplayName: "", Symbol: "CFA"}, "BMD": ut.Currency{Currency: "BMD", DisplayName: "", Symbol: "$"}, "BND": ut.Currency{Currency: "BND", DisplayName: "", Symbol: "$"}, "CUP": ut.Currency{Currency: "CUP", DisplayName: "", Symbol: "$"}, "CRC": ut.Currency{Currency: "CRC", DisplayName: "", Symbol: "₡"}, "KMF": ut.Currency{Currency: "KMF", DisplayName: "", Symbol: "CF"}, "LRD": ut.Currency{Currency: "LRD", DisplayName: "", Symbol: "$"}, "PKR": ut.Currency{Currency: "PKR", DisplayName: "", Symbol: "Rs"}, "STD": ut.Currency{Currency: "STD", DisplayName: "", Symbol: "Db"}, "BOB": ut.Currency{Currency: "BOB", DisplayName: "", Symbol: "Bs"}, "BWP": ut.Currency{Currency: "BWP", DisplayName: "", Symbol: "P"}, "CAD": ut.Currency{Currency: "CAD", DisplayName: "", Symbol: "CA$"}, "LTL": ut.Currency{Currency: "LTL", DisplayName: "", Symbol: "Lt"}, "TRY": ut.Currency{Currency: "TRY", DisplayName: "", Symbol: "₺"}, "XPF": ut.Currency{Currency: "XPF", DisplayName: "", Symbol: "CFPF"}, "BBD": ut.Currency{Currency: "BBD", DisplayName: "", Symbol: "$"}, "GIP": ut.Currency{Currency: "GIP", DisplayName: "", Symbol: "£"}, "KRW": ut.Currency{Currency: "KRW", DisplayName: "", Symbol: "₩"}, "MNT": ut.Currency{Currency: "MNT", DisplayName: "", Symbol: "₮"}, "XAF": ut.Currency{Currency: "XAF", DisplayName: "", Symbol: "FCFA"}, "ZAR": ut.Currency{Currency: "ZAR", DisplayName: "", Symbol: "R"}, "CLP": ut.Currency{Currency: "CLP", DisplayName: "", Symbol: "$"}, "FJD": ut.Currency{Currency: "FJD", DisplayName: "", Symbol: "$"}, "HKD": ut.Currency{Currency: "HKD", DisplayName: "", Symbol: "HK$"}, "MUR": ut.Currency{Currency: "MUR", DisplayName: "", Symbol: "Rs"}, "NIO": ut.Currency{Currency: "NIO", DisplayName: "", Symbol: "C$"}, "USD": ut.Currency{Currency: "USD", DisplayName: "", Symbol: "US$"}, "EGP": ut.Currency{Currency: "EGP", DisplayName: "", Symbol: "E£"}, "ISK": ut.Currency{Currency: "ISK", DisplayName: "", Symbol: "kr"}, "LVL": ut.Currency{Currency: "LVL", DisplayName: "", Symbol: "Ls"}, "MGA": ut.Currency{Currency: "MGA", DisplayName: "", Symbol: "Ar"}, "MMK": ut.Currency{Currency: "MMK", DisplayName: "", Symbol: "K"}, "NOK": ut.Currency{Currency: "NOK", DisplayName: "", Symbol: "kr"}, "RUR": ut.Currency{Currency: "RUR", DisplayName: "", Symbol: "р."}, "HRK": ut.Currency{Currency: "HRK", DisplayName: "", Symbol: "kn"}, "KYD": ut.Currency{Currency: "KYD", DisplayName: "", Symbol: "$"}, "LKR": ut.Currency{Currency: "LKR", DisplayName: "", Symbol: "Rs"}, "NPR": ut.Currency{Currency: "NPR", DisplayName: "", Symbol: "Rs"}, "SGD": ut.Currency{Currency: "SGD", DisplayName: "", Symbol: "$"}, "VND": ut.Currency{Currency: "VND", DisplayName: "", Symbol: "₫"}, "BAM": ut.Currency{Currency: "BAM", DisplayName: "", Symbol: "KM"}, "BDT": ut.Currency{Currency: "BDT", DisplayName: "", Symbol: "৳"}, "HNL": ut.Currency{Currency: "HNL", DisplayName: "", Symbol: "L"}, "GTQ": ut.Currency{Currency: "GTQ", DisplayName: "", Symbol: "Q"}, "MYR": ut.Currency{Currency: "MYR", DisplayName: "", Symbol: "RM"}, "SYP": ut.Currency{Currency: "SYP", DisplayName: "", Symbol: "£"}, "VEF": ut.Currency{Currency: "VEF", DisplayName: "", Symbol: "Bs"}, "ZMW": ut.Currency{Currency: "ZMW", DisplayName: "", Symbol: "ZK"}, "BYR": ut.Currency{Currency: "BYR", DisplayName: "", Symbol: "р."}, "KPW": ut.Currency{Currency: "KPW", DisplayName: "", Symbol: "₩"}, "MXN": ut.Currency{Currency: "MXN", DisplayName: "", Symbol: "MX$"}, "ILS": ut.Currency{Currency: "ILS", DisplayName: "", Symbol: "₪"}, "JMD": ut.Currency{Currency: "JMD", DisplayName: "", Symbol: "$"}, "LAK": ut.Currency{Currency: "LAK", DisplayName: "", Symbol: "₭"}, "NAD": ut.Currency{Currency: "NAD", DisplayName: "", Symbol: "$"}, "NZD": ut.Currency{Currency: "NZD", DisplayName: "", Symbol: "NZ$"}, "AOA": ut.Currency{Currency: "AOA", DisplayName: "", Symbol: "Kz"}, "CUC": ut.Currency{Currency: "CUC", DisplayName: "", Symbol: "$"}, "GBP": ut.Currency{Currency: "GBP", DisplayName: "", Symbol: "£"}, "SEK": ut.Currency{Currency: "SEK", DisplayName: "", Symbol: "kr"}, "SHP": ut.Currency{Currency: "SHP", DisplayName: "", Symbol: "£"}, "UAH": ut.Currency{Currency: "UAH", DisplayName: "", Symbol: "₴"}, "BRL": ut.Currency{Currency: "BRL", DisplayName: "", Symbol: "R$"}, "GYD": ut.Currency{Currency: "GYD", DisplayName: "", Symbol: "$"}, "RUB": ut.Currency{Currency: "RUB", DisplayName: "", Symbol: "₽"}, "RWF": ut.Currency{Currency: "RWF", DisplayName: "", Symbol: "RF"}, "AUD": ut.Currency{Currency: "AUD", DisplayName: "", Symbol: "A$"}, "GEL": ut.Currency{Currency: "GEL", DisplayName: "", Symbol: "₾"}, "HUF": ut.Currency{Currency: "HUF", DisplayName: "", Symbol: "Ft"}, "PYG": ut.Currency{Currency: "PYG", DisplayName: "", Symbol: "₲"}, "SBD": ut.Currency{Currency: "SBD", DisplayName: "", Symbol: "$"}, "SRD": ut.Currency{Currency: "SRD", DisplayName: "", Symbol: "$"}, "SSP": ut.Currency{Currency: "SSP", DisplayName: "", Symbol: "£"}, "THB": ut.Currency{Currency: "THB", DisplayName: "", Symbol: "฿"}, "ESP": ut.Currency{Currency: "ESP", DisplayName: "", Symbol: "₧"}, "KZT": ut.Currency{Currency: "KZT", DisplayName: "", Symbol: "₸"}, "PLN": ut.Currency{Currency: "PLN", DisplayName: "", Symbol: "zł"}, "TTD": ut.Currency{Currency: "TTD", DisplayName: "", Symbol: "$"}, "FKP": ut.Currency{Currency: "FKP", DisplayName: "", Symbol: "£"}, "IDR": ut.Currency{Currency: "IDR", DisplayName: "", Symbol: "Rp"}, "INR": ut.Currency{Currency: "INR", DisplayName: "", Symbol: "₹"}, "KHR": ut.Currency{Currency: "KHR", DisplayName: "", Symbol: "៛"}, "LBP": ut.Currency{Currency: "LBP", DisplayName: "", Symbol: "L£"}, "CNY": ut.Currency{Currency: "CNY", DisplayName: "", Symbol: "CN¥"}, "COP": ut.Currency{Currency: "COP", DisplayName: "", Symbol: "$"}, "DKK": ut.Currency{Currency: "DKK", DisplayName: "", Symbol: "kr"}, "PHP": ut.Currency{Currency: "PHP", DisplayName: "", Symbol: "₱"}, "TWD": ut.Currency{Currency: "TWD", DisplayName: "", Symbol: "NT$"}, "UYU": ut.Currency{Currency: "UYU", DisplayName: "", Symbol: "$"}, "DOP": ut.Currency{Currency: "DOP", DisplayName: "", Symbol: "$"}, "TOP": ut.Currency{Currency: "TOP", DisplayName: "", Symbol: "T$"}, "CZK": ut.Currency{Currency: "CZK", DisplayName: "", Symbol: "Kč"}, "XCD": ut.Currency{Currency: "XCD", DisplayName: "", Symbol: "EC$"}, "ARS": ut.Currency{Currency: "ARS", DisplayName: "", Symbol: "$"}, "BSD": ut.Currency{Currency: "BSD", DisplayName: "", Symbol: "$"}, "BZD": ut.Currency{Currency: "BZD", DisplayName: "", Symbol: "$"}}
