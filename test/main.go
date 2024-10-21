package main

import (
	"encoding/hex"
	"fmt"
	asn1 "github.com/kitandara/asn1ber"
)

func main() {

	d := asn1.NewBerEnum(4)
	w := asn1.NewReversedIOWriter()

	d.Encode(w)

	xd := w.GetBytes()

	fmt.Println("WE got: " + hex.EncodeToString(xd))
}
