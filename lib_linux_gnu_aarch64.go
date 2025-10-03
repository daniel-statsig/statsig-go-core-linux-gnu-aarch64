package main

import (
	_ "embed"
)

//go:embed libstatsig_ffi.so
var binaryData []byte

//go:embed libstatsig_ffi.so.sig
var signatureData []byte

func GetBinaryData() []byte {
	return binaryData
}

func GetSignatureData() []byte {
	return signatureData
}
