package configuration_loader

import (
	"bytes"
	"strings"
	"testing"
)

type testJsonConfig struct {
	Name     string      `json:"name"`
	Type     int         `json:"type"`
	Comments []string    `json:"comments"`
	Real     bool        `json:"real"`
	Tester   interface{} `json:"tester"`
}

func Test_jsonConfiguration_Load(t *testing.T) {
	type fields struct {
		scheme interface{}
	}
	type args struct {
		json string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "empty",
			fields:  fields{scheme: &jsonConfiguration{}},
			args:    args{json: "{}"},
			wantErr: false,
		},
		{
			name:    "good",
			fields:  fields{scheme: &jsonConfiguration{}},
			args:    args{json: "{\"name\":\"Ok\",\"type\":0,\"comments\":[\"Test\"],\"real\":true,\"tester\":null}"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jsonConfiguration{
				scheme: tt.fields.scheme,
			}
			reader := strings.NewReader(tt.args.json)
			if err := j.Load(reader); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
