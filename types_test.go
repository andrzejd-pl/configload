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
			fields:  fields{scheme: &testJsonConfig{}},
			args:    args{json: "{}"},
			wantErr: false,
		},
		{
			name:    "good",
			fields:  fields{scheme: &testJsonConfig{}},
			args:    args{json: "{\"name\":\"Ok\",\"type\":0,\"comments\":[\"Test\"],\"real\":true,\"tester\":null}"},
			wantErr: false,
		},
		{
			name:    "good",
			fields:  fields{scheme: &testJsonConfig{}},
			args:    args{json: "{\"name\":[],\"type\":0,\"comments\":[\"Test\"],\"real\":true,\"tester\":null}"},
			wantErr: true,
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

func Test_jsonConfiguration_Save(t *testing.T) {
	type fields struct {
		scheme interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		wantWriter string
		wantErr    bool
	}{
		{
			name: "empty",
			fields: fields{
				scheme: testJsonConfig{
					Name:     "",
					Type:     0,
					Comments: []string{},
					Real:     false,
					Tester:   nil,
				},
			},
			wantWriter: "{\"name\":\"\",\"type\":0,\"comments\":[],\"real\":false,\"tester\":null}",
			wantErr:    false,
		},
		{
			name: "good",
			fields: fields{
				scheme: testJsonConfig{
					Name:     "Test",
					Type:     3,
					Comments: []string{"www"},
					Real:     true,
					Tester: testJsonConfig{
						Name:     "",
						Type:     0,
						Comments: []string{},
						Real:     false,
						Tester:   nil,
					},
				},
			},
			wantWriter: "{\"name\":\"Test\",\"type\":3,\"comments\":[\"www\"],\"real\":true," +
				"\"tester\":{\"name\":\"\",\"type\":0,\"comments\":[],\"real\":false,\"tester\":null}}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jsonConfiguration{
				scheme: tt.fields.scheme,
			}
			writer := &bytes.Buffer{}
			err := j.Save(writer)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Save() gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
