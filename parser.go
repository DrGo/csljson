// Package csljson parses the csl-json format used by citeproc-js
// https://citeproc-js.readthedocs.io/en/latest/csl-json/markup.html
package csljson

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ParseFile(path string) (*CslData, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var data *CslData
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
