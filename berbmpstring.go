package asn1ber

import (
	"io"
)

type BerBmpString struct {
	BerOctetString
}

var bmpstringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, BMP_STRING_TAG)

func NewBerBmpString(value []byte) *BerBmpString {
	return &BerBmpString{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerBmpString) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(bmpstringTag, reversedWriter, withTagList...)
}

func (b *BerBmpString) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(bmpstringTag, input, withTagList...)
}
func (b *BerBmpString) GetTag() *BerTag {
	return bmpstringTag
}
