[![Go Report Card](https://goreportcard.com/badge/github.com/mskrha/ripestat)](https://goreportcard.com/report/github.com/mskrha/ripestat)

## ripestat

### Description
Golang library for accessing RIPE Stat API (https://stat.ripe.net/docs/data_api).

### Important note
This project is at the very beginning, so coverage of the API is very low.

### Installation
`go get github.com/mskrha/ripestat`

### Example usage
```go
package main

import (
	"fmt"

	"github.com/mskrha/ripestat"
)

func main() {
	rs := ripestat.New()

	/*
		If you have registered your sourceapp
	*/
	rs.SetSourceApp("your sourceapp")

	ni, err := rs.GetNetworkInfo("8.8.8.8")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ni)
}
```
