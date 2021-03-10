# slidingwindow

[![GitHub release](https://img.shields.io/github/release/gaius-qi/slidingwindow.svg)](https://github.com/gaius-qi/slidingwindow/releases)
[![Github Build Status](https://github.com/gaius-qi/slidingwindow/workflows/Go/badge.svg?branch=main)](https://github.com/gaius-qi/slidingwindow/actions?query=workflow%3AGo+branch%3Amain)
[![GoDoc](https://godoc.org/github.com/gaius-qi/slidingwindow?status.svg)](https://godoc.org/github.com/gaius-qi/slidingwindow)

Sliding window implementation for Golang

## Installation

```shell
$ go get github.com/gaius-qi/slidingwindow
```

## Usage

```js
package main

import (
    "fmt"

    "github.com/gaius-qi/slidingwindow"
)

func main() {
    window := slidingwindow.NewWindow(10, 4)
    window.AddCount(0)
    fmt.Printf("Window moved to %d, window.Count())
}
```

## Issues

- [Open an issue in GitHub](https://github.com/gaius-qi/slidingwindow/issues)

## License

[MIT](LICENSE)
