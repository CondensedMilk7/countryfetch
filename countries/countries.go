package countries

import (
	"errors"
	"fmt"
	"github.com/CondensedMilk7/countryfetch-go/color"
	"strings"
)

type Config struct {
	Url   string `json:"url"`
	Query string `json:"query"`
}

type CurrencyInfo struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Country struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	//          Lord have mercy
	Currencies  map[string]CurrencyInfo `json:"currencies"`
	Capital     []string                `json:"capital"`
	Flag        string                  `json:"flag"`
	Population  int                     `json:"population"`
	Languages   map[string]string       `json:"languages"`
	Region      string                  `json:"region"`
	Subregion   string                  `json:"subregion"`
	Timezones   []string                `json:"timezones"`
	Latlng      []float32               `json:"latlng"`
	CapitalInfo struct {
		Latlng []float32 `json:"latlng"`
	} `json:"capitalInfo"`
	Tld   []string `json:"tld"`
	Flags struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags"`
}

// TODO: Make fields and values of different color
func (c Country) Print() {
	output :=
		`
Name: %s %s
Lat/Lng: %.1f/%.1f
Populaiton: %s
Languages: %s
Capital: %s
Capital Lat/Lng: %.2f/%.2f
Region: %s
Subregion: %s
Timezones: %s
Top Level Domain: %s
Currencies: %s
`
	fmt.Printf(
		color.WrapInColor(color.Green, output),
		c.Name.Common,
		c.Flag,
		c.Latlng[0],
		c.Latlng[1],
		FormatInt(c.Population),
		FormatLanguages(c.Languages),
		c.Capital[0],
		c.CapitalInfo.Latlng[0],
		c.CapitalInfo.Latlng[1],
		c.Region,
		c.Subregion,
		c.Timezones,
		c.Tld,
		FormatCurrencies(c.Currencies),
	)
}

func FindByName(countries []Country, exp string) (Country, error) {
	country, err := ExactMatch(countries, exp)
	if err != nil {
		country, err = FuzzyMatch(countries, exp)
	}
	return country, err
}

// func FindByCapital(countries []Country, exp string) (Country, err) {

// }

func ExactMatch(countries []Country, exp string) (Country, error) {
	for _, c := range countries {
		if strings.ToLower(c.Name.Common) == strings.ToLower(exp) {
			return c, nil
		}
	}
	return Country{}, errors.New("Could not find exact match for " + exp)
}

func FuzzyMatch(countries []Country, exp string) (Country, error) {
	for _, c := range countries {
		if strings.Contains(strings.ToLower(c.Name.Common), strings.ToLower(exp)) {
			return c, nil
		}
	}
	return Country{}, errors.New("Could not find fuzzy match for " + exp)
}
