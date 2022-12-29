package countries

import (
	"fmt"
	"strconv"
)

func FormatCurrencies(currencies map[string]CurrencyInfo) string {
	str := ""
	fmt.Println(currencies)
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
