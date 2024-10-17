package asn1ber

import (
	"io"
)

type BerUTF8String struct {
	BerOctetString
}

var utf8stringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, UTF8_STRING_TAG)

func NewBerUTF8String(value []byte) *BerUTF8String {
	return &BerUTF8String{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerUTF8String) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(utf8stringTag, reversedWriter, withTagList...)
}

func (b *BerUTF8String) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(utf8stringTag, input, withTagList...)
}
