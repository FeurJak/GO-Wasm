// See https://github.com/golang/go/issues/26366.
package lib

import (
	_ "github.com/FeurJak/GO-Wasm/pkg/wasmer/packaged/lib/darwin-amd64"
	_ "github.com/FeurJak/GO-Wasm/pkg/wasmer/packaged/lib/linux-aarch64"
	_ "github.com/FeurJak/GO-Wasm/pkg/wasmer/packaged/lib/linux-amd64"
)
