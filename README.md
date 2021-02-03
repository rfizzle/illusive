# Illusive API Library

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/rfizzle/illusive)

A golang API library for Illusive appliances.


## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contributing](#contributing)
- [License](#license)

## Install

This project uses [go modules](https://golang.org/ref/mod)

You can install it locally by running:

```sh
$ go get github.com/rfizzle/illusive
```

## Usage

Import the library into your project, setup the client, and call the API with valid parameters:

```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/rfizzle/illusive"
	"log"
	"strings"
)

func main() {
	client, err := illusive.NewClient(
		"https://illusive.example.com",
		"Basic <Token>",
		illusive.ClientDisableTLSValidation,
		illusive.ClientTimeout(60),
	)

	if err != nil {
		log.Fatal(err)
	}

	results, err := client.ForensicsIncidentSummary("1")
	summary := make([]string, 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range results {
		tmpText := v.Template
		for k, v2 := range v.TemplateData {
			tmpText = strings.ReplaceAll(tmpText, k, v2)
		}
		summary = append(summary, tmpText)
	}

	for k, v := range summary {
		fmt.Printf("%s:\n%s\n--------\n", results[k].AnalysisType, v)
	}
}
```

## Maintainers

[@rfizzle](https://github.com/rfizzle)

## Contributing

Feel free to dive in! [Open an issue](https://github.com/rfizzle/illusive/issues/new) or submit PRs.

## License

[MIT](LICENSE) © Coleton Pierson