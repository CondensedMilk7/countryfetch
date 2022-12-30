package countries

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

func ReadData(cacheDir string, fileName string) ([]Country, error) {
	path := path.Join(cacheDir, fileName)
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("Could not read file: " + path + "\nTry --sync.")
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

	body, err := ioutil.ReadAll(resp.Body)

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
