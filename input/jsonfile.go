package input

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/forbole/bookkeeper/types"
)

func ImportJsonInput(path string) (*types.Data, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var input types.Data
	json.Unmarshal(byteValue, &input)

	return &input, nil

}
