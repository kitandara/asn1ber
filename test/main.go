package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	asn1 "github.com/kitandara/asn1ber"
	"io"
)

var (
	_ = io.Discard
	_ = fmt.Sprint
	_ = hex.Decode
	_ = errors.New
	_ = bytes.NewBuffer
)

func main() {

	_ = bytes.NewBuffer
	d := asn1.NewBerEnum(4)
	w := asn1.NewReversedIOWriter()

	n := new(asn1.BerOctetString).GetTag()
	var b []*asn1.BerTag = nil

	b = append(b, nil)
	d.Encode(w)

	xd := w.GetBytes()

	fmt.Println("WE got: %s, %d ", hex.EncodeToString(xd), n)
}
