package main

import (
	"encoding/hex"
	"fmt"
	asn1 "github.com/kitandara/asn1ber"
)

func main() {

	d := asn1.NewBerEnum(4)
	w := asn1.NewReversedIOWriter()

	n := new(asn1.BerOctetString).GetTag()
	var b []int = nil

	b = append(b, 4)
	d.Encode(w)

	xd := w.GetBytes()

	fmt.Println("WE got: %s, %d ", hex.EncodeToString(xd), n)
}
