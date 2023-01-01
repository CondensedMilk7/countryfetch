package countries

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"path"

	"image/png"

	"github.com/qeesung/image2ascii/convert"
)

func ReadFlag(fileName string, cacheDir string) (string, error) {
	flagsPath := path.Join(cacheDir, "flags")
	entries, err := os.ReadDir(flagsPath)
	if err != nil {
		return "", errors.New("Could not read directory: " + flagsPath + "\nTry -sync -flag")
	}
	for _, entry := range entries {
		if entry.Name() == fileName {
			file, err := os.ReadFile(path.Join(flagsPath, entry.Name()))

			return string(file), err
		}
	}
	return "", errors.New("Failed executing a loop over cache file entries")
}

func FlagAscii(countries []Country, cacheDir string, cb func(current int, total int, countryName string)) error {
	flagsPath := path.Join(cacheDir, "flags")
	os.MkdirAll(flagsPath, os.ModePerm)

	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 40
	convertOptions.FixedHeight = 12
	convertOptions.FitScreen = true
	converter := convert.NewImageConverter()

	for i, country := range countries {
		cb(i+1, len(countries), country.Name.Common)

		resp, err := http.Get(country.Flags.Png)
		if err != nil {
			return err
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		image, err := png.Decode(bytes.NewReader(data))

		asciiFlag := converter.Image2ASCIIString(image, &convertOptions)
		fileName := FormatFlagFileName(country.Name.Common)
		os.WriteFile(path.Join(flagsPath, fileName), []byte(asciiFlag), os.ModePerm)
	}

	return nil
}

func ReadData(cacheDir string, fileName string) ([]Country, error) {
	path := path.Join(cacheDir, fileName)
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("Could not read file: " + path + "\nTry -sync")
	}
	var data []Country
	err = json.Unmarshal(f, &data)
	if err != nil {
		return nil, errors.New("Error during unmarshal of: " + path)
	}
	return data, nil
}

func SaveData(data []Country, cacheDir string, fileName string) error {
	result, err := json.MarshalIndent(&data, "", " ")
	if err != nil {
		return err
	}
	err = os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		return err
	}
	path := path.Join(cacheDir, fileName)
	os.WriteFile(path, result, os.ModePerm)
	return nil
}

func GetData(url string, query string) ([]Country, error) {
	resp, err := http.Get(url + query)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var data []Country
	err = json.Unmarshal(body, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}
