package go_convert

import (
	"reflect"
	"testing"
)

func TestUnsortedMap_ToXml(t *testing.T) {
	tests := []struct {
		name    string
		field   func() *UnsortedMap
		want    []byte
		wantErr bool
	}{
		{
			name: "unsorted map to xml",
			field: func() *UnsortedMap {
				uMap := New()
				uMap.Set("CreationDateStart", "0000-00-00")
				uMap.Set("CreationDateEnd", "0000-00-00")
				return uMap
			},
			want:    []byte(`<CreationDateStart>0000-00-00</CreationDateStart><CreationDateEnd>0000-00-00</CreationDateEnd>`),
			wantErr: false,
		},
		{
			name: "unsorted map to xml with header",
			field: func() *UnsortedMap {
				header := New()
				body := New()
				body.Set("CreationDateStart", "0000-00-00")
				body.Set("Advertiser", 1)
				body.Set("SecondField", "1")
				header.Set("Header", body)
				return header
			},
			want:    []byte(`<Header><CreationDateStart>0000-00-00</CreationDateStart><Advertiser>1</Advertiser><SecondField>1</SecondField></Header>`),
			wantErr: false,
		},
		{
			name: "unsorted map to xml with array",
			field: func() *UnsortedMap {
				advertiseOne := New()
				advertiseOne.Set("ID", "1")
				advertiseSecond := New()
				advertiseSecond.Set("ID", "1")
				header := New()
				body := New()
				body.Set("CreationDateStart", "0000-00-00")
				body.Set("Advertiser", []*UnsortedMap{
					advertiseOne,
					advertiseSecond,
				})
				body.Set("SecondField", "1")
				header.Set("Header", body)
				return header
			},
			want:    []byte(`<Header><CreationDateStart>0000-00-00</CreationDateStart><Advertiser><ID>1</ID><ID>1</ID></Advertiser><SecondField>1</SecondField></Header>`),
			wantErr: false,
		},
		{
			name: "unsorted map to xml with header one attribute",
			field: func() *UnsortedMap {
				attributes := New()
				attributes.Set("xmlns:xsi", "\"http://www.w3.org/2001/XMLSchema-instance\"")
				header := New()
				body := New()
				body.Set("CreationDateStart", "0000-00-00")
				body.Set("CreationDateEnd", "0000-00-00")
				header.Set("Header", body)
				header.Set("attributes", attributes)
				return header
			},
			want:    []byte(`<Header xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><CreationDateStart>0000-00-00</CreationDateStart><CreationDateEnd>0000-00-00</CreationDateEnd></Header>`),
			wantErr: false,
		},
		{
			name: "unsorted map to xml with header many attribute",
			field: func() *UnsortedMap {
				attributes := New()
				attributes.Set("xmlns:xsi", "\"http://www.w3.org/2001/XMLSchema-instance\"")
				header := New()
				body := New()
				bodyAttributes := New()
				bodyAttributes.Set("VM", "100")
				bodyAttributes.Set("VR", "200")
				body.Set("CreationDateStart", "0000-00-00")
				body.Set("CreationDateEnd", "0000-00-00")
				body.Set("attributes", bodyAttributes)
				header.Set("Header", body)
				header.Set("attributes", attributes)
				return header
			},
			want:    []byte(`<Header xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"><CreationDateStart VM=100 VR=200>0000-00-00</CreationDateStart><CreationDateEnd VM=100 VR=200>0000-00-00</CreationDateEnd></Header>`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uMap := tt.field()
			got, err := uMap.ToXml()
			if (err != nil) != tt.wantErr {
				t.Errorf("ToXml() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToXml() got = %v, want %v", got, tt.want)
			}
		})
	}
}
