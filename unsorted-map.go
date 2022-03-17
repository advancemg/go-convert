package go_convert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"
)

const (
	arrayStart     = '['
	arrayEnd       = ']'
	objectStart    = '{'
	objectEnd      = '}'
	objectValue    = ':'
	objectEndValue = ','
	attributesKey  = "attributes"
	contentKey     = "_content"
)

type UnsortedPair struct {
	key   string
	value interface{}
}

// Key UnsortedPair.
func (kv *UnsortedPair) Key() string {
	return kv.key
}

// Value UnsortedPair.
func (kv *UnsortedPair) Value() interface{} {
	return kv.value
}

type SortedPair struct {
	Pairs    []*UnsortedPair
	LessFunc func(a *UnsortedPair, j *UnsortedPair) bool
}

// Len SortedPair.
func (a SortedPair) Len() int { return len(a.Pairs) }

// Swap SortedPair.
func (a SortedPair) Swap(i, j int) { a.Pairs[i], a.Pairs[j] = a.Pairs[j], a.Pairs[i] }

// Less SortedPair.
func (a SortedPair) Less(i, j int) bool { return a.LessFunc(a.Pairs[i], a.Pairs[j]) }

type UnsortedMap struct {
	keys       []string
	buffer     bytes.Buffer
	values     map[string]interface{}
	escapeHTML bool
}

// New UnsortedMap.
func New() *UnsortedMap {
	result := &UnsortedMap{}
	result.keys = []string{}
	result.values = map[string]interface{}{}
	result.escapeHTML = true
	result.buffer = bytes.Buffer{}
	return result
}

// Get value UnsortedMap.
func (uMap *UnsortedMap) Get(key string) (val interface{}, exists bool) {
	val, exists = uMap.values[key]
	return
}

// Set value UnsortedMap.
func (uMap *UnsortedMap) Set(key string, value interface{}) {
	_, exists := uMap.values[key]
	if !exists {
		uMap.keys = append(uMap.keys, key)
	}
	uMap.values[key] = value
}

// Keys UnsortedMap.
func (uMap *UnsortedMap) Keys() []string {
	return uMap.keys
}

// XmlAttributes UnsortedMap.
func (uMap *UnsortedMap) XmlAttributes() []byte {
	attributeInterface, exists := uMap.values[attributesKey]
	if !exists {
		return nil
	}
	var buf bytes.Buffer
	switch attributeInterface.(type) {
	case *UnsortedMap:
		unsortedMap := attributeInterface.(*UnsortedMap)
		for _, k := range unsortedMap.keys {
			v := unsortedMap.values[k]
			buf.WriteString(fmt.Sprintf(" %v=%v", k, v))
		}
	case UnsortedMap:
		unsortedMap := attributeInterface.(UnsortedMap)
		for _, k := range unsortedMap.keys {
			v := unsortedMap.values[k]
			buf.WriteString(fmt.Sprintf(" %v=%v", k, v))
		}
	}
	return buf.Bytes()
}

// Remove value UnsortedMap.
func (uMap *UnsortedMap) Remove(key string) {
	_, exists := uMap.values[key]
	if !exists {
		return
	}
	for i, k := range uMap.keys {
		if k == key {
			uMap.keys = append(uMap.keys[:i], uMap.keys[i+1:]...)
			break
		}
	}
	delete(uMap.values, key)
}

// SortKeys map keys using your sort func
func (uMap *UnsortedMap) SortKeys(sortFunc func(keys []string)) {
	sortFunc(uMap.keys)
}

// SortPairs using your sort func
func (uMap *UnsortedMap) SortPairs(lessFunc func(a *UnsortedPair, b *UnsortedPair) bool) {
	pairs := make([]*UnsortedPair, len(uMap.keys))
	for i, key := range uMap.keys {
		pairs[i] = &UnsortedPair{key, uMap.values[key]}
	}
	sort.Sort(SortedPair{pairs, lessFunc})
	for i, pair := range pairs {
		uMap.keys[i] = pair.key
	}
}

// UnmarshalJSON UnsortedMap.
func (uMap *UnsortedMap) UnmarshalJSON(b []byte) error {
	if uMap.values == nil {
		uMap.values = map[string]interface{}{}
	}
	err := json.Unmarshal(b, &uMap.values)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(bytes.NewReader(b))
	if _, err = dec.Token(); err != nil { // skip '{'
		return err
	}
	uMap.keys = make([]string, 0, len(uMap.values))
	return decodeMap(dec, uMap)
}

// MarshalJSON UnsortedMap.
func (uMap UnsortedMap) MarshalJSON() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte(objectStart)
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(uMap.escapeHTML)
	for i, k := range uMap.keys {
		if i > 0 {
			buf.WriteByte(objectEndValue)
		}
		if err := encoder.Encode(k); err != nil {
			return nil, err
		}
		buf.WriteByte(objectValue)
		if err := encoder.Encode(uMap.values[k]); err != nil {
			return nil, err
		}
	}
	buf.WriteByte(objectEnd)
	return buf.Bytes(), nil
}

