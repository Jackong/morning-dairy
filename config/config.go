/**
 * User: jackong
 * Date: 11/5/13
 * Time: 5:22 PM
 */
package config

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"os"
)

type Config map[string]interface {}

func NewConfig(file string) Config {
	config := make(Config)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	json.Unmarshal(data, &config)
	return config
}

func (this Config) String(keys ... string) string {
	tmp := this
	for _, key := range keys {
		val := tmp[key]
		if val == nil {
			return ""
		}
		switch val.(type) {
		case string:
			return val.(string)
		default:
			tmp = val.(map[string] interface {})
		}
	}
	return ""
}

func (this Config) Slice(keys ... string) (array []string) {
    tmp := this
	for _, key := range keys {
		val := tmp[key]
		if val == nil {
			return array
		}
		switch val.(type) {
		case ([]interface {}):
			for _, element := range val.([]interface {}) {
				array = append(array, element.(string))
			}
			return array
		default:
			tmp = val.(map[string] interface {})
		}
	}
	return array
}
