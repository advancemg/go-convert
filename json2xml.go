package go_convert

import (
	"encoding/json"
)

// JsonToXml covert json to xml.
// Output []byte and error.
func JsonToXml(input []byte) ([]byte, error) {
	var inputMap UnsortedMap
	err := json.Unmarshal(input, &inputMap)
	if err != nil {
		return nil, err
	}
	return inputMap.ToXml()
}
