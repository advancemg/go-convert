package go_convert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/clbanning/mxj/v2/j2x"
	"strings"
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

func JsonToXml1(input []byte) ([]byte, error) {
	/*	var json = jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            false,
		ValidateJsonRawMessage: false,
	}.Froze()*/
	tmpMap := make(map[string]interface{})
	resultMap := make(map[string]interface{})
	/*inputMap, err := j2x.JsonToMap(input)*/
	var inputMap map[string]interface{}
	err := json.Unmarshal(input, &inputMap)
	if err != nil {
		return nil, err
	}
	rootKey := ""
	attrKey := ""
	for key, value := range inputMap {
		rootKey = key
		tmpMap = addField(key, value)
	}
	for key, value := range tmpMap {
		attrKey = key
		resultMap[rootKey] = value
	}
	xmlBuilder := strings.Builder{}
	for k, v := range resultMap {

		xmlBuilder.WriteString(fmt.Sprintf("<%s>%s</%s>", k, v, k))
	}
	//json, err := j2x.MapToJson(resultMap)
	jsonData, err := json.Marshal(resultMap)
	if err != nil {
		return nil, err
	}
	/*json reader*/
	// empty or nil begets empty
	if len(jsonData) == 0 {
		return nil, err
	}
	if jsonData[0] == '[' {
		jsonData = []byte(`{"object":` + string(jsonData) + `}`)
	}
	m := make(map[string]interface{})
	// err := json.Unmarshal(jsonVal, &m)
	buf := bytes.NewReader(jsonData)
	dec := json.NewDecoder(buf)

	err = dec.Decode(&m)
	toMap, err := j2x.JsonToMap(jsonData)
	if err != nil {
		return nil, err
	}
	for k, v := range toMap {
		println(k, v)
	}
	xml, err := j2x.JsonToXml(jsonData)
	if err != nil {
		return nil, err
	}
	//Add attributes to rootKey
	res := strings.Replace(string(xml), rootKey, attrKey, 1)
	return []byte(res), nil
}

func addField(key string, value interface{}) map[string]interface{} {
	fields := make(map[string]interface{})
	result := make(map[string]interface{})
	attrKey := key
	if arr, ok := value.([]interface{}); ok {
		var tmpArr []interface{}
		for _, vv := range arr {
			if mp, ok := vv.(map[string]interface{}); ok {
				for k, v := range mp {
					if _, ok := fields[k]; !ok {
						tmpArr = []interface{}{}
					}
					tmpArr = append(tmpArr, v)
					fields[k] = tmpArr
				}
			}
		}
	}
	if v, ok := value.(map[string]interface{}); ok {
		for k, val := range v {
			if arr, ok := val.([]interface{}); ok {
				var tmpArr []interface{}
				for _, vv := range arr {
					if mp, ok := vv.(map[string]interface{}); ok {
						for keyL1, v := range mp {
							if _, ok := fields[k]; !ok {
								tmpArr = []interface{}{}
							}
							tmpArr = append(tmpArr, v)
							fields[k] = map[string]interface{}{keyL1: tmpArr}
						}
					}
				}
			}
			if mp, ok := val.(map[string]interface{}); ok {
				if k == "attributes" {
					for attKey, attValue := range mp {
						attrKey = fmt.Sprintf(`%s xmlns:%s="%v"`, key, attKey, attValue)
					}
				}
				if k != "attributes" {
					fields[k] = val
				}
			}
			if str, ok := val.(string); ok {
				fields[k] = str
			}
		}
	}
	if v, ok := value.(string); ok {
		if v != "" {
			fields[key] = v
		}
	}
	result[attrKey] = fields
	return result
}
