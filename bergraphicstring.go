package asn1ber

import (
	"io"
)

type BerGraphicString struct {
	BerOctetString
}

var graphicStringTag = NewBerTag(UNIVERSAL_CLASS, PRIMITIVE, GRAPHIC_STRING_TAG)

func NewBerGraphicString(value []byte) *BerGraphicString {
	return &BerGraphicString{BerOctetString: *NewBerOctetString(value)}
}

func (b *BerGraphicString) Encode(reversedWriter io.Writer, withTagList ...bool) (int, error) {
	return b.EncodeUsingTag(graphicStringTag, reversedWriter, withTagList...)
}

func (b *BerGraphicString) Decode(input io.Reader, withTagList ...bool) (int, error) {
	return b.DecodeUsingTag(graphicStringTag, input, withTagList...)
}
func (b *BerGraphicString) GetTag() *BerTag {
	return graphicStringTag
}
