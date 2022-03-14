package go_convert

import (
	"fmt"
	"github.com/clbanning/mxj/v2/j2x"
	"strings"
)

// JsonToXml covert json to xml.
// Output []byte and error.
func JsonToXml(input []byte) ([]byte, error) {
	tmpMap := make(map[string]interface{})
	resultMap := make(map[string]interface{})
	inputMap, err := j2x.JsonToMap(input)
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
	json, err := j2x.MapToJson(resultMap)
	if err != nil {
		return nil, err
	}
	xml, err := j2x.JsonToXml(json)
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
