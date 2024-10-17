package asn1ber

import (
	"io"
)

type BerNumericString struct {
	BerOctetString
}

var numericstringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, NUMERIC_STRING_TAG)

func NewBerNumericString(value []byte) *BerNumericString {
	return &BerNumericString{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerNumericString) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(numericstringTag, reversedWriter, withTagList...)
}

func (b *BerNumericString) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(numericstringTag, input, withTagList...)
}
