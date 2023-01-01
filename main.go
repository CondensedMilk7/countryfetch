package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/CondensedMilk7/countryfetch/color"
	"github.com/CondensedMilk7/countryfetch/countries"
)

var config = countries.Config{
	Url:       "https://restcountries.com/v3.1/",
	Query:     "all?fields=name,capital,currencies,population,flag,languages,region,subregion,timezones,latlng,capitalInfo,tld,flags",
	CacheFile: "countries.json",
	CacheDir: func() (string, error) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return path.Join(homeDir, ".cache", "countryfetch"), nil
	},
}

var sync bool
var countryName string
var capital string
var withFlag bool
var flagOnly bool
var flagRemote bool

func main() {
	flag.BoolVar(&sync, "sync", false, "Fetch and save data to cache.")
	flag.StringVar(&countryName, "name", "", "Find country by given name.")
	flag.StringVar(&capital, "capital", "", "Find country by given capital.")
	flag.BoolVar(&withFlag, "flag", false, "Include ASCII flag in the operation. Can be used in combination with -sync & -name.")
	flag.BoolVar(&flagOnly, "flagonly", false, "Print flag only. Used in combination with -name.")
	flag.BoolVar(&flagRemote, "flagremote", false, "Print flag only via remote URL. Used in combination -name.")
	flag.Parse()

	cacheDir, err := config.CacheDir()
	checkErr(err)

	if sync {
		fmt.Println("Fetching countries data...")
		data, err := countries.GetData(config.Url, config.Query)
		checkErr(err)
		fmt.Println("Saving cache...")
		err = countries.SaveData(data, cacheDir, config.CacheFile)
		checkErr(err)
		fmt.Println(color.WrapInColor(color.Green, "Cache saved."))

		if withFlag {
			fmt.Println("Generating ASCII art for each country flag. This may take a minute...")
			data, err := countries.ReadData(cacheDir, config.CacheFile)
			checkErr(err)
			err = countries.CacheFlags(data, cacheDir, func(current int, total int, cName string) {
				fmt.Println(fmt.Sprintf("(%d/%d) %s", current, total, cName))
			})
			checkErr(err)
			fmt.Println(color.WrapInColor(color.Green, "Done."))
		}
	}

	if countryName != "" {
		data, err := countries.ReadData(cacheDir, config.CacheFile)
		checkErr(err)
		c, err := countries.FindByName(data, countryName)
		checkErr(err)
		if withFlag {
			err := c.PrintFlag(cacheDir)
			checkErr(err)
			c.Print()
		} else if flagOnly {
			err := c.PrintFlag(cacheDir)
			checkErr(err)
		} else if flagRemote {
			err := c.PrintFlagRemote()
			checkErr(err)
		} else {
			fmt.Println()
			c.Print()
		}
	}

	if capital != "" {
		data, err := countries.ReadData(cacheDir, config.CacheFile)
		checkErr(err)
		c, err := countries.FindByCapital(data, capital)
		checkErr(err)
		c.Print()
	}

	if len(os.Args) == 1 {
		fmt.Println("USAGE:")
		flag.PrintDefaults()
		fmt.Println("EXAMPLE:")
		fmt.Println(`  countryfetch -name italy -flag
      Fetch information about Italy, including its flag.
  countryfetch -sync -flag
      Store information of all countries in cache, including generated flag ASCII art.
  countryfetch -capital \"kuala lumpur\"
      Fetch information about the country of given capital.
  countryfetch -flagonly -name "united states"
      Fetch just the flag of USA.`)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(fmt.Sprintf("%s%s", color.Red, err))
		os.Exit(2)
	}
}
