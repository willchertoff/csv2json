# csv2json

**Simple Binary to convert csv to arbitrary JSON. Many CSV go libraries require a struct to marshall json into. This small library simply returns JSON without having to provide a struct.**

Use it in bash or use it in go.

## CLI

```sh
# Install binary
go get github.com/willchertoff/csv2json/cmd/csv2json

# File Path Usage
csv2json [filepath] [spearator=comma|tab]

# Stdin Usage
cat test.csv | csv2json
```

## Go

```go
package mytool

import (
	//...
    "gitlab.com/linc.cloud/metadata-extractor/metadata"
    //...
)

path := "file.csv"
file, err := os.Open(path)
fileBytes, err := ioutil.ReadAll(file)
jsonBytes := converter.ConvertCSVToJSON(fileBytes, '\t')

// ... Do something with jsonBytes

```
