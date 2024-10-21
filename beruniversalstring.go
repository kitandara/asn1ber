package asn1ber

import (
	"io"
)

type BerUniversalString struct {
	BerOctetString
}

var universalstringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, UNIVERSAL_STRING_TAG)

func NewBerUniversalString(value []byte) *BerUniversalString {
	return &BerUniversalString{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerUniversalString) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(universalstringTag, reversedWriter, withTagList...)
}

func (b *BerUniversalString) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(universalstringTag, input, withTagList...)
}
func (b *BerUniversalString) GetTag() *BerTag {
	return universalstringTag
}
