# configuration-loader

This is a [golang](https://golang.org/) package to simple load and save configuration for application.

## How to use

### Importing
```go
import "github.com/andrzejd-pl/configuration-loader"
```

### Json files

#### Importing
```go
package main
//...
import "github.com/andrzejd-pl/configuration-loader"
//...
configLoader := configuration-loader.NewJsonFileConfiguration("config.json", &configStruct)
configLoader.LoadFromFile()
//...
```

#### Saving
```go
package main
//...
import "github.com/andrzejd-pl/configuration-loader"
//...
configLoader := configuration-loader.NewJsonFileConfiguration("config.json", &configStruct)
configLoader.SaveToFile()
//...
```

## TODO

- load from yaml
- load from xml
- load from `.env` file
- load to go env