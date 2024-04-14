package bot_service

import (
	"fmt"
	"strconv"
	"strings"
)

func StringToFloat64(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Error parsing to float 64:", err)
		return 0
	}
	return f
}

func Float64ToString(f float64) string {
	//n are decimals
	//d := strconv.FormatFloat(f, 'f', n, 64)

	strValue := fmt.Sprintf("%.7f", f)
	strValue = strings.TrimRight(strValue, "0")
	strValue = strings.TrimRight(strValue, ".")

	return strValue
}
