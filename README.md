# Deluge RPC

A Go Library for simplified usage of DelugeRPC

Created and documented based off a combination of the DelugeRPC documentation found here:
https://deluge-bendikro.readthedocs.io/en/latest/core/rpc.html

And manual testing using the `system.listMethod` method.

Where possible the library returns the `result` field of the JSON RPC response.
If the `error` field is present then the error is returned as a go error.

## Usage
`go get github.com/jt28828/delugerpc`

### Example

```go
package main

import (
	"fmt"
	"github.com/jt28828/delugerpc/deluge"
	"github.com/jt28828/delugerpc/dto"
)

func main() {
	client, err := deluge.NewClient("localhost", 8112, "yourPasswordHere")
	if err != nil {
		panic(err.Error())
	}

	torrents, err := client.CoreListTorrents(nil)
	if err != nil {
		panic(err.Error())
	}

	for _, torrent := range torrents {
		fmt.Println(torrent.Name)
	}
}
```

## What's Complete

The below methods have been wrapped in client methods:

- Most of the `core` methods
- All `auth` methods

## What's Missing

- some `core` methods that took dicts as an input and had little to no documentation about what those dicts consist of.
- all `daemon` methods
- all `webui` methods
- all `web` methods