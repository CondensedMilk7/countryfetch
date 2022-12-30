package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CondensedMilk7/countryfetch-go/color"
	"github.com/CondensedMilk7/countryfetch-go/countries"
)

var config = countries.Config{
	Url:   "https://restcountries.com/v3.1/",
	Query: "all?fields=name,capital,currencies,population,flag,languages,region,subregion,timezones,latlng,capitalInfo,tld,flags",
}

var syncFlag bool
var nameFlag string
var capitalFlag string

func main() {
	flag.BoolVar(&syncFlag, "sync", false, "Fetch and save data to cache")
	flag.StringVar(&nameFlag, "name", "", "Find country by given name")
	flag.StringVar(&capitalFlag, "capital", "", "Find country by given capital")
	flag.Parse()

	if syncFlag {
		fmt.Println("Fetching countries data...")
		data, err := countries.GetData(&config)
		checkErr(err)
		fmt.Println("Saving cache...")
		err = countries.SaveData(data)
		checkErr(err)
		fmt.Println("Cache saved.")
	}

	if nameFlag != "" {
		data, err := countries.ReadData("./countries.json")
		checkErr(err)
		c, err := countries.FindByName(data, nameFlag)
		checkErr(err)
		c.Print()
	}

	if capitalFlag != "" {
		data, err := countries.ReadData("./countries.json")
		checkErr(err)
		c, err := countries.FindByCapital(data, capitalFlag)
		checkErr(err)
		c.Print()
	}

	if len(os.Args) == 1 {
		fmt.Println("USAGE:")
		flag.PrintDefaults()
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(fmt.Sprintf("%s%s", color.Red, err))
		os.Exit(2)
	}
}
