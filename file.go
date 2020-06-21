package configuration_loader

import "os"

type fileConfiguration struct {
	filePath string
}

type jsonFileConfiguration struct {
	fileConfiguration
	jsonConfiguration
}

func NewJsonFileConfiguration(filePath string, jsonScheme interface{}) FileConfigurable {
	return &jsonFileConfiguration{
		fileConfiguration: fileConfiguration{filePath: filePath},
		jsonConfiguration: jsonConfiguration{scheme: jsonScheme},
	}
}

func (j *jsonFileConfiguration) LoadFromFile() error {
	file, err := os.Open(j.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return j.Load(file)
}

func (j *jsonFileConfiguration) SaveToFile() error {
	file, err := os.OpenFile(j.filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return j.Save(file)
}
