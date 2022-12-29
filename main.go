package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/CondensedMilk7/countryfetch-go/countries"
)

var config = countries.Config{
	Url:   "https://restcountries.com/v3.1/",
	Query: "all?fields=name,capital,currencies,population,flag,languages,region,subregion,timezones,latlng,capitalInfo,tld,flags",
}

var syncFlag bool
var nameFlag string

func main() {
	flag.BoolVar(&syncFlag, "sync", false, "Fetch and save data to cache")
	flag.StringVar(&nameFlag, "name", "", "Find country by given name")
	flag.Parse()

	if syncFlag {
		fmt.Println("Fetching countries data...")
		data, err := countries.GetData(&config)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Saving cache...")
		err = countries.SaveData(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Cache saved.")
	}

	if nameFlag != "" {
		data, err := countries.ReadData("./countries.json")
		if err != nil {
			fmt.Println(err)
		}
		c, err := countries.FindByName(data, nameFlag)
		if err != nil {
			fmt.Println(err)
		}
		c.Print()
	}

	if len(os.Args) == 1 {
		fmt.Println("USAGE:")
		flag.PrintDefaults()
	}
}
