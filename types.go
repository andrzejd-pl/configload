package configuration_loader

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type jsonConfiguration struct {
	scheme interface{}
}

func (j *jsonConfiguration) Load(reader io.Reader) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, j.scheme)
}

func (j *jsonConfiguration) Save(writer io.Writer) error {
	data, err := json.Marshal(j.scheme)
	if err != nil {
		return err
	}

	_, err = writer.Write(data)
	return err
}
