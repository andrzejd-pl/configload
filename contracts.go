package configuration_loader

import "io"

type Configurable interface {
	Load(reader io.Reader) error
	Save(writer io.Writer) error
	ToEnv() error
}

type FileConfigurable interface {
	LoadFromFile() error
	SaveToFile() error
}
