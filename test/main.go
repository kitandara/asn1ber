package main

import (
	 
	types"
	"encoding/hex"
	"fmt"
)

func main() {

	d := NewBerEnum(4)
	w := .NewReversedIOWriter()

	d.Encode(w)

	xd := w.GetBytes()

	fmt.Println("WE got: " + hex.EncodeToString(xd))
}
