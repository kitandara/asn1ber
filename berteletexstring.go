package asn1ber

import (
	"io"
)

type BerTeletexString struct {
	BerOctetString
}

var teletexstringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, TELETEX_STRING_TAG)

func NewBerTeletexString(value []byte) *BerTeletexString {
	return &BerTeletexString{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerTeletexString) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(teletexstringTag, reversedWriter, withTagList...)
}

func (b *BerTeletexString) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(teletexstringTag, input, withTagList...)
}