// ToXml convert UnsortedMap to xml.
func (uMap *UnsortedMap) ToXml() ([]byte, error) {
	uMap.buffer = bytes.Buffer{}
	attributes := uMap.XmlAttributes()
	attr := string(attributes)
	for _, key := range uMap.keys {
		if key == attributesKey {
			continue
		}
		value, _ := uMap.Get(key)
		uMap.buffer.WriteString(fmt.Sprintf("<%v%s>", key, attr))
		switch value.(type) {
		case UnsortedMap:
			unsortedMap := value.(UnsortedMap)
			xmlObject, err := unsortedMap.ToXml()
			if err != nil {
				return nil, err
			}
			uMap.buffer.Write(xmlObject)
		case *UnsortedMap:
			unsortedMap := value.(*UnsortedMap)
			xmlObject, err := unsortedMap.ToXml()
			if err != nil {
				return nil, err
			}
			uMap.buffer.Write(xmlObject)
		case []UnsortedMap:
			unsortedMapList := value.([]UnsortedMap)
			for _, unsortedMap := range unsortedMapList {
				xmlObject, err := unsortedMap.ToXml()
				if err != nil {
					return nil, err
				}
				uMap.buffer.Write(xmlObject)
			}
		case []interface{}:
			marshal, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			var list []UnsortedMap
			err = json.Unmarshal(marshal, &list)
			if err != nil {
				return nil, err
			}
			for _, unsortedMap := range list {
				xmlObject, err := unsortedMap.ToXml()
				if err != nil {
					return nil, err
				}
				uMap.buffer.Write(xmlObject)
			}
		case []*UnsortedMap:
			unsortedMapList := value.([]*UnsortedMap)
			for _, unsortedMap := range unsortedMapList {
				xmlObject, err := unsortedMap.ToXml()
				if err != nil {
					return nil, err
				}
				uMap.buffer.Write(xmlObject)
			}
		default:
			if value == "nil" {
				uMap.buffer.Truncate(uMap.buffer.Len() - 1)
				uMap.buffer.WriteString(fmt.Sprintf(" xsi:nil=\"true\">"))
			} else {
				uMap.buffer.WriteString(fmt.Sprintf("%v", value))
			}
		}
		uMap.buffer.WriteString(fmt.Sprintf("</%v>", key))
	}
	return uMap.buffer.Bytes(), nil
}

func decodeMap(dec *json.Decoder, uMap *UnsortedMap) error {
	hasKey := make(map[string]bool, len(uMap.values))
	for {
		token, err := dec.Token()
		if err != nil {
			return err
		}
		if delim, ok := token.(json.Delim); ok && delim == '}' {
			return nil
		}
		key := token.(string)
		if hasKey[key] {
			for j, k := range uMap.keys {
				if k == key {
					copy(uMap.keys[j:], uMap.keys[j+1:])
					break
				}
			}
			uMap.keys[len(uMap.keys)-1] = key
		} else {
			hasKey[key] = true
			uMap.keys = append(uMap.keys, key)
		}
		token, err = dec.Token()
		if err != nil {
			return err
		}
		if delim, ok := token.(json.Delim); ok {
			switch delim {
			case objectStart:
				if values, ok := uMap.values[key].(map[string]interface{}); ok {
					newMap := UnsortedMap{
						keys:       make([]string, 0, len(values)),
						values:     values,
						escapeHTML: uMap.escapeHTML,
					}
					if err = decodeMap(dec, &newMap); err != nil {
						return err
					}
					uMap.values[key] = newMap
				} else if oldMap, ok := uMap.values[key].(UnsortedMap); ok {
					newMap := UnsortedMap{
						keys:       make([]string, 0, len(oldMap.values)),
						values:     oldMap.values,
						escapeHTML: uMap.escapeHTML,
					}
					if err = decodeMap(dec, &newMap); err != nil {
						return err
					}
					uMap.values[key] = newMap
				} else if err = decodeMap(dec, &UnsortedMap{}); err != nil {
					return err
				}
			case arrayStart:
				if values, ok := uMap.values[key].([]interface{}); ok {
					if err = decodeSlice(dec, values, uMap.escapeHTML); err != nil {
						return err
					}
				} else if err = decodeSlice(dec, []interface{}{}, uMap.escapeHTML); err != nil {
					return err
				}
			}
		}
	}
}

func decodeSlice(dec *json.Decoder, s []interface{}, escapeHTML bool) error {
	for index := 0; ; index++ {
		token, err := dec.Token()
		if err != nil {
			return err
		}
		if delim, ok := token.(json.Delim); ok {
			switch delim {
			case objectStart:
				if index < len(s) {
					if values, ok := s[index].(map[string]interface{}); ok {
						newMap := UnsortedMap{
							keys:       make([]string, 0, len(values)),
							values:     values,
							escapeHTML: escapeHTML,
						}
						if err = decodeMap(dec, &newMap); err != nil {
							return err
						}
						s[index] = newMap
					} else if oldMap, ok := s[index].(UnsortedMap); ok {
						newMap := UnsortedMap{
							keys:       make([]string, 0, len(oldMap.values)),
							values:     oldMap.values,
							escapeHTML: escapeHTML,
						}
						if err = decodeMap(dec, &newMap); err != nil {
							return err
						}
						s[index] = newMap
					} else if err = decodeMap(dec, &UnsortedMap{}); err != nil {
						return err
					}
				} else if err = decodeMap(dec, &UnsortedMap{}); err != nil {
					return err
				}
			case arrayStart:
				if index < len(s) {
					if values, ok := s[index].([]interface{}); ok {
						if err = decodeSlice(dec, values, escapeHTML); err != nil {
							return err
						}
					} else if err = decodeSlice(dec, []interface{}{}, escapeHTML); err != nil {
						return err
					}
				} else if err = decodeSlice(dec, []interface{}{}, escapeHTML); err != nil {
					return err
				}
			case arrayEnd:
				return nil
			}
		}
	}
}
