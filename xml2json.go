package go_convert

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

type xmlNode struct {
	XMLName xml.Name
	Content string     `xml:",innerxml"`
	Nodes   []*xmlNode `xml:",any"`
	Attrs   []xml.Attr `xml:",any,attr"`
}

// XmlToJson covert xmt to json.
// Output []byte and error.
func XmlToJson(input []byte) ([]byte, error) {
	n := &xmlNode{}
	xml.Unmarshal(input, n)
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
				fields["attributes"] = map[string]interface{}{"xsi": v.Value}
			}
			if v.Name.Local != "nil" {
				attrFields[v.Name.Local] = v.Value
				fields["attributes"] = attrFields
			}
		}
		isAttributes = true
	}
	if len(fields) == 0 {
		result[n.XMLName.Local] = n.Content
	} else {
		if len(n.Attrs) > 0 && len(n.Nodes) == 0 {
			if len(n.Content) > 0 {
				addFields("_content", n.Content)
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
