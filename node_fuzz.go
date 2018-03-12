// +build gofuzz

package cbornode

import (
	"bytes"

	mh "github.com/multiformats/go-multihash"
)

func Fuzz(data []byte) int {
	nd1, err := Decode(data, mh.SHA2_256, -1)
	if err != nil {
		return 0
	}

	nd2, err := Decode(nd1.RawData(), mh.SHA2_256, -1)
	if err != nil {
		panic(err)
	}

	if !nd2.Cid().Equals(nd1.Cid()) || !bytes.Equal(nd2.RawData(), nd1.RawData()) {
		panic("re-decoding a canonical node should be idempotent")
	}

	return 1
}
