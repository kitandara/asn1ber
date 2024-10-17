package asn1ber

import (
	"io"
)

type BerPrintableString struct {
	BerOctetString
}

var printablestringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, PRINTABLE_STRING_TAG)

func NewBerPrintableString(value []byte) *BerPrintableString {
	return &BerPrintableString{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerPrintableString) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(printablestringTag, reversedWriter, withTagList...)
}

func (b *BerPrintableString) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(printablestringTag, input, withTagList...)
}
