package bitmedia_api

import (
	"strings"
)

func HexaNumberToInteger(hexaString string) string {
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}
