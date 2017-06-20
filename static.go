// +build static_liblzma

package xz

/*
#cgo LDFLAGS: -Wl,-unresolved-symbols=ignore-all
*/
import "C"

import _ "github.com/akmistry/go-liblzma/internal"
