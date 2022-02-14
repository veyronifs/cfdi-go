package curconv

import (
	"github.com/shopspring/decimal"
)

// decimals maps currency codes to the number of decimals they have.
// Any currency not in this map will have 2 decimals.
var decimals = map[string]int{
	"CLF": 4,
	"BHD": 3,
	"IQD": 3,
	"JOD": 3,
	"KWD": 3,
	"LYD": 3,
	"OMR": 3,
	"TND": 3,
	"BIF": 0,
	"BYR": 0,
	"CLP": 0,
	"DJF": 0,
	"GNF": 0,
	"ISK": 0,
	"JPY": 0,
	"KMF": 0,
	"KRW": 0,
	"PYG": 0,
	"RWF": 0,
	"UGX": 0,
	"UYI": 0,
	"VND": 0,
	"VUV": 0,
	"XAF": 0,
	"XAG": 0,
	"XAU": 0,
	"XBA": 0,
	"XBB": 0,
	"XBC": 0,
	"XBD": 0,
	"XDR": 0,
	"XOF": 0,
	"XPD": 0,
	"XPF": 0,
	"XPT": 0,
	"XSU": 0,
	"XTS": 0,
	"XUA": 0,
	"XXX": 0,
}

func getDecimals(curr string) int {
	n, ok := decimals[curr]
	if !ok {
		n = 2
	}
	return n
}

func RoundToMax[S ~string](v decimal.Decimal, curr S) decimal.Decimal {
	dec := getDecimals(string(curr))
	return RoundToDec(v, dec)
}

func RoundToMaxStr[S ~string](v decimal.Decimal, curr S) string {
	dec := getDecimals(string(curr))
	return RoundToDecStr(RoundToDec(v, dec), dec)
}

func RoundToDec(v decimal.Decimal, dec int) decimal.Decimal {
	return v.Round(int32(dec))
}

func RoundToDecStr(v decimal.Decimal, dec int) string {
	return RoundToDec(v, dec).StringFixed(int32(dec))
}
