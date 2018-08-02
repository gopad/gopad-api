package swagger

import (
	_ "golang.org/x/net/webdav"
)

//go:generate swagger generate spec -o swagger.json
//go:generate fileb0x ab0x.yaml
