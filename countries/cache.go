package countries

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func ReadData(path string) ([]Country, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var data []Country
	err = json.Unmarshal(f, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func SaveData(data []Country) error {
	result, err := json.MarshalIndent(&data, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile("countries.json", result, 0644)
	return nil
}

func GetData(config *Config) ([]Country, error) {
	resp, err := http.Get(config.Url + config.Query)
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
