package go_convert

import (
	"encoding/base64"
	"encoding/json"
	"github.com/advancemg/go-convert/test-data"
	"reflect"
	"testing"
)

func TestFileXmlToJson(t *testing.T) {
	xml := []string{
		"./test-data/xml/GetAdvMessagesI.xml",
		"./test-data/xml/GetAdvMessagesO.xml",
		"./test-data/xml/GetBudgetsI.xml",
		"./test-data/xml/GetBudgetsO.xml",
		"./test-data/xml/GetChannelsI.xml",
		"./test-data/xml/GetChannelsO.xml",
		"./test-data/xml/GetCustomersWithAdvertisersI.xml",
		"./test-data/xml/GetCustomersWithAdvertisersO.xml",
		"./test-data/xml/GetDeletedSpotInfoI.xml",
		"./test-data/xml/GetDeletedSpotInfoO.xml",
		"./test-data/xml/GetRanksI.xml",
		"./test-data/xml/GetRanksO.xml",
		"./test-data/xml/GetSpotsI.xml",
		"./test-data/xml/GetSpotsO.xml",
		"./test-data/xml/AddSpotI.xml",
		"./test-data/xml/AddSpotO.xml",
		"./test-data/xml/DeleteSpotI.xml",
		"./test-data/xml/DeleteSpotO.xml",
		"./test-data/xml/ChangeSpotI.xml",
		"./test-data/xml/ChangeSpotO.xml",
		"./test-data/xml/SetSpotPositionI.xml",
		"./test-data/xml/SetSpotPositionO.xml",
		"./test-data/xml/ChangeFilmsI.xml",
		"./test-data/xml/ChangeFilmsO.xml",
		"./test-data/xml/AddMPlanI.xml",
		"./test-data/xml/AddMPlanO.xml",
		"./test-data/xml/AddMPlanFilmI.xml",
		"./test-data/xml/AddMPlanFilmO.xml",
		"./test-data/xml/DeleteMPlanFilmI.xml",
		"./test-data/xml/DeleteMPlanFilmO.xml",
		"./test-data/xml/ChangeMPlanFilmPlannedInventoryI.xml",
		"./test-data/xml/ChangeMPlanFilmPlannedInventoryO.xml",
		"./test-data/xml/GetMPLansI.xml",
		"./test-data/xml/GetMPLansO.xml",
		"./test-data/xml/GetProgramBreaksLight.xml",
		"./test-data/xml/GetProgramBreaksO.xml",
	}
	jsonItems := []string{
		test_data.JsGetAdvMessagesI,
		test_data.JsGetAdvMessagesO,
		test_data.JsGetBudgetsI,
		test_data.JsGetBudgetsO,
		test_data.JsGetChannelsI,
		test_data.JsGetChannelsO,
		test_data.JsGetCustomersWithAdvertisersI,
		test_data.JsGetCustomersWithAdvertisersO,
		test_data.JsGetDeletedSpotInfoI,
		test_data.JsGetDeletedSpotInfoO,
		test_data.JsGetRanksI,
		test_data.JsGetRanksO,
		test_data.JsGetSpotsI,
		test_data.JsGetSpotsO,
		test_data.JsAddSpotI,
		test_data.JsAddSpotO,
		test_data.JsDeleteSpotI,
		test_data.JsDeleteSpotO,
		test_data.JsChangeSpotI,
		test_data.JsChangeSpotO,
		test_data.JsSetSpotPositionI,
		test_data.JsSetSpotPositionO,
		test_data.JsChangeFilmsI,
		test_data.JsChangeFilmsO,
		test_data.JsAddMPlanI,
		test_data.JsAddMPlanO,
		test_data.JsAddMPlanFilmI,
		test_data.JsAddMPlanFilmO,
		test_data.JsDeleteMPlanFilmI,
		test_data.JsDeleteMPlanFilmO,
		test_data.JsChangeMPlanFilmPlannedInventoryI,
		test_data.JsChangeMPlanFilmPlannedInventoryO,
		test_data.JsGetMPLansI,
		test_data.JsGetMPLansO,
		test_data.JsGetProgramBreaksLight,
		test_data.JsGetProgramBreaksO,
	}

	for i, v := range xml {
		js, err := FileXmlToJson(v)
		if err != nil {
			t.Errorf("FileXmlToJson() error = %v, wantErr %v", err, false)
			return
		}
		js2 := jsonItems[i]
		if js2 != "" {
			got, err := jsonBytesEqual(js, []byte(js2))
			if err != nil {
				t.Errorf("FileXmlToJson() error = %v, wantErr %v", err, false)
				return
			}
			if got != true {
				t.Errorf("FileXmlToJson() got = %v, want %v", got, true)
			}
		}
	}
}

func BenchmarkFileXmlToJson(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := FileXmlToJson("./test-data/xml/GetProgramBreaksLight.xml")
		if err != nil {
			panic(err)
		}
	}
}

func jsonBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}

func TestZipXmlToJson(t *testing.T) {
	decodeBytes, _ := base64.StdEncoding.DecodeString(`H4sIAAAAAAAA/7SRQUrEMBSGr5IbvBfpjBUe2TQgWehGcZ8ZAx2IaWl6geIJ3IqLHkEEQQr2DK83Eq22togr3eV9P/z8H6HKxbII0Z26OsttCM5HRZ8vRUYriTJNjrcERiu6yIuqPrc3TnHLd3wv+IH7oeGOn/iRYI7JRH2IdufdtUKCbxed2UPI8uDHZkSCidA++MyFurLeaJUc4ZZggWi3L+crlQQL8B6P6xazvijBJLY23MhkbXh5JSQiCm6H26H5aOv5WXDHL/zK/R+on/yivkn/Sx1++vK3AAAA//+RE7FMCAIAAA==`)
	tests := []struct {
		name    string
		field   []byte
		want    []byte
		wantErr bool
	}{
		{
			name:    "read gzip Xml data to Json",
			field:   decodeBytes,
			want:    []byte(`{"responseGetChannels":[{"Channel":{"ID":"1018476","IsDisabled":"0","MainChnl":"1018400","ShortName":"РБК Москва","bcpCentralID":"81","bcpName":"Москва","cnlCentralID":"4206"}},{"Channel":{"ID":"1018514","IsDisabled":"0","MainChnl":"1018409","ShortName":"TV 1000 Русское кино Москва","bcpCentralID":"81","bcpName":"Москва","cnlCentralID":"4258"}}]}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ZipXmlToJson(tt.field)
			if (err != nil) != tt.wantErr {
				t.Errorf("ZipXmlToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ZipXmlToJson() got = %v, want %v", got, tt.want)
			}
		})
	}
}
