package countries

import (
	"fmt"
	"strconv"
)

func FormatCurrencies(currencies map[string]CurrencyInfo) string {
	str := ""
	for abr, info := range currencies {
		str += fmt.Sprintf("%s [%s][%s]", info.Name, abr, info.Symbol)
	}
	return str
}

func FormatLanguages(langs map[string]string) string {
	str := ""
	i := 1
	for _, value := range langs {
		if i < len(langs) {
			str += fmt.Sprintf("%s | ", value)
		} else {
			str += fmt.Sprintf("%s", value)
		}
		i++
	}
	return str
}

func FormatLatLng(latlng []float32) string {
	return fmt.Sprintf("%.2f/%.2f", latlng[0], latlng[1])
}

func FormatTz(tzs []string) string {
	result := ""
	for i, tz := range tzs {
		if i+1 < len(tzs) {
			result += tz + " | "
		} else {
			result += tz
		}
	}
	return result
}

func FormatInt(number int) string {
	output := strconv.Itoa(number)
	startOffset := 3
	if number < 0 {
		startOffset++
	}
	for outputIndex := len(output); outputIndex > startOffset; {
		outputIndex -= 3
		output = output[:outputIndex] + "," + output[outputIndex:]
	}
	return output
}
