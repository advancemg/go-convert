package go_convert

import (
	"bytes"
	test_data "github.com/advancemg/go-convert/test-data"
	"github.com/ajankovic/xdiff"
	"github.com/ajankovic/xdiff/parser"
	"io/ioutil"
	"testing"
)

func TestJsonToXml(t *testing.T) {
	xml := []string{
		"./test-data/xml/GetAdvMessagesI.xml",
		"./test-data/xml/GetBudgetsI.xml",
		"./test-data/xml/GetChannelsI.xml",
		"./test-data/xml/GetCustomersWithAdvertisersI.xml",
		"./test-data/xml/GetDeletedSpotInfoI.xml",
		"./test-data/xml/GetRanksI.xml",
		"./test-data/xml/GetSpotsI.xml",
		"./test-data/xml/DeleteSpotI.xml",
		"./test-data/xml/ChangeSpotI.xml",
		"./test-data/xml/SetSpotPositionI.xml",
		"./test-data/xml/ChangeFilmsI.xml",
		"./test-data/xml/DeleteMPlanFilmI.xml",
		"./test-data/xml/ChangeMPlanFilmPlannedInventoryI.xml",
		"./test-data/xml/GetMPLansI.xml",
	}
	json := []string{
		test_data.JsGetAdvMessagesI,
		test_data.JsGetBudgetsI,
		test_data.JsGetChannelsI,
		test_data.JsGetCustomersWithAdvertisersI,
		test_data.JsGetDeletedSpotInfoI,
		test_data.JsGetRanksI,
		test_data.JsGetSpotsI,
		test_data.JsDeleteSpotI,
		test_data.JsChangeSpotI,
		test_data.JsSetSpotPositionI,
		test_data.JsChangeFilmsI,
		test_data.JsDeleteMPlanFilmI,
		test_data.JsChangeMPlanFilmPlannedInventoryI,
		test_data.JsGetMPLansI,
	}
	for i, v := range json {
		toXml, err := JsonToXml([]byte(v))
		if err != nil {
			t.Errorf("JsonToXml() error = %v, wantErr %v", err, false)
			return
		}
		baseXml, err := ioutil.ReadFile(xml[i])
		if err != nil {
			t.Errorf("ReadFile() error = %v", err)
			return
		}
		got, err := xmlDiff(toXml, baseXml)
		if err != nil {
			t.Errorf("JsonToXml() error = %v, wantErr %v", err, false)
			return
		}
		if got != true {
			t.Errorf("JsonToXml() got = %v, want %v", got, true)
		}
	}
}

func BenchmarkJsonToXml(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := JsonToXml([]byte(test_data.JsGetAdvMessagesI))
		if err != nil {
			panic(err)
		}
	}
}

func xmlDiff(a, b []byte) (bool, error) {
	p := parser.New()
	left, err := p.ParseBytes(a)
	if err != nil {
		return false, err
	}
	right, err := p.ParseBytes(b)
	if err != nil {
		return false, err
	}
	diff, err := xdiff.Compare(left, right)
	if err != nil {
		return false, err
	}
	buf := new(bytes.Buffer)
	enc := xdiff.NewTextEncoder(buf)
	if err := enc.Encode(diff); err != nil {
		return false, err
	}
	if buf.String() == "No difference.\n" {
		return true, nil
	}
	return false, nil
}
