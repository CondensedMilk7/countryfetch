package countries

import (
	"bytes"
	"errors"
	"fmt"
	"image/png"
	"io"
	"net/http"
	"strings"

	"github.com/CondensedMilk7/countryfetch/internal/color"
	"github.com/qeesung/image2ascii/convert"
)

type CacheDir func() (string, error)

type FlagDimensions struct {
	Width  int
	Height int
}

type Config struct {
	Url        string
	Query      string
	CacheFile  string
	CacheDir   CacheDir
	Dimensions FlagDimensions
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

// TODO: Make getters for Country properties that are troublesome
// (like capital that is sometimes empty slice)
func (c Country) Print() {
	output :=
		`Name: %s %s
Lat/Lng: %s
Population: %s
Languages: %s
Capital: %s
Capital Lat/Lng: %s
Region: %s
Subregion: %s
Timezones: %s
Top Level Domain: %s
Currencies: %s
`
	fmt.Printf(
		output,
		color.WrapInColor(color.Cyan, c.Name.Common),
		c.Flag,
		color.WrapInColor(color.Cyan, FormatLatLng(c.Latlng)),
		color.WrapInColor(color.Cyan, FormatInt(c.Population)),
		color.WrapInColor(color.Cyan, FormatLanguages(c.Languages)),
		color.WrapInColor(color.Cyan, c.Capital[0]),
		color.WrapInColor(color.Cyan, FormatLatLng(c.CapitalInfo.Latlng)),
		color.WrapInColor(color.Cyan, c.Region),
		color.WrapInColor(color.Cyan, c.Subregion),
		color.WrapInColor(color.Cyan, FormatTz(c.Timezones)),
		color.WrapInColor(color.Cyan, c.Tld[0]),
		color.WrapInColor(color.Cyan, FormatCurrencies(c.Currencies)),
	)
}

func (c Country) PrintFlagRemote(width int, height int) error {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = width
	convertOptions.FixedHeight = height
	converter := convert.NewImageConverter()

	ascii, err := c.FlagAscii(converter, &convertOptions)
	if err != nil {
		return err
	}
	fmt.Println(ascii)

	return nil
}

// Generate flag ASCII from remote url image and return it as string
func (c Country) FlagAscii(converter *convert.ImageConverter, options *convert.Options) (string, error) {
	resp, err := http.Get(c.Flags.Png)
	if err != nil {
		return "", nil
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()

	image, err := png.Decode(bytes.NewReader(data))

	asciiFlag := converter.Image2ASCIIString(image, options)
	return asciiFlag, nil
}

// Print flag from cache
func (c Country) PrintFlag(cacheDir string) error {
	searchStr := FormatFlagFileName(c.Name.Common)
	flag, err := ReadFlag(searchStr, cacheDir)
	if err != nil {
		return err
	}
	fmt.Println(flag)
	return nil
}

func FindByName(countries []Country, exp string) (Country, error) {
	country, err := ExactMatch(countries, exp)
	if err != nil {
		country, err = FuzzyMatch(countries, exp)
	}
	return country, err
}

func FindByCapital(countries []Country, exp string) (Country, error) {
	for _, c := range countries {
		if len(c.Capital) > 0 &&
			strings.Contains(strings.ToLower(c.Capital[0]), strings.ToLower(exp)) {
			return c, nil
		}
	}
	return Country{}, errors.New("Could not find country of the given capital " + exp)
}

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
