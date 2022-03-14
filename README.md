Go package that converts XML to JSON

### Install

    go get -u github.com/advancemg/go-convert

### Importing

    import github.com/advancemg/go-convert

### Usage

**Code example**

```go
package main

import (
	"fmt"
	converter "github.com/advancemg/go-convert"
)

func main() {
	xml := []byte(`<?xml version="1.0" encoding="UTF-8"?><hello>world</hello>`)
	json, err := converter.XmlToJson(xml)
	if err != nil {
		panic("Error converter.XmlToJson...")
	}
	fmt.Println(string(json))
	// {"hello":"world"}

	toXml, err := converter.JsonToXml(json)
	if err != nil {
		panic("Error converter.JsonToXml...")
	}
	fmt.Println(string(toXml))
	//<hello>world</hello>
}
```

**Input**

```xml
  <responseProgramBreaksV2Light xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <BreakList>
        <b BlockID="117952183" HasAucSpots="false" VM="180" VR="0">
            <Booked RankID="2" VM="175" VR="0"/>
        </b>
        <b BlockID="117952235" HasAucSpots="false" VM="180" VR="0"/>
        <b BlockID="117952251" HasAucSpots="false" VM="180" VR="0">
            <Booked RankID="2" VM="125" VR="0"/>
            <Booked RankID="3" VM="10" VR="0"/>
        </b>
    </BreakList>
</responseProgramBreaksV2Light>
```

**Output**

```json
{
  "responseProgramBreaksV2Light": {
    "BreakList": [
      {
        "b": {
          "Booked": {
            "attributes": {
              "RankID": "2",
              "VM": "175",
              "VR": "0"
            }
          },
          "attributes": {
            "BlockID": "117952183",
            "HasAucSpots": "false",
            "VM": "180",
            "VR": "0"
          }
        }
      },
      {
        "b": {
          "attributes": {
            "BlockID": "117952235",
            "HasAucSpots": "false",
            "VM": "180",
            "VR": "0"
          }
        }
      },
      {
        "b": {
          "Booked": [
            {
              "attributes": {
                "RankID": "2",
                "VM": "125",
                "VR": "0"
              }
            },
            {
              "attributes": {
                "RankID": "3",
                "VM": "10",
                "VR": "0"
              }
            }
          ],
          "attributes": {
            "BlockID": "117952251",
            "HasAucSpots": "false",
            "VM": "180",
            "VR": "0"
          }
        }
      }
    ],
    "attributes": {
      "xsi": "http://www.w3.org/2001/XMLSchema-instance"
    }
  }
}
```

**Input**

```json
{
  "GetAdvMessages": {
    "Advertisers": {
      "ID": "1"
    },
    "AdvertisingMessageIDs": [
      {
        "ID": "1"
      },
      {
        "ID": "2"
      }
    ],
    "Aspects": {
      "ID": "2"
    },
    "CreationDateEnd": "2019-03-02",
    "CreationDateStart": "2019-03-01",
    "FillMaterialTags": "true",
    "attributes": {
      "xsi": "http://www.w3.org/2001/XMLSchema-instance"
    }
  }
}
```

**Output**

```xml
<GetAdvMessages xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <Advertisers>
        <ID>1</ID>
    </Advertisers>
    <AdvertisingMessageIDs>
        <ID>1</ID>
        <ID>2</ID>
    </AdvertisingMessageIDs>
    <Aspects>
        <ID>2</ID>
    </Aspects>
    <CreationDateEnd>2019-03-02</CreationDateEnd>
    <CreationDateStart>2019-03-01</CreationDateStart>
    <FillMaterialTags>true</FillMaterialTags>
</GetAdvMessages>

```

## Benchmark

```shell
go test -bench=. -benchmem
```
XmlToJson

| Op    | ns/op | B/op  | allocs/op |
|-------|-------|-------|-----------|
| 21218 | 54522 | 41222 | 581       |

JsonToXml

| Op    | ns/op | B/op  | allocs/op |
|-------|-------|-------|-----------|
| 68661 | 15192 | 18615 | 226       |

### Contributing
Feel free to contribute to this project if you want to fix/extend/improve it.

## Licence
[MIT](https://opensource.org/licenses/MIT)
