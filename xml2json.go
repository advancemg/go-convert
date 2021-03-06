package go_convert

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
)

type xmlNode struct {
	XMLName xml.Name
	Content string     `xml:",innerxml"`
	Nodes   []*xmlNode `xml:",any"`
	Attrs   []xml.Attr `xml:",any,attr"`
}

// ZipXmlToJson covert xmt to json.
// Input gzip []byte.
// Output []byte and error.
func ZipXmlToJson(input []byte) ([]byte, error) {
	in := bytes.NewReader(input)
	reader, err := gzip.NewReader(in)
	if err != nil {
		return nil, fmt.Errorf("not correct zip data, %v", err)
	}
	defer reader.Close()
	unZipBuffer := new(bytes.Buffer)
	_, err = io.Copy(unZipBuffer, reader)
	if err != nil {
		return nil, fmt.Errorf("not copy unzip data, %v", err)
	}
	return XmlToJson(unZipBuffer.Bytes())
}

// XmlToJson covert xmt to json.
// Output []byte and error.
func XmlToJson(input []byte) ([]byte, error) {
	n := &xmlNode{}
	err := xml.Unmarshal(input, n)
	if err != nil {
		return nil, err
	}
	return n.marshalJson()
}

// FileXmlToJson covert xml file to json.
// Output []byte and error.
func FileXmlToJson(file string) ([]byte, error) {
	readFile, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return XmlToJson(readFile)
}

func (n *xmlNode) marshalJson() ([]byte, error) {
	return json.Marshal(n.xmlToMap())
}

func (n *xmlNode) xmlToMap() map[string]interface{} {
	result := make(map[string]interface{})
	fields := make(map[string]interface{})
	addFields := func(key string, value interface{}) {
		if key != "nil" {
			if v, ok := fields[key]; ok {
				if array, ok := v.([]interface{}); ok {
					array = append(array, value)
					fields[key] = array
				} else {
					arrayValue := make([]interface{}, 2)
					arrayValue[0] = v
					arrayValue[1] = value
					fields[key] = arrayValue
				}
			} else {
				fields[key] = value
			}
		}
	}
	var fieldArray []interface{}
	ifNodeArray := n.isNodeArray()
	if len(n.Nodes) > 0 {
		for _, node := range n.Nodes {
			if ifNodeArray {
				fieldArray = append(fieldArray, node.xmlToMap())
			}
			addFields(node.XMLName.Local, node.xmlToMap()[node.XMLName.Local])
		}
	}
	isAttributes := false
	attrFields := map[string]interface{}{}
	if len(n.Attrs) > 0 {
		for _, v := range n.Attrs {
			if v.Name.Local == "xsi" {
				fields[attributesKey] = map[string]interface{}{"xsi": v.Value}
			}
			if v.Name.Local != "nil" {
				attrFields[v.Name.Local] = v.Value
				fields[attributesKey] = attrFields
			}
		}
		isAttributes = true
	}
	if len(fields) == 0 {
		result[n.XMLName.Local] = n.Content
	} else {
		if len(n.Attrs) > 0 && len(n.Nodes) == 0 {
			if len(n.Content) > 0 {
				addFields(contentKey, n.Content)
			}
			result[n.XMLName.Local] = fields
		} else {
			result[n.XMLName.Local] = fields
		}
	}
	if ifNodeArray && !isAttributes {
		result[n.XMLName.Local] = fieldArray
	}
	return result
}

func (n *xmlNode) isNodeArray() bool {
	if len(n.Nodes) <= 1 {
		return false
	}
	nodeName := n.Nodes[0].XMLName.Local
	for i, nd := range n.Nodes {
		if nd.XMLName.Local != nodeName {
			return false
		} else {
			if i >= 1 {
				return true
			}
		}
	}
	return true
}
