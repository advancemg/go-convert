package go_convert

import (
	"bytes"
	"github.com/ajankovic/xdiff"
	"github.com/ajankovic/xdiff/parser"
	"reflect"
	"testing"
)

func TestJsonToXml(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "json to xml",
			args: args{
				input: []byte(`{"attributes": {"xmlns:xsi": "\"http://www.w3.org/2001/XMLSchema-instance\""}, "GetAdvMessages": {"CreationDateEnd": "2019-03-02", "CreationDateStart": "2019-03-01","Advertisers": {"ID": "1" }, "AdvertisingMessageIDs": [{"ID": "1"}, {"ID": "2"}], "Aspects": {"ID": "2" }, "FillMaterialTags": "true"}}`),
			},
			want:    []byte(`<GetAdvMessages xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><CreationDateEnd>2019-03-02</CreationDateEnd><CreationDateStart>2019-03-01</CreationDateStart><Advertisers><ID>1</ID></Advertisers><AdvertisingMessageIDs><ID>1</ID><ID>2</ID></AdvertisingMessageIDs><Aspects><ID>2</ID></Aspects><FillMaterialTags>true</FillMaterialTags></GetAdvMessages>`),
			wantErr: false,
		},
		{
			name: "json to xml",
			args: args{
				input: []byte(`{"GetBudgets": {"AdvertiserList": {"ID": "700064621"},"ChannelList": {"ID": "1018574"},"EndMonth": "20180115","SellingDirectionID": "21","StartMonth": "201801"}}`),
			},
			want:    []byte(`<GetBudgets><AdvertiserList><ID>700064621</ID></AdvertiserList><ChannelList><ID>1018574</ID></ChannelList><EndMonth>20180115</EndMonth><SellingDirectionID>21</SellingDirectionID><StartMonth>201801</StartMonth></GetBudgets>`),
			wantErr: false,
		},
		{
			name: "json to xml",
			args: args{
				input: []byte(`{"GetChannels": {"SellingDirectionID": "21"}}`),
			},
			want:    []byte(`<GetChannels><SellingDirectionID>21</SellingDirectionID></GetChannels>`),
			wantErr: false,
		},
		{
			name: "json to xml",
			args: args{
				input: []byte(`{"attributes": {"xmlns:xsi": "\"http://www.w3.org/2001/XMLSchema-instance\""}, "AddSpot": {"BlockID": "1","FilmID": "1","Position": "nil","FixedPosition": "true","AuctionBidValue": "nil"}}`),
			},
			want:    []byte(`<AddSpot xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><BlockID>1</BlockID><FilmID>1</FilmID><Position xsi:nil="true"></Position><FixedPosition>true</FixedPosition><AuctionBidValue xsi:nil="true"></AuctionBidValue></AddSpot>`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := JsonToXml(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("JsonToXml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JsonToXml() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkJsonToXml(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := JsonToXml([]byte(`{"attributes": {"xmlns:xsi": "\"http://www.w3.org/2001/XMLSchema-instance\""}, "GetAdvMessages": {"CreationDateEnd": "2019-03-02", "CreationDateStart": "2019-03-01","Advertisers": {"ID": "1" }, "AdvertisingMessageIDs": [{"ID": "1"}, {"ID": "2"}], "Aspects": {"ID": "2" }, "FillMaterialTags": "true"}}`))
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
