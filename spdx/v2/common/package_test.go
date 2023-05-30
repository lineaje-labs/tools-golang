// SPDX-License-Identifier: Apache-2.0 OR GPL-2.0-or-later

package common

import (
	"reflect"
	"testing"
)

func TestOriginator_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name           string
		args           args
		wantOriginator *Originator
		wantErr        bool
	}{
		{
			name: "Test for NOASSERTION",
			args: args{data: []byte("NOASSERTION")},
			wantOriginator: &Originator{
				Originator:     "NOASSERTION",
				OriginatorType: "",
			},
			wantErr: false,
		},
		{
			name: "Test for Package originator being a Person",
			args: args{data: []byte("Person: John Doe (jdoe@example.com)")},
			wantOriginator: &Originator{
				Originator:     "John Doe (jdoe@example.com)",
				OriginatorType: "Person",
			},
			wantErr: false,
		},
		{
			name: "Test for Package originator being an Organization",
			args: args{data: []byte("Organization: John Doe, Inc.")},
			wantOriginator: &Originator{
				Originator:     "John Doe, Inc.",
				OriginatorType: "Organization",
			},
			wantErr: false,
		},
		{
			name: "Test for Package originator being an Organization generated by SPDX Maven Plugin",
			args: args{data: []byte("Organization:FasterXML")},
			wantOriginator: &Originator{
				Originator:     "FasterXML",
				OriginatorType: "Organization",
			},
			wantErr: false,
		},
		{
			name: "Test for Package originator without the correct type",
			args: args{data: []byte("UnknownType: John Doe (jdoe@example.com)")},
			wantOriginator: &Originator{
				Originator:     "",
				OriginatorType: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := new(Originator)
			if err := o.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(tt.wantOriginator, o) {
				t.Fatalf("unexpected value: %v != %v", tt.wantOriginator, o)
			}
		})
	}
}
